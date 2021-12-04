package main

import (
	"bytes"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

func main() {
	var buf bytes.Buffer

	for i := 0 ; i < 100; i++{
		go func() {
			time.Sleep(2 * time.Second)
		}()
	}

	err := pprof.Lookup("goroutine").WriteTo(&buf,2)
	if err != nil{
		log.Fatal(err)
	}
	file,err := os.OpenFile("goroutine" , os.O_RDWR | os.O_CREATE | os.O_TRUNC,0644 )
	if err != nil{
		log.Fatal(err)
	}
	defer file.Close()

	file.Write(buf.Bytes())

}
