package pkgs

import (
	"github.com/5imili/reboot/pkg/task/scheduler"
	"sync"
	"github.com/5imili/reboot/pkg/dao"
	"github.com/leopoldxx/go-utils/trace"
	"context"
	"github.com/5imili/reboot/pkg/task/taskscheduler"
)

var (
	manager *scheduler.Manager
	managerOnce sync.Once
)

func GetScheduler(sto dao.Storage) *scheduler.Manager{
	managerOnce.Do(func() {
		var err error
		ctx := trace.WithTraceForContext(context.TODO(), "task-scheduler")
		manager , err = scheduler.NewManager(ctx, sto)
		if err !=nil{
			panic(err)
		}
		if err = manager.InitSchedulers(taskscheduler.Scheduler());err != nil{
			panic(err)
		}
	})
	return manager
}
