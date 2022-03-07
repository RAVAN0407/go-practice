package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type file__attributes struct {
	name      string
	size      int
	file_perm string
	date      string
	time1     string
	gmt       string
}

var fp [100]file__attributes

func main() {

	//sliced_data := make([]string,100,100)
	var arg string
	cmd1 := "ls"
	cmd2 := "-la"
	cmd3 := "--time-style=full-iso"
	fmt.Println("Enter the path of file")
	fmt.Scan(&arg)
	cmd_data, err := exec.Command(cmd1, cmd2, cmd3, arg).Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	split := strings.Split(string(cmd_data), "\n")
	sliced_data := split_data(string(cmd_data))
	//fmt.Println(sliced_data[15][8])
	for i := 1; i < (len(split) - 1); i++ {
		fp[i].date = sliced_data[i][5]
		fp[i].time1 = sliced_data[i][6]
		fp[i].gmt = sliced_data[i][7]
		fp[i].name = sliced_data[i][8]
		fp[i].size, err = strconv.Atoi(sliced_data[i][4])
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fp[i].file_perm = sliced_data[i][0]
	}
	sizesort(fp)
	//timesort(fp)

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
	fmt.Println("sorted based on size")
	for i := 0; i < len(a); i++ {
		if a[i].size != 0 {

			fmt.Println(a[i].size, a[i].name)
		}
	}

}

func timesort(a [100]file__attributes) {  //error
	var b []int64
	for i := 0; i < len(a); i++ {

		str := a[i].date
		//fmt.Println(str)
		tt, er := time.Parse("2006-01-02", str)
		if er != nil {
			panic("Can't parse time format")
		}
		epoch := tt.Unix()
		b[i] = epoch

	}

	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if b[j] > b[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	fmt.Println("sorted based on size")
	for i := 0; i < len(a); i++ {
		if a[i].size != 0 {

			fmt.Println(a[i].size, a[i].name, a[i].date, a[i].time1)
		}
	}

}
