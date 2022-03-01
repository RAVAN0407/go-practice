package main

import (
    "bufio"
    "flag"
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
    fptr := flag.String("fpath", "grades.csv", "file path to read from")
    flag.Parse()
    var student Student
    f, err := os.Open(*fptr)
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
        student.last_name = studen[0]
        student.first_name = studen[1]
        student.SSN = studen[2]
        student.Test1 = studen[3]
        student.Test2 = studen[4]
        student.Test3 = studen[5]
        student.Test4 = studen[6]
        student.Final = studen[7]
        student.Grades = studen[8]
        fmt.Println(studen)
    }
    err = s.Err()
    if err != nil {
        log.Fatal(err)
    }
}

func Appendfile() {
    var str string
    f, err := os.OpenFile("grades.csv",
        os.O_APPEND|os.OWRONLY, 0644)
    if err != nil {
        log.Println(err)
    }
    fmt.Println("Enter the data to be appended")
    fmt.Scan(&str)
    defer f.Close()
    if , err := f.WriteString(str); err != nil {
        log.Println(err)
    }

}
