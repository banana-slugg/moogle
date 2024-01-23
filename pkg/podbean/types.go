package podbean

import "time"

type PodBeanInfo struct {
	Title       string
	Description string
	Link        string
	Copyright   string
	Episodes    []Episode
}

type Episode struct {
	Title       string
	Description string
	Published   *time.Time
	Link        string
	DirectLink  string
}
