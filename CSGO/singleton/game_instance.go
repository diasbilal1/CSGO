package singleton

import (
	"sync"
)

type GameInstance struct {
	Players []string
}

var mu = &sync.Mutex{}

var instance *GameInstance

func GetGameInstance() *GameInstance {
	if instance == nil {
		mu.Lock()
		defer mu.Unlock()
		if instance == nil {
			instance = &GameInstance{}
		}
	}

	return instance
}
