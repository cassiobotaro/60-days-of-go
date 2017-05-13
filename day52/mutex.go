package main

import (
	"sync"
	"time"
)

// code extracted by http://stackoverflow.com/questions/36167200/how-safe-are-golang-maps-for-concurrent-read-write-operations
var m = map[string]int{"a": 1}
var lock = sync.RWMutex{}

// read and write maps are not thread safe.
func main() {
	go Read()
	time.Sleep(1 * time.Second)
	go Write()
	time.Sleep(1 * time.Minute)
}

// Read reads an element of the map
func Read() {
	for {
		read()
	}
}

// Write writes an element in the map
func Write() {
	for {
		write()
	}
}

func read() {
	// if you want to see the race condition,
	// comment the lines below and run `go run -race mutex.go`
	lock.RLock()
	defer lock.RUnlock()
	_ = m["a"]
}

func write() {
	// if you want to see the race condition,
	// comment the lines below and run `go run -race mutex.go`
	lock.Lock()
	defer lock.Unlock()
	m["b"] = 2
}
