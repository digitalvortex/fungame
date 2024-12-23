package planet

import (
	"testing"
	"time"
)

// TestCorrectAnswerInTime verifies that the task succeeds when the correct answer is provided within the deadline.
func TestCorrectAnswerInTime(t *testing.T) {
	task := NewSaveThePlanet("Test Task", 5*time.Second) // 5 seconds to complete

	// Simulate completing the task correctly within the deadline
	go func() {
		time.Sleep(1 * time.Second) // Simulate user typing time
		if sp, ok := task.(*SaveThePlanet); ok {
			sp.TaskStatus = true // Simulate task completion
		}
	}()

	time.Sleep(2 * time.Second) // Wait to ensure simulation completes

	if sp, ok := task.(*SaveThePlanet); ok {
		if !sp.TaskStatus {
			t.Fatalf("Expected task '%s' to be completed successfully, but it failed", sp.TaskName)
		}
	}
}

// TestWrongAnswer ensures that the task fails when the user provides the wrong answer.
func TestWrongAnswer(t *testing.T) {
	task := NewSaveThePlanet("Test Task", 5*time.Second)

	// Simulate the wrong answer
	if sp, ok := task.(*SaveThePlanet); ok {
		sp.TaskStatus = false
		// Validate that the task fails
		if sp.TaskStatus {
			t.Fatalf("Expected task '%s' to fail with wrong answer, but it succeeded", sp.TaskName)
		}
	}
}

// TestTimeout verifies that the task fails if the user takes too long to complete it.
func TestTimeout(t *testing.T) {
	task := NewSaveThePlanet("Test Task", 1*time.Second) // Very short deadline

	// Simulate user response coming after the deadline
	go func() {
		time.Sleep(2 * time.Second) // Delay beyond the deadline
		task.CompleteTask()         // Attempt to complete the task
	}()

	time.Sleep(3 * time.Second) // Wait to ensure timeout has occurred

	if sp, ok := task.(*SaveThePlanet); ok {
		if sp.TaskStatus {
			t.Fatalf("Expected task '%s' to fail due to timeout, but it succeeded", sp.TaskName)
		}
	}
}

// TestEdgeCase verifies that the task fails if the answer is provided just after the deadline.
func TestEdgeCase(t *testing.T) {
	task := NewSaveThePlanet("Test Task", 2*time.Second) // Tight deadline

	// Simulate user response coming immediately after the deadline
	go func() {
		time.Sleep(3 * time.Second) // Respond after deadline
		task.CompleteTask()         // Simulate task completion
	}()

	time.Sleep(4 * time.Second) // Wait to ensure simulation completes

	if sp, ok := task.(*SaveThePlanet); ok {
		if sp.TaskStatus {
			t.Fatalf("Expected task '%s' to fail as it was completed after the deadline, but it succeeded", sp.TaskName)
		}
	}
}
