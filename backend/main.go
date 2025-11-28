package main

import (
	"gin-quickstart/db"
	"gin-quickstart/handlers"
	"gin-quickstart/utils"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	ARTICLE_FILE = "jawiki-latest-pages-articles.xml"
	MAX_ARTICLES = 4e6
)

var (
	Graph     map[int][]int
	sexualIDs mapset.Set[int]
	idToTitle map[int]string
	titleToID map[string]int
)

func main() {
	err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/api/articles", handlers.SearchArticles)


	go func() {
		idToTitle, titleToID, sexualIDs, err := utils.BuildMap(ARTICLE_FILE, MAX_ARTICLES, db.DB, false)
		if err != nil {
			panic(err)
		}
		Graph, err := utils.BuildGraph(ARTICLE_FILE, MAX_ARTICLES, titleToID, db.DB, false)
		if err != nil {
			panic(err)
		}
		handlers.Graph = Graph
		handlers.SexualIDs = sexualIDs
		handlers.IdToTitle = idToTitle
		handlers.TitleToID = titleToID
	}()

	r.GET("/api/path", func(c *gin.Context) {
		if handlers.Graph == nil {
			c.JSON(503, gin.H{"error": "グラフ構築中です。しばらくお待ちください。"})
			return
		}
		handlers.FindShortestPath(c)
	})

	r.Run(":8080")
}
