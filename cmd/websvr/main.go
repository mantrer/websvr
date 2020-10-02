// main
package main

import (
	"log"
	"net/http"
	"context"
	"os"
	"os/signal"
	"syscall"
	//"fmt"
	"127.0.0.1/version"
	"go.uber.org/zap"
)

const root string = "/site/sunsea.ru/public"

type config struct {
	Port string
}

func New() *config {
	return &config{Port: getEnv("PORT")}
}


func getEnv(key string) string {

	value, exists := os.LookupEnv(key)
	if ! exists {
		zap.S().Fatalf("Variable %s is not initialized", key)
	}
	if value == "" {
		zap.S().Fatalf("Variable %s is empty", key)
	}
	return value
}

func loggerInit() {
		// init logger & replace global logger
		logger, err := zap.NewDevelopment()
		zap.ReplaceGlobals(logger)
		if err != nil {
			log.Fatalf("can't initialize zap logger: %v", err)
		  }
		defer logger.Sync()
}

func healthz(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	loggerInit()
	sugar := zap.S()

	log.Printf(
		"Starting the service...\ncommit: %s, build time: %s, release: %s",
		version.Commit, version.BuildTime, version.Release,
	)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	handler := http.FileServer(http.Dir(root))
	http.Handle("/", handler)
	http.HandleFunc("/healthz", healthz)
	cfg := New()
	port:=cfg.Port
	sugar.Infof ("Starting server at :%s", port)

	svr := http.Server{
		Addr: ":" + port,
	}
	go func() {
		//sugar.Fatal(http.ListenAndServe(":"+port, nil))
		//sugar.Fatal(svr.ListenAndServe())
		svr.ListenAndServe()
	}()
	killSignal := <-interrupt

	switch killSignal {
	case os.Interrupt:
		sugar.Info("Got SIGINT...")
	case syscall.SIGTERM:
		sugar.Info("Got SIGTERM...")
	}
	sugar.Info("The service is shutting down...")
	svr.Shutdown(context.Background())
	sugar.Info("Done")

}
