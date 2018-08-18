package task

import (
	"github.com/5imili/reboot/server/controller"
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"github.com/leopoldxx/go-utils/trace"
	"github.com/leopoldxx/go-utils/middleware"
)

type task struct{
	opt *controller.Options
}

func New(opt *controller.Options) controller.Controller{
	return &task{opt:opt}
}

func (t *task) Register(router *mux.Router){
	///reboot/api/v1/namespace/{xx}/tasks/{task}
	subrouter := router.PathPrefix("/namespaces/{namespace}").Subrouter()

	subrouter.Methods("GET").Path("/tasks/{task}").HandlerFunc(
		middleware.RecoverWithTrace("getTask").HandlerFunc(t.getTask),
		)

	subrouter.Methods("GET").Path("/tasks").HandlerFunc(
		middleware.RecoverWithTrace("listTask").HandlerFunc(t.listTask),
		)
	subrouter.Methods("POST").Path("/tasks").HandlerFunc(
		middleware.RecoverWithTrace("createTask").HandlerFunc(t.createTask),
		)

	subrouter.Methods("DELETE").Path("/tasks/{task}").HandlerFunc(
		middleware.RecoverWithTrace("deleteTask").HandlerFunc(t.deleteTask),
		)
	subrouter.Methods("PUT").Path("/tasks/{task}").HandlerFunc(
		middleware.RecoverWithTrace("updateTask").HandlerFunc(t.updateTask),
		)
}

func (t *task) getTask(w http.ResponseWriter, r *http.Request){
	tracer := trace.GetTraceFromRequest(r)
	tracer.Info("call getTask")
	fmt.Fprintln(w,"call getTask")
}

func (t *task) listTask(w http.ResponseWriter, r *http.Request){
	tracer := trace.GetTraceFromRequest(r)
	tracer.Info("call listTask")
	fmt.Fprintln(w,"call listTask")
}

func (t *task) createTask(w http.ResponseWriter, r *http.Request){
	tracer := trace.GetTraceFromRequest(r)
	tracer.Info("call createTask")
	fmt.Fprintln(w,"call createTask")
}

func (t *task) deleteTask(w http.ResponseWriter, r *http.Request){
	tracer := trace.GetTraceFromRequest(r)
	tracer.Info("call deleteTask")
	fmt.Fprintln(w,"call deleteTask")
}

func (t *task) updateTask(w http.ResponseWriter, r *http.Request){
	tracer := trace.GetTraceFromRequest(r)
	tracer.Info("call updateTask")
	fmt.Fprintln(w,"call updateTask")
}