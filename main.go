package main

import (
	"fmt"
)

func greet(n string) {
	fmt.Printf("Good morning %v \n", n)
}

func sayBye(n string) {
	fmt.Printf("Bye %v \n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, name := range n {
		f(name)
	}
}

func main() {
	names := []string{"Emmanuel", "James", "Helloi"}
	cycleNames(names, greet)
	greet("Emmanuel")
}
