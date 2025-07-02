package main

import (
	"flag"
	"fmt"
	"math/rand"

	"os"
	"log"
	"encoding/json"
)

const (
	taskAdd      = "add"
	taskDelete   = "delete"
	taskComplete = "complete"
	taskList     = "list"
)

var version = "dev"
var dbfile = "tasksdb.json"

type task struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Desc      string `json:"desc"`
	Completed bool   `json:"completed"`
}

func logErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func saveTasks(tasks map[int]task) {
	var file, err = os.Create(dbfile)
	logErr(err)
	defer file.Close()

	var encoder = json.NewEncoder(file)
	err = encoder.Encode(tasks)
	if err != nil {
		log.Fatal(err)
	}
}

func loadTasks() map[int]task {
	if _, err := os.Stat(dbfile); os.IsNotExist(err) {
		return map[int]task{}
	}

	var tasks map[int]task

	var file, err = os.Open(dbfile)
	logErr(err)
	defer file.Close()

	var decoder = json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	logErr(err)
	return tasks
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
		log.Fatal("Validate Errors: ", validationErrors)
	}

	var taskId = rand.Int() % 1000
	fmt.Println("Adding task: ", taskId)
	tasks[taskId] = task{taskId, name, desc, false}
	return taskId
}

func deleteTask(taskId int, tasks map[int]task) {
	fmt.Println("\nDeleting task: ", taskId)
	delete(tasks, taskId)
}

func completeTask(taskId int, tasks map[int]task) {
	fmt.Println("\nCompleting task: ", taskId)
	if tsk, ok := tasks[taskId]; ok {
		tasks[taskId] = task{taskId, tsk.Name, tsk.Desc, true}
	} else {
		log.Fatalf("Task %d does not exists", taskId)
	}
}

func listTasks(tasks map[int]task) {
	if len(tasks) == 0 {
		log.Fatal("There are not available tasks")
	}

	fmt.Printf("\n==================== START Tasks List (%d) ====================\n", len(tasks))
	for k, v := range tasks {
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
		log.Fatalf("Command (%s|%s|%s|%s) is required", taskAdd, taskDelete, taskComplete, taskList)
	}

	var tasks = loadTasks()

	switch cmd {
	case taskAdd:
		addTask(name, desc, tasks)
	case taskList:
		listTasks(tasks)
	default:
		if taskId == 0 {
			log.Fatal("You must provide a Task ID --id")
		}
		switch cmd {
		case taskDelete:
			deleteTask(taskId, tasks)
		case taskComplete:
			completeTask(taskId, tasks)
		}
	}

	saveTasks(tasks)
}
