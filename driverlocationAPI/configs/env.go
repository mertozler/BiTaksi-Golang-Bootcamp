package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"path"
	"runtime"
)

func EnvMongoURI() string {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
	err1 := os.Chdir(dir)
	if err1 != nil {
		log.Fatal("Error loading env file")
	}
	err := godotenv.Load(dir + "\\.env")
	if err != nil {
		log.Fatal("Error loading env file")
	}
	return os.Getenv("MONGOURI")

}
