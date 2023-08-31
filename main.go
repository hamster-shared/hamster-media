package main

import (
	"getNews/handler"
	"getNews/initialization"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()
	// init default news
	initialization.Init()
	handler.GetNewsDefault()
	httpHandler := handler.NewHandlerServer()
	// 定时拉取在线的新闻
	go func() {
		for {
			// 如果爬新闻出错，则使用默认的新闻
			NewsTmp, err := handler.GetOnlineMedia()
			if err != nil {
				log.Println(err)
			}
			if NewsTmp != nil {
				handler.News = NewsTmp
			}
			time.Sleep(time.Hour * 3)
		}
	}()

	// api
	r.GET("/articles", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
			"data":    handler.News,
		})
	})

	r.POST("/JoinMiddleware", httpHandler.JoinMiddleware)

	api := r.Group("/api")

	api.GET("/activity/:id/status", httpHandler.GetActivityStatus)
	api.GET("/check/deploy", httpHandler.CheckDeploy)
	api.GET("/deploy/info", httpHandler.GetDeployContractInfo)
	api.POST("/nft/airdrop", httpHandler.SaveNftAirdrop)

	api.GET("/contact/platform", httpHandler.GetContactPlatform)
	api.POST("/contact/ecosystems", httpHandler.SaveEcosystemsContact)
	// 监听并在 0.0.0.0:8888 上启动服务
	r.Run(":8888")
}
