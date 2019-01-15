package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"
)

func child() {
	fmt.Printf("I'm child! my pid is %v.\n", syscall.Getpid())
	binary, lookErr := exec.LookPath("echo")
	if lookErr != nil {
		panic(lookErr)
	}
	args := []string{"echo", "hello"}
	env := os.Environ()
	syscall.Exec(binary, args, env)
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
		time.Sleep(3 * time.Second)
	}

}
