package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Student struct {
	first_name string
	last_name  string
	SSN        string
	Test1      string
	Test2      string
	Test3      string
	Test4      string
	Final      string
	Grades     string
}

func main() {

	var value int
	for true {
		fmt.Println("Press 1 to read grades.csv")
		fmt.Println("Press 2 to append grades.csv")
		fmt.Println("Press 3 to read exit")
		fmt.Scan(&value)

		switch {
		case value == 1:
			Readfile()
		case value == 2:
			Appendfile()
		case value == 3:
			os.Exit(0)
		default:
			fmt.Println("Enter valid Input")

		}
	}

}

func Readfile() {

	var student [100]Student
	f, err := os.Open("grades.csv")
	var i int = 0
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	s := bufio.NewScanner(f)
	for s.Scan() {
		data := s.Text()
		studen := strings.Split(data, ",")
		student[i].last_name = studen[0]
		student[i].first_name = studen[1]
		student[i].SSN = studen[2]
		student[i].Test1 = studen[3]
		student[i].Test2 = studen[4]
		student[i].Test3 = studen[5]
		student[i].Test4 = studen[6]
		student[i].Final = studen[7]
		student[i].Grades = studen[8]
		i++
	}

	for i = 0; i < 100; i++ {
		if student[i].last_name == "" {
			break
		} else {
			fmt.Println(student[i])
		}

	}

	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func Appendfile() {
	var student Student
	var str string
	f, err := os.OpenFile("grades.csv",
		os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Enter the data to be appended in the following format")
	fmt.Println("Enter the last name")
	fmt.Scan(&str)
	student.last_name = str + ","
	fmt.Println("Enter the first name")
	fmt.Scan(&str)
	student.first_name = str + ","
	fmt.Println("Enter the SSN")
	fmt.Scan(&str)
	student.SSN = str + ","
	fmt.Println("Enter test 1 score")
	fmt.Scan(&str)
	student.Test1 = str + ","
	fmt.Println("Enter test 2 score")
	fmt.Scan(&str)
	student.Test2 = str + ","
	fmt.Println("Enter test 3 score")
	fmt.Scan(&str)
	student.Test3 = str + ","
	fmt.Println("Enter test 4 score")
	fmt.Scan(&str)
	student.Test4 = str + ","
	fmt.Println("Enter the final score")
	fmt.Scan(&str)
	student.Final = str + ","
	fmt.Println("Enter the grade")
	fmt.Scan(&str)
	student.Grades = str + ","
	str = student.last_name + student.first_name + student.SSN + student.Test1 + student.Test2 + student.Test3 + student.Test4 + student.Final + student.Grades
	defer f.Close()
	if _, err := f.WriteString(str); err != nil {
		log.Println(err)
	}

}
