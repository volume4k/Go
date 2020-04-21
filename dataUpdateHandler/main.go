package main

import (
	"fmt"
	"time"
)

// First

func main(){
	update(time.Minute)
}

func update(t time.Duration){
	// This function should be called with a secondary go routine to always run in the background.
	// The function updates the existingArticles.md file by adding single strings separated by a comma (,).
	// You can add as many PathToFile's as you want, as long as none of them already exists in existingArticles.md.
	// To add a PathToFile simply type the path from the execution directory to the file in the following format:
	// ./path/to/file.tmp

	timer := time.NewTimer(t)
	for {
		<- timer.C
		fmt.Println("Timer done.")
		go handleUpdateCircle()
		timer.Reset(t)
	}
}
