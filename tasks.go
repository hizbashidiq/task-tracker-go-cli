package main

import(
	"encoding/json"
	"time"
	"log"
	"os"
	"fmt"
)

type Tasks struct{
	Tasks []Task `json:"tasks"`
}

func Load()Tasks{
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
	return tasks
}

func(t *Tasks)Save(){
	jsontasks, err := json.MarshalIndent(&t, "", "  ")
	if err!=nil{
		log.Fatal(err)
	}

	err = os.WriteFile("tasks.json", jsontasks, 0666)
	if err!=nil{
		log.Fatal(err)
	}
}

func(t *Tasks)Add(flag *string){
	if *flag != ""{
		task := Task{
			Id : len(t.Tasks)+1,
			Description: *flag,
			Status: "todo",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
	
		t.Tasks = append(t.Tasks, task)
		t.Save()
		fmt.Printf("Task added successfully (ID):%d", task.Id)
	}
}

func(t *Tasks)Update(flag *int, description string){
	if *flag != 0{
		i := 0
		for i:=0;i<len(t.Tasks);i++{
			if t.Tasks[i].Id == *flag{
				break
			}
		}
		if t.Tasks[i].Id != *flag{
			log.Fatalf("There's no task with ID: %d\n", *flag)
		}

		t.Tasks[i].Description = description
		t.Tasks[i].UpdatedAt = time.Now()

		t.Save()
		fmt.Printf("Task updated successfully (ID: %d)", *flag)
	}
}