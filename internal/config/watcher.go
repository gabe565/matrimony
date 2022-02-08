package config

import (
	"github.com/fsnotify/fsnotify"
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

	if err = watcher.Add(Filename); err != nil {
		log.Fatalln(err)
	}

	log.Println("Watching for config changes")

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Op&fsnotify.Write == fsnotify.Write {
				mu.Lock()
				log.Println("Reloading config")
				err = ReadConfig()
				mu.Unlock()
				if err != nil {
					log.Printf("could not reload config: %v", err)
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
