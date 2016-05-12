package main

import "flag"

func main() {
	srcPtr := flag.String("src", ".", "the folder contaning series to sort")
	dstPtr := flag.String("dst", "../Series", "the folder where the series should go ")

	flag.Parse()
	Sort(*srcPtr, *dstPtr)
	Clean(*srcPtr)

}
