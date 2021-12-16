## deriving π given a normal distribution

Try it yourself:
```go
package main

import (
    "fmt"

	"github.com/cpmech/gosl/rnd"
)

// normal returns a number according to the normal distribution in range [0, 1].
func normal() float64 {
	return rnd.Normal(0.5, 0.12)
}

func main() {
    // Your solution here

	fmt.Println("hello, world")
}
```

This repository contains a possible solution to this problem. But unfortunately, it derives π only up to the second digit. These parameters were the most decent: `n=10_000_000` and `r=0.01`
