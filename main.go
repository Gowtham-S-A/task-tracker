package main

import (
	"fmt"
	"os"
)

const SubDir = "taskList"

func main() {
	if len(os.Args) > 2 {
		switch os.Args[1] {
		case "add":
			created := createTask(os.Args[2])
			if created {
				fmt.Println("task created successfully")
			}
		case "update":
			if len(os.Args) != 4 || len(os.Args) > 4 {
				fmt.Println("usage: add <task name> <task description>")
			} else {
				updated := updateTask(os.Args[2], os.Args[3], "description")
				if updated {
					fmt.Println("task updated successfully")
				}
			}
		case "mark-in-progress":
			updated := updateTask(os.Args[2], "", "inProgress")
			if updated {
				fmt.Println("task marked in progress")
			}
		case "mark-done":
			updated := updateTask(os.Args[2], "", "done")
			if updated {
				fmt.Println("task marked done successfully")
			}
		case "delete":
			deleted := deleteTask(os.Args[2])
			if deleted {
				fmt.Println("task deleted successfully")
			}
		case "list":
			switch os.Args[2] {
			case "toDo":
				fmt.Println("<========== List of to do tasks ==========>")
				getTask("toDo")
			case "inProgress":
				fmt.Println("<========== List of in progress tasks ==========>")
				getTask("inProgress")
			case "done":
				fmt.Println("<========== List of done tasks ==========>")
				getTask("done")
			}
		}
	} else if len(os.Args) == 2 {
		switch os.Args[1] {
		case "list":
			fmt.Println("<========== List of tasks ==========>")
			getTask("")
		default:
			fmt.Println("Usage: add|update|delete|mark-in-progress|mark-done|list")
		}
	} else {
		fmt.Println("Usage: add|update|delete|mark-in-progress|mark-done|list")
	}
}
