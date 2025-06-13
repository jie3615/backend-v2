package minio

import (
	"context"
	"metaLand/app/api/internal/config"
	"time"

	"github.com/minio/minio-go/v7"
)

func PreSignUpload(client *minio.Client, file string, c config.Config) (url string, err error) {
	u, err := client.PresignedPutObject(
		context.TODO(),
		c.Minio.Bucket,
		file,
		time.Minute*10,
	)
	if err != nil {
		return
	}
	url = u.String()
	return
}
