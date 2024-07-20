package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TODOs struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}

func main() {
	// initialize a counter to keep track of TODO item IDs
	counter := 0

	// create a slice to store TODO items
	data := make([]TODOs, 0)

	// create a new scanner to read input from keyboard
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Commands: show create remove done clear")
	fmt.Print("> ")

	// continuously read input from user and handle the command
	for scanner.Scan() {
		line := scanner.Text()
		handleInput(line, &data, &counter)
		fmt.Print("> ")
	}
}

// process the user's command and modifies the TODO list accordingly
func handleInput(input string, data *[]TODOs, counter *int) {
	split := strings.Split(input, " ")
	cmd := split[0]

	switch cmd {
	case "create":
		createTODO(strings.Join(split[1:], " "), data, counter)
	case "remove":
		// ensure an ID is provided for the remove command
		if len(split) < 2 {
			fmt.Println("Error: Missing ID for remove command")
			return
		}
		removeTODO(split[1], data)
	case "show":
		// display all TODO items
		showTODOs(data)
	case "done":
		// ensure an ID is provided for the done command
		if len(split) < 2 {
			fmt.Println("Error: Missing ID for done command")
			return
		}
		// mark the TODO item with the specified ID as done
		markTODOAsDone(split[1], data)
	case "clear":
		clearTODOs(data)
	default:
		// inform the user if the command is invalid
		fmt.Println("Invalid command")
	}
}

func createTODO(description string, data *[]TODOs, counter *int) {
	*data = append(*data, TODOs{
		Id:          *counter,
		Description: description,
		Status:      false,
	})
	*counter++ 
	fmt.Println("Created TODO item")
}

func removeTODO(idStr string, data *[]TODOs) {
	id, err := strconv.Atoi(idStr) // convert the ID string to an integer
	if err != nil {
		fmt.Println("Error: Invalid ID")
		return
	}
	for i, todo := range *data {
		if todo.Id == id {
			*data = append((*data)[:i], (*data)[i+1:]...)
			fmt.Println("Removed TODO item")
			return
		}
	}
	fmt.Println("TODO item not found")
}

// showTODOs displays all TODO items in the list
func showTODOs(data *[]TODOs) {
	if len(*data) == 0 {
		fmt.Println("No TODO items") // inform the user if there are no TODO items
		return
	}
	for _, todo := range *data {
		fmt.Println("------")
		statusLine := "[ ]"
		if todo.Status {
			statusLine = "[x]"
		}
		fmt.Printf("%s %s ID: %d\n", statusLine, todo.Description, todo.Id)
	}
}

// set the status of the TODO item with the specified ID to done
func markTODOAsDone(idStr string, data *[]TODOs) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Error: Invalid ID")
		return
	}
	for i, todo := range *data {
		if todo.Id == id {
			(*data)[i].Status = true // mark the item as done
			fmt.Println("Marked TODO item as done")
			return
		}
	}
	fmt.Println("TODO item not found")
}

// clearTODOs removes all TODO items from the list, clearing the slice by reinitializing it
func clearTODOs(data *[]TODOs) {
	*data = []TODOs{}
	fmt.Println("Cleared all TODO items")
}
