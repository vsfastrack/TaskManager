package data

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/vsfastrack/TaskManager/models"
)

//NoteRepository resource
type NoteRepository struct {
	C *mgo.Collection
}

//Create handler
func (r *NoteRepository) Create(note *models.TaskNote) error {
	objID := bson.NewObjectId()
	note.ID = objID
	note.CreatedOn = time.Now()
	err := r.C.Insert(&note)
	return err
}

//Update handler
func (r *NoteRepository) Update(note *models.TaskNote) error {
	// partial update on MogoDB
	err := r.C.Update(bson.M{"_id": note.ID},
		bson.M{"$set": bson.M{
			"description": note.Description,
		}})
	return err
}

//
//Delete handler
func (r *NoteRepository) Delete(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

//GetByTask handler
func (r *NoteRepository) GetByTask(id string) []models.TaskNote {
	var notes []models.TaskNote
	taskid := bson.ObjectIdHex(id)
	iter := r.C.Find(bson.M{"taskid": taskid}).Iter()
	result := models.TaskNote{}
	for iter.Next(&result) {
		notes = append(notes, result)
	}
	return notes
}

//GetAll handler
func (r *NoteRepository) GetAll() []models.TaskNote {
	var notes []models.TaskNote
	iter := r.C.Find(nil).Iter()
	result := models.TaskNote{}
	for iter.Next(&result) {
		notes = append(notes, result)
	}
	return notes
}

//GetByID handler
func (r *NoteRepository) GetByID(id string) (note models.TaskNote, err error) {
	err = r.C.FindId(bson.ObjectIdHex(id)).One(&note)
	return
}
