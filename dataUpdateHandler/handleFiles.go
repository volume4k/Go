package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func handleUpdateCircle(){
	alreadyExists, toWrite := handleFileIndex()
	if alreadyExists {
		fmt.Println("The File is already loaded.")
		cleanupInit()
		return
	}
	if toWrite == "" {
		return
	}
	fmt.Println("Applying file loading circle.")
	addArticlePathToExistsFile(toWrite)
	cleanupInit()
}

func handleFileIndex() (bool, string){
	initFile, err := os.Open("initArticle.md")
	if err != nil {
		panic(err)
	}
	defer initFile.Close()
	byteRaw, err := ioutil.ReadAll(initFile)
	if err != nil {
		panic(err)
	}
	result := string(byteRaw)
	return checkForDuplicates(result), result
}

func checkForDuplicates(s string) bool {
	existsFile, err := os.Open("existingArticles.md")
	if err != nil {
		panic(err)
	}
	defer existsFile.Close()

	byteRaw, err := ioutil.ReadAll(existsFile)
	if err != nil {
		panic(err)
	}

	ss := strings.Split(s,",")
	result := string(byteRaw)
	res := strings.Split(result, ",")
	for _, inp := range ss{
		for _, stc := range res {
			if stc == inp {
				return true
			}
		}
	}
	return false
}

func addArticlePathToExistsFile (s string){
	file, err := os.Open("existingArticles.md")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	byteRaw, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	res := byteSliceToStringSlice(byteRaw)
	toWrite := addPathToSliceToString(res, s)
	fmt.Println(toWrite)
	writeFile(toWrite, "existingArticles.md")
}

func writeFile(s string, location string) {
	var file, err = os.OpenFile(location, os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString(s)
	if err != nil {
		panic(err)
	}

	err = file.Sync()
	if err != nil {
		panic(err)
	}
}

func byteSliceToStringSlice(b []byte) []string {
	str := string(b)
	ret := strings.Split(str,",")
	return ret
}

func addPathToSliceToString (ss []string, s string) string {
	conv := append(ss, s)
	ret := strings.Join(conv,",")
	return ret
}

func cleanupInit (){
	err := os.Remove("initArticle.md")
	if err != nil {
		panic(err)
	}
	_, err = os.Create("initArticle.md")
	if err != nil {
		panic(err)
	}
}