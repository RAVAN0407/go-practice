package main

import (
	"fmt"
	"os/exec"
	"strings"
)

type file__attributes struct {
	name      string
	size      string
	file_perm string
}

func main() {
	cm := "ls"
	cm2 := "-l"
	var arg string
	//var arr [100]int
	var fa [100]file__attributes
	var split2 [100][]string
	fmt.Println("Enter the file path")
	fmt.Scan(&arg)
	cmd1, err := exec.Command(cm, cm2, arg).Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	split := strings.Split(string(cmd1), "\n")
	for i := 0; i < len(split); i++ {
		split2[i] = strings.Split(split[i], " ")
	}

	for j := 0; j < 10; j++ {

		fmt.Println(split2[1][j])
	}

	for i := 1; i < (len(split) - 2); i++ {
		fa[i].file_perm = split2[i][0]
		fa[i].size = split2[i][5]
		fa[i].name = split2[i][8]
	}

	//for i := 0; i < len(split); i++ {
	//fmt.Println(split[i], i)
	//}
	//fmt.Println(split2[8][9])
	//for i := 0; i < len(split); i++ {
	//arr[i], err = strconv.Atoi(fa[i].size)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//fmt.Println(arr[i])
	//}

}
