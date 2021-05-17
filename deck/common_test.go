package deck

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func containsString(slice []string, e string) bool {
	if len(slice) == 0 {
		return false
	}
	for _, se := range slice {
		if e == se {
			return true
		}
	}
	return false
}

func areAllStringsUnique(slice []string) (bool, string) {
	set := make(map[string]int)
	for _, e := range slice {
		if _, ok := set[e]; ok {
			return false, e
		} else {
			set[e] = 0
		}
	}
	return true, ""
}

func set_env_map() map[string]string {
	env_map, err := godotenv.Read("../test.env")
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}
	log.Println(env_map)
	return env_map
}

func removeFile(fileName string) {
	err := os.Remove(fileName)
	if err != nil {
		log.Fatal("Error removing the file: ", err)
	}
}
