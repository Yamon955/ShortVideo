# Video--视频模块

## 简介

该模块部署两个 server.service 

- publish 发布服务-- http 协议
- video 视频基础服务-- trpc 协议

## Publish--视频发布服务

### 1. 协议

视频发布采用**HTTP协议**，而非TRPC，原因是前端通过HTTP的**form-data形式**携带视频文件数据，在网关处HTTP->TRPC时TRPC框架反序列化form-data格式会有点问题，不能映射到指定字段上，因此采用HTTP协议。

采用HTTP协议之后，我们需要**自己注册form-data的序列化器**，重写反序列化方法为空，即不序列化请求内容，直接从 HTTP 请求包取数据即可

- 定义 form-data 序列化器

```go
const (
	SerializationTypeFormData = 131 // trpc 框架 formdata 格式序列化对应的键值
)
// FormDataSerialization form-data 格式数据序列化处理器
type FormDataSerialization struct {
	codec.JSONPBSerialization
}
// Unmarshal multipart/form-data 类型解码方法, 重写 Unmarshal 方法，对于 form-data 数据无需反序列化
func (s *FormDataSerialization) Unmarshal(in []byte, body interface{}) error {

	return nil
}
```

- Main 函数注册

```go
func init() {
	// 对于 Content-Type=multipart/form-data 的请求无需反序列化，因此注册对应的反序列化处理对象，重写其 Unmarshal 方法
	codec.RegisterSerializer(SerializationTypeFormData, &FormDataSerialization{})
}
```

### 2. 下载数据[无用，去除]

支持普通下载和分片下载，默认采用分片下载

```go
// 指定读取的文件偏移，及读取的字节数
sectionReader := io.NewSectionReader(file, start, partSize)
```

这里下载类似于读取数据，根据前端传来的视频信息，读取具体的视频二进制数据。

**TODO:实现支持传入 url 去分片下载视频**

### 3. 上传数据

#### 方案

拟采用MinIO对象存储来存储视频及图片数据

方案一：使用 kafka 解耦，将要上传的视频数据及信息分片读到 kafka 中，写入mq成功后，服务直接返回成功，后续的转码、存储异步完成。

- 视频数据过大不适合直接存到 kafka
- 如果分片存到 kafka，合并时需要消费一条信息才能读到 key，不能指定key读取，合并困难

方案二：依然使用 kafka 解耦转码/上传操作，减少用户等待时间，但是只需要将视频必要信息写入 kafka 即可，视频数据暂时写到临时存储器中，这就要求临时存储器必须具备快速写入的功能(KV键值数据库)

- 首先想到使用 redis 暂存，但是还是同样问题，视频过大，不适合 redis 存储
- LevelDB，这个感觉不错，可以把视频数据切分为 1KB 左右的分片，批量写入 leveldb 中。它是单机部署的，之后docker部署需要挂载宿主机文件。但是其写入速度有待考量，毕竟它不支持并发写入，可以考虑按照 key 取模运算设置多个 leveldb 数据库。

方案三：直接存储到 MinIO 中，取出对应的 URL 填充视频信息 ,完成后同步视频信息到 MySQL 的 videos 表中

**拟采用方案三**

 (其实**最优方案**应该是前端直接分片上传到MinIO中，上传完毕后调用后端服务进行合并（一般的对象服务都有支持分块上传的API，不需要手动合并），并且后端可以利用 redis的bitmap 存储各分片的是否上传完成，提供接口供前端访问，从而实现断点续传。参考https://www.cnblogs.com/jsonq/p/18186340，由于本项目针对后端开发，因此未采用该方案)

#### 方案三实现步骤

##### 1.从 http 协议解析出 form-data 数据【视频标签、标题、视频文件】

##### 2.文件上传前置处理

1. 雪花算法生成视频 vid

2. 判断上传的文件类型是否为视频文件

3. 定义上传的 url 文件名【http://127.0.0.1:9000/bucket_name/...】

   **注意：文件后缀匹配，视频文件后缀使用 `.mp4` 图片使用`.jpeg`**

```go
	urlPrefix := conf.GetAppConfig().CommConfig.UrlPrefix
	playUrl := fmt.Sprintf("%s/video/VID_%d/video.mp4", urlPrefix, vid)
	coverUrl := fmt.Sprintf("%s/video/VID_%d/cover.jpeg", urlPrefix, vid)
```

##### 3.调用minio API上传

```go
	_, err = u.MinIOClient.PutObject(ctx,
		def.MinIOBucketName,
		coverUrl,
		imgBuffer,
		int64(imgBuffer.Len()),
		minio.PutObjectOptions{
			ContentType: "image/jpeg",
		},
	)
```

##### 4.视频信息写入数据库

```go
	videoInfo := &base.Video{
		UID:         req.UID,
		VID:         vid,
		VideoURL:    playUrl,
		CoverURL:    coverUrl,
		Title:       req.Title,
		PublishTime: time.Now().Unix(),
	}
	// 视频信息存储到视频表中
	err = h.DB.InsertVideo(ctx, videoInfo)
```



## Video--基础视频服务

### 1. GetFeeds
获取推荐视频

### 2. GetPublishList
索引配合分页查询，每次查询返回10条数据，同时返回下一次开始查询的偏移量

`videos` 增加了自增主键 id 用于实现分页查询效果
```sql
select * from videos id > ? and uid = ? limit 10;
```

根据查询语句建立 (id, uid) 的联合索引，时查询更高效
```sql
CREATE INDEX idx_id_uid ON videos(id, uid);
```

TODO: 目前只实现了主态访问，客态访问待实现
