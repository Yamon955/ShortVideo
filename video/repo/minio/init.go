package minio

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"trpc.group/trpc-go/trpc-go/log"
)

var minIOCli *minio.Client

// Init 初始化 MinIO 连接对象
func Init() (err error) {
	endpoint := "play.min.io"
	accessKey := "Q3AM3UQ867SPQQA43P2F"
	secretKey := "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG"
	useSSL := true
	minIOCli, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Errorf("MinIO init failed, err: %v", err)
	}
	return err
}

// GetClient 获取连接对象
func GetClient() *minio.Client {
	return minIOCli
}
