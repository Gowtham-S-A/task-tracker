package handler

import (
	"encoding/json"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"os"
	"path/filepath"
	"task-tracker/model"
	"time"
)

const (
	TaskStatusToDo       = "toDo"
	TaskStatusInProgress = "inProgress"
	TaskStatusDone       = "done"
	TaskDescription      = "description"
	SubDir               = "taskList"
)

func CreateTask(description string) bool {
	var task model.Task
	timeNow := time.Now().Format(time.RFC3339)
	id := uuid.NewV4().String()

	task.Id = id
	task.Description = description
	task.Status = TaskStatusToDo
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

func UpdateTask(Id string, Description string, Type string) bool {
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
	case TaskDescription:
		task.Description = Description
	case TaskStatusInProgress:
		task.Status = TaskStatusInProgress
	case TaskStatusDone:
		task.Status = TaskStatusDone
	}
	task.UpdatedAt = timeNow
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

func DeleteTask(Id string) bool {
	fileName := filepath.Join(SubDir, Id+".json")
	err := os.Remove(fileName)
	if err != nil {
		fmt.Println("Error deleting task ===> ", err)
		return false
	}
	return true
}

func GetTask(status string) error {
	files, err := os.ReadDir(SubDir)
	if err != nil {
		fmt.Println("Error reading dir ===> ", err)
		return err
	}
	fmt.Printf("<========== List of %s tasks ==========>\n", status)
	for _, file := range files {
		match := false
		fileName := filepath.Join(SubDir, file.Name())
		fileData, err := os.ReadFile(fileName)
		if err != nil {
			fmt.Println("Error reading task ===> ", err)
			return err
		}
		var task model.Task
		err = json.Unmarshal(fileData, &task)
		if err != nil {
			fmt.Println("Error unmarshalling task ===> ", err)
			return err
		}
		if task.Status == status {
			match = true
		} else if status == "" {
			match = true
		}
		if match {
			fmt.Println(string(fileData))
		}
	}
	return nil
}
