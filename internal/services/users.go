package services

import (
	"context"
	"fmt"
	d "my_life/internal/domain"
	"my_life/internal/repository"
)

type UserService struct {
	repo repository.Repository
}

func NewUserService(repo repository.Repository) *UserService {
	return &UserService{repo: repo}
}

func (s UserService) GetUserById(ctx context.Context, UId int32) (*d.User, error) {
	return s.repo.GetUserById(ctx, UId)
}

func (s UserService) GetFullUserInfo(ctx context.Context, UId int32) (*d.Reply, error) {
	var reply d.Reply

	subtasks, err := s.repo.GetSubTasksByUId(ctx, UId)
	if err != nil {
		return nil, err
	}
	tasks, err := s.repo.GetTasksByUId(ctx, UId)
	if err != nil {
		return nil, err
	}
	sections, err := s.repo.GetSectionsByUId(ctx, UId)
	if err != nil {
		return nil, err
	}
	lists, err := s.repo.GetListsByUId(ctx, UId)
	if err != nil {
		return nil, err
	}
	replLists := bindAndTrimIds(lists, sections, tasks, subtasks)

	reply.Title = fmt.Sprintf("User %d", UId)
	reply.Body = replLists

	return &reply, nil
}

func bindAndTrimIds(lists []d.TaskList, sections []d.TaskSection, tasks []d.Task, subtasks []d.Subtask) []d.ReplTaskList {
	var replLists []d.ReplTaskList

	for _, l := range lists {
		var replList d.ReplTaskList
		replList.Emoji = l.Emoji
		replList.Title = l.Title
		replList.Order = l.Order
		replList.Sections = makeReplSections(l.Id, sections, tasks, subtasks)
		replLists = append(replLists, replList)
	}
	return replLists
}

func makeReplSections(listId int32, sections []d.TaskSection, tasks []d.Task, subtasks []d.Subtask) []d.ReplTaskSection {
	var replSections []d.ReplTaskSection

	for _, s := range sections {
		if s.ListId != listId {
			continue
		}
		var replSection d.ReplTaskSection
		replSection.Title = s.Title
		replSection.Order = s.Order
		replSection.Tasks = makeReplTasks(s.Id, tasks, subtasks)

		replSections = append(replSections, replSection)
	}
	return replSections
}

func makeReplTasks(sectionId int32, tasks []d.Task, subtasks []d.Subtask) []d.ReplTask {
	var replTasks []d.ReplTask

	for _, t := range tasks {
		if t.SectionId != sectionId {
			continue
		}
		var replTask d.ReplTask
		replTask.Title = t.Title
		replTask.IsCompleted = t.IsCompleted
		replTask.CompletedDays = t.CompletedDays
		replTask.Note = t.Note
		replTask.Order = t.Order
		replTask.RepeatType = t.RepeatType
		replTask.DaysOfWeek = t.DaysOfWeek
		replTask.DaysOfMonth = t.DaysOfMonth
		replTask.ConcreteDate = t.ConcreteDate
		replTask.DateStart = t.DateStart
		replTask.DateEnd = t.DateEnd
		replTask.DateReminder = t.DateReminder
		replTask.Subtasks = makeReplSubtasks(t.Id, subtasks)

		replTasks = append(replTasks, replTask)
	}
	return replTasks
}

func makeReplSubtasks(taskId int32, subtasks []d.Subtask) []d.ReplSubtask {
	var replSubtasks []d.ReplSubtask

	for _, subtask := range subtasks {
		if subtask.TaskId == taskId {
			replSubtasks = append(replSubtasks, trimIdsFromSubtask(subtask))
		}
	}
	return replSubtasks
}

func trimIdsFromSubtask(subtask d.Subtask) d.ReplSubtask {
	var replSubtask d.ReplSubtask

	replSubtask.Title = subtask.Title
	replSubtask.Order = subtask.Order
	replSubtask.IsCompleted = subtask.IsCompleted

	return replSubtask
}
