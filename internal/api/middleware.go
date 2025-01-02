package api

import (
	"log"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 记录请求信息
		log.Printf("请求 %s %s", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
