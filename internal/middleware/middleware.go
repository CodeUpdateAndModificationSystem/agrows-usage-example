package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/felixge/httpsnoop"
	"github.com/gorilla/handlers"
)

type Middleware func(http.Handler) http.Handler

func CreateStack(xs ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(xs) - 1; i >= 0; i-- {
			x := xs[i]
			next = x(next)
		}

		return next
	}
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cap := httpsnoop.CaptureMetrics(next, w, r)
		// datetime endpoint method status duration
		starttime := time.Now().Add(-cap.Duration)
		fmt.Printf("%s %s %s %d %s %.4fms\n", starttime.Format("2006-01-02 15:04:05"), r.URL.Path, r.Method, cap.Code, cap.Duration, float64(cap.Duration.Abs().Microseconds())/1000.0)

		// size := humanize.Bytes(uint64(cap.Written))
		// log.Info("%4s %-20s %9s %3d %8s %.4fms", r.Method, r.URL.Path, r.Proto, cap.Code, size, float64(cap.Duration.Abs().Microseconds())/1000.0)
	})
}

func Recovery(next http.Handler) http.Handler {
	return handlers.RecoveryHandler()(next)
}
