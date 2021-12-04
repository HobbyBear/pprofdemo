package main

import (
	"net/http"
	_ "net/http/pprof"
)

func main() {
	go func() {
		for i := 0; i < 10; i++ {
			go func() {}()
		}
	}()
	http.ListenAndServe(":39090", nil)
}
