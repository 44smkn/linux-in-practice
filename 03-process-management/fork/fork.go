package main

import (
	"fmt"
	"syscall"
)

func child() {
	fmt.Printf("I'm child! my pid is %v.\n", syscall.Getpid())
}

func parent(child uintptr) {
	fmt.Printf("I'm parent! my pid is %v and the pid of my child is %v.\n", syscall.Getpid(), child)
}

func main() {
	ret, _, _ := syscall.RawSyscall(syscall.SYS_FORK, 0, 0, 0)
	if ret == 0 {
		child()
	} else {
		parent(ret)
	}

}
