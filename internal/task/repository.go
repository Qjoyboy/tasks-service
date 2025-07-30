package task

import "gorm.io/gorm"

type TaskRepository interface {
	CreateTask(t *Task) error
	GetTasks() ([]Task, error)
	GetTaskByID(id string) (Task, error)
	GetTasksForUser(id string) ([]Task, error)
	UpdateTask(t Task) error
	DeleteTask(id string) error
}

type taskRepo struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepo{db: db}
}

func (r *taskRepo) CreateTask(t *Task) error {
	return r.db.Create(&t).Error
}

func (r *taskRepo) GetTasks() ([]Task, error) {
	var task []Task
	err := r.db.Find(&task).Error
	return task, err
}

func (r *taskRepo) GetTaskByID(id string) (Task, error) {
	var task Task
	err := r.db.First(&task, "id=?", id).Error
	return task, err
}
func (r *taskRepo) GetTasksForUser(id string) ([]Task, error) {
	var userTasks []Task
	err := r.db.Where("user_id=?", id).Find(&userTasks).Error
	return userTasks, err
}
func (r *taskRepo) UpdateTask(t Task) error {
	return r.db.Save(&t).Error
}

func (r *taskRepo) DeleteTask(id string) error {
	return r.db.Delete(&Task{}, "id=?", id).Error
}
