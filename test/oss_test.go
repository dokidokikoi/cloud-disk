package test

import (
	"cloud-disk/core/define"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func TestUpload(t *testing.T) {
	client, err := oss.New(define.BucketName, define.AccessId, define.Secret)
	if err != nil {
		t.Fatal(err)
	}

	// Create bucket
	// err = client.CreateBucket(bucketName)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// Get bucket
	bucket, err := client.Bucket("harukaze-cloud-disk")
	if err != nil {
		t.Fatal(err)
	}

	err = bucket.PutObjectFromFile("objectKey.png", "./img/20211217162025.png")
	if err != nil {
		t.Fatal(err)
	}
}

func TestPartUpload(t *testing.T) {
	// 创建OSSClient实例。
	// yourEndpoint填写Bucket所在地域对应的Endpoint。以华东1（杭州）为例，Endpoint填写为https://oss-cn-hangzhou.aliyuncs.com。
	// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
	client, err := oss.New(define.BucketName, define.AccessId, define.Secret)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	// 填写存储空间名称。
	bucketName := "harukaze-cloud-disk"
	// 填写Object完整路径。Object完整路径中不能包含Bucket名称。
	objectName := "exampleobject.png"
	// 填写本地文件的完整路径。如果未指定本地路径，则默认从示例程序所属项目对应本地路径中上传文件。
	locaFilename := "./img/20211217162025.png"

	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	// 将本地文件分片，且分片数量指定为3。
	chunks, err := oss.SplitFileByPartNum(locaFilename, 3)
	fd, err := os.Open(locaFilename)
	defer fd.Close()

	// 指定过期时间。
	expires := time.Date(2049, time.January, 10, 23, 0, 0, 0, time.UTC)
	// 如果需要在初始化分片时设置请求头，请参考以下示例代码。
	options := []oss.Option{
		oss.MetadataDirective(oss.MetaReplace),
		oss.Expires(expires),
		// 指定该Object被下载时的网页缓存行为。
		// oss.CacheControl("no-cache"),
		// 指定该Object被下载时的名称。
		// oss.ContentDisposition("attachment;filename=FileName.txt"),
		// 指定该Object的内容编码格式。
		// oss.ContentEncoding("gzip"),
		// 指定对返回的Key进行编码，目前支持URL编码。
		// oss.EncodingType("url"),
		// 指定Object的存储类型。
		// oss.ObjectStorageClass(oss.StorageStandard),
	}

	// 步骤1：初始化一个分片上传事件，并指定存储类型为标准存储。
	imur, err := bucket.InitiateMultipartUpload(objectName, options...)
	// 步骤2：上传分片。
	var parts []oss.UploadPart
	for _, chunk := range chunks {
		fd.Seek(chunk.Offset, os.SEEK_SET)
		// 调用UploadPart方法上传每个分片。
		part, err := bucket.UploadPart(imur, fd, chunk.Size, chunk.Number)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(-1)
		}
		parts = append(parts, part)
	}

	// 指定Object的读写权限为公共读，默认为继承Bucket的读写权限。
	objectAcl := oss.ObjectACL(oss.ACLPublicRead)

	// 步骤3：完成分片上传，指定文件读写权限为公共读。
	cmur, err := bucket.CompleteMultipartUpload(imur, parts, objectAcl)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	fmt.Println("cmur:", cmur)
}
