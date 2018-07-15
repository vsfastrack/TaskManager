package data

import (
	"github.com/vsfastrack/TaskManager/models"
	mgo "gopkg.in/mgo.v2"
)

//TaskRepository resource methods
type TaskRepository struct {
	C *mgo.Collection
}

//CreateTask handler
func (t *TaskRepository) CreateTask(task *models.Task) {

}
