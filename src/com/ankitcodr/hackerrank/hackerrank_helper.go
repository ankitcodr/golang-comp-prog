package hackerrank

import (
	"bufio"
	"os"
	"fmt"
	"io"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	input := readLine(reader)

	fmt.Printf("%s\n", input)//Solution(input))
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

