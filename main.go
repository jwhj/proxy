package main

import "flag"

func main() {
	var isLocal = flag.Bool("l", false, "")
	var isMid = flag.Bool("m", false, "")
	flag.Parse()
	if *isLocal {
		localServer()
	} else if *isMid {
		midServer()
	}
}
