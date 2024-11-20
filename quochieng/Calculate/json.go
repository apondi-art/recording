package Calculate

import (
	"encoding/json"
	"fmt"
	"os"
)

var Record = make(map[string]*Status)

const Jsonfile = "record.json"

func Save() {
	file, err := os.Create(Jsonfile)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close() // Ensure the file is closed after writing

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(Record); err != nil {
		fmt.Println("Error encoding JSON to file:", err)
	}
}


func Read() {
	file, err := os.Open(Jsonfile)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("JSON file doesn't exist. Starting with an empty record.")
			return
		}
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() // Ensure the file is closed after reading

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&Record); err != nil {
		fmt.Println("Error decoding JSON file:", err)
	}
}


