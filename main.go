package main

import (
	"getNews/handler"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()
	// init default news
	handler.GetNewsDefault()

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

	// 定时清除访问记录
	go func() {
		for {
			handler.VisitedIP = make(map[string]struct{})
			time.Sleep(time.Hour * 24)
		}
	}()

	// api
	r.GET("/articles", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
			"data":    handler.News,
		})
	})
	r.POST("/JoinMiddleware", handler.JoinMiddleware)

	// 监听并在 0.0.0.0:8888 上启动服务
	r.Run(":8080")
}
