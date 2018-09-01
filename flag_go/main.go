package main

import (
	"flag"
	"fmt"
)

var procfile = flag.String("f", "Procfile", "proc file")
var deamon = flag.Bool("c", false, "start as deamon")
var num = flag.Int("n", 123, "num of thread")

func main() {
	flag.Parse()
	flag.NewFlagSet().Init()
	fmt.Println(*procfile)
	fmt.Println(*deamon)
	fmt.Println(*num)

}
