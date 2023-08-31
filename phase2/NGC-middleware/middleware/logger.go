package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

// logger
func LoggerMiddleware(next httprouter.Handle) httprouter.Handle {
    return func(w http.ResponseWriter, r *http.Request,p httprouter.Params) {
        // Waktu sekarang
        currentTime := time.Now().Format("2006/01/02 15:04:05")

        // Format log
        logMessage := fmt.Sprintf("[%s] - HTTP request sent to %s %s", currentTime, r.Method, r.URL.Path)

        // Menampilkan log
        fmt.Println(logMessage)

        //
        next(w, r,p)
    }
}

func Log(router http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		currentTime := time.Now()
		log := fmt.Sprintf("[%v] - HTTP Request sent to %v %v", currentTime.Format("2006/01/02 15:04:05"), r.Method, r.URL.Path)
		fmt.Println(log)

		router.ServeHTTP(w, r)
	})
}
