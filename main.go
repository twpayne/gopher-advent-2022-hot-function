package main

import "fmt"

var result int

// FIXME tool discussion first, i.e. how to measure, how to diagnose
// FIXME byte-by-byte read or too many syscalls
// FIXME distributed system performance, global etcd, stale connections?

func main() {
	fmt.Println(result)
}
