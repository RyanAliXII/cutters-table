package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"test/cutters/data"
)

const ALPHABET_STRING string = "abcdefghijklmnopqrstuvwxyz"

var ALPHABET_ARR []string = strings.Split(ALPHABET_STRING, "")

func main() {
	fmt.Println(GenerateCutter("Yasser", "ali"))
}
func LoadFromJSON() map[string]int {

	jsonFile, jsonReadErr := os.Open("data/data.json")

	if jsonReadErr != nil {
		fmt.Println(jsonReadErr.Error())
	}
	defer jsonFile.Close()
	jsonByte, _ := ioutil.ReadAll(jsonFile)

	cutters := map[string]int{}
	json.Unmarshal(jsonByte, &cutters)
	fmt.Println(cutters)
	return cutters
}

func GenerateCutter(firstname string, lastname string) string {

	firstname = strings.ToLower(firstname)
	lastname = strings.ToLower(lastname)
	firstnameInitialChar := firstname[0:1]
	lastnameInitialChar := lastname[0:1]
	endIndex := strings.Index(ALPHABET_STRING, firstnameInitialChar)

	concatenatedAlphabet := ALPHABET_ARR[0 : endIndex+1]
	alphabetLength := len(concatenatedAlphabet)
	for alphabetLength > 0 {
		alphabetLength--
		letter := concatenatedAlphabet[alphabetLength]
		var key string = fmt.Sprintf("%s, %s.", strings.Title(lastname), strings.Title(letter))
		number := data.CUTTERS_TABLE[key]
		if number != 0 {
			return fmt.Sprint(strings.Title(lastnameInitialChar), number)
		}

	}

	var key string = strings.Title(lastname)
	for len(key) != 0 {
		fmt.Println(key)
		number := data.CUTTERS_TABLE[key]
		if number != 0 {
			return fmt.Sprint(strings.Title(lastnameInitialChar), number)
		}
		key = key[0 : len(key)-1]
	}

	return ""

}
