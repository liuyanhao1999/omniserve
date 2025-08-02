package handler

import (
	"elegance-gateway/service/handler"
)

func InitHandler() {

	//启动http
	handler.RunHttp()

	//启动websocket
	handler.RunWebsocket()

	//启动grpc
	handler.RunGrpc()
}
