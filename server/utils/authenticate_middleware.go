package utils

import (
	"github.com/leopoldxx/go-utils/middleware"
	"net/http"
	"github.com/leopoldxx/go-utils/trace"
	"fmt"
)

var tokenUsers = map[string]string {
	"123":"chenchao",
	"456":"songjiang",
	"88e89def2a79407ab8902a1d5506581a":"wangfei",
}

func LoggingMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		tracer := trace.GetTraceFromRequest(r)
		tracer.Info(r.RequestURI)
		next.ServeHTTP(w,r)
	})
}

func Authorization() middleware.Middleware{
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request){
			tracer := trace.GetTraceFromRequest(r)
			tracer.Info("call Authorization")
			tokens ,_ := r.Header["Authorization"]
			if len(tokens) == 0{
				fmt.Fprintln(w,"Authorization error")
				return
			}
			if _ ,ok := tokenUsers[tokens[0]] ; !ok{
				fmt.Fprintf(w,"There is no such token:%s\n",tokens[0])
				return
			}
			next(w,r)
		}
	}
}
