package service

import (
	"context"
	"github.com/leopoldxx/go-utils/trace"
)

func (s *service) GetTask(ctx context.Context) error{
	tracer := trace.GetTraceFromContext(ctx)
	tracer.Info("s.GetTask")
	return nil
}

func (s *service) ListTask(ctx context.Context) error{
	tracer := trace.GetTraceFromContext(ctx)
	tracer.Info("s.ListTask")
	return nil
}

func (s *service) CreateTask(ctx context.Context) error{
	tracer := trace.GetTraceFromContext(ctx)
	tracer.Info("s.CreateTask")
	return nil
}

func (s *service) UpdateTask(ctx context.Context) error{
	tracer := trace.GetTraceFromContext(ctx)
	tracer.Info("s.UpdateTask")
	return nil
}

func (s *service) DeleteTask(ctx context.Context) error{
	tracer := trace.GetTraceFromContext(ctx)
	tracer.Info("s.UpdateTask")
	return nil
}