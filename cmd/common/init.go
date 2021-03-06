package common

import (
	"fmt"

	"oss-aliyun-cli/config"
	bu "oss-aliyun-cli/domain/bucket"
	ob "oss-aliyun-cli/domain/object"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var (
	bucket   string
	object   string
	filepath string
	suffix   string

	bClient *bu.Bucket
	oClient *ob.Object
)

func init() {
	getBucketClient()
	getObjectClient()

	// Up
	CmdUp.Flags().StringVarP(&bucket, "bucket", "b", "", "aliyun-oss上的存储空间命名")
	CmdUp.Flags().StringVarP(&object, "object", "o", "", "aliyun-oss上的文件对象命名")
	CmdUp.Flags().StringVarP(&filepath, "filepath", "f", "", "指定需要上传文件或者目录的绝对路径")
	CmdUp.Flags().StringVarP(&suffix, "suffix", "s", "", "文件后缀名过滤条件")

	// Download
	CmdDownload.Flags().StringVarP(&bucket, "bucket", "b", "", "aliyun-oss上的存储空间命名")
	CmdDownload.Flags().StringVarP(&object, "object", "o", "", "aliyun-oss上的文件对象命名")
	CmdDownload.Flags().StringVarP(&filepath, "filepath", "f", "", "指定需要上传文件或者目录的绝对路径")

	// Get
	// CmdGet.Flags().StringVarP(&bucket, "bucket", "b", "", "aliyun-oss上的存储空间命名")
	// CmdGet.Flags().StringVarP(&object, "object", "o", "", "aliyun-oss上的文件对象命名")

	// Creat
	CmdCreate.Flags().StringVarP(&bucket, "bucket", "b", "", "aliyun-oss上的存储空间命名")
	CmdCreate.Flags().StringVarP(&object, "object", "o", "", "aliyun-oss上的文件对象命名")

	// Delete
	CmdDelete.Flags().StringVarP(&bucket, "bucket", "b", "", "aliyun-oss上的存储空间命名")
	CmdDelete.Flags().StringVarP(&object, "object", "o", "", "aliyun-oss上的文件对象命名")
}

func getBucketClient() {
	client, err := oss.New(config.Endpoint, config.AccessKeyID, config.AccessKeySecret)
	if err != nil {
		fmt.Printf("创建 aliyun-oss-clint 出错，具体错误如下：")
		fmt.Println("")
		fmt.Println(err)
	}

	bClient = &bu.Bucket{
		Client: client,
	}
}

func getObjectClient() {
	client, err := oss.New(config.Endpoint, config.AccessKeyID, config.AccessKeySecret)
	if err != nil {
		fmt.Printf("创建 aliyun-oss-clint 出错，具体错误如下：")
		fmt.Println("")
		fmt.Println(err)
	}

	oClient = &ob.Object{
		Client: client,
	}
}
