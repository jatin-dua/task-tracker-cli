package main

import (
	"log"
	"os"
)

/*
  -  Add, Update, and Delete tasks
  -  Mark a task as in progress or done
  -  List all tasks
  -  List all tasks that are done
  -  List all tasks that are not done
  -  List all tasks that are in progress
*/

func main() {
	argv := os.Args
	if len(argv) != 3 {
		log.Fatal("usage: todo <command>")
	}
}
