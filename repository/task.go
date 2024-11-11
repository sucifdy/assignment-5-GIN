package repository

import (
	"a21hc3NpZ25tZW50/db/filebased"
	"a21hc3NpZ25tZW50/model"
)

type TaskRepository interface {
	Store(task *model.Task) error
	Update(task *model.Task) error
	Delete(id int) error
	GetByID(id int) (*model.Task, error)
	GetList() ([]model.Task, error)
	GetTaskCategory(id int) ([]model.TaskCategory, error)
}

type taskRepository struct {
	filebased *filebased.Data
}

func NewTaskRepo(filebasedDb *filebased.Data) *taskRepository {
	return &taskRepository{
		filebased: filebasedDb,
	}
}

func (t *taskRepository) Store(task *model.Task) error {
	return t.filebased.StoreTask(*task)
}

func (t *taskRepository) Update(task *model.Task) error {
	// Perbarui task berdasarkan ID-nya
	return t.filebased.UpdateTask(task.ID, *task)
}

func (t *taskRepository) Delete(id int) error {
	// Hapus task berdasarkan ID menggunakan DeleteTask dari filebased
	return t.filebased.DeleteTask(id)
}

func (t *taskRepository) GetByID(id int) (*model.Task, error) {
	// Ambil task berdasarkan ID menggunakan GetTaskByID dari filebased
	return t.filebased.GetTaskByID(id)
}

func (t *taskRepository) GetList() ([]model.Task, error) {
	// Ambil semua task menggunakan GetTasks dari filebased
	return t.filebased.GetTasks()
}

func (t *taskRepository) GetTaskCategory(id int) ([]model.TaskCategory, error) {
	// Ambil task berdasarkan kategori menggunakan GetTaskListByCategory dari filebased
	return t.filebased.GetTaskListByCategory(id)
}
