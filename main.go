package main

import (
	"LF-KV/utils/mmap"
	"fmt"
	"os"
)

func main() {
	f, _ := os.OpenFile("D:\\log.txt", os.O_CREATE|os.O_RDWR, 0644)
	buf, _ := mmap.Map(f, mmap.RDWR)
	copy(buf, []byte("呀iu"))
	fmt.Println(string(buf))
	for true {

	}
}
