package domain

import "time"

type User struct {
	UId           int32     `json:"id"`
	Name          string    `json:"name"`
	Phone         string    `json:"phone"`
	Email         string    `json:"email"`
	Password      string    `json:"password"`
	RelevanceTime time.Time `json:"relevanceTime"`
}

type TaskList struct {
	Id            int32     `json:"id"`
	UId           int32     `json:"uId"`
	Emoji         string    `json:"emoji"`
	Title         string    `json:"title"`
	Order         uint8     `json:"order"`
	RelevanceTime time.Time `json:"relevanceTime"`
}

type TaskSection struct {
	Id            int32     `json:"id"`
	UId           int32     `json:"uId"`
	ListId        int32     `json:"listId"`
	Title         string    `json:"title"`
	Order         uint8     `json:"order_"`
	RelevanceTime time.Time `json:"relevanceTime"`
}

type Task struct {
	Id            int32     `json:"id"`
	UId           int32     `json:"iId"`
	ListId        int32     `json:"listId"`
	SectionId     int32     `json:"sectionId"`
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
	Id            int32     `json:"id"`
	UId           int32     `json:"iId"`
	ListId        int32     `json:"listId"`
	SectionId     int32     `json:"sectionId"`
	TaskId        int32     `json:"taskId"`
	Title         string    `json:"title"`
	IsCompleted   bool      `json:"isCompleted"`
	Order         uint8     `json:"order_"`
	RelevanceTime time.Time `json:"relevanceTime"`
}
