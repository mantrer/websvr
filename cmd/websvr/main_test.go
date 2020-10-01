package main

import (
	//"net/http/httptest"
	"testing"
	"os"
	"os/exec"

)

func TestGetEnvEmpty(t * testing.T) {
	key := "PORT"
	if os.Getenv("BE_CRASHER") == "1" {
        getEnv(key)
        return
    }

	cmd := exec.Command(os.Args[0], "-test.run=TestGetEnvEmpty")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1", "PORT=")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}


func TestGetEnvNotInit(t * testing.T) {
	
	key := "PORTTxYv6cuv94TxGJsW"
	if os.Getenv("BE_CRASHER") == "1" {
		getEnv(key)
		return
	}
	
	cmd := exec.Command(os.Args[0], "-test.run=TestGetEnvNotInit")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}		
	
	

	// defer  func() {	
	// 	if err := recover(); err != nil {
	// 		fmt.Println ("XXXXXXXXX")
	// 	} else {
	// 		fmt.Println("WWWWWWW")
	// 	}
	// }()
	
	//getEnv(key)