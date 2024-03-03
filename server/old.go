package main

//
//import (
//	"net/http"
//
//	"contrib.go.opencensus.io/exporter/prometheus"
//	log "github.com/sirupsen/logrus"
//	"go.opencensus.io/plugin/ochttp"
//	"go.opencensus.io/stats/view"
//)
//
//func main() {
//	stop := make(chan struct{})
//
//	server := &http.Server{Addr: ":8080"}
//
//	statsMux := http.NewServeMux()
//	statsServer := &http.Server{Handler: statsMux, Addr: ":57475"}
//
//	if err := view.Register(ochttp.DefaultServerViews...); err != nil {
//		log.WithError(err).Fatal("register HTTP metrics view")
//	}
//
//	exporter, err := prometheus.NewExporter(prometheus.Options{
//		Namespace: "default",
//	})
//	if err != nil {
//		log.WithError(err).Fatal("create Prometheus exporter")
//	}
//
//	view.RegisterExporter(exporter)
//
//	statsMux.Handle("/metrics", exporter)
//
//	originalHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		w.Write([]byte("Hello, World!"))
//	})
//	och := &ochttp.Handler{
//		Handler: originalHandler,
//	}
//
//	server.Handler = och
//
//	go func() {
//		log.Info("Starting stats server...")
//		if err := statsServer.ListenAndServe(); err != nil {
//			log.WithError(err).Fatal("listen and serve stats")
//		}
//	}()
//
//	go func() {
//		log.Info("Starting server...")
//		if err := server.ListenAndServe(); err != nil {
//			log.WithError(err).Fatal("listen and serve service endpoints")
//		}
//	}()
//
//	<-stop
//}
