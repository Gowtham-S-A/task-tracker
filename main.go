package main

import (
	"fmt"
	"os"
	"task-tracker/handler"
)

func main() {
	if len(os.Args) > 2 {
		switch os.Args[1] {
		case "add":
			if len(os.Args) == 3 {
				created := handler.CreateTask(os.Args[2])
				if created {
					fmt.Println("task created successfully")
				}
			} else {
				fmt.Println("usage: add <task description>")
			}
		case "update":
			if len(os.Args) == 4 {
				updated := handler.UpdateTask(os.Args[2], os.Args[3], handler.TaskDescription)
				if updated {
					fmt.Println("task updated successfully")
				}
			} else {
				fmt.Println("usage: add <task name> <task description>")
			}
		case "mark-in-progress", "mark-done":
			if len(os.Args) == 3 {
				status := handler.TaskStatusInProgress
				if os.Args[1] == "mark-done" {
					status = handler.TaskStatusDone
				}
				updated := handler.UpdateTask(os.Args[2], "", status)
				if updated {
					fmt.Println("Task marked successfully.")
				}
			} else {
				fmt.Println("usage: mark-in-progress|mark-done <task ID>")
			}
		case "delete":
			if len(os.Args) == 3 {
				deleted := handler.DeleteTask(os.Args[2])
				if deleted {
					fmt.Println("Task deleted successfully.")
				}
			} else {
				fmt.Println("usage: delete <task ID>")
			}
		case "list":
			if len(os.Args) == 3 {
				taskType := os.Args[2]
				err := handler.GetTask(taskType)
				if err != nil {
					fmt.Println("Unable to get task:", err.Error())
				}
			} else {
				fmt.Println("usage: list <toDo|inProgress|done>")
			}
		default:
			fmt.Println("Invalid command, usage: add|update|delete|mark-in-progress|mark-done|list")
		}
	} else if len(os.Args) == 2 {
		switch os.Args[1] {
		case "list":
			err := handler.GetTask("")
			if err != nil {
				fmt.Println("Unable to get task :", err.Error())
			}
		default:
			fmt.Println("Invalid command, usage: add|update|delete|mark-in-progress|mark-done|list")
		}
	} else {
		fmt.Println("Invalid command, usage: add|update|delete|mark-in-progress|mark-done|list")
	}
}
