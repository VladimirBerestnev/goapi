package files

import (
	"encoding/json"
	"errors"
	"go/rest/internal/app/models"
	"log"
	"os"
)

func LoadFromFile() {

	data := OpenYaml()

	_, error := os.Stat(data["filename"])
	isFileExist := !errors.Is(error, os.ErrNotExist)

	if isFileExist {
		data, err := os.ReadFile(data["filename"])
		if err != nil {
			log.Fatal(err)
		}
		if err = json.Unmarshal(data, &models.Tasklist); err != nil {
			log.Fatal(err)
		}
	}
}
