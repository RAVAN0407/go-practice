package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

type file__attributes struct {
	name      string
	size      int
	file_perm string
	seconds   int
}

var fp [100]file__attributes

func main() {

	var arg string
	cmd1 := "ls"
	cmd2 := "-la"
	cmd3 := "--time-style=+%s"
	fmt.Println("Enter the path of file")
	fmt.Scan(&arg)
	cmd_data, err := exec.Command(cmd1, cmd2, cmd3, arg).Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	split := strings.Split(string(cmd_data), "\n")
	sliced_data := split_data(string(cmd_data))
	for i := 1; i < (len(split) - 1); i++ {
		fp[i].seconds, _ = strconv.Atoi(sliced_data[i][5])
		fp[i].name = sliced_data[i][6]
		fp[i].size, err = strconv.Atoi(sliced_data[i][4])
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fp[i].file_perm = sliced_data[i][0]
	}
	sizesort(fp)
	timesort(fp)

}

func split_data(a string) [100][]string {

	split := strings.Split(string(a), "\n")
	var split2 [100][]string
	singleSpacePattern := regexp.MustCompile(`\s+`)
	for i := 1; i < 15; i++ {
		str := split[i]
		single_str := singleSpacePattern.ReplaceAllString(str, " ")
		split2[i] = strings.Split(single_str, " ")
	}
	return split2

}
func sizesort(a [100]file__attributes) {
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j].size > a[j+1].size {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	fmt.Println()
	fmt.Println("sorted based on size")
	fmt.Println()
	for i := 0; i < len(a); i++ {
		if a[i].size != 0 {

			fmt.Println(a[i].seconds, "\t\t\t\t", a[i].size, "\t\t\t\t", a[i].name)
		}
	}

}

func timesort(a [100]file__attributes) {

	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j].seconds > a[j+1].seconds {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	fmt.Println()
	fmt.Println("sorted based on time")
	fmt.Println()
	for i := 0; i < len(a); i++ {
		if a[i].size != 0 {

			fmt.Println(a[i].seconds, "\t\t\t\t", a[i].size, "\t\t\t\t", a[i].name)
		}
	}

}
