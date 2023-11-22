package main

import (
	"fmt"

	"github.com/jboursiquot/go-proverbs"
)

func main() {
	fmt.Println(getProverbs())
}

func getProverbs() (proverb string) {
	proverb = proverbs.Random().Saying
	return
}
