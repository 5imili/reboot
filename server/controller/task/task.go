package task

import (
	"github.com/5imili/reboot/server/controller"
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"github.com/leopoldxx/go-utils/trace"
	"github.com/leopoldxx/go-utils/middleware"
	"github.com/5imili/reboot/server/utils"
	"io/ioutil"
	"github.com/5imili/reboot/server/service"
	"encoding/json"
)

type task struct{
	opt *controller.Options
}

/*
{
"code":200,
"msg" :"suceess"
}
{
"code":400,
"msg" :"error message",
}
**/

func New(opt *controller.Options) controller.Controller{
	return &task{opt:opt}
}

func (t *task) Register(router *mux.Router){
	///reboot/api/v1/namespace/{xx}/tasks/{task}
	subrouter := router.PathPrefix("/namespaces/{namespace}").Subrouter()
	subrouter.Use(utils.LoggingMiddleware)
	//subrouter.Methods("GET").Path("/tasks/{task}").HandlerFunc(
	//	middleware.RecoverWithTrace("getTask").HandlerFunc(
	//		utils.Authorization().HandlerFunc(t.getTask),
	//	),
	//)
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
	err := t.opt.Service.GetTask(r.Context())
	if err !=nil {
		tracer.Error(err)
		utils.CommReply(w,r,http.StatusBadRequest,err.Error())
		return
	}
	utils.CommReply(w,r, http.StatusOK, "success")
}

func (t *task) listTask(w http.ResponseWriter, r *http.Request){
	tracer := trace.GetTraceFromRequest(r)
	tracer.Info("call listTask")
	fmt.Fprintln(w,"call listTask")
}

func (t *task) createTask(w http.ResponseWriter, r *http.Request){
	tracer := trace.GetTraceFromRequest(r)
	vars := mux.Vars(r)
	ns := vars["namespace"]
	data , err := ioutil.ReadAll(r.Body)
	if err != nil{
		tracer.Error(err)
		utils.CommReply(w,r,http.StatusBadRequest,err.Error())
		return
	}
	tracer.Info(string(data))
	info := &service.Task{}
	err = json.Unmarshal(data, info)
	if err != nil{
		tracer.Error(err)
		utils.CommReply(w,r,http.StatusBadRequest,err.Error())
		return
	}

	//Todo: validate
	task , err := t.opt.Service.CreateTask(r.Context(),ns,info.Resource)
	if err != nil{
		tracer.Error(err)
		utils.CommReply(w,r,http.StatusBadRequest,err.Error())
		return
	}
	js,_ := json.Marshal(task)
	utils.CommReply(w,r,http.StatusAccepted,string(js))
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