package main

import (
	"myweb/app"
	config "myweb/config"
)

type Response struct {
	Code    int         `json:"code" form:"code"`
	Message string      `json:"message" form:"message"`
	Data    interface{} `json:"data"`
}

func main() {

	//Connection DB
	config.ConnectDB()

	//Handlers
	app.StartApp()

}
