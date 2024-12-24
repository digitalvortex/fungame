# Save The Planet Game

## Overview

The "Save The Planet" game is a fun and interactive coding challenge where users are tasked with completing a coding problem within a specified deadline. The goal is to reverse a string and return the correct output before time runs out. This project is implemented in Go and includes a set of tests to ensure functionality.

## Features

- **Task Creation**: Create a new task with a name and a deadline.
- **Task Completion**: Attempt to complete the task by providing the correct output.
- **Deadline Enforcement**: The task must be completed within the specified time limit.
- **Error Handling**: Provides feedback for incorrect answers or timeouts.

## Getting Started

### Prerequisites

- Go (version 1.16 or higher) installed on your machine.
- A terminal or command prompt.

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/savetheplanet.git
   cd savetheplanet
   ```

2. Install dependencies (if any):
   ```bash
   go mod tidy
   ```

### Running the Game

To run the game, execute the following command in your terminal:
```bash
go run cmd/main.go
```

### User Input

During the game, you will be prompted to complete a coding challenge. The challenge will ask you to reverse a specific string. You will need to type the exact reversed string as your answer. For example:

- **Challenge**: Write a function in Go to reverse the string `'save_the_planet'` and return it.
- **Example Input**: `save_the_planet`
- **Expected Output**: `tenalp_eht_evas`

Make sure to provide your answer within the specified time limit to successfully complete the task.

### Running Tests

To ensure everything is working correctly, you can run the tests included in the project. Use the following command:

```bash
go test ./internal/planet
```

This will execute all the test cases defined in `planet_code_test.go` and report any failures.

## Code Structure

- `cmd/main.go`: Entry point for the application.
- `internal/planet/planet_code.go`: Contains the main logic for the game, including the `SaveThePlanet` struct and methods.
- `internal/planet/planet_code_test.go`: Contains unit tests for the game logic.

## Contributing

Contributions are welcome! If you have suggestions for improvements or new features, feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Inspired by coding challenges and interactive learning experiences.
