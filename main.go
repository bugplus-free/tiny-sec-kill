package main

import "tiny-sec-kill/models"

func main() {
	models.InitDB()
	models.InitCache()
}
