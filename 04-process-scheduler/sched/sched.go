package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"syscall"
	"time"
)

const (
	nloopForEstimation = 1000000000
	nsecsPerMsec       = 1000000
)

func loopsPerMsec() uint64 {
	start := time.Now()
	for i := 0; i < nloopForEstimation; i++ {
	}
	end := time.Now()
	diff := end.Sub(start).Nanoseconds()
	return nloopForEstimation * nsecsPerMsec / uint64(diff)
}

func load(nloop uint64) {
	for i := uint64(0); i < nloop; i++ {
	}
}

func childFn(id int, buf []time.Time, nrecord int, nLoopPerResol uint64, start time.Time) {
	for i := 0; i < nrecord; i++ {
		load(nLoopPerResol)
		buf[i] = time.Now()
	}
	for i := 0; i < nrecord; i++ {
		diff := buf[i].Sub(start).Nanoseconds()
		fmt.Printf("%v\t%v\t%v\n", id, diff/nsecsPerMsec, (i+1)*100/nrecord)
	}
	os.Exit(0)
}

func main() {
	if len(os.Args) < 4 {
		log.Fatalf("usage: %v <nproc> <total[ms]> <resolution[ms]>\n", os.Args[0])
	}

	nproc, _ := strconv.Atoi(os.Args[1])
	total, _ := strconv.Atoi(os.Args[2])
	resol, _ := strconv.Atoi(os.Args[3])

	if nproc < 1 {
		log.Fatalf("<nproc>(%v) should be >= 1\n", nproc)
	}
	if total < 1 {
		log.Fatalf("<total>(%v) should be >= 1\n", total)
	}
	if resol < 1 {
		log.Fatalf("<resol>(%v) should be >= 1\n", resol)
	}
	if total%resol != 0 {
		log.Fatalf("<total>(%v) should be multiple of <resolution>(%v)\n", total, resol)
	}

	nrecord := total / resol
	logbuf := make([]time.Time, nrecord)

	fmt.Println("estimating workload which takes just one milisecond")
	nLoopPerResol := loopsPerMsec() * uint64(resol)
	fmt.Println("end estimation")

	//pids := make([]uintptr, nproc)
	start := time.Now()

	for i := 0; i < nproc; i++ {
		ppid, pid, _ := syscall.RawSyscall(syscall.SYS_FORK, 0, 0, 0)
		if ppid == 0 {
			childFn(i, logbuf, nrecord, nLoopPerResol, start)
		} else if pid == 0 {
			time.Sleep(time.Duration(2*total) * time.Millisecond)
		}
	}
}
