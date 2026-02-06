package main

import (
	"flag"
)

func main(){

	// load json file
	tasks := Load()

	// flag setup
	addFlag := flag.String("add","",`add your task to JSON file. Usage -add task_description`)
	updateFlag := flag.Int("update",0,`update your available task. Usage -update task_id`)

	flag.Parse()
	

	// Add task
	tasks.Add(addFlag)
	
	// Update Task
	tasks.Update(updateFlag, flag.Arg(0))
}