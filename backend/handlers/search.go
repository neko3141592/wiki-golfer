package handlers

import (
	"gin-quickstart/db"
	"gin-quickstart/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SearchArticles(c *gin.Context) {
	title := c.Query("title")
	limitString := c.DefaultQuery("limit", "10") 
	limit, _ := strconv.Atoi(limitString)
	database := db.DB

	var articles []models.Article

	if len(title) < 3 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "タイトルが短すぎます",
		})
		return
	}

	if err := database.
		Where("title LIKE ?", "%"+title+"%").
		Order("LENGTH(title) ASC").
		Limit(limit). 
		Find(&articles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "検索に失敗しました",
		})
		return
	}

	c.JSON(http.StatusOK, articles)
}
