package main

import (
	"fmt"
	"time"

	"github.com/digitalvortex/fungame/internal/planet"
)

func main() {
	// Use the constructor to create an impossible coding task
	task := planet.NewSaveThePlanet("Write a Go Function Under Pressure", 5*time.Second)

	// Attempt to complete the task
	err := task.CompleteTask()
	if err != nil {
		// Print humorous error messages if the task fails
		fmt.Printf("Mission Failed: %s\n", err.Error())
		fmt.Println("💀 You let the planet down. Maybe try coding without a ticking clock next time.")
	} else {
		fmt.Println("🎆 You're a coding wizard! The world celebrates your quick fingers and genius! 🦸‍♂️")
	}
}
