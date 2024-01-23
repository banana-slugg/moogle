package podbean

import (
	"html/template"
	"time"
)

type PodBeanInfo struct {
	Title       string
	Description string
	Link        string
	Copyright   string
	Logo        string
	Episodes    []Episode
}

type Episode struct {
	Title       string
	Description template.HTML
	Published   *time.Time
	Link        string
	DirectLink  string
}
