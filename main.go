package main

import (
	_ "errors"
	"io"
	"log"
	"os"

	"gamch1k/todo/database"
	env_manager "gamch1k/todo/envmanager"
	"gamch1k/todo/server"
)


func main() {
	log_file, err := os.OpenFile("logs/logs.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v\n", err)
	}
	defer log_file.Close()
	
	mw := io.MultiWriter(os.Stdout, log_file)
	log.SetOutput(mw)

	log.Println("----------- PROGRAM STARTED -----------")

	database.InitDatabase()

	// database.InsertTask(db_path, "Some test text")

	// fmt.Println(database.GetTasks(db_path))

	// database.UpdateTaskDone(db_path, 4, false)

	// fmt.Println(database.CheckTaskExist(db_path, 2))
	// fmt.Println(database.CheckTaskExist(db_path, 25))
	// database.DeleteTaskById(db_path, 7)

	server.Start(env_manager.GetEnvVariable("HOST_NAME") + ":" + env_manager.GetEnvVariable("PORT"))
	
}

