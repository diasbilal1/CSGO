package strategy

import "fmt"

type PlayerStrategy interface {
	Execute()
	GetTeam() string
}

type Terrorist struct {
	Name string
}

func (p *Terrorist) Execute() {
	fmt.Println(p.Name + " is performing actions.")
}

func (p *Terrorist) GetTeam() string {
	return "Your team is Terrorist!"
}

type CounterTerrorist struct {
	Name string
}

func (p *CounterTerrorist) Execute() {
	fmt.Println(p.Name + " is performing actions.")
}

func (p *CounterTerrorist) GetTeam() string {
	return "Your team is CounterTerrorist!"
}
