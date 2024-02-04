package main

import (
	"flag"
	"fmt"
	"os"
)

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
	tm.Tasks = append(tm.Tasks[:index], tm.Tasks[index+1:]...)
}

func (tm *TaskManager) ListTasks() {
	for i, task := range tm.Tasks {
		fmt.Printf("Task %d:\n", i)
		fmt.Printf("Title: %s\n", task.Title)
		fmt.Printf("Description: %s\n", task.Description)
		fmt.Printf("Complete: %v\n", task.Complete)
	}
}

func main() {
	tm := &TaskManager{}

	addCommand := flag.NewFlagSet("add", flag.ExitOnError)
	addTitle := addCommand.String("title", "", "Title of the task")
	addDescription := addCommand.String("description", "", "Description of the task")

	completeCommand := flag.NewFlagSet("complete", flag.ExitOnError)
	completeIndex := completeCommand.Int("index", -1, "Index of the task to complete")

	deleteCommand := flag.NewFlagSet("delete", flag.ExitOnError)
	deleteIndex := deleteCommand.Int("index", -1, "Index of the task to delete")

	listCommand := flag.NewFlagSet("list", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("expected 'add' or 'complete' or 'delete' or 'list' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		addCommand.Parse(os.Args[2:])

		tm.AddTask(*addTitle, *addDescription)
		tm.ListTasks()

	case "complete":
		completeCommand.Parse(os.Args[2:])
		tm.CompleteTask(*completeIndex)

	case "delete":
		deleteCommand.Parse(os.Args[2:])
		tm.DeleteTask(*deleteIndex)

	case "list":
		listCommand.Parse(os.Args[2:])
		tm.ListTasks()

	default:
		flag.PrintDefaults()
		os.Exit(1)
	}
}
