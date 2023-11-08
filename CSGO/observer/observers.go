package observer

import "fmt"

type ConsoleLogger struct{}

func (cl *ConsoleLogger) Update(event string) {
	fmt.Println("Console Log: " + event)
}
