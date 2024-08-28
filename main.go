package main

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

/*
  -  Add, Update, and Delete tasks
  -  Mark a task as in progress or done
  -  List all tasks
  -  List all tasks that are done
  -  List all tasks that are not done
  -  List all tasks that are in progress
*/

type Task struct {
	Id		int
	Description	string
	Status		string
	CreatedAt	time.Time
	UpdatedAt	time.Time
}

func generateId() int {
	data, err := os.ReadFile("db/counter")
	if err != nil {
		log.Fatal("unable to read counter")
	}
	id, err := strconv.Atoi(strings.TrimSpace(string(data)))
	if err != nil {
		log.Fatal("id cannot be generated\n", err)
	}

	err = os.WriteFile("db/counter", []byte(strconv.Itoa(id + 1)), 0666)
	if err != nil {
		log.Println(err)
	}

	return id
}

func addTask(desc, filename string) int {
	id := generateId()
	curTime := time.Now()
	task := Task{
		Id: id,
		Description: desc,
		Status: "todo",
		CreatedAt: curTime,
		UpdatedAt: curTime,
	}
	// Add to json file
	if err := ensureFileExists(filename); err != nil {
		log.Println(err)
	}
	tasks, err := readJSONFile(filename)
	if err != nil {
		log.Println(err)
	}

	tasks = append(tasks, task)
	if err := writeJSONFile(filename, tasks); err != nil {
		log.Println(err)
	}

	return id
}

func ensureFileExists(filename string) error {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		initialData := []Task{}
		return writeJSONFile(filename, initialData)
	}
	return nil
}

func readJSONFile(filename string) ([]Task, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return []Task{}, err
	}
	
	var tasks []Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return []Task{}, err
	}
	return tasks, nil
}

func writeJSONFile(filename string, tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(filename, data, 0666); err != nil {
		return err
	}
	return nil
}

func main() {
	argv := os.Args
	if len(argv) != 3 {
		log.Fatal("usage: todo <command>")
	}

	command := argv[1]
	// option := argv[2]
	filename := "db/tasks.json"

	switch command {
	case "add":
		task := argv[2]
		taskId := addTask(task, filename)
		log.Printf("Task added successfully (ID: %d)", taskId)
	case "update":
	case "delete":
	case "mark-in-progress":
	case "mark-done":
	case "list":
	default:
		log.Fatal("invalid command")
	}
}
