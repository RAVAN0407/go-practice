package main

import (
    "encoding/csv"
    "fmt"
    "os"
    "strconv"
)

type Student struct {
    first_name string
    lastname  string
    SSN        string
    Test1      int
    Test2      int
    Test3      int
    Test4      int
    Final      int
    Grades     string
}

func main() {
    records, err := ReadCsv("grades.csv")
    if err != nil {
        panic(err)
    }

    var student []Student
    for , record := range records {
        test1,  := strconv.Atoi(record[3])
        test2,  := strconv.Atoi(record[4])
        test3,  := strconv.Atoi(record[5])
        test4,  := strconv.Atoi(record[6])
        final, _ := strconv.Atoi(record[7])
        data := Student{
            first_name: record[0],
            last_name:  record[1],
            SSN:        record[2],
            Test1:      test1,
            Test2:      test2,
            Test3:      test3,
            Test4:      test4,
            Final:      final,
            Grades:     record[8],
        }
        student = append(student, data)
    }
    fmt.Println(student)
}

func ReadCsv(filename string) ([][]string, error) {
    file, err := os.Open(filename)
    if err != nil {
        return [][]string{}, err
    }
    defer file.Close()

    records, err := csv.NewReader(file).ReadAll()
    if err != nil {
        return [][]string{}, err
    }

    return records, nil
}
