package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
	"time"
)

type New struct {
	Title string `json:"title,omitempty"`
	Link  string `json:"link,omitempty"`
	Cover string `json:"cover,omitempty"`
}

var (
	News [4]New
	Url  = "https://hamsternet.medium.com"
)

func main() {
	r := gin.Default()

	go func() {
		for {
			// 如果爬新闻出错，则使用默认的新闻
			err := getNewsNet()
			if err != nil {
				log.Println(err)
				getNewsDefault()
			}
			time.Sleep(time.Hour * 24)
		}
	}()

	r.GET("/articles", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
			"data":    News,
		})
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8888")
}

func getNewsNet() error {
	res, err := http.Get(Url)
	if err != nil {
		log.Println(err)
		return err
	}
	defer res.Body.Close()

	var (
		titles []string
		links  []string
		covers []string
	)

	if res.StatusCode != 200 {
		log.Printf("status code error: %d %s\n", res.StatusCode, res.Status)
		return err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Println(err)
		return err
	}

	// 获取文章标题列表
	doc.Find("main article a[aria-label='Post Preview Title'] h2").Each(func(idx int, selection *goquery.Selection) {
		titles = append(titles, selection.Text())
	})

	// 获取文章链接列表
	doc.Find("main article a[aria-label='Post Preview Title']").Each(func(idx int, selection *goquery.Selection) {
		val, exits := selection.Attr("href")
		if exits {
			links = append(links, val)
		}
	})

	// 获取文章图片链接列表
	doc.Find("main article a[aria-label='Post Preview Image']").Each(func(idx int, selection *goquery.Selection) {
		val, exits := selection.Find("img").Attr("src")
		if exits {
			tmp := strings.Split(val, "/")
			var newCover string
			for i := range tmp {
				if strings.HasPrefix(tmp[i], "resize:") {
					continue
				}
				newCover = newCover + tmp[i] + "/"
			}

			covers = append(covers, newCover[:len(newCover)-1])
		}
	})

	for i := range News {
		News[i] = New{
			Title: titles[i],
			Link:  Url + links[i],
			Cover: covers[i],
		}
	}

	return nil
}

func getNewsDefault() {
	New0 := New{
		Title: "KNN3 x Hamster — to liberate developers with multi-chain aggregated data graph services",
		Link:  "https://hamsternet.medium.com/knn3-x-hamster-to-liberate-developers-with-multi-chain-aggregated-data-graph-services-e190dfc21bfc?source=user_profile---------0----------------------------",
		Cover: "https://miro.medium.com/v2/0*NswMnN70uVLh4x9v",
	}

	New1 := New{
		Title: "CESS x Hamster —Empowering Web3 Developers with Revolutionary Decentralized Storage",
		Link:  "https://hamsternet.medium.com/cess-x-hamster-empowering-web3-developers-with-revolutionary-decentralized-storage-aa84134eeac6?source=user_profile---------1----------------------------",
		Cover: "https://miro.medium.com/v2/0*-WOsONjpryfTXLNp",
	}

	New2 := New{
		Title: "Hamster Network Monthly Report",
		Link:  "https://hamsternet.medium.com/hamster-network-monthly-report-3495ecd4d047?source=user_profile---------2----------------------------",
		Cover: "https://miro.medium.com/v2/1*hbHMuP_iuAHhvzCTcWSIUQ.png",
	}

	New3 := New{
		Title: "MetaTrust Labs x Hamster — providing developers with a more secured and automated Web3.0 development environment",
		Link:  "https://hamsternet.medium.com/metatrust-labs-x-hamster-providing-developers-with-a-more-secured-and-automated-web3-0-40688e815ecf?source=user_profile---------3----------------------------",
		Cover: "https://miro.medium.com/v2/0*wyxb-RFM8Jy7wPXP",
	}

	News[0] = New0
	News[1] = New1
	News[2] = New2
	News[3] = New3
}
