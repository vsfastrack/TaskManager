package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/vsfastrack/TaskManager/common"
	"github.com/vsfastrack/TaskManager/data"
)

//CreateTask handler
func CreateTask(w http.ResponseWriter, r http.Request) {
	var dataResource TaskResource

	err := json.NewDecoder(r.Body).Decode(&dataResource)

	if err != nil {
		common.DisplayAppError(w, err, "Invalid Task Data", 500)
		return
	}

	context := NewContext()
	defer context.Close()
	c := context.DbCollection("tasks")

	repo := &data.TaskRepository{c}

	task := &dataResource.Data

	repo.CreateTask(task)

}
