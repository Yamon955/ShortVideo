package logic

import (
	"bytes"
	"context"
	"io/ioutil"
	"mime/multipart"
	"strconv"

	"github.com/Yamon955/ShortVideo/protocol/video/pb"
	"github.com/Yamon955/ShortVideo/video/entity/errcode"
	"github.com/Yamon955/ShortVideo/video/repo/uploader"
	"trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/errs"
	thttp "trpc.group/trpc-go/trpc-go/http"
	"trpc.group/trpc-go/trpc-go/log"
)

const (
	reqTitleName = "video_title"
	reqTagName   = "video_tags" // form-data 携带视频标签的字段名
	reqFileName  = "data"
)

func (h *handlerImpl) HandlePublish(ctx context.Context, req *pb.PublishReq) (*pb.PublishRsp, error) {
	uid, _ := strconv.ParseUint(string(trpc.GetMetaData(ctx, "sv_login_uid")), 10, 64)
	tags, title, file, _, err := parseHTTPForm(ctx)
	defer file.Close()
	if err != nil {
		log.ErrorContextf(ctx, "parseHTTPForm failed, err:%v", err)
		return nil, errs.New(errcode.ErrPublishFailed, "上传失败,请稍后重试!")
	}
	uploadReq := &uploader.MediaFileUploadReq{
		UID:   uid,
		File:  file,
		Title: title,
		Tags:  tags,
	}
	videoInfo, err := h.Uploader.VideoUpload(ctx, uploadReq)
	if err != nil {
		log.ErrorContextf(ctx, "VideoUpload failed, err:%v", err)
		return nil, errs.New(errcode.ErrPublishFailed, "上传失败,请稍后重试!")
	}
	log.Infof("uid:%d, VideoUpload success. vid:%v", uid, videoInfo.VID)

	err = h.DB.InsertVideo(ctx, videoInfo)
	if err != nil {
		log.ErrorContextf(ctx, "Insert video failed, err:%v", err)
		return nil, errs.New(errcode.ErrInsertVideo, "视频信息存储失败")
	}

	return &pb.PublishRsp{
		StatusMsg: "上传成功",
		Vid:       videoInfo.VID,
	}, nil
}

// parseHTTPForm 解析 http_form-data 请求
func parseHTTPForm(ctx context.Context) ([]string, string, multipart.File, *multipart.FileHeader, error) {
	head := thttp.Head(ctx)
	head.Request.Body = ioutil.NopCloser(bytes.NewReader(head.ReqBody))

	// 最多存储 32MB 内容到内存中，其余存储在磁盘的临时文件中
	if err := head.Request.ParseMultipartForm(32 << 20); err != nil {
		log.ErrorContextf(ctx, "ParseMultipartForm failed, err:%v", err)
		return nil, "", nil, nil, err
	}
	// 获取 form-data 值参数（kv结构）
	formValues := head.Request.MultipartForm.Value
	tags := formValues[reqTagName]
	var title string
	if len(formValues[reqTagName]) != 0 {
		title = formValues[reqTitleName][0]
	}

	// 获取 form-data 文件参数
	file, fileHeader, err := head.Request.FormFile(reqFileName) // file = fileHeader.Open()
	if err != nil {
		log.ErrorContextf(ctx, "read FormFile failed, err:%v:", err)
		return nil, "", nil, nil, err
	}
	return tags, title, file, fileHeader, nil
}

// transCode 视频转码 是不是应该在kafka消费者实现，从而保证转码可以重试直到成功
func transCode(data []byte) {

}
