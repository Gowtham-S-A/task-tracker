package main

import (
	"encoding/json"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"os"
	"path/filepath"
	"task-tracker/model"
	"time"
)

func createTask(description string) bool {
	var task model.Task
	timeNow := time.Now().Format(time.RFC3339)
	id := uuid.NewV4().String()

	task.Id = id
	task.Description = description
	task.Status = "toDo"
	task.CreatedAt = timeNow
	task.UpdatedAt = timeNow

	jsonData, err := json.Marshal(task)
	if err != nil {
		fmt.Println("Error marshalling task ===> ", err)
		return false
	}

	err = os.MkdirAll(SubDir, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating directory ===> ", err)
	}

	fileName := filepath.Join(SubDir, id+".json")

	err = os.WriteFile(fileName, jsonData, 0644)
	if err != nil {
		fmt.Println("Error creating task ===> ", err)
		return false
	}

	return true
}

func updateTask(Id string, Description string, Type string) bool {
	timeNow := time.Now().Format(time.RFC3339)
	fileName := filepath.Join(SubDir, Id+".json")
	jsonData, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading task ===> ", err)
		return false
	}
	var task model.Task
	err = json.Unmarshal(jsonData, &task)
	if err != nil {
		fmt.Println("Error unmarshalling task ===> ", err)
		return false
	}
	switch Type {
	case "description":
		task.Description = Description
		task.UpdatedAt = timeNow
	case "inProgress":
		task.Status = "inProgress"
		task.UpdatedAt = timeNow
	case "done":
		task.Status = "done"
		task.UpdatedAt = timeNow
	}
	jsonData, err = json.Marshal(task)
	if err != nil {
		fmt.Println("Error marshalling task ===> ", err)
		return false
	}
	err = os.WriteFile(fileName, jsonData, 0644)
	if err != nil {
		fmt.Println("Error updating task ===> ", err)
		return false
	}
	return true
}

func deleteTask(Id string) bool {
	fileName := filepath.Join(SubDir, Id+".json")
	err := os.Remove(fileName)
	if err != nil {
		fmt.Println("Error deleting task ===> ", err)
		return false
	}
	return true
}

func getTask(status string) {
	files, err := os.ReadDir(SubDir)
	if err != nil {
		fmt.Println("Error reading dir ===> ", err)
	}
	for _, file := range files {
		fileName := filepath.Join(SubDir, file.Name())
		fileData, err := os.ReadFile(fileName)
		if err != nil {
			fmt.Println("Error reading task ===> ", err)
		}
		match := checkTaskStatus(fileData, status)
		if match {
			fmt.Println(string(fileData))
		}
	}
}

func checkTaskStatus(file []byte, status string) bool {
	var task model.Task
	err := json.Unmarshal(file, &task)
	if err != nil {
		fmt.Println("Error unmarshalling task ===> ", err)
		return false
	}
	if task.Status == status {
		return true
	} else if status == "" {
		return true
	} else {
		return false
	}
}
