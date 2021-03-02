package main

import (
	"fmt"
	"gin_one/app/blog"
	"gin_one/app/shop"
	"gin_one/app/login"
	"gin_one/app/appall"
	"gin_one/routers"
)

func main() {
	// 加载多个APP的路由配置
	routers.Include(shop.Routers, blog.Routers, login.Routers, appall.Routers)
	// 初始化路由
	r := routers.Init()
	if err := r.Run(":8090"); err != nil {
	fmt.Println("startup service failed, err:%v\n", err)
	}
}