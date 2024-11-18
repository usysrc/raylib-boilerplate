# Default target: run the game
.PHONY: all
all: run 

# Run the game
.PHONY: run
run: 
	@echo "Running the game..."
	go run main.go
