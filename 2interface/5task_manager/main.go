package main

import (
	"errors"
	"fmt"
)

var ErrTaskNotFound = errors.New("задача не найдена")

type Task struct {
	ID    int
	Title string
	Done  bool
}
type TaskStorage interface {
	Add(title string) Task
	Complete(id int) error
	List() []Task
}
type InMemoryStorage struct {
	tasks []Task
}

type LoggingStorage struct {
	inner TaskStorage
}

func (ls *LoggingStorage) Add(title string) Task {
	task := ls.inner.Add(title)
	fmt.Println("[LOG]: Добавлена задача:", task.Title)
	return task

}
func (ls *LoggingStorage) Complete(id int) error {
	return ls.inner.Complete(id)
}
func (ls *LoggingStorage) List() []Task {
	return ls.inner.List()

}
func PrintAllTasks(ts TaskStorage) {
	for _, task := range ts.List() {
		fmt.Println(task.ID, task.Title, task.Done)
	}

}
func (st *InMemoryStorage) Add(title string) Task {

	new_task := Task{
		ID:    len(st.tasks) + 1,
		Title: title,
		Done:  false,
	}
	st.tasks = append(st.tasks, new_task)
	return new_task

}
func (st *InMemoryStorage) Complete(id int) error {
	flag := 0
	for i := range st.tasks {
		if id == st.tasks[i].ID {
			flag = 1
			st.tasks[i].Done = true

		}

	}
	if flag == 0 {
		return ErrTaskNotFound

	}

	return nil
}
func (st *InMemoryStorage) List() []Task {
	return st.tasks

}

func main() {
	storage := &LoggingStorage{inner: &InMemoryStorage{}}

	_ = storage.Add("Учить англиский")

	_ = storage.Add("Сходить в зал")

	_ = storage.Add("Медитация")

	err := storage.Complete(1)
	if errors.Is(err, ErrTaskNotFound) {
		fmt.Println("ID не найден!")
	} else {
		fmt.Println("Задача выполнена!")
	}
	err2 := storage.Complete(999)
	if errors.Is(err2, ErrTaskNotFound) {
		fmt.Println("ID не найден!")
	} else {
		fmt.Println("Задача выполнена!")
	}

	result := storage.List()
	for i := range result {
		fmt.Println(result[i].ID, ".", result[i].Title, result[i].Done)

	}
	PrintAllTasks(storage)

}
