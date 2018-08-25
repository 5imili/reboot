package service

import (
	"context"
	"github.com/5imili/reboot/pkg/dao"
)

type Options struct{
	DB dao.Storage
}

type Operation interface {
	GetTask(ctx context.Context) error
	ListTask(ctx context.Context)error
	DeleteTask(ctx context.Context) error
	CreateTask(ctx context.Context,namespace string, resource string) error
	UpdateTask(ctx context.Context) error
}

type service struct{
	opt *Options
}

func New(opt *Options) Operation{
	return &service{
		opt:opt,
	}
}

type Task struct{
	Resource string
}