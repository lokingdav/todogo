package main

import (
	"flag"
	"fmt"
	"math/rand"
	// "os"
)

const (
	taskAdd      = "add"
	taskDelete   = "delete"
	taskComplete = "complete"
	taskList     = "list"
)

var version = "dev"

type task struct {
	id   int
	name string
	desc string
	done bool
}

func addTask(name, desc string, tasks map[int]task) int {
	var validationErrors []string = make([]string, 0, 2)

	if name == "" {
		validationErrors = append(validationErrors, "Task name is required: Use --name")
	}

	if desc == "" {
		validationErrors = append(validationErrors, "Task description is required: Use --desc")
	}

	if len(validationErrors) > 0 {
		fmt.Println("There are some validation errors:")
		for _, err := range validationErrors {
			fmt.Println("\t" + err)
		}
		return 0
	}

	var taskId = rand.Int() % 1000
	tasks[taskId] = task{taskId, name, desc, false}
	return taskId
}

func deleteTask(taskId int, tasks map[int]task) {
	fmt.Println("\nDeleting task: ", taskId)
	delete(tasks, taskId)
}

func completeTask(taskId int, tasks map[int]task) {
	fmt.Println("\nCompleting task: ", taskId)
	tasks[taskId] = task{taskId, tasks[taskId].name, tasks[taskId].desc, true}
}

func listTasks(tasks map[int]task) {
	if len(tasks) == 0 {
		fmt.Println("There are not available tasks")
		return
	}

	fmt.Printf("\n==================== START Tasks List (%d) ====================\n", len(tasks))
	for k, v := range tasks {
		var status string
		if v.done {
			status = "Yes"
		} else {
			status = "No"
		}
		fmt.Printf("\nID: %d\nName: %s\nDesc: %s\nDone: %s\n", k, v.name, v.desc, status)
	}
	fmt.Println("\n==================== END Tasks List ====================")
}

func main() {
	var taskId int
	var showVersion bool
	var name, desc, cmd string

	flag.BoolVar(&showVersion, "version", false, "Show version number")
	flag.IntVar(&taskId, "id", 0, "The Task ID")
	flag.StringVar(&name, "name", "", "The name of the Task Name")
	flag.StringVar(&desc, "desc", "", "The description of the Task")
	flag.Parse()

	if showVersion {
		fmt.Println("Version:", version)
		if len(flag.Args()) == 0 {
			return
		}
	}

	if args := flag.Args(); len(args) > 0 && (args[0] == taskAdd ||
		args[0] == taskDelete ||
		args[0] == taskComplete ||
		args[0] == taskList) {
		cmd = args[0]
	} else {
		fmt.Printf("Command (%s|%s|%s|%s) is required", taskAdd, taskDelete, taskComplete, taskList)
		return
	}

	var tasks = make(map[int]task, 0)

	switch cmd {
	case taskAdd:
		addTask(name, desc, tasks)
	case taskList:
		listTasks(tasks)
	default:
		if taskId == 0 {
			fmt.Println("You must provide a Task ID --id")
		}
		switch cmd {
		case taskDelete:
			deleteTask(taskId, tasks)
		case taskComplete:
			completeTask(taskId, tasks)
		}
	}
}
