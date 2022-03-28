package domain

type TaskList struct {
	Id            uint8  `json:"id"`
	Emoji         string `json:"emoji"`
	Title         string `json:"title"`
	Order         uint8  `json:"order"`
	RelevanceTime string `json:"relevanceTime"`
}
