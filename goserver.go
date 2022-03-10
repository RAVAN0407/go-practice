package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"
)

const portNumber = ":8080"

var kernel_v []byte
var cpu []string

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "kernel : ")
	w.Write([]byte(kernel_v))
	w.Write([]byte(cpu[12]))
	fmt.Fprintf(w, "\n")
	w.Write([]byte(cpu[4]))
}

func main() {

	cmd1 := "uname"
	cmd2 := "-r"
	kernel_v, _ = exec.Command(cmd1, cmd2).Output()
	cmd3 := "cat"
	cmd4 := "/proc/cpuinfo"
	cpuinfo, _ := exec.Command(cmd3, cmd4).Output()
	cpu = strings.Split(string(cpuinfo), "\n")
	http.HandleFunc("/", HomePageHandler)
	fmt.Printf("Starting application on port %v\n", portNumber)
	http.ListenAndServe(portNumber, nil)
}
