package domain

import "time"

type TaskList struct {
	Id            uint64    `json:"id"`
	UId           uint64    `json:"UId"`
	Emoji         string    `json:"emoji"`
	Title         string    `json:"title"`
	Order         uint8     `json:"order"`
	RelevanceTime time.Time `json:"relevanceTime"`
}
