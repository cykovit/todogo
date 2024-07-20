# todogo

## Overview

todogo is a command-line TODO list application written in Go. It allows users to manage a list of TODO items through simple commands. 

## Features

- **create**: add a new TODO item with a description 
- **remove**: remove a TODO item by its ID 
- **show**: display all TODO items with their ID and status
- **done**: mark a TODO item as completed by its ID 
- **clear**: remove all TODO items from the list
- **exit**: quit and close todogo

## Usage

1. **Running the Program**

   ```sh
   go run main.go
   ```

2. **Commands**

   - **create** `<description>` -> Descriptions can be separated by commas to add multiple items at once
     ```
     > create buy groceries, cry, meditate, cry again
     Created TODO item
     ```

   - **remove** `<id>` -> IDs can be separated by commas to remove multiple items at once
     ```
     > remove 0, 1, 3
     Removed TODO item
     ```

   - **show**
     ```
     > show
     ------
     [ ] buy groceries ID: 0
     ```

   - **done** `<id>` -> IDs can be seperated by commas to mark multiple items as done
     ```
     > done 2, 3
     Marked TODO item as done
     ```

   - **clear**
     ```
     > clear
     Cleared all TODO items
     ```

   - **exit**
     ```
     > exit
     Exiting todogo...
     ```

## Dependencies

- Go (version 1.18 or later)

## License

This project is licensed under the MIT License.

Feel free to adjust the content based on your specific needs or any additional information you might want to include.
