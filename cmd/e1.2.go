// echo prints its command-line arguments
// with each's index number first followed by a space and each argument on a new line
package main

import (
	"fmt"
	"os"
)

func main() {
// avoids string concatentaion
	for i, v := range os.Args[:] {
		fmt.Print(i)
		fmt.Print(" ")
		fmt.Println(v)
	}
}
