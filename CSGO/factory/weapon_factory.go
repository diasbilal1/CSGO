package factory

import "CSGO/strategy"

type WeaponFactory struct{}

func (wf WeaponFactory) CreateWeapon(weaponType string) strategy.WeaponStrategy {
	switch weaponType {
	case "AK-47":
		return &strategy.AK47{}
	case "M4A1-S":
		return &strategy.M4A1S{}
	default:
		return nil
	}
}
