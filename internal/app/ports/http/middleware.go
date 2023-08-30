package http

//func (l *LoggerMiddleware) LoggerMiddleware() mux.MiddlewareFunc {
//	return func(next http.Handler) http.Handler {
//		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//			start := time.Now()
//			path := r.URL.Path
//			raw := r.URL.RawQuery
//
//			lrw := negroni.NewResponseWriter(w)
//			next.ServeHTTP(lrw, r)
//
//			timeStamp := time.Now()
//			latency := timeStamp.Sub(start)
//
//			status := lrw.Status()
//
//			method := r.Method
//
//			bodySize := lrw.Size()
//			if raw != "" {
//				path = path + "?" + raw
//			}
//
//			if status >= 500 {
//				l.logger.SetPrefix("ERROR:\t")
//			} else {
//				l.logger.SetPrefix("INFO:\t")
//			}
//
//			l.logger.Printf("method: %s, path: %s, status_code: %d, body_size: %dB, latency: %s\n",
//				method,
//				path,
//				status,
//				bodySize,
//				latency,
//			)
//		})
//	}
//}
