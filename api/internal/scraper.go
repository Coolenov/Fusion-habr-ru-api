package internal

import (
	"fmt"
	"github.com/Coolenov/Fusion-library/lib"
	"github.com/microcosm-cc/bluemonday"
	"github.com/mmcdole/gofeed"
	"strings"
	"time"
)

func HabrScraper() []lib.Post {
	var fp = gofeed.NewParser()
	var feed, _ = fp.ParseURL("https://habr.com/ru/rss/all/all/?limit=100")

	var posts []lib.Post
	for i := 0; i < 100; i++ {
		item := lib.Post{
			Title:          feed.Items[i].Title,
			Link:           feed.Items[i].Link,
			Description:    sanitizeText(feed.Items[i].Description),
			Tags:           feed.Items[i].Categories,
			Source:         "HabrRu",
			PublishingTime: formPubTime(feed.Items[i].Published),
			ImageUrl:       "",
		}
		posts = append(posts, item)
	}
	return posts
}

func sanitizeText(str string) string {
	p := bluemonday.StripTagsPolicy()
	result := p.Sanitize(str)
	result = strings.ReplaceAll(result, "Читать далее", "")
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
