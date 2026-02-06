package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"time"
)

func main(){

	// load json file
	var tasks Tasks

	jsontasks, err := os.ReadFile("tasks.json")
	if err!=nil{
		if os.IsNotExist(err){
			os.Create("tasks.json")
		}
	}else{
		err = json.Unmarshal(jsontasks, &tasks)
		if err!=nil{
			log.Fatal(err)
		}
	}

	// flag setup
	addFlag := flag.String("add","",`add your task to JSON file. Usage -add "task_description"`)

	flag.Parse()
	

	// Add task
	if *addFlag == ""{
		log.Fatal("Task description can't be empty")
	}
	task := Task{
		Id : len(tasks.Tasks)+1,
		Description: *addFlag,
		Status: "todo",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	tasks.Tasks = append(tasks.Tasks, task)

	jsontasks, err = json.Marshal(&tasks)
	if err!=nil{
		log.Fatal(err)
	}

	err = os.WriteFile("tasks.json", jsontasks, 0666)
	if err!=nil{
		log.Fatal(err)
	}
}