package blog

/*
连接不同的数据库都需要导入对应数据的驱动程序，GORM已经贴心的为我们包装了一些驱动程序，只需要按如下方式导入需要的数据库驱动即可：
// import _ "github.com/jinzhu/gorm/dialects/postgres"
// import _ "github.com/jinzhu/gorm/dialects/sqlite"
// import _ "github.com/jinzhu/gorm/dialects/mysql"
*/

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

// UserInfo 用户信息
type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

// gorm简单使用
func postHandler(c *gin.Context) {
	db, err := gorm.Open("mysql", "root:root1234@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 自动迁移
	//db.AutoMigrate(&UserInfo{})

	//u1 := UserInfo{1, "wzy", "男", "篮球"}
	//u2 := UserInfo{2, "xw", "女", "足球111"}
	// 创建记录
	//db.Create(&u1)
	//db.Create(&u2)
	// 查询
	var u = new(UserInfo)
	db.First(u)
	fmt.Printf("%#v\n", u)

	var uu UserInfo
	db.Find(&uu, "hobby=?", "足球")
	fmt.Printf("%#v\n", uu)

	// 更新
	db.Model(&u).Update("hobby", "双色球")
	// 删除
	//db.Delete(&u)

	c.JSON(http.StatusOK, gin.H{
		"message": "postHandler!",
	})
}

func commentHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "commentHandler!",
	})
}
