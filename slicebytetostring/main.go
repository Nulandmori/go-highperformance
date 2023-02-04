package slicebytetostring

import "fmt"

func main() {
	s := string([]byte{65, 66, 67, 226, 130, 172})
	fmt.Println(s)
}
