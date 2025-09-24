package main

import (
	"bufio"
	"fmt"
	"os"
)

func prompt() int {

	var number int

	fmt.Printf("\n")
	fmt.Println("Welcome to the todo list app.")
	fmt.Printf("\n")
	fmt.Println("Please select one of the following option: ")
	fmt.Printf("\n")
	fmt.Println("1. Add a new task, 2. Delete a task, 3. View the list of tasks, 4. Exit the program ")
	fmt.Printf("\n")
	fmt.Printf("Option: ")
	fmt.Scan(&number)
	fmt.Printf("\n")

	return number
}

func new_task() {

	reader := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter the new task: ")
	reader.Scan()

	line := reader.Text()
	storing_tasks(line)
}

func storing_tasks(task string) {

	f, err := os.OpenFile("file.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file: ", err)
	}
	defer f.Close()

	_, err = f.WriteString(task + "\n")
	if err != nil {
		fmt.Println("Error writing to file: ", err)
	}
}

func num_to_delete() int {

	fmt.Println("Which number tasks you want to delete: ")
	var i int
	fmt.Scanf("%d", &i)

	return i
}

func delete_tasks() {

	number := num_to_delete()
	fmt.Println("The number is: ", number)

	file, err := os.OpenFile("file.txt", os.O_RDWR, 0666)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if number < 1 || number > len(lines) {
		fmt.Println("Invalid line number")
		return
	}

	lines = append(lines[:number-1], lines[number:]...)

	file.Truncate(0)
	file.Seek(0, 0)

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(writer, line)
	}
	writer.Flush()
}

func view_tasks() {

	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var i int
	for scanner.Scan() {
		i++
		fmt.Printf("%d: %s\n", i, scanner.Text())
	}
}

func main() {

	for {
		number := prompt()

		switch number {

		case 1:
			new_task()
		case 2:
			delete_tasks()
		case 3:
			view_tasks()
		case 4:
			fmt.Println("Exiting the program. Goodbye!")
			return
		default:
			prompt()

		}
	}

}
