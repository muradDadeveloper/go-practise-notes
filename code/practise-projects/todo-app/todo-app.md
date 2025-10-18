# Simple Todo List Application in Go

This is a command-line todo list application written in Go that allows users to manage their tasks through a simple interface.

## Features

- Add new tasks to the list
- Delete tasks by their number
- View all tasks in the list
- Persistent storage using a text file

## Implementation Details

The application is built using Go's standard library and implements the following key features:

### File Operations
- Uses `os` package for file handling
- Stores tasks in a `file.txt` file
- Implements read and write operations with proper error handling

### User Interface
- Console-based menu system
- Clear numbered options for task management
- Interactive prompt for user input

### Core Functions

1. `prompt()`: Displays the main menu and returns user's choice
2. `new_task()`: Handles the addition of new tasks
3. `storing_tasks()`: Manages the file writing operation for new tasks
4. `delete_tasks()`: Removes specified tasks from the list
5. `view_tasks()`: Displays all current tasks with numbering
6. `num_to_delete()`: Gets user input for task deletion

## How to Use

1. Run the program
2. Choose from the following options:
   - 1: Add a new task
   - 2: Delete a task
   - 3: View all tasks
   - 4: Exit the program

## Technical Implementation

The program uses a simple text file (`file.txt`) for persistent storage. Each task is stored on a new line, and the application manages the file operations with proper error handling and resource cleanup using Go's `defer` statements.

### Error Handling
- File operations are wrapped with error checking
- Invalid input handling for task deletion
- Proper file resource management

## Learning Outcomes

This project demonstrates understanding of:
- Basic Go syntax and control structures
- File I/O operations in Go
- User input handling
- Error management
- Basic data structures (slices)
- Command-line interface design

## Future Improvements

Potential enhancements could include:
- Task prioritization
- Due dates for tasks
- Categories/tags for tasks
- Task completion status
- Better data persistence (database instead of text file)
- Input validation
- Task editing functionality
