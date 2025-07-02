package main

import (
	"flag"
	"fmt"
	"log"
	"github.com/lokingdav/todogo/internal/storage"
	taskspkg "github.com/lokingdav/todogo/internal/tasks"
)

const (
	taskAdd      = "add"
	taskDelete   = "delete"
	taskComplete = "complete"
	taskList     = "list"
)

var version = "dev"


func main() {
	var taskId, completed int
	var showVersion bool
	var name, desc, cmd string

	flag.BoolVar(&showVersion, "version", false, "Show version number")
	flag.IntVar(&taskId, "id", 0, "The Task ID")
	flag.StringVar(&name, "name", "", "The name of the Task Name")
	flag.StringVar(&desc, "desc", "", "The description of the Task")
	flag.IntVar(&completed, "completed", -1, "Completed (0|1)")
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

	var  autoInc, tasks = storage.LoadTasks()

	switch cmd {
	case taskAdd:
		autoInc++
		taskspkg.AddTask(autoInc, name, desc, tasks)
	case taskList:
		taskspkg.ListTasks(tasks, completed)
	default:
		if taskId == 0 {
			log.Fatal("You must provide a Task ID --id")
		}
		switch cmd {
		case taskDelete:
			taskspkg.DeleteTask(taskId, tasks)
		case taskComplete:
			taskspkg.CompleteTask(taskId, tasks)
		}
	}

	storage.SaveTasks(autoInc, tasks)
}
