package strategy

import "fmt"

type WeaponStrategy interface {
	Execute()
}

type AK47 struct {
	// Implementation for AK47
}

func (w *AK47) Execute() {
	fmt.Println("Firing AK-47")
}

type M4A1S struct {
	// Implementation for M4A1S
}

func (w *M4A1S) Execute() {
	fmt.Println("Firing M4A1-S")
}
