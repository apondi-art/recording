package Calculate

import (
	"encoding/json"
	"fmt"
	"os"
)

var Record = make(map[string]*Status)

const jsonfile = "record.json"

func Read() {
	file, err := os.Open(jsonfile)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println()
		}
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	decoder.Decode(&Record)
}

func Save() {
	file, err := os.Create("jsonfile")
	if err != nil {
		fmt.Println(err)
	}
	encoder := json.NewEncoder(file)
	encoder.Encode(&Record)
}
