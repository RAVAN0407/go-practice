package main

import "fmt"

func binarysearch(arr []int, s int) {
    var left int = 0
    var right int = len(arr) - 1
    for left <= right {
        var m int = left + (right-left)/2
        if arr[m] == s {
            fmt.Println("postion", m+1)
            break
        } else if arr[m] < s {
            left = m + 1
        } else {
            right = m - 1
        }
    }

}

func main() {
    arr := []int{1, 5, 7, 9, 13, 23, 45, 76}
    s := 9
    binarysearch(arr, s)

}
