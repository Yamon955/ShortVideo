# Gateway--网关服务

## 功能简介

gateway服务主要负责路由寻址、用户鉴权、HTTP请求转化成TRPC请求。

## 路由寻址

1. **路由信息写入配置文件**

```json
{
   "path": "/shortvideo/user/register",
   "is_regex": "0",
   "service": "trpc.shortvideo.user.User",
   "method": "/trpc.shortvideo.user.User/Register",
   "target": "ip://127.0.0.1:20001",
   "timeout": "2s",
   "protocol": "trpc",
   "is_allow_no_auth": true
 }
```

2. **构造基准树**

在程序开始之前读取配置文件，根据`path`字段构造`radixTree`基准树，支持精确匹配、最长前缀匹配。

另外，也可保存正则规则的路由，支持正则匹配，项目实现了但整体并未使用正则匹配。

3. **定义路由拦截器RouterFilter**

在`RouterFilter`拦截器内实现路由匹配，按照精确匹配、最长前缀匹配的先后顺序进行路由匹配。

## 用户鉴权

HTTP协议头部 Authorization 字段携带用户登录态 token 信息

```go
token := thttp.Head(ctx).Request.Header.Get("Authorization")
```

取出token调用验证函数验证其是否正确，同时返回正确的用户 uid，并将其放到 ctx 中透传

```go
return next(context.WithValue(ctx, "sv_login_uid", strconv.FormatUint(loginUID, 10)), req)
```

在 invoke 函数调用下游服务之前，从 ctx 中取出 uid 放到 client.Option 中透传到下游服务

```go
loginUID := ctx.Value("sv_login_uid")
opts = []client.Option{
		client.WithMetaData("sv_login_uid", []byte(uid)),
		client.WithMetaData("sv_trace_id", []byte(tranceID)),
	}
if err := client.DefaultClient.Invoke(ctx, reqBody, rspBody, opts...); err != nil {
		// err != nil 时，rsp.Body 会被清空
		rsp := &errRsp{
			Code: int(errs.Code(err)),
			Msg:  errs.Msg(err),
		}
		rspBody.Data, _ = json.Marshal(rsp)
	}
```

**trpc 链路透传：**

- **client 透传数据到 server**

 client 发起请求时，通过增加 option 来设置透传字段，可以增加多个透传字段。

``` go
options := []client.Option{
    client.WithMetaData("key1", []byte("v1")),
    client.WithMetaData("key2", []byte("v2")),
    client.WithMetaData("key3", []byte("v3")),
}
rsp, err := proxy.SayHello(ctx, options...) // 注意：框架传过来的 ctx
```

下游 server 通过框架的 ctx 可以获取到 client 的透传字段

```go
trpc.GetMetaData(ctx, "key1") // 注意使用框架的 ctx
```

- **server 透传数据到 client**

  server 在回包的时候可以通过 ctx 设置透传字段返回给上游调用方

```go
trpc.SetMetaData(ctx, "key1", []byte("val1")) // 注意使用框架的 ctx
```

上游 client 可以通过设置各协议的 rsp head 获取

```go
head := &trpc.ResponseProtocol{}
opts := []client.Option{
 client.WithRspHead(head),
}

rsp, err := proxy.SayHello(ctx, opts...) // 注意：框架传过来的 ctx
head.TransInfo // 框架透传回来的信息 key-value 对（map[string][]byte）
```



## HTTP转TRPC

**请求包：HTTP -> TRPC** 

**返回包：TRPC -> HTTP**

回包时注意，如果下游服务接口返回 err != nil，则 trpc.Body 会被清空，因此，为了给前端展示出错信息，这里定义了一个错误返回结构体，会将该错误信息写入到 http.rsp.body 中返回给前端

```go
// 当调用的服务返回 err != nil 时，rsp.Body 会被清空，因此，需要将 err 信息转成 json 形式放到 http.rsp 中
type errRsp struct {
    Code int    `json:"code"`
    Msg  string `json:"msg"`
}
```

## HTTP转HTTP：补充透传信息

