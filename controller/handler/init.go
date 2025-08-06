package handler

import (
	"omniserve/service/handler"
)

func InitHandler() {

	//启动http
	handler.RunHttp()

	//启动websocket
	handler.RunWebsocket()

	//启动grpc
	handler.RunGrpc()
}
