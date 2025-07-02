package storage

import (
	"encoding/json"
	"log"
	"os"

	"github.com/lokingdav/todogo/internal/tasks"
)

var DbFile = "tasksdb.json"

type dbFormat struct {
	AutoInc int            `json:"autoInc"`
	Tasks   tasks.TasksMap `json:"tasks"`
}

func logErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func SaveTasks(autoInc int, tasks tasks.TasksMap) {
	var file, err = os.Create(DbFile)
	logErr(err)
	defer file.Close()

	var encoder = json.NewEncoder(file)
	var dbdata = map[string]interface{}{
		"autoInc": autoInc,
		"tasks":   tasks,
	}
	err = encoder.Encode(dbdata)

	if err != nil {
		log.Fatal(err)
	}
}

func LoadTasks() (int, tasks.TasksMap) {
	if _, err := os.Stat(DbFile); os.IsNotExist(err) {
		return 0, tasks.TasksMap{}
	}

	var dbdata dbFormat

	var file, err = os.Open(DbFile)
	logErr(err)
	defer file.Close()

	var decoder = json.NewDecoder(file)
	err = decoder.Decode(&dbdata)
	logErr(err)

	return dbdata.AutoInc, dbdata.Tasks
}
