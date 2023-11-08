package factory

import "CSGO/strategy"

type EquipmentFactory struct{}

var bomb *strategy.Bomb

func (ef EquipmentFactory) CreateEquipment(equipmentType string) strategy.EquipmentStrategy {
	switch equipmentType {
	case "Bomb":
		bomb = &strategy.Bomb{Seconds: 10}
		return bomb
	case "Defuse Kit":
		bomb.Seconds = 20
		return &strategy.DefuseKit{Bomb: bomb, Exist: true}
	case "No Defuse Kit":
		bomb.Seconds = 10
		return &strategy.DefuseKit{Bomb: bomb, Exist: true}
	default:
		return nil
	}
}
