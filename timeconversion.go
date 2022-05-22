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
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(s string) string {
	// Write your code here

	h := []string{}
	k := ""
	f := []string{}
	for i, v := range s {
		if i > 1 && string(v) != string(s[len(s)-2]) && string(v) != string(s[len(s)-1]) {
			h = append(h, string(v))
		}
	}
	for i, v := range s {
		if i <= 1 && string(s[1]) != ":" {
			f = append(f, string(v))
		}
	}
	a := strings.Join(f, "")
	if strings.Contains(s, "PM") {

		sl, _ := strconv.Atoi(a)
		if sl == 12 {

			k = strconv.Itoa(sl) + strings.Join(h, "")

			//fmt.Println(k)
		} else {
			j := sl + 12
			k = strconv.Itoa(j) + strings.Join(h, "")

			//fmt.Println(k)
		}

	} else if strings.Contains(s, "AM") {

		sl, _ := strconv.Atoi(a)
		if sl == 12 {

			k = "00" + strings.Join(h, "")

			//fmt.Println(k)
		} else if sl < 10 {

			k = "0" + strconv.Itoa(sl) + strings.Join(h, "")

			//fmt.Println(k)
		}
	}
	return k
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := timeConversion(s)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
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
