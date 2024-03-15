package utils

import (
	"gateway_go/global"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func Upload2Ali(imgName string) (url string, err error) {
	//provider, err := oss.NewEnvironmentVariableCredentialsProvider()
	//if err != nil {
	//	return "", err
	//}
	client, err := oss.New(global.App.Config.Storage.Disks.AliOss.Endpoint, global.App.Config.Storage.Disks.AliOss.AccessKeyId, global.App.Config.Storage.Disks.AliOss.AccessKeySecret)
	if err != nil {
		return "", err
	}

	// 指定图片所在Bucket的名称，例如examplebucket。
	bucketName := global.App.Config.Storage.Disks.AliOss.Bucket
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return "", err
	}
	// 指定图片名称。如果图片不在Bucket根目录，需携带文件完整路径，例如exampledir/example.jpg。
	ossImageName := global.App.Config.Storage.Disks.LocalStorage.RootFileDir + imgName
	// 生成带签名的URL，并指定过期时间为600s。
	signedURL, err := bucket.SignURL(ossImageName, oss.HTTPGet, 600, oss.Process("image/format,png"))
	if err != nil {
		return "", err
	}
	// 删除当前文件
	// _ = os.Remove(ossImageName)
	return signedURL, err
}
