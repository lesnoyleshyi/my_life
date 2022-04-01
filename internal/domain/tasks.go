package domain

import "time"

type User struct {
	UId int
}

type TaskList struct {
	Id            uint64    `json:"id"`
	UId           uint64    `json:"uId"`
	Emoji         string    `json:"emoji"`
	Title         string    `json:"title"`
	Order         uint8     `json:"order"`
	RelevanceTime time.Time `json:"relevanceTime"`
}

type TaskSection struct {
	Id            uint64    `json:"id"`
	UId           uint64    `json:"uId"`
	ListId        uint64    `json:"listId"`
	Title         string    `json:"title"`
	Order         uint8     `json:"order_"`
	RelevanceTime time.Time `json:"relevanceTime"`
}

type Task struct {
	Id            uint64    `json:"id"`
	UId           uint64    `json:"iId"`
	ListId        uint64    `json:"listId"`
	SectionId     uint64    `json:"sectionId"`
	Title         string    `json:"title"`
	IsCompleted   bool      `json:"isCompleted"`
	CompletedDays string    `json:"completedDays"`
	Note          string    `json:"note"`
	Order         uint8     `json:"order_"`
	RepeatType    string    `json:"repeatType"`
	DaysOfWeek    string    `json:"daysOfWeek"`
	DaysOfMonth   string    `json:"daysOfMonth"`
	ConcreteDate  time.Time `json:"concreteDate"`
	DateStart     time.Time `json:"dateStart"`
	DateEnd       time.Time `json:"dateEnd"`
	DateReminder  time.Time `json:"dateReminder"`
	RelevanceTime time.Time `json:"relevanceTime"`
}

type Subtask struct {
	Id            uint64    `json:"id"`
	UId           uint64    `json:"iId"`
	ListId        uint64    `json:"listId"`
	SectionId     uint64    `json:"sectionId"`
	TaskId        uint64    `json:"taskId"`
	Title         string    `json:"title"`
	IsCompleted   bool      `json:"isCompleted"`
	Order         uint8     `json:"order_"`
	RelevanceTime time.Time `json:"relevanceTime"`
}
