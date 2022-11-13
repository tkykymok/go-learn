package configs

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

var lock = &sync.Mutex{}

var fieldNamesInstance map[string]string

func GetFieldNames() map[string]string {
	if fieldNamesInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		file, _ := os.Open("/app/configs/field-names.json")
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				log.Fatal(err)
			}
		}(file)

		decoder := json.NewDecoder(file)
		if err := decoder.Decode(&fieldNamesInstance); err != nil {
			log.Fatal(err)
		}
	}
	return fieldNamesInstance
}
