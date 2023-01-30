package conf

import (
	"github.com/shagnyin/gin-scaffold/internal/config"
	"gopkg.in/yaml.v3"
	"io"
	"os"
)

func MustLoad(file string, cfg *config.Config) {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	bytes, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(bytes, cfg)
	if err != nil {
		panic(err)
	}
}
