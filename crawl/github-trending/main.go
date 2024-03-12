package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

type trending struct {
	repo  string
	stars int
	date  time.Time
	desc  string
}

func main() {
	c := colly.NewCollector()
	// 对于每个目标元素，使用OnHTML方法设置要执行的操作
	c.OnHTML("article.Box-row", func(e *colly.HTMLElement) {
		// 仓库名
		repoName := e.ChildAttr("a[data-hydro-click]", "href")
		// 仓库简介
		description := e.ChildText("p.col-9")
		// 收藏数
		stars := strings.TrimSpace(e.DOM.Find("a.Link--muted.d-inline-block.mr-3").First().Text())

		if strings.HasPrefix(repoName, "/login?return_to=%2") {
			repoName = strings.Replace(repoName[20:], "%2", "/", 1)
		}
		fmt.Printf("仓库名: %s\n", repoName)
		fmt.Printf("简介: %s\n", description)
		fmt.Printf("收藏数: %s\n", stars)
		fmt.Println("----------------------------")
	})

	// 设置Colly的User Agent
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3"

	// 访问页面
	c.Visit("https://github.com/trending/go?since=daily")
}
