package config

import (
	"errors"
	"github.com/gabe565/matrimony/internal/config/models"
	flag "github.com/spf13/pflag"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"sync"
)

var (
	Path     string
	Recreate bool
	Watch    bool
	mu       sync.Mutex
)

var Config *models.Config

func init() {
	flag.StringVarP(&Path, "config", "c", "./matrimony.yaml", "YAML configuration path")
	flag.BoolVar(&Recreate, "recreate-config", false, "Force create example config file")
	flag.BoolVar(&Watch, "watch", false, "Watch config file for updates")
}

func Setup() (err error) {
	if Watch {
		defer func() {
			go Watcher()
		}()
	}

	if Recreate {
		log.Println("Recreating config")
		Recreate = false
		return createConfig(Path)
	}

	if _, err = os.Stat(Path); err == nil {
		if err = readConfig(Path); err != nil {
			return err
		}
	} else if errors.Is(err, os.ErrNotExist) {
		return createConfig(Path)
	}
	return err
}

func CreateConfig() error {
	return createConfig(Path)
}

func createConfig(path string) error {
	conf := ExampleConfig()
	err := saveConfig(conf, path)
	if err != nil {
		return err
	}
	Config = conf
	return nil
}

func ReadConfig() error {
	return readConfig(Path)
}

func readConfig(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	decoder := yaml.NewDecoder(f)
	conf := &models.Config{}
	err = decoder.Decode(conf)
	if err != nil {
		return err
	}
	Config = conf

	return nil
}

func SaveConfig() error {
	return saveConfig(Config, Path)
}

func saveConfig(conf *models.Config, path string) error {
	mu.Lock()
	defer mu.Unlock()

	out, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer func(out *os.File) {
		_ = out.Close()
	}(out)

	encoder := yaml.NewEncoder(out)
	err = encoder.Encode(conf)
	if err != nil {
		return err
	}

	err = out.Close()
	if err != nil {
		return err
	}

	return nil
}
