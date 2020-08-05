package actor

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
)

type Watch struct {
	Id       Id
	Filename string
}

func NewWatch(id Id, filename string) *Watch {
	return &Watch{Id: id, Filename: filename}
}

type Watcher struct {
	watcher *fsnotify.Watcher
	watches map[string]*Watch
	running bool
	quit    chan bool
	changes chan *Watch
}

func NewWatcher() (*Watcher, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, fmt.Errorf("error while creating fsnotify watcher: %s", err)
	}

	return &Watcher{
		watcher: watcher,
		watches: map[string]*Watch{},
		quit:    make(chan bool),
		changes: make(chan *Watch, 10),
	}, nil
}

func (w *Watcher) Watch(id Id, filename string) error {
	if _, ok := w.watches[filename]; ok {
		return fmt.Errorf("already watching file %s", filename)
	}

	err := w.watcher.Add(filename)
	if err != nil {
		return err
	}
	w.watches[filename] = NewWatch(id, filename)

	return nil
}

func (w *Watcher) Unwatch(filename string) error {
	if _, ok := w.watches[filename]; !ok {
		return fmt.Errorf("file %s is not being watched", filename)
	}

	err := w.watcher.Remove(filename)
	if err != nil {
		return err
	}

	delete(w.watches, filename)

	return nil
}

func (w *Watcher) Start() {
	if w.running {
		return
	}
	w.running = true

	go func() {
		defer func() {
			w.running = false
		}()

		for {
			select {
			case event, ok := <-w.watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					fmt.Printf("actor file modified: %s\n", event.Name)
					w.changes <- w.watches[event.Name]
				}
			case err, ok := <-w.watcher.Errors:
				if !ok {
					return
				}
				fmt.Printf("actor watcher error: %s\n", err)
			case <-w.quit:
				fmt.Printf("stopping actor watcher\n")
				return
			}
		}
	}()
}

func (w *Watcher) Stop() {
	if w.running {
		w.quit <- true
	}
}

func (w *Watcher) GetChangedActors() []*Watch {
	var list []*Watch
	for len(w.changes) > 0 {
		list = append(list, <-w.changes)
	}

	return list
}
