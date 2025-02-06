package uploader

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Yamon955/ShortVideo/comm/utils"
	"github.com/Yamon955/ShortVideo/video/entity/def"
	"github.com/Yamon955/ShortVideo/video/entity/errcode"
	"github.com/minio/minio-go/v7"
	"trpc.group/trpc-go/trpc-go/errs"
	"trpc.group/trpc-go/trpc-go/log"
)

// uploaderImpl 上传器实现
type uploaderImpl struct {
	MinIOClient *minio.Client
}

// VideoUpload 视频上传
func (u *uploaderImpl) VideoUpload(ctx context.Context, req *MediaFileUploadReq) (int64, error) {
	var (
		contentBytes []byte
		contentType  string
		videoBuffer  *bytes.Buffer
		imgBuffer    *bytes.Buffer
		err          error
	)
	contentBytes, err = io.ReadAll(req.File)
	if err != nil {
		log.ErrorContextf(ctx, "ioutil.ReadAll failed, err: %v", err)
		return 0, err
	}
	// 雪花算法生成 vid
	vid := utils.GenID()
	// 判断文件类型
	contentType = http.DetectContentType(contentBytes)
	if !strings.HasPrefix(contentType, "video") {
		return 0, errs.New(errcode.ErrNotVideoType, "非视频文件")
	}
	videoBuffer = bytes.NewBuffer(contentBytes)
	imgBuffer = bytes.NewBuffer(nil)

	// ffmpeg 提取首帧图片

	_, err = u.MinIOClient.PutObject(ctx,
		def.MinIOBucketName,
		fmt.Sprintf("video/%d/video", vid),
		videoBuffer,
		int64(videoBuffer.Len()),
		minio.PutObjectOptions{
			ContentType: contentType,
		},
	)
	if err != nil {
		log.ErrorContextf(ctx, "minioClient.PutObject failed, err: %v", err)
		return 0, err
	}

	_, err = u.MinIOClient.PutObject(ctx,
		def.MinIOBucketName,
		fmt.Sprintf("video/%d/cover", vid),
		imgBuffer,
		int64(imgBuffer.Len()),
		minio.PutObjectOptions{
			ContentType: "image/jpeg",
		},
	)
	if err != nil {
		log.ErrorContextf(ctx, "minioClient.PutObject failed, err: %v", err)
		return 0, err
	}
	return 0, nil
}
