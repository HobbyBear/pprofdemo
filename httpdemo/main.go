package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	arr := make([]string, 0, 10000000)

	go func() {
		for {
			time.Sleep(200 * time.Millisecond)
			arr = append(arr, "aaaaaaaaa")
			log.Println(1)
		}
	}()
	http.ListenAndServe(":39090", nil)
}
