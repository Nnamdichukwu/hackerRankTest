package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func plusMinus(arr []int32) {
	// Write your code here
	var i []int32
	var j []int32
	var y []int32
	g := float32(len(arr))
	for _, v := range arr {
		if v > 0 {
			i = append(i, v)

		} else if v < 0 {
			j = append(j, v)

		} else if v == 0 {
			y = append(y, v)
		}

	}
	h := float32(float32(len(i)) / g)
	k := float32(float32(len(j)) / g)
	m := float32(float32(len(y)) / g)
	fmt.Printf("%.6f \n", h)
	fmt.Printf("%.6f \n", k)
	fmt.Printf("%.6f \n", m)
}
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	plusMinus(arr)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
