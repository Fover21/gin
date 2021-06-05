package appall

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"unsafe"
)

type dataInfo struct {
	Txt string `form:"txt" json:"txt"`
	Sim int    `form:"sim" json:"sim"`
}

func getTextHandler(c *gin.Context) {
	// var data dataInfo

	//song := make(map[string]interface{})
	//song["txt"] = data.Txt
	//song["sim"] = data.Sim
	txt := c.PostForm("txt")
	sim, _ := strconv.Atoi(c.PostForm("sim"))
	song := dataInfo{
		Txt: txt,
		Sim: sim,
	}
	fmt.Println("song", song)
	// struct -> json string
	bytesData, err := json.Marshal(&song)
	//fmt.Println("bytesData", bytesData)
	if err != nil {
		fmt.Println("struct -> json string", err.Error())
		return
	}
	reader := bytes.NewReader(bytesData)
	fmt.Printf("reader:- %T\n", reader)
	url_item := "http://apis.5118.com/wyc/rewrite?txt=发电方式大&sim=0"
	var data = make(url.Values)
	data.Set("txt", "发电方式大-v") // 组装数据
	data.Set("sim", "0")       // 组装数据
	//request, err := http.NewRequest("POST", url, reader)
	request, err := http.NewRequest("POST", url_item, strings.NewReader(data.Encode()))
	if err != nil {
		fmt.Println("http.NewRequest", err.Error())
		return
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	request.Header.Set("Authorization", "1CE3007628754D0E9DCAB85F33766D10")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("client.Do", err.Error())
		return
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ioutil.ReadAll", err.Error())
		return
	}
	//byte数组直接转成string，优化内存
	str := (*string)(unsafe.Pointer(&respBytes))
	fmt.Println(*str)

	c.JSON(http.StatusOK, gin.H{
		"message": *str,
	})
}
