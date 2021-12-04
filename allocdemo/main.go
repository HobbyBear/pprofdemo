package main

import (
	"bytes"
	"log"
	"os"
	"runtime/pprof"
)

func main() {
	var buf bytes.Buffer

	arr := make([]int,10000000000)
	arr[1] = 0

	err := pprof.Lookup("allocs").WriteTo(&buf,0)
	if err != nil{
		log.Fatal(err)
	}
	file, err := os.OpenFile("allocs", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.Write(buf.Bytes())
}
