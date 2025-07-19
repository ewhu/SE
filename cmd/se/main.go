// cmd/se/main.go
package main

import (
"flag"
"log"
"os"

"se/internal/se"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := se.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
