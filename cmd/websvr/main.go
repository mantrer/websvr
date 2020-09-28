// main
package main

import (
	"log"
	"net/http"
	"os"
	"go.uber.org/zap"
)

const root string = "/site/hugo/sunsea.ru/public/"

type config struct {
	Port string
}

func (c *config) New() {
	c.Port = getEnv("PORT")
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

func main() {
	loggerInit()
	sugar := zap.S()
	
	handler := http.FileServer(http.Dir(root))
	http.Handle("/", handler)
	cfg := new(config)
	cfg.New()
	port:=cfg.Port
	sugar.Infof ("Starting server at :%s", port)
	sugar.Fatal(http.ListenAndServe(":"+port, nil))

}
