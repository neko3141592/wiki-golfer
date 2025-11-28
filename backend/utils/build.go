package utils

import (
	"fmt"
	"gin-quickstart/models"
	"regexp"
	"strconv"
	"sync"

	mapset "github.com/deckarep/golang-set/v2"
	"gorm.io/gorm"
)

func notUnion(s string) bool {
	reDate := regexp.MustCompile(`^\d{1,2}月\d{1,2}日$`)
	reYear := regexp.MustCompile(`^\d{4}年$`)
	return reDate.MatchString(s) || reYear.MatchString(s)
	// return false
}

func BuildMap(filename string, maxCount int, db *gorm.DB, saveDB bool) (map[int]string, map[string]int, mapset.Set[int], error) {
	idToTitle := make(map[int]string)
	titleToID := make(map[string]int)
	sexualIDs := mapset.NewSet[int]()
	var mu sync.Mutex
	var wg sync.WaitGroup

	fmt.Println("Building title to ID map...")
	err := ProcessArticles(
		filename,
		maxCount,
		func(_ []string, title, id string, _ []string, judge bool) error {
			if title != "" && id != "" {
				wg.Add(1)
				go func(title, id string, judge bool) {
					defer wg.Done()
					mu.Lock()
					titleToID[title] = atoi(id)
					idToTitle[atoi(id)] = title
					if judge {
						sexualIDs.Add(atoi(id))
					}
					mu.Unlock()
					if saveDB {
						article := models.Article{
							Title:    title,
							WikiID:   atoi(id),
							IsSexual: judge,
						}
						db.FirstOrCreate(&article, models.Article{WikiID: article.WikiID})
					}
				}(title, id, judge)
			}
			return nil
		},
	)
	wg.Wait()
	if err != nil {
		return nil, nil, nil, err
	}
	fmt.Println("Title to ID map built.")
	return idToTitle, titleToID, sexualIDs, nil
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func BuildGraph(filename string, maxCount int, titleToID map[string]int, db *gorm.DB, saveDB bool) (map[int][]int, error) {
	fmt.Println("Building graph...")
	graph := make(map[int][]int)
	var mu sync.Mutex
	var wg sync.WaitGroup

	err := ProcessArticles(
		filename,
		maxCount,
		func(_ []string, title, id string, links []string, judge bool) error {
			if id == "" {
				return nil
			}
			wg.Add(1)
			go func(id string, links []string) {
				defer wg.Done()
				for _, linkTitle := range links {
					linkedID, ok := titleToID[linkTitle]
					if !ok || notUnion(linkTitle) {
						continue
					}
					mu.Lock()
					graph[atoi(id)] = append(graph[atoi(id)], linkedID)
					mu.Unlock()
					if saveDB {
						link := models.Link{
							FromID: atoi(id),
							ToID:   linkedID,
						}
						db.FirstOrCreate(&link, models.Link{FromID: link.FromID, ToID: link.ToID})
					}
				}
			}(id, links)
			return nil
		},
	)
	wg.Wait()
	if err != nil {
		return nil, err
	}
	fmt.Println("Graph built.")
	return graph, nil
}
