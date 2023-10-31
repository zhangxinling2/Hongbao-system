package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("666"))
	})
	server := http.Server{Addr: "8080", Handler: mux}
	go func() {
		server.ListenAndServe()
	}()
	sig := make(chan os.Signal)
	signal.Notify(sig)
	c := <-sig
	fmt.Println(c.String())
	fmt.Println("退出")
}
