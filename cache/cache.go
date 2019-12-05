package cache

import (
	"log"
	"time"
)

type Entry struct {
	Text string
	Time time.Time
}

type Cache map[string]*Entry

const MaxEntryAge = 30 * time.Minute

var instance = make(Cache)

func GetOrSet(key string, retrieve func() string) string {
	log.Printf("cache, key: %s", key)

	entry, ok := instance[key]
	if ok {
		log.Printf("Found! Returning contents ...")
		log.Printf("(%s ago)", time.Since(entry.Time).String())
		return entry.Text
	} else {
		log.Printf("Not found, adding to cache.")
		entry = &Entry{
			Text: retrieve(),
			Time: time.Now(),
		}
		instance[key] = entry
		return entry.Text
	}
}

func deleteOldEntries() {
	log.Println("Deleting old entries...")

	for key, entry := range instance {
		if time.Since(entry.Time) > MaxEntryAge {
			log.Printf("Key \"%s\" too old, deleting", key)
			delete(instance, key)
		}
	}
}

func DeletionLoop() {
	for {
		deleteOldEntries()

		time.Sleep(5 * time.Second)
	}
}