package main

import (
	"bytes"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
	"time"
)

func main() {
	var (
		buf  bytes.Buffer
		lock sync.Mutex
	)
	runtime.SetMutexProfileFraction(1)
	defer runtime.SetMutexProfileFraction(0)
	go func() {
		for i := 0; i < 10000000; i++ {
			go func() {
				lock.Lock()
				//time.Sleep(1 * time.Second)
				lock.Unlock()
			}()
		}
	}()
	time.Sleep(2 * time.Second)
	err := pprof.Lookup("mutex").WriteTo(&buf, 0)
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.OpenFile("mutex", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.Write(buf.Bytes())
}
