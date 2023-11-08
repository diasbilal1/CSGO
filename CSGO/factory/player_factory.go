package factory

import "CSGO/strategy"

type PlayerFactory struct{}

func (pf PlayerFactory) CreatePlayer(name string) strategy.PlayerStrategy {
	switch name {
	case "Terrorist":
		return &strategy.Terrorist{}
	case "CounterTerrorist":
		return &strategy.CounterTerrorist{}
	default:
		return nil
	}

}
