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
	err := pprof.StartCPUProfile(&buf)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for  {

		}
	}()
	time.Sleep(10 * time.Second)
	pprof.StopCPUProfile()

	file, err := os.OpenFile("cpu", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.Write(buf.Bytes())
}
