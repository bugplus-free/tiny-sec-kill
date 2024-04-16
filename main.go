package main

import (
	"tiny-sec-kill/models"
	"tiny-sec-kill/routers"
)

// @title 微商品抢购
// @version 1.0
// @description 一个简洁的商品抢购网站

// @contact.name dhw
// @contact.email 

func main() {
	models.InitDB()
	models.InitCache()
	routers.InitRouter()
}
