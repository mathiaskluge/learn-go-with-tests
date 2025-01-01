package main

import (
	"fmt"
	"io"
)

func main() {
	//Countdown(os.Stdout)
}

func Countdown(out io.Writer) {
	for i := 3; i > 0; i-- {
		fmt.Fprint(out, i)
	}
	fmt.Fprint(out, "Go!")

}
