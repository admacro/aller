// https://golang.google.cn/ref/spec#Go_statements
package main

import (
	"fmt"
	"time"
)

// A Go routine is a lightweight thread managed by Go runtime.
// go f(x, y, z) starts a new goroutine running f(x, y, z).

// The evaluation of f(x, y, z) happens in the current goroutine and the
// execution of f(x, y, z) happens in the new goroutine.

// Goroutines run in the same address space, so access to shared memory must be
// synchronized. The `sync` package provides useful primitives, although you
// don't need them much in Go as there is better choices (channel).

func sayHi(from string, to string) {
	fmt.Printf("%v: Hi %v!\n", from, to)
}

func main() {
	newKid := "Tom"
	fmt.Printf("Mrs. Helen: boys and girls, this is %v. Can you say hi to %v?\n",
		newKid, newKid)
	kids := []string{"Jack", "Emma", "Johnny", "Cindy"}
	for i, kid := range kids {
		// time.Millisecond is of type time.Duration
		time.Sleep(time.Duration(200*i) * time.Millisecond)
		go sayHi(kid, newKid)
	}

	// hardcoded waiting time (0.5s)
	// This waiting is necessary, as gorotunes are terminated when main() exists.
	// In this case the main() almost always exits before the goroutines ever start,
	// and you won't be able to see any output from the goroutines without explicit waiting.
	time.Sleep(500 * time.Millisecond)

	// Yet, there are better (or correct) ways to wait for all goroutines to complete
	// before the main program exits, such as:
	//   1. sync.WaitGroup
	//   2. channel
}
