package domain

import (
	"time"
)

type ReplTaskList struct {
	Emoji    string            `json:"emoji"`
	Title    string            `json:"title"`
	Order    uint8             `json:"order"`
	Sections []ReplTaskSection `json:"sections"`
}

type ReplTaskSection struct {
	Title string     `json:"title"`
	Order uint8      `json:"order"`
	Tasks []ReplTask `json:"tasks"`
}

type ReplTask struct {
	Title         string        `json:"title"`
	IsCompleted   bool          `json:"isCompleted"`
	CompletedDays string        `json:"completedDays"`
	Note          string        `json:"note"`
	Order         uint8         `json:"order"`
	RepeatType    string        `json:"repeatType"`
	DaysOfWeek    string        `json:"daysOfWeek"`
	DaysOfMonth   string        `json:"daysOfMonth"`
	ConcreteDate  time.Time     `json:"concreteDate"`
	DateStart     time.Time     `json:"dateStart"`
	DateEnd       time.Time     `json:"dateEnd"`
	DateReminder  time.Time     `json:"dateReminder"`
	Subtasks      []ReplSubtask `json:"subtasks"`
}

type ReplSubtask struct {
	Title       string `json:"title"`
	IsCompleted bool   `json:"isCompleted"`
	Order       uint8  `json:"order"`
}

type UsernamePasswd struct {
	Name   string `json:"name"`
	Passwd string `json:"password"`
}

type Response struct {
	Success  bool              `json:"success"`
	ErrCode  int               `json:"errCode,omitempty"`
	ErrMsg   string            `json:"errMsg,omitempty"`
	Token    string            `json:"token,omitempty"`
	Lists    []ReplTaskList    `json:"lists,omitempty"`
	Sections []ReplTaskSection `json:"sections,omitempty"`
	Tasks    []ReplTask        `json:"tasks,omitempty"`
	Subtasks []ReplSubtask     `json:"subtasks,omitempty"`
}
