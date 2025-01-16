package downloader

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"time"

	"github.com/Yamon955/ShortVideo/video/entity/def"
	"trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/log"
)

// downloaderImpl 下载器实现
type downloaderImpl struct {
}

// PieceDownloadVideoFile 分片下载
func (d *downloaderImpl) PieceDownloadVideoFile(
	ctx context.Context,
	file multipart.File,
	fileHeader *multipart.FileHeader,
) (*MediaFile, error) {
	// 计算分片数量
	pieceNum := fileHeader.Size / def.DefaultPieceSize
	if fileHeader.Size%def.DefaultPieceSize != 0 {
		pieceNum += 1
	}
	data := make([][]byte, pieceNum)
	var handlers []func() error
	for i := 0; i < int(pieceNum); i++ {
		index := i
		start := int64(index * def.DefaultPieceSize)
		partSize := int64(def.DefaultPieceSize)
		if i == int(pieceNum)-1 {
			partSize = def.DefaultPieceSize - 1
		}
		handlers = append(handlers, func() error {
			sectionReader := io.NewSectionReader(file, start, partSize)
			buffer := make([]byte, partSize)
			bytesRead, err := sectionReader.Read(buffer)
			if err != nil && err != io.EOF {
				return err
			}
			data[index] = buffer[:bytesRead]
			return nil
		})
	}
	if err := trpc.GoAndWait(handlers...); err != nil {
		log.ErrorContextf(ctx, "GoAndWait cocurrent download failed, err:%v", err)
		return nil, err
	}
	res := make([]byte, 0, fileHeader.Size)
	// 合并数据
	for _, partData := range data {
		res = append(res, partData...)
	}
	videoFile := &MediaFile{
		Data: res,
		Size: fileHeader.Size,
	}
	return videoFile, nil
}

// DownloadVideoFile 普通下载
func (d *downloaderImpl) DownloadVideoFile(
	ctx context.Context,
	file multipart.File,
	fileHeader *multipart.FileHeader,
) (*MediaFile, error) {
	st := time.Now()
	// 使用缓冲区逐块读取文件内容并写入 req.Data
	var buffer bytes.Buffer
	_, err := io.Copy(&buffer, file)
	if err != nil {
		log.Errorf("err:%v", err)
		return nil, err
	}
	res := buffer.Bytes()
	elapsed := time.Since(st)
	log.Infof("函数执行时间: %d毫秒", elapsed.Milliseconds())
	size := float64(len(res) * 1.0 / 1024 / 1024)
	log.Infof("len(res) = %fMB", size)
	log.Infof("res:%v", res[:def.DefaultPieceSize/256])
	videoFile := &MediaFile{
		Data: res,
		Size: fileHeader.Size,
	}
	return videoFile, nil
}
