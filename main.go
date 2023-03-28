package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/JFernandoHerreraC/go-cli-crud/tasks"
)

func main() {
	file, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var Tasks []tasks.Task
	info, err := file.Stat()

	if err != nil {
		panic(err)
	}

	if info.Size() != 0 {
		bytes, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(bytes, &Tasks)
		if err != nil {
			panic(err)
		}
	} else {
		Tasks = []tasks.Task{}
	}

	if len(os.Args) < 2 {
		printUsage()
	}

	switch os.Args[1] {
	case "list":
		tasks.ListTasks(Tasks)
	case "add":
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("What is your homework? ")
		name, _ := reader.ReadString('\n')
		strings.TrimSpace(name)
		Tasks = tasks.AddTask(Tasks, name)
		tasks.SaveTask(file, Tasks)
		tasks.ListTasks(Tasks)
	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("Enter the ID: ")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("ID must be an integer!")
			return
		}
		Tasks = tasks.CompleteTask(Tasks, id)
		tasks.SaveTask(file, Tasks)
		tasks.ListTasks(Tasks)
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Enter the ID of the task you want to delete: ")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("ID must be an integer!")
			return
		}
		Tasks = tasks.DeleteTask(Tasks, id)
		tasks.SaveTask(file, Tasks)
		tasks.ListTasks(Tasks)
	default:
		printUsage()
	}
}

func printUsage() {
	fmt.Println("Usage: go-task [list | add | complete | delete]")
}
