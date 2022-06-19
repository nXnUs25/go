package main

import (
	"sync"
)

type FileRW struct {
	sync.RWMutex
}

func main() {

}

func ReadFile(path string, frw *FileRW) (err error) {
	frw.RLock()

}
