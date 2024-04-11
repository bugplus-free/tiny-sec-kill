package main

import (
	"tiny-sec-kill/models"
	"tiny-sec-kill/routers"
)
func main() {
	models.InitDB()
	models.InitCache()
	routers.InitRouter()
}
