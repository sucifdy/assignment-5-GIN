package service

import (
	"a21hc3NpZ25tZW50/model"
	repo "a21hc3NpZ25tZW50/repository"
)

type TaskService interface {
	Store(task *model.Task) error
	Update(task *model.Task) error
	Delete(id int) error
	GetByID(id int) (*model.Task, error)
	GetList() ([]model.Task, error)
	GetTaskCategory(id int) ([]model.TaskCategory, error)
}

type taskService struct {
	taskRepository repo.TaskRepository
}

func NewTaskService(taskRepository repo.TaskRepository) TaskService {
	return &taskService{taskRepository}
}

func (s *taskService) Store(task *model.Task) error {
	return s.taskRepository.Store(task)
}

func (s *taskService) Update(task *model.Task) error {
	// Memperbarui task dengan memanggil taskRepository.Update
	return s.taskRepository.Update(task)
}

func (s *taskService) Delete(id int) error {
	// Menghapus task berdasarkan ID dengan memanggil taskRepository.Delete
	return s.taskRepository.Delete(id)
}

func (s *taskService) GetByID(id int) (*model.Task, error) {
	// Mengambil task berdasarkan ID dengan memanggil taskRepository.GetByID
	return s.taskRepository.GetByID(id)
}

func (s *taskService) GetList() ([]model.Task, error) {
	// Mengambil daftar semua task dengan memanggil taskRepository.GetList
	return s.taskRepository.GetList()
}

func (s *taskService) GetTaskCategory(id int) ([]model.TaskCategory, error) {
	// Mengambil task berdasarkan kategori dengan memanggil taskRepository.GetTaskCategory
	return s.taskRepository.GetTaskCategory(id)
}
