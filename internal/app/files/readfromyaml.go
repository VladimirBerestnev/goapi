package files

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func OpenYaml() map[string]string {

	yfile, err := os.ReadFile("config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	data := make(map[string]string)
	err2 := yaml.Unmarshal(yfile, &data)

	if err2 != nil {
		log.Fatal(err2)
	}

	return data
}
