package main

import (
	"github.com/fsnotify/fsnotify"
	"log"
)

func watcherThread(w *fsnotify.Watcher) {
	for {
		select {
		case ev := <-w.Events:
			{
				if ev.Op&fsnotify.Write == fsnotify.Write {
					log.Printf("Configure write:%s", ev.Name)
				}

				if ev.Op&fsnotify.Create == fsnotify.Create {

					log.Printf("Configure create:%s", ev.Name)
				}

				if ev.Op&fsnotify.Rename == fsnotify.Rename {
					log.Printf("Configure rename:%s", ev.Name)
				}

				if ev.Op&fsnotify.Remove == fsnotify.Remove {
					log.Printf("Configure delete:%s", ev.Name)
				}
			}
		case err := <-w.Errors:
			{
				log.Printf("watcherThread Error:%s", err.Error())
			}
		}
	}
}

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	go watcherThread(watcher)

	log.Printf("add to watcher 1st: %s\n", "/waf/config/misc")
	err = watcher.Add("/waf/config/misc")
	if err != nil {
		log.Fatal(err)
	}

	// to keep waiting forever, to prevent main exit
	// this is to replace the done channel
	select {}
}
