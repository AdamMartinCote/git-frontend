package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type AppConfig struct {
	Port    string
	GitRoot string
}

func NewConfig() *AppConfig {
	c := AppConfig{}
	p := os.Getenv("APP_PORT")
	if len(p) == 0 {
		p = "80"
	}
	c.Port = p
	gr := os.Getenv("GIT_ROOT")
	c.GitRoot = gr
	fmt.Println(gr)
	return &c
}

func main() {
	conf := NewConfig()
	fmt.Printf("Starting front end service on port %s\n", conf.Port)

	srv := &http.Server{
		Addr:    fmt.Sprintf("localhost:%s", conf.Port),
		Handler: conf.routes(),
	}

	// Start the server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
