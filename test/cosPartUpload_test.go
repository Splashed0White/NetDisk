package test

import (
	"NetDisk/core/define"
	"bytes"
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
	"os"
	"testing"
)

func TestInitPartUpload(t *testing.T) {
	u, _ := url.Parse("https://szluyu99-1259132563.cos.ap-nanjing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.CosID,
			SecretKey: define.CosKey,
		},
	})

	key := "cloud-disk/exampleobject.jpeg"
	v, _, err := client.Object.InitiateMultipartUpload(context.Background(), key, nil)
	if err != nil {
		t.Fatal(err)
	}
	UploadID := v.UploadID // 165320975357e514f921b93de385fc75fb5c1f9702a5da1bacd086900c6c1a613805b15df1
	fmt.Println(UploadID)
}

// 分片上传
func TestPartUpload(t *testing.T) {
	u, _ := url.Parse("https://szluyu99-1259132563.cos.ap-nanjing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.CosID,
			SecretKey: define.CosKey,
		},
	})

	key := "cloud-disk/exampleobject.jpeg"
	UploadID := "165320975357e514f921b93de385fc75fb5c1f9702a5da1bacd086900c6c1a613805b15df1"
	f, err := os.ReadFile("0.chunk") // md5 : 8f10a58845f83846ca5aa422edc0087d
	if err != nil {
		t.Fatal(err)
	}
	// opt可选
	resp, err := client.Object.UploadPart(
		context.Background(), key, UploadID, 1, bytes.NewReader(f), nil,
	)
	if err != nil {
		t.Fatal(err)
	}
	PartETag := resp.Header.Get("ETag")
	fmt.Println(PartETag)
}

// 分片上传完成
func TestPartUploadComplete(t *testing.T) {
	u, _ := url.Parse("https://szluyu99-1259132563.cos.ap-nanjing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.CosID,
			SecretKey: define.CosKey,
		},
	})

	key := "cloud-disk/exampleobject.jpeg"
	UploadID := "165320975357e514f921b93de385fc75fb5c1f9702a5da1bacd086900c6c1a613805b15df1"

	opt := &cos.CompleteMultipartUploadOptions{}
	opt.Parts = append(opt.Parts, cos.Object{
		PartNumber: 1, ETag: "8f10a58845f83846ca5aa422edc0087d"},
	)
	_, _, err := client.Object.CompleteMultipartUpload(
		context.Background(), key, UploadID, opt,
	)
	if err != nil {
		t.Fatal(err)
	}
}
