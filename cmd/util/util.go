package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func JsonStringify(m map[string]interface{}) (string, error) {
	result, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(result), nil
}
func ReadJsonFile(path string) ([]byte, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	return byteValue, nil
}
func CreateFile(fileName string, value string) {
	f, err := os.Create(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(value)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("done")
}
func StripFileType(path string) string {
	if strings.Contains(path, ".") {
		splitted := strings.Split(path, ".")
		fileType := "." + splitted[len(splitted)-1]
		return strings.Replace(path, fileType, "", -1)
	}
	return path
}
