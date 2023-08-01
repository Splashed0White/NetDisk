package test

//测试Gorm连接情况
import (
	"NetDisk/core/models" //导入模型包
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"                  //导入gorm包
	_ "github.com/jinzhu/gorm/dialects/mysql" //导入mysql驱动包
	"testing"
)

func TestGorm(t *testing.T) {
	//创建Gorm连接
	db, err := gorm.Open("mysql", "root:123456@/netdisk?charset=utf8mb4&parseTime=True&loc=Local") //打开mysql数据库连接
	if err != nil {
		t.Fatal(err) //如果连接失败，输出错误信息并退出程序
	}
	defer db.Close() //函数结束时关闭数据库连接

	data := make([]*models.UserBasic, 0) //定义一个UserBasic类型的切片data，长度为0
	db.Find(&data)                       //查询数据并将结果存储到data中

	b, err := json.Marshal(data) //将data转换为json格式的字节数组b
	if err != nil {
		t.Fatal(err) //如果转换失败，输出错误信息并退出程序
	}

	dst := new(bytes.Buffer)          //创建一个bytes.Buffer类型的变量dst
	err = json.Indent(dst, b, "", "") //将b格式化后存储到dst中
	if err != nil {
		t.Fatal(err) //如果格式化失败，输出错误信息并退出程序
	}

	fmt.Println(dst.String()) //输出dst中的内容
}
