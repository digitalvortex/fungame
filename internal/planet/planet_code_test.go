// internal/planet/planet_code_test.go
package planet

import (
	"testing"
	"time"
)

// TestCorrectAnswerInTime verifies that the task succeeds when the correct answer is provided within the deadline.
func TestCorrectAnswerInTime(t *testing.T) {
	task := NewSaveThePlanet("Test Task", 5*time.Second) // 5 seconds to complete

	// Simulate completing the task correctly within the deadline
	err := task.CompleteTask("tenalp_eht_evas") // Directly provide the correct answer
	if err != nil {
		t.Fatalf("Expected task '%s' to be completed successfully, but it failed: %s", task.(*SaveThePlanet).TaskName, err)
	}
}

// TestWrongAnswer ensures that the task fails when the user provides the wrong answer.
func TestWrongAnswer(t *testing.T) {
	task := NewSaveThePlanet("Test Task", 5*time.Second)

	// Simulate the wrong answer
	err := task.CompleteTask("wrong_answer") // Provide an incorrect answer
	if err == nil {
		t.Fatalf("Expected task '%s' to fail with wrong answer, but it succeeded", task.(*SaveThePlanet).TaskName)
	}
}

// TestTimeout verifies that the task fails if the user takes too long to complete it.
func TestTimeout(t *testing.T) {
	task := NewSaveThePlanet("Test Task", 1*time.Second) // Very short deadline

	// Simulate user response coming after the deadline
	time.Sleep(2 * time.Second)                 // Delay beyond the deadline
	err := task.CompleteTask("tenalp_eht_evas") // Attempt to complete the task
	if err == nil {
		t.Fatalf("Expected task '%s' to fail due to timeout, but it succeeded", task.(*SaveThePlanet).TaskName)
	}
}

// TestEdgeCase verifies that the task fails if the answer is provided just after the deadline.
func TestEdgeCase(t *testing.T) {
	task := NewSaveThePlanet("Test Task", 2*time.Second) // Tight deadline

	// Simulate user response coming immediately after the deadline
	time.Sleep(6 * time.Second)                 // Respond after deadline
	err := task.CompleteTask("tenalp_eht_evas") // Simulate task completion

	if err == nil {
		t.Fatalf("Expected task '%s' to fail as it was completed after the deadline, but it succeeded", task.(*SaveThePlanet).TaskName)
	}
}
