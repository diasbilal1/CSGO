package main

import (
	"CSGO/factory"
	"CSGO/observer"
	"CSGO/singleton"
	"CSGO/strategy"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	game := singleton.GetGameInstance()
	game.Players = []string{"Terrorist", "CounterTerrorist"}

	playerFactory := factory.PlayerFactory{}
	weaponFactory := factory.WeaponFactory{}
	equipmentFactory := factory.EquipmentFactory{}
	eventNotifier := observer.EventNotifier{}
	consoleLogger := observer.ConsoleLogger{}
	eventNotifier.RegisterObserver(&consoleLogger)

	var ak74 strategy.WeaponStrategy
	var m4a1s strategy.WeaponStrategy
	var bomb strategy.EquipmentStrategy
	var defuseKit strategy.EquipmentStrategy
	var tCanShoot = false
	var ctCanShoot = false
	fmt.Println("Welcome to the Bomb Defusal Game!")

	for _, playerName := range game.Players {
		var side int
		var player strategy.PlayerStrategy

		if playerName == "Terrorist" {
			player = playerFactory.CreatePlayer(playerName)
			bomb = equipmentFactory.CreateEquipment("Bomb")
			fmt.Println("Buy weapon for Terrorist:\n  Yes (1)\n  No (2)")
			fmt.Print("Your input: ")
			fmt.Scan(&side)
			if side == 1 {
				ak74 = weaponFactory.CreateWeapon("AK-47")
				tCanShoot = true
			}
			eventNotifier.Notify(playerName + " performed actions!")
		} else {
			player = playerFactory.CreatePlayer(playerName)
			fmt.Println("Buy weapon for CounterTerrorist:\n  Yes (1)\n  No (2)")
			fmt.Print("Your input: ")
			fmt.Scan(&side)
			if side == 1 {
				m4a1s = weaponFactory.CreateWeapon("M4A1-S")
				ctCanShoot = true
			}

			fmt.Println("Buy defuse kits:\n  Yes (1)\n  No (2)")
			fmt.Print("Your input: ")
			fmt.Scan(&side)
			if side == 1 {
				defuseKit = equipmentFactory.CreateEquipment("Defuse Kit")
			} else {
				defuseKit = equipmentFactory.CreateEquipment("No Defuse Kit")
			}

			eventNotifier.Notify(playerName + " performed actions!")

			if tCanShoot == true {
				ak74.Execute()
				rand.Seed(time.Now().UnixNano())
				randomNumber1 := rand.Intn(3) + 1
				fmt.Print("Write number between 1 and 3(to shoot the Counter-Terrorist: ")
				fmt.Scan(&side)
				fmt.Printf("Random number is: %v\n", randomNumber1)
				if randomNumber1 == side {
					fmt.Println("Terrorist killed the Counter-Terrorist")
					fmt.Println("Terrorists won!")
					break
				} else {
					fmt.Println("Terrorist shooted, but missed")
				}

				eventNotifier.Notify("Terrorist performed actions!")
			}

			if ctCanShoot == true {
				m4a1s.Execute()
				rand.Seed(time.Now().UnixNano())
				randomNumber3 := rand.Intn(3) + 1
				fmt.Print("Write numeber between 1 and 3(to shoot the Terrorist: ")
				fmt.Scan(&side)
				fmt.Printf("Random number is: %v\n", randomNumber3)
				if randomNumber3 == side {
					fmt.Printf("%v killed the Terrorist\n", playerName)
					fmt.Printf("%v won!", player.GetTeam())
					break
				} else {
					fmt.Println("Counter-Terrorist shooted, but missed")
				}
				eventNotifier.Notify(playerName + " performed actions!")
			}
			bomb.Execute()
			eventNotifier.Notify("Terrorist planted bomb!")
			defuseKit.Execute()
			eventNotifier.Notify("Counter-Terrorist performed actions!")

		}
	}
}
