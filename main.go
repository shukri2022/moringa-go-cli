
package main

import (
    "flag"
    "fmt"
    "time"
)

func main() {
    name := flag.String("name", "Moringa", "Name to greet")
    count := flag.Int("count", 1, "How many times to print")
    flag.Parse()

    for i := 0; i < *count; i++ {
        fmt.Printf("Hello, %s! Time: %s\n", *name, time.Now().Format(time.RFC3339))
    }
}




