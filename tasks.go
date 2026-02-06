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
			Id : t.Tasks[len(t.Tasks)-1].Id+1,
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
	if *flag != 0 && description!=""{
		i := 0
		found := false
		for ;i<len(t.Tasks);i++{
			if t.Tasks[i].Id == *flag{
				found = true
				break
			}
		}
		if !found{
			log.Fatalf("There's no task with ID: %d\n", *flag)
		}

		t.Tasks[i].Description = description
		t.Tasks[i].UpdatedAt = time.Now()

		t.Save()
		fmt.Printf("Task updated successfully (ID: %d)", *flag)
	}
}

func(t *Tasks)Delete(flag *int){
	if *flag != 0{
		i:=0
		found := false
		for ;i<len(t.Tasks);i++{
			if t.Tasks[i].Id == *flag{
				found = true
				break
			}
		}
		if !found {
			log.Fatalf("There's no task with ID: %d\n", *flag)
		}
		t.Tasks = append(t.Tasks[:i], t.Tasks[i+1:]...)

		t.Save()
		fmt.Printf("Task deleted successfully (ID: %d)\n", *flag)
	}
}

func(t *Tasks)Mark(flag *int, status string){
	if *flag!=0 && status!=""{
		if status == "in-progress" || status == "done"{
			i := 0
			found := false
			for ;i<len(t.Tasks);i++{
				if t.Tasks[i].Id == *flag{
					found = true
					break
				}
			}
			if !found{
				log.Fatalf("There's no task with ID: %d\n", *flag)
			}
	
			t.Tasks[i].Status = status

			t.Save()
			fmt.Printf("Mark task as %s successfully (ID: %d)\n",status, *flag)
		}
	}
	
}