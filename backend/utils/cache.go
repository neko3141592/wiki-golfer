package utils

import (
	"encoding/json"
	"gin-quickstart/models"

	"gorm.io/gorm"
)

func GetCache(startID, endID int, db *gorm.DB) (data []int) {
	var cache models.Cache

	result := db.Where("start_id = ? AND end_id = ?", startID, endID).First(&cache)
	if result.Error != nil {
		return nil
	}

	var arr []int
	json.Unmarshal([]byte(cache.Data), &arr)
	return arr
}

func SaveCache(startID, endID int, arr []int, db *gorm.DB) {

	jsonStr, _ := json.Marshal(arr)
	cache := models.Cache{
		StartID: startID,
		EndID:   endID,
		Data:    string(jsonStr),
	}
	db.FirstOrCreate(&cache, models.Cache{StartID: startID, EndID: endID})
}
