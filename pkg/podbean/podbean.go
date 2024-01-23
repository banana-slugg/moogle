package podbean

import (
	"fmt"
	"html/template"
	"strings"

	"github.com/mmcdole/gofeed"
)

func GetPodbean() (*PodBeanInfo, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL("https://feed.podbean.com/themusicalmoogle/feed.xml")
	if err != nil {
		return nil, err
	}

	info := &PodBeanInfo{
		Title:       feed.Title,
		Description: feed.Description,
		Link:        feed.Link,
		Copyright:   feed.Copyright,
		Logo:        feed.Image.URL,
	}

	for _, item := range feed.Items {
		description := item.Description
		description = strings.ReplaceAll(description, "<p>\u00a0</p>", "")
		description = strings.TrimSpace(description)
		fmt.Println(description)
		info.Episodes = append(info.Episodes, Episode{
			Title:       item.Title,
			Description: template.HTML(description),
			Published:   item.PublishedParsed,
			Link:        item.Link,
			DirectLink:  item.Enclosures[0].URL,
		})

	}
	return info, nil
}
