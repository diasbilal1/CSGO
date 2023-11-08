// strategy/equipment_strategy.go
package strategy

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type EquipmentStrategy interface {
	Execute()
}

type Bomb struct {
	babah   bool
	Code    int
	Seconds time.Duration
	timer   *time.Timer
}

func (b *Bomb) Execute() {
	fmt.Println("Planting the bomb")

	var code string
LOOP1:
	for {
		fmt.Print("Input code for bomb (10 digit): ")
		fmt.Scan(&code)
		if len(code) != 10 {
			fmt.Println("Code must be 10 digit number.")
			continue
		}
		codeInt, err := strconv.Atoi(strings.Trim(code, " "))
		if err != nil {
			continue
		}
		b.Code = codeInt
		break LOOP1
	}
	timer := time.NewTimer(b.Seconds * time.Second)
	b.timer = timer

	go func(b *Bomb) {
		for {
			select {
			case _ = <-timer.C:
				b.babah = true
				fmt.Println("Bomb was exploaded(vzarvalos)(babah koroche).\nTerrorist WON!!!")
			}
		}
	}(b)

}

type DefuseKit struct {
	Bomb  *Bomb
	Exist bool
}

func (e *DefuseKit) Execute() {
	if e.Exist {
		fmt.Println("Using the defuse kit")
	} else {
		fmt.Println("Without defuse kit")

	}
	bombCode := e.Bomb.Code
	fmt.Println("BombCode", bombCode)
	var code string

	for {
		fmt.Printf("Input code:\t %v\n", bombCode)
		fmt.Scan(&code)
		codeInt, err := strconv.Atoi(code)
		if err != nil {
			continue
		}
		if codeInt == bombCode {
			if !e.Bomb.babah {
				e.Bomb.timer.Stop()
				fmt.Println("Bomb was defused\nCounter-Terrorist WON!!!")
				return
			} else {
				fmt.Println("Bomb was bababahbhashdhsaklfhio asncirwnurioasnfo jsakoaj\nTerrorist WON!!!")
				return
			}
		} else if e.Bomb.babah {
			fmt.Println("Bomb was bababahbhashdhsaklfhio asncirwnurioasnfo jsakoaj YOu lose the game\nTerrorist WON!!!")
			return
		}
	}

}
