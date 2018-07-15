package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/vsfastrack/TaskManager/common"
)

//CreateTask handler
func CreateTask(w http.ResponseWriter, r http.Request) {
	var dataResource TaskResource

	err := json.Decoder(r.Body).decode(&dataResource)

	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid Task Data",
			500,
		)
		return
	}
}
