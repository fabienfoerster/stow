package main

import (
	"flag"
	"log"
)

func main() {
	srcPtr := flag.String("src", ".", "the folder contaning series to sort")
	dstPtr := flag.String("dst", "../Series", "the folder where the series should go ")

	flag.Parse()
	Sort(*srcPtr, *dstPtr)
	err := CleanSubDir(*srcPtr)
	if err != nil {
		log.Printf("Unable to clean the dir %s : %s", *srcPtr, err)
	}

}
