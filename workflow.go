package main

import (
  "flag"
  "fmt"
)

var specMode bool

func init() {
  const (
    defaultSpecMode = false
    usage = "enable spec mode"
  )
  flag.BoolVar(&specMode, "mode-spec", defaultSpecMode, usage)
  flag.BoolVar(&specMode, "ms", defaultSpecMode, usage+" (shorthand)")
}

func main() {
  flag.Parse()
  fmt.Printf("specMode: %v\n", specMode)
}
