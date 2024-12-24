package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/digitalvortex/fungame/internal/planet"
)

func main() {
	// Set up a channel to listen for interrupt signals (Ctrl+C)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		// Create a new task
		task := planet.NewSaveThePlanet("Save the Planet", 5*time.Second)

		// Start the task
		fmt.Println("🚨 ALERT: The world is in grave danger! 🚨")
		fmt.Printf("You have %v seconds to complete this task or the planet goes BOOM 💥!\n", task.(*planet.SaveThePlanet).Duration.Seconds())
		fmt.Println("To save the planet, write a Go function to reverse a string and type the exact solution!")
		fmt.Println("Challenge: Write a function in Go to reverse the string 'save_the_planet' and return it.")
		fmt.Println("Example input: 'save_the_planet'")
		fmt.Println("Expected output: 'tenalp_eht_evas'")
		fmt.Println("Quick! Write the function, compile it, and provide the output:")

		// Capture user input
		var input string
		done := make(chan bool)

		// Start a goroutine to capture user input
		go func() {
			fmt.Scanln(&input)
			done <- true
		}()

		// Wait for user input or timeout
		select {
		case <-done:
			// User provided input
			err := task.CompleteTask(input) // Pass the user input to CompleteTask
			if err != nil {
				fmt.Println(err)
			}
		case <-time.After(task.(*planet.SaveThePlanet).Duration):
			fmt.Println("⏰ Time's up! You couldn't complete the task in time, and now it's GAME OVER. BOOM 💥")
		}

		// Prompt for exit
		fmt.Println("Type 'exit' to quit or press Enter to continue...")
		var exitInput string
		fmt.Scanln(&exitInput)

		// Check for exit command
		if exitInput == "exit" {
			fmt.Println("Exiting the game. Goodbye!")
			break
		}

		// Check for interrupt signal
		select {
		case <-sigChan:
			fmt.Println("\nExiting the game. Goodbye!")
			return
		default:
			// Continue the loop
		}
	}
}
