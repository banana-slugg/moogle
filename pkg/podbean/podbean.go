package podbean

import (
	"html/template"

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
	}

	for _, item := range feed.Items {
		info.Episodes = append(info.Episodes, Episode{
			Title:       item.Title,
			Description: template.HTML(item.Description),
			Published:   item.PublishedParsed,
			Link:        item.Link,
			DirectLink:  item.Enclosures[0].URL,
		})

	}

	return info, nil
}
