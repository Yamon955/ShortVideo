package router

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"sync/atomic"

	"github.com/Yamon955/ShortVideo/gateway/entity"
	"github.com/armon/go-radix"
	"trpc.group/trpc-go/trpc-go/errs"
	"trpc.group/trpc-go/trpc-go/log"
)

const (
	exactMatch         = 1 // 精确匹配
	longestPrefixMatch = 2 // 最长前缀匹配
	regexMatch         = 3 // 正则匹配
)

var (
	radixTree     atomic.Value
	regexRuleConf atomic.Value
)

type readerImpl struct {
	tree       atomic.Value
	regexRules atomic.Value
}

// Init 初始化路由配置
func Init() error {
	// Open the router.json file
	jsonFile, err := os.Open("./conf/router.json")
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	var routeConfs []entity.RouteConf
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &routeConfs)

	// build radixTree
	tree := radix.New()
	regexRules := make([]*entity.RouteConf, 0)
	for _, r := range routeConfs {
		if r.Path == "" {
			continue
		}
		if r.IsRegex == 1 {
			r.Regexp = regexp.MustCompile(r.Path)
			regexRules = append(regexRules, &r)
			continue
		}
		tree.Insert(r.Path, &r)
	}
	log.Debugf("parse router success, regex conf: %+v\n", regexRules)
	log.Debugf("parse router success, radix tree conf: %+v\n", tree.ToMap())
	radixTree.Store(tree)
	regexRuleConf.Store(regexRules)
	return nil
}

func newReaderImpl() *readerImpl {
	return &readerImpl{
		tree:       radixTree,
		regexRules: regexRuleConf,
	}
}

// GetMatchRouter 匹配路由
func (r *readerImpl) GetMatchRouter(ctx context.Context, path string) (*entity.RouteConf, error) {
	tree, err := r.loadRadixTree()
	if err != nil {
		return nil, err
	}
	regexRules, err := r.loadRegexRules()
	if err != nil {
		return nil, err
	}
	// 精确匹配
	if node, ok := tree.Get(path); ok {
		return parseRouter(node, path, exactMatch)
	}
	// 最长前缀匹配
	longestPrefix, node, ok := tree.LongestPrefix(path)
	if ok && longestPrefix != "/" {
		return parseRouter(node, path, longestPrefixMatch)
	}
	// 正则匹配
	for _, rule := range regexRules {
		if rule.Regexp.MatchString(path) {
			return parseRouter(rule, path, regexMatch)
		}
	}
	return nil, errs.New(entity.ErrMatchRoute, "not find match route")
}

func (r *readerImpl) loadRadixTree() (*radix.Tree, error) {
	val, ok := r.tree.Load().(*radix.Tree)
	if !ok {
		return nil, errs.New(entity.ErrLoadRadixTree, "radix tree load err")
	}
	return val, nil
}

func (r *readerImpl) loadRegexRules() ([]*entity.RouteConf, error) {
	val, ok := r.regexRules.Load().([]*entity.RouteConf)
	if !ok {
		return nil, errs.New(entity.ErrLoadRegexRules, "regex rules load err")
	}
	return val, nil
}

func parseRouter(node interface{}, path string, matchType int) (*entity.RouteConf, error) {
	routeConf, ok := node.(*entity.RouteConf)
	if !ok {
		return nil, errs.New(entity.ErrGetRouteConfFromNode, "get routeConf from node err")
	}
	var calleeMethod string
	switch matchType {
	case exactMatch:
		calleeMethod = routeConf.Method
	case longestPrefixMatch:
		calleeMethod = strings.TrimPrefix(path, routeConf.Method)
	case regexMatch:
		res := routeConf.Regexp.FindStringSubmatch(path)
		if len(res) == 0 {
			return nil, errs.New(entity.ErrFindRouteConfFromRegex, "get routeConf from regexp err")
		}
		calleeMethod = res[len(res)-1]
	default:
	}
	routeConf.CalleeMethod = calleeMethod
	log.Infof("match routeConf:%v", routeConf)
	return routeConf, nil
}
