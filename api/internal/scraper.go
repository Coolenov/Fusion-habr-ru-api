package internal

import (
	"fmt"
	"github.com/Coolenov/Fusion-library/types"
	"github.com/microcosm-cc/bluemonday"
	"github.com/mmcdole/gofeed"
	"strings"
	"time"
)

func HabrScraper() []types.Post {
	var fp = gofeed.NewParser()
	var feed, _ = fp.ParseURL("https://habr.com/ru/rss/all/all/")

	var posts []types.Post
	for i := 0; i < 100; i++ {
		item := types.Post{
			Title:          feed.Items[i].Title,
			Link:           feed.Items[i].Link,
			Description:    sanitizeText(feed.Items[i].Description),
			Tags:           feed.Items[i].Categories,
			Source:         "HabrRu",
			PublishingTime: formPubTime(feed.Items[i].Published),
			ImageUrl:       "",
		}
		fmt.Println(item.Title)

		posts = append(posts, item)
	}
	return posts
}

func sanitizeText(str string) string {
	p := bluemonday.StripTagsPolicy()
	result := p.Sanitize(str)
	result = strings.ReplaceAll(result, "Читать далее", "")
	result = strings.ReplaceAll(result, "Читать дальше", "")
	return result
}

func formPubTime(timeStr string) int64 {
	timeString := timeStr
	timeLayout := "Mon, 02 Jan 2006 15:04:05 MST"
	timeObj, err := time.Parse(timeLayout, timeString)
	if err != nil {
		fmt.Println("Ошибка при парсинге времени:", err)
	}
	timestamp := timeObj.Unix()
	return timestamp
}
