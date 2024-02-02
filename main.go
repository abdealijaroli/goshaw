package main

type Task struct {
	Title       string
	Description string
	Complete    bool
}

type TaskManager struct {
	Tasks []Task
}

func (tm *TaskManager) AddTask(title, description string) {
	tm.Tasks = append(tm.Tasks, Task{Title: title, Description: description})
}

func (tm *TaskManager) CompleteTask(index int) {
	tm.Tasks[index].Complete = true
}

func (tm *TaskManager) DeleteTask(index int) {
	tm.Tasks = append(tm.Tasks[:index], tm.Tasks[index+1:]... )
}

func main() {

}
