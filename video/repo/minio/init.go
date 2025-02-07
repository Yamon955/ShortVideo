package minio

import (
	"github.com/Yamon955/ShortVideo/video/entity/conf"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"trpc.group/trpc-go/trpc-go/log"
)

var minIOCli *minio.Client

// Init 初始化 MinIO 连接对象
func Init() (err error) {
	endpoint := conf.GetAppConfig().MinIOConf.Endpoint
	accessKey := conf.GetAppConfig().MinIOConf.AccessKeyID
	secretKey := conf.GetAppConfig().MinIOConf.SecretAccessKey
	useSSL := false
	minIOCli, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Errorf("MinIO init failed, err: %v, conf:%v", err, conf.GetAppConfig().MinIOConf)
	}
	return err
}

// GetClient 获取连接对象
func GetClient() *minio.Client {
	return minIOCli
}
