// package controllers

// import (
// 	"encoding/json"
// 	"net/http"

// 	"github.com/vsfastrack/TaskManager/common"
// 	"github.com/vsfastrack/TaskManager/data"
// )

// //CreateNote handler
// func CreateNote(w http.ResponseWriter, r *http.Request) {

// 	var noteResource NoteResource

// 	err := json.NewDecoder(r.Body).Decode(&noteResource)

// 	if err != nil {
// 		common.DisplayAppError(w, err, "Invalid Note Data", 500)
// 		return
// 	}

// 	context := NewContext()
// 	defer context.Close()
// 	c := context.DbCollection("notes")

// 	repo := &data.NoteRepository{c}
// 	note := &noteResource.Data

// 	repo.Create(noteResource)
// 	if j, err := json.Marshal(TaskResource{Data: *note}); err != nil {
// 		common.DisplayAppError(
// 			w,
// 			err,
// 			"An unexpected error has occurred",
// 			500,
// 		)
// 		return
// 	} else {
// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusCreated)
// 		w.Write(j)
// 	}

// }
