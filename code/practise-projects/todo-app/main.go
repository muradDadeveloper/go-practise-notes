package main

import (
	"fmt"
)

/* func new_task() {

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
} */

func num_to_delete() int {

	fmt.Println("Which number tasks you want to delete: ")
	var i int
	fmt.Scanf("%d", &i)

	return i
}

func delete_tasks(num int) {

}

/* func prompt() int {

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
} */

func main() {

	// number := prompt()

	new_task()

}
