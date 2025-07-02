package tasks

import (
	"log"
	"fmt"
)

type Task struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Desc      string `json:"desc"`
	Completed bool   `json:"completed"`
}

type TasksMap map[int]Task

func AddTask(taskId int, name, desc string, tasks TasksMap) int {
	var validationErrors []string = make([]string, 0, 2)

	if name == "" {
		validationErrors = append(validationErrors, "Task name is required: Use --name")
	}

	if desc == "" {
		validationErrors = append(validationErrors, "Task description is required: Use --desc")
	}

	if len(validationErrors) > 0 {
		log.Fatal("Validate Errors: ", validationErrors)
	}

	fmt.Println("Adding task: ", taskId)
	tasks[taskId] = Task{Id: taskId, Name: name, Desc: desc, Completed: false}
	return taskId
}

func DeleteTask(taskId int, tasks TasksMap) {
	fmt.Println("\nDeleting task: ", taskId)
	delete(tasks, taskId)
}

func CompleteTask(taskId int, tasks TasksMap) {
	fmt.Println("\nCompleting task: ", taskId)
	if tsk, ok := tasks[taskId]; ok {
		tasks[taskId] = Task{Id: taskId, Name: tsk.Name, Desc: tsk.Desc, Completed: true}
	} else {
		log.Fatalf("Task %d does not exists", taskId)
	}
}

func ListTasks(tasks TasksMap, completed int) {
	if len(tasks) == 0 {
		log.Fatal("There are not available tasks")
	}

	fmt.Printf("\n==================== START Tasks List (%d) ====================\n", len(tasks))
	for k, v := range tasks {
		if (completed == 0 && v.Completed) || (completed == 1 && !v.Completed) {
			continue
		}

		var status string
		if v.Completed {
			status = "Yes"
		} else {
			status = "No"
		}
		fmt.Printf("\nID: %d\nName: %s\nDesc: %s\nCompleted: %s\n", k, v.Name, v.Desc, status)
	}
	fmt.Println("\n==================== END Tasks List ====================")
}