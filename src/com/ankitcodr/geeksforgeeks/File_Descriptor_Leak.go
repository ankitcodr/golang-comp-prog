package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

// At least on OSX 10.10 (Yosemite, didn't check on Mavericks),
// you can get the list of open files by process via the default
// activity monitor application. Just double click on the relevant
// process on the list and select "Open Files and Ports" tab on the popup.

func main() {
	fmt.Printf("Start: Outside GoRoutines Open FD: %v\n", countOpenFiles())

	var done chan bool
	writeGoRoutine := 20
	for i := 0; i < writeGoRoutine; i++ {
		go func() {
			fileHandle, _ := os.OpenFile("/tmp/dat2", os.O_APPEND|os.O_WRONLY, 0666)
			writer := bufio.NewWriter(fileHandle)
			defer fileHandle.Close()

			fmt.Fprintln(writer, "String I want to append")
			err := writer.Flush()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("GoRoutine %d Open FD: %v\n", i, countOpenFiles())
			fmt.Println(<-done)
		}()
	}

	time.Sleep(3 * time.Second)
	fmt.Printf("End: Outside GoRoutines Open FD: %v\n", countOpenFiles())

	<-done
}

func countOpenFiles() int {
	out, err := exec.Command("/bin/sh", "-c", fmt.Sprintf("lsof -p %v", os.Getpid())).Output()
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(out), "\n")
	return len(lines) - 1
}
