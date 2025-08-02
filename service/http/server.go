package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var WEB *gin.Engine

func InitServer() {

	WEB = gin.New()

	// 中间件插入
	WEB.Use(
		gin.Logger(),
		gin.Recovery(),
	)

	//路由
	web()

	//启动http server
	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      WEB,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			log.Fatalf("start error: %+v", err)
		}
	}()

	log.Printf("HTTP Server started on %s", httpServer.Addr)

	// 等待中断信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// 设置关闭的超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")

}
