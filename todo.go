package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	description string
	completed   bool
}

func main() {
	var tasks []Task
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("To-Do List Manager")
	fmt.Println("===================")
	fmt.Println("Commands:")
	fmt.Println(" - add <task>: Add a new task")
	fmt.Println(" - list      : List all tasks")
	fmt.Println(" - complete <number>: Mark a task as completed")
	fmt.Println(" - delete <number>: Delete a task")
	fmt.Println(" - exit      : Quit the application")

	for {
		fmt.Print("\nEnter a command: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if strings.ToLower(input) == "exit" {
			fmt.Println("Goodbye!")
			break
		}

		parts := strings.SplitN(input, " ", 2)
		command := strings.ToLower(parts[0])
		arg := ""
		if len(parts) > 1 {
			arg = parts[1]
		}

		switch command {
		case "add":
			if arg == "" {
				fmt.Println("Error: You must provide a task description.")
			} else {
				tasks = append(tasks, Task{description: arg, completed: false})
				fmt.Printf("Task added: %s\n", arg)
			}

		case "list":
			if len(tasks) == 0 {
				fmt.Println("No tasks in your to-do list.")
			} else {
				fmt.Println("Your To-Do List:")
				for i, task := range tasks {
					status := "[ ]"
					if task.completed {
						status = "[x]"
					}
					fmt.Printf("%d. %s %s\n", i+1, status, task.description)
				}
			}

		case "complete":
			index := parseTaskNumber(arg, len(tasks))
			if index == -1 {
				fmt.Println("Error: Invalid task number.")
			} else {
				tasks[index].completed = true
				fmt.Printf("Task %d marked as completed.\n", index+1)
			}

		case "delete":
			index := parseTaskNumber(arg, len(tasks))
			if index == -1 {
				fmt.Println("Error: Invalid task number.")
			} else {
				fmt.Printf("Task %d deleted: %s\n", index+1, tasks[index].description)
				tasks = append(tasks[:index], tasks[index+1:]...)
			}

		default:
			fmt.Println("Unknown command. Use 'add', 'list', 'complete', 'delete', or 'exit'.")
		}
	}
}

func parseTaskNumber(input string, max int) int {
	num, err := strconv.Atoi(input)
	if err != nil || num < 1 || num > max {
		return -1
	}
	return num - 1
}
