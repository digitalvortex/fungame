// internal/planet/planet_code.go
package planet

import (
	"fmt"
	"time"
)

// PlanetSaver is an interface that defines the methods for saving the planet.
type PlanetSaver interface {
	CompleteTask(input string) error
}

// SaveThePlanet implements the PlanetSaver interface.
// It represents the task of saving the world with a deadline.
type SaveThePlanet struct {
	TaskName   string
	Deadline   time.Time
	Duration   time.Duration
	TaskStatus bool
}

// NewSaveThePlanet is a constructor function to create a new SaveThePlanet object.
func NewSaveThePlanet(taskName string, duration time.Duration) PlanetSaver {
	return &SaveThePlanet{
		TaskName:   taskName,
		Deadline:   time.Now().Add(duration),
		Duration:   duration,
		TaskStatus: false,
	}
}

// CompleteTask allows the user to attempt to complete an impossible coding task.
// Returns an error if the task is not completed in time.
func (s *SaveThePlanet) CompleteTask(input string) error {
	start := time.Now() // Record the start time of the task

	// Print the task instructions
	fmt.Printf("🚨 ALERT: The world is in grave danger! 🚨\n")
	fmt.Printf("Your mission, should you choose to accept it: '%s'\n", s.TaskName)
	fmt.Printf("You have %v seconds to complete this task or the planet goes BOOM 💥!\n", s.Duration.Seconds())
	fmt.Println("To save the planet, write a Go function to reverse a string and type the exact solution!")

	// Display the coding challenge
	fmt.Println("Challenge: Write a function in Go to reverse the string 'save_the_planet' and return it.")
	fmt.Println("Example input: 'save_the_planet'")
	fmt.Println("Expected output: 'tenalp_eht_evas'")
	fmt.Println("Quick! Write the function, compile it, and provide the output:")

	// debug
	fmt.Println("start ->", start)
	fmt.Println("deadline ->", s.Deadline)
	fmt.Println("duration ->", s.Duration)
	fmt.Println("Time since start ->", time.Since(start))

	// Check if the task was completed within the deadline
	if time.Since(start) > s.Duration {
		return fmt.Errorf("🤦 You took too long! The planet goes BOOM 💥")
	}

	if time.Now().After(s.Deadline) {
		return fmt.Errorf("task '%s' failed due to timeout", s.TaskName)
	}

	// Impossible task: User must provide the exact reversed string
	if input == "tenalp_eht_evas" {
		if time.Since(start) <= s.Duration {
			s.TaskStatus = true
			fmt.Println("🌍🎉 Unbelievable! You actually did it! The planet is safe... for now.")
			return nil
		} else {
			return fmt.Errorf("🤦 You took too long! The planet goes BOOM 💥")
		}
	}

	// User failed to complete the coding task or provided incorrect output
	return fmt.Errorf("🤦 Incorrect or incomplete! You just doomed us all. BOOM 💥")
}
