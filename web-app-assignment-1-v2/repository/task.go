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

func (r *taskRepository) Store(task *model.Task) error {
	r.filebased.StoreTask(*task)

	return nil
}

func (r *taskRepository) Update(task *model.Task) error {
	err := r.filebased.UpdateTask(task.ID, *task) 
	if err != nil {
		return err
	}
	return nil
}


func (r *taskRepository) Delete(id int) error {
	err := r.filebased.DeleteTask(id)
	if err != nil {
		return err
	}
	return nil
}

func (r *taskRepository) GetByID(id int) (*model.Task, error) {
	task, err := r.filebased.GetTaskByID(id)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (r *taskRepository) GetList() ([]model.Task, error) {
	tasks, err := r.filebased.GetTasks()
	if err != nil {
		return tasks, err
	}
	return tasks, nil
}

func (r *taskRepository) GetTaskCategory(id int) ([]model.TaskCategory, error) {
	categories, err := r.filebased.GetTaskListByCategory(id)
	if err != nil {
		return nil, err
	}
	return categories, nil
}
