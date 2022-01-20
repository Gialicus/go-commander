package util

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func JsonStringify(m map[string]interface{}) (string, error) {
	result, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(result), nil
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
