package dao

import (
	"context"
	"github.com/5imili/reboot/pkg/dao/mysql/types"
)

//Storage xxx
type Storage interface {
	ListTask(ctx context.Context)
	CreateTask(ctx context.Context, task *types.Task) (int64, error)
	GetTask(ctx context.Context)
	DeleteTask(ctx context.Context)
	UpdateTask(ctx context.Context, task *types.Task) error

	ListOpenTasks(ctx context.Context) ([]types.Task, error)
	GetOpenTaskByTaskID(ctx context.Context, id int64) (*types.Task, error)
}