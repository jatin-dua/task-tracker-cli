package main

import (
	"log"
	"os"
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

func main() {
	argv := os.Args
	if len(argv) != 3 {
		log.Fatal("usage: todo <command>")
	}
}
