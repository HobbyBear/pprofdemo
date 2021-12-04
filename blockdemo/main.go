package main

import (
	"bytes"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

func main() {
	var (
		buf bytes.Buffer
	)
	runtime.SetBlockProfileRate(1)
	defer runtime.SetBlockProfileRate(0)
	go func() {
		for i := 0; i < 10000000; i++ {
			go func() {
				<-time.After(2  * time.Second)
			}()
		}
	}()
	time.Sleep(2 * time.Second)
	err := pprof.Lookup("block").WriteTo(&buf, 0)
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.OpenFile("block", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.Write(buf.Bytes())
}
