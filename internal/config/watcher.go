package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/gabe565/matrimony/internal/datadir"
	"log"
)

func Watcher() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalln(err)
	}

	defer func(watcher *fsnotify.Watcher) {
		_ = watcher.Close()
	}(watcher)

	if err = watcher.Add(datadir.DataDir); err != nil {
		log.Fatalln(err)
	}

	log.Println("Watching for config changes")

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			switch {
			case event.Op&fsnotify.Write == fsnotify.Write,
				event.Op&fsnotify.Create == fsnotify.Create:
				{
					mu.Lock()
					log.Println("Reloading config")
					err = ReadConfig()
					if err != nil {
						log.Printf("could not reload config: %v", err)
					}
					mu.Unlock()
				}
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Fatalln(err)
		}
	}
}
