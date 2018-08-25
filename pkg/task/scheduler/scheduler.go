package scheduler

import (
	"context"
	"github.com/5imili/reboot/pkg/dao/mysql/types"
	taskTypes "github.com/5imili/reboot/pkg/task/types"
	"github.com/5imili/reboot/pkg/dao"
	"time"
	"sync"
	"github.com/leopoldxx/go-utils/trace"
	"errors"
	"fmt"
	"github.com/5imili/reboot/pkg/task/utils"
	"github.com/zieckey/etcdsync"
)
const (
	defaultPrefix = "/github.com/reboot/pkg/task/scheduler/lock"
	defaultScheduleCycle  = time.Second * 1
)

type Scheduler interface {
	GetName ()string
	Init(taskTypes.InitConfigs)error
	Schedule(ctx context.Context, task *taskTypes.Task) error
}

type Manager struct{
	schedulers map[string] Scheduler
	dao dao.Storage
	lockPrefix string
	ticker *time.Ticker
	ctx context.Context
	cancel context.CancelFunc
	wg sync.WaitGroup
}

type options struct{
	lockPrefix string
	ScheduleCycle time.Duration
}

//NewManager create schedule manager
func NewManager(ctx context.Context,dao dao.Storage)(*Manager , error){
	ops := &options{
		lockPrefix:defaultPrefix,
		ScheduleCycle:defaultScheduleCycle,
	}
	ctx, cancel := context.WithCancel(trace.WithTraceForContext(ctx, "task-scheduler"))
	return &Manager{
		schedulers:map[string] Scheduler{},
		ctx : ctx,
		cancel:cancel,
		dao:dao,
		lockPrefix:ops.lockPrefix,
		ticker:time.NewTicker(ops.ScheduleCycle),
	}, nil
}

func (m *Manager) InitSchedulers(schedules ...Scheduler) error{
	for _, scheduler := range schedules{
		if err := scheduler.Init(taskTypes.InitConfigs{
			Dao: m.dao,
		});err != nil{
			return err
		}
		m.schedulers[scheduler.GetName()] = scheduler
	}
	return nil
}

func (m *Manager) Stop(){
	if m == nil{
		return
	}
	tracer := trace.GetTraceFromContext(m.ctx)
	m.ticker.Stop()
	m.cancel()
	tracer.Info("stopping scheduler manager")
	m.wg.Wait()
	tracer.Info("stopped  scheduler manager")
}

func (m *Manager) Schedule()error{
	if m == nil{
		return errors.New("task scheduler is not created")
	}
	if m.dao == nil{
		return errors.New("task scheduler dao is nil")
	}
	tracer := trace.GetTraceFromContext(m.ctx)
	tracer.Info("start scheduling the pending tasks")
	for range m.ticker.C{
		tasks , err := m.dao.ListOpenTasks(m.ctx)
		if err != nil{
			if err == context.Canceled{
				tracer.Info("scheduler has been stoppd.")
				break
			}
			tracer.Warnf("list open task failed: %s\n", err)
			time.Sleep(time.Second)
			continue
		}
		tracer.Infof("total get %d pending tasks \n", len(tasks))
		for _, task := range tasks{
			scheduler , exits := m.schedulers[task.Resource]
			if !exits{
				tracer.Warn("resource [%s]'s scheduler not exists: %v\n",task.Resource,task)
				continue
			}
			m.wg.Add(1)
			go func(task types.Task){
				defer m.wg.Done()
				newCtx := trace.WithTraceForContext(m.ctx,
					fmt.Sprintf("schedTask:%s:%s:%d", task.Resource, task.Type,task.ID))
				tracer := trace.GetTraceFromContext(newCtx)
				tracer.Info("task start...")
				//Todo etcd lock
				lockkey := fmt.Sprintf("%s/%d", m.lockPrefix, task.ID)
				locker , err := etcdsync.New(lockkey,20, []string{"http://127.0.0.1:2379"})
				if locker == nil || err != nil {
					tracer.Error("etcdsync.New failed")
					return
				}
				err = locker.Lock()
				if err != nil {
					tracer.Errorf("lock error %v\n",err)
					return
				}

				defer func(){
					time.Sleep(time.Second)
					err := locker.Unlock()
					if err != nil {
						tracer.Errorf("unlock error %v\n",err)
						return
					}
				}()

				newTask, err := m.dao.GetOpenTaskByTaskID(newCtx, task.ID)
				if err != nil{
					if err.Error() == "task not found"{
						return
					}
					tracer.Errorf("task can't be scheduled now:%d,err: %v\n",task.ID, err)
					return
				}

				err = scheduler.Schedule(newCtx, utils.ConvertDBTaskToSchedulerTask(newTask))
				if err != nil{
					if err == context.Canceled{
						tracer.Info("scheduler has been stoppd.")
						return
					}
					tracer.Warn("schedule task failed: %v\n",err)
				}
			}(task)
		}
	}
	return nil
}