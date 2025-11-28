package handlers

import (
	"gin-quickstart/db"
	"gin-quickstart/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SearchArticles(c *gin.Context) {
	title := c.Query("title")
	database := db.DB


	var articles []models.Article

	if len(title) < 3 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "タイトルが短すぎます",
		})
	} else {
		if err := database.
		Where("title LIKE ?", "%"+title+"%").
		Order("LENGTH(title) ASC").
		Find(&articles).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "検索に失敗しました",
			})
			return
		}
	}

	c.JSON(http.StatusOK, articles)
}
