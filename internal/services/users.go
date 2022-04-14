package services

import (
	"context"
	"my_life/internal/domain"
	"my_life/internal/repository"
)

type UserService struct {
	repo repository.Repository
}

func NewUserService(repo repository.Repository) *UserService {
	return &UserService{repo: repo}
}

func (s UserService) GetUserById(ctx context.Context, UId int32) (*domain.User, error) {
	return s.repo.GetUserById(ctx, UId)
}

func (s UserService) GetFullUserInfo(ctx context.Context, UId int32) ([]domain.ReplTask, error) {
	subtasks, err := s.repo.GetSubTasksByUId(ctx, UId)
	if err != nil {
		return nil, err
	}
	tasks, err := s.repo.GetTasksByUId(ctx, UId)
	if err != nil {
		return nil, err
	}
	replTasks := trimIdsFromTasks(tasks, subtasks)
	return replTasks, nil
}

func trimIdsFromTasks(tasks []domain.Task, subtasks []domain.Subtask) []domain.ReplTask {
	var replTasks []domain.ReplTask

	for _, task := range tasks {
		var replTask domain.ReplTask
		replTask.Title = task.Title
		replTask.IsCompleted = task.IsCompleted
		replTask.CompletedDays = task.CompletedDays
		replTask.Note = task.Note
		replTask.Order = task.Order
		replTask.RepeatType = task.RepeatType
		replTask.DaysOfWeek = task.DaysOfWeek
		replTask.DaysOfMonth = task.DaysOfMonth
		replTask.ConcreteDate = task.ConcreteDate
		replTask.DateStart = task.DateStart
		replTask.DateEnd = task.DateEnd
		replTask.DateReminder = task.DateReminder
		replTask.Subtasks = getSubtasksByTaskId(subtasks)
		replTasks = append(replTasks, replTask)
	}
	return replTasks
}

func getSubtasksByTaskId(subtasks []domain.Subtask) []domain.ReplSubtask {
	var replSubtasks []domain.ReplSubtask

	for _, subtask := range subtasks {
		replSubtasks = append(replSubtasks, trimIdsFromSubtask(subtask))
	}

	return replSubtasks
}

func trimIdsFromSubtask(subtask domain.Subtask) domain.ReplSubtask {
	var replSubtask domain.ReplSubtask

	replSubtask.Title = subtask.Title
	replSubtask.Order = subtask.Order
	replSubtask.IsCompleted = subtask.IsCompleted

	return replSubtask
}