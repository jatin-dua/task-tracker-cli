package main

import (
	"encoding/json"
	"flag"
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
	Id		int `json:"id"`
	Description	string `json:"desc"`
	Status		string `json:"status"`
	CreatedAt	time.Time `json:"created_at"`
	UpdatedAt	time.Time `json:"updated_at"`
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

func addTask(filename, desc string) int {
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

func deleteTask(filename string, taskId int) error {
	var updatedTasks []Task
	tasks, err := readJSONFile(filename)
	if err != nil {
		return err
	}

	for _, task := range tasks {
		if task.Id == taskId {
			continue
		}
		updatedTasks = append(updatedTasks, task)
	}
	if err := writeJSONFile(filename, updatedTasks); err != nil {
		return err
	}
	return nil
}

func updateTaskStatus(filename string, taskId int, newStatus string) error {
	var updatedTasks []Task
	tasks, err := readJSONFile(filename)
	if err != nil {
		return err
	}

	for _, task := range tasks {
		if task.Id == taskId {
			task.Status = newStatus
		}
		updatedTasks = append(updatedTasks, task)
	}
	if err := writeJSONFile(filename, updatedTasks); err != nil {
		return err
	}
	return nil
}

func main() {
	taskPtr := flag.String("add", "", "Task description")
	deleteIdPtr := flag.Int("delete", 0, "Task Id to delete")
	markInProgressIdPtr := flag.Int("mark-in-progress", 0, "Task Id to mark-in-progress")
	markDoneIdPtr := flag.Int("mark-done", 0, "Task Id to mark-done")

	filename := "db/tasks.json"

	flag.Parse()

	if *taskPtr != "" {
		taskId := addTask(filename, *taskPtr)
		log.Printf("Task added successfully (ID: %d)", taskId)
	}
	
	if *deleteIdPtr != 0 {
		if err := deleteTask(filename, *deleteIdPtr); err != nil {
			log.Fatal(err)
		}
		log.Printf("Task deleted successfully (ID: %d)", *deleteIdPtr)
	}

	if *markInProgressIdPtr != 0 {
		if err := updateTaskStatus(filename, *markInProgressIdPtr, "in-progress"); err != nil {
			log.Fatal(err)
		}
		log.Printf("Task marked-in-progress successfully (ID: %d)", *markInProgressIdPtr)
	}

	if *markDoneIdPtr != 0 {
		if err := updateTaskStatus(filename, *markDoneIdPtr, "done"); err != nil {
			log.Fatal(err)
		}
		log.Printf("Task marked-done successfully (ID: %d)", *markDoneIdPtr)
	}
}
