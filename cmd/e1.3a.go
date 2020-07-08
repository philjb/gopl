// echo prints its command-line arguments
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("%dus elapsed\n", time.Since(start).Microseconds())
}
