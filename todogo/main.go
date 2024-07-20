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

	fmt.Println("Commands: show create remove done clear exit")
	fmt.Print("> ")

	// continuously read input from user and handle the command
	for scanner.Scan() {
		line := scanner.Text()
		if handleInput(line, &data, &counter) {
			break
		}
		fmt.Print("> ")
	}
	fmt.Println("Exiting the application...")
}

// process the user's command and modifies the TODO list accordingly
func handleInput(input string, data *[]TODOs, counter *int) bool {
	split := strings.SplitN(input, " ", 2)
	cmd := split[0]

	// handle different commands based on the user's input
	switch cmd {
	case "create":
		if len(split) < 2 {
			fmt.Println("Error: Missing description for create command")
			return false
		}
		description := split[1]
		createTODO(description, data, counter)
	case "remove":
		if len(split) < 2 {
			fmt.Println("Error: Missing IDs for remove command")
			return false
		}
		removeTODO(split[1], data)
	case "show":
		// display all TODO items
		showTODOs(data)
	case "done":
		if len(split) < 2 {
			fmt.Println("Error: Missing IDs for done command")
			return false
		}
		ids := split[1]
		markTODOAsDone(ids, data)
	case "clear":
		clearTODOs(data)
	case "exit":
		return true
	default:
		// inform the user if the command is invalid
		fmt.Println("Invalid command")
	}
	return false
}

func createTODO(description string, data *[]TODOs, counter *int) {
	// split the description with commas to create multiple TODO items
	items := strings.Split(description, ",")
	for _, item := range items {
		item = strings.TrimSpace(item)
		if item != "" {
			*data = append(*data, TODOs{
				Id:          *counter,
				Description: item,
				Status:      false,
			})
			*counter++
		}
	}
	fmt.Println("Created TODO items")
}

// remove TODO items with the specified IDs from the list
func removeTODO(ids string, data *[]TODOs) {
	idStrs := strings.Split(ids, ",")
	idsToRemove := make(map[int]struct{})
	for _, idStr := range idStrs {
		idStr = strings.TrimSpace(idStr)
		if idStr == "" {
			continue
		}
		id, err := strconv.Atoi(idStr) // convert the ID string to an integer
		if err != nil {
			fmt.Println("Error: Invalid ID", idStr)
			continue
		}
		idsToRemove[id] = struct{}{}
	}
	newData := make([]TODOs, 0)
	for _, todo := range *data {
		if _, exists := idsToRemove[todo.Id]; !exists {
			newData = append(newData, todo)
		}
	}
	*data = newData
	fmt.Println("Removed TODO items")
}

// display all TODO items in the list
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

// set the status of the TODO items with the specified IDs to done
func markTODOAsDone(ids string, data *[]TODOs) {
	idStrs := strings.Split(ids, ",")
	for _, idStr := range idStrs {
		idStr = strings.TrimSpace(idStr)
		if idStr == "" {
			continue
		}
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("Error: Invalid ID", idStr)
			continue
		}
		for i, todo := range *data {
			if todo.Id == id {
				(*data)[i].Status = true // mark the item as done
				fmt.Println("Marked TODO item as done", id)
				break
			}
		}
	}
}

// remove all TODO items from the list, clear the slice by reinitializing it
func clearTODOs(data *[]TODOs) {
	*data = []TODOs{}
	fmt.Println("Cleared all TODO items")
}
