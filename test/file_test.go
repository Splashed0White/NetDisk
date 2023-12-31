package test

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"testing"
)

var chunkSize = 1024 * 1024

// 文件分片：
func TestGenerateChunkFile(t *testing.T) {
	// 读取文件
	fileInfo, err := os.Stat("test.mp4")
	if err != nil {
		t.Fatal(err)
	}

	// 分片个数 = 文件大小 / 分片大小
	// 390 / 100 ==> 4, 向上取整
	chunkNum := math.Ceil(float64(fileInfo.Size()) / float64(chunkSize))
	// 只读方式打开文件
	myFile, err := os.OpenFile("test.mp4", os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}
	// 存放每一次的分片数据
	b := make([]byte, chunkSize)
	// 遍历所有分片
	for i := 0; i < int(chunkNum); i++ {
		// 指定读取文件的起始位置
		myFile.Seek(int64(i*chunkSize), 0)
		// 最后一次的分片数据不一定是整除下来的数据
		// 例如: 文件 120M, 第一次读了 100M, 剩下只有 20M
		if int64(chunkSize) > fileInfo.Size()-int64(i*chunkSize) {
			b = make([]byte, fileInfo.Size()-int64(i*chunkSize))
		}
		myFile.Read(b)

		f, err := os.OpenFile("./"+strconv.Itoa(i)+".chunk", os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}
		f.Write(b)
		f.Close()
	}
	myFile.Close()
}

// 分片文件的合并：
func TestMergeChunkFile(t *testing.T) {
	myFile, err := os.OpenFile("test2.mp4", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}

	// 计算分片个数, 正常应该由前端传来, 这里测试时自行计算
	fileInfo, err := os.Stat("test.mp4")
	if err != nil {
		t.Fatal(err)
	}
	// 分片个数 = 文件大小 / 分片大小
	chunkNum := math.Ceil(float64(fileInfo.Size()) / float64(chunkSize))

	// 合并分片
	for i := 0; i < int(chunkNum); i++ {
		f, err := os.OpenFile("./"+strconv.Itoa(i)+".chunk", os.O_RDONLY, os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}
		b, err := ioutil.ReadAll(f)
		if err != nil {
			t.Fatal(err)
		}
		myFile.Write(b)
		f.Close()
	}
	myFile.Close()
}

// 文件一致性校验：
func TestCheck(t *testing.T) {
	// 获取第一个文件的信息
	f1, err := os.OpenFile("test.mp4", os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}
	b1, err := ioutil.ReadAll(f1)
	if err != nil {
		t.Fatal(err)
	}

	// 获取第二个文件的信息
	f2, err := os.OpenFile("test2.mp4", os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}
	b2, err := ioutil.ReadAll(f2)
	if err != nil {
		t.Fatal(err)
	}

	s1 := fmt.Sprintf("%x", md5.Sum(b1))
	s2 := fmt.Sprintf("%x", md5.Sum(b2))

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s1 == s2)
}
