package config

import (
	"errors"
	"github.com/gabe565/matrimony/internal/config/models"
	flag "github.com/spf13/pflag"
	"gopkg.in/yaml.v3"
	"io/fs"
	"log"
	"os"
	"path"
	"sync"
)

const DataDirDefault = "data"

var (
	DataDir  string
	Filename string
	Recreate bool
	Watch    bool
	mu       sync.Mutex
)
var Config *models.Config

func init() {
	flag.StringVarP(&DataDir, "data", "d", DataDirDefault, "Data directory")
	flag.StringVarP(&Filename, "config", "c", DataDirDefault+"/matrimony.yaml", "Config filename")
	flag.BoolVar(&Recreate, "recreate-config", false, "Force create example config file")
	flag.BoolVar(&Watch, "watch", true, "Watch config file for updates")
}

func UpdateFilename() {
	if DataDir != DataDirDefault && Filename[0:len(DataDirDefault)] == DataDirDefault {
		Filename = path.Join(DataDir, Filename[len(DataDirDefault):])
	}
}

func Setup() (err error) {
	UpdateFilename()

	if Watch {
		defer func() {
			go Watcher()
		}()
	}

	if Recreate {
		log.Println("Recreating config")
		Recreate = false
		return createConfig(Filename)
	}

	if _, err = os.Stat(Filename); err == nil {
		if err = readConfig(Filename); err != nil {
			return err
		}
	} else if errors.Is(err, os.ErrNotExist) {
		return createConfig(Filename)
	}
	return err
}

func CreateConfig() error {
	return createConfig(Filename)
}

func createConfig(f string) error {
	conf := ExampleConfig()
	err := saveConfig(conf, f)
	if err != nil {
		return err
	}
	Config = conf
	return nil
}

func ReadConfig() error {
	return readConfig(Filename)
}

func readConfig(confpath string) error {
	f, err := os.Open(confpath)
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
	return saveConfig(Config, Filename)
}

func saveConfig(conf *models.Config, confpath string) (err error) {
	dir := path.Dir(confpath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.Mkdir(dir, 0755)
		if err != nil {
			return err
		}
	}

	mu.Lock()
	defer mu.Unlock()

	out, err := os.OpenFile(confpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
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

func DataFS() fs.FS {
	return os.DirFS(path.Join(DataDir, "public"))
}
