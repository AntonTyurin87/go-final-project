package main

import (
	"fmt"

	"go_final_project/handlers"
	"go_final_project/sqlite"
	"net/http"
	"os"

	_ "modernc.org/sqlite"
)

func main() {

	port := os.Getenv("TODO_PORT")
	todoDB := os.Getenv("TODO_DBFILE")

	if port == "" {
		port = ":7540"
	}

	//TODO Поменять или убрать после отладки
	if todoDB == "" {
		todoDB = "/home/anton/go_final_project/scheduler.db" ///home/anton/go_final_project/scheduler.db
	}

	//Проверить существует ли файл БД. Если его нет, то создать БД.
	dbURL, err := sqlite.FindOrCreateDB(todoDB)
	if err != nil {
		fmt.Println("Ошибка с базой данных ", err)
	}

	//Подключаемся к БД
	db, err := sqlite.InitDB(dbURL)
	if err != nil {
		fmt.Println("Ошибка инициализации БД ", err)
	}

	fmt.Println(dbURL) //TODO Убрать после отработки

	storage := sqlite.NewStorage(db)

	//Запускаем Web интерфейс
	http.Handle("/", http.FileServer(http.Dir("./web")))

	//Выводим значение новой даты
	http.HandleFunc("/api/nextdate", handlers.NextDateHandler)

	//Работаем с задачами
	http.HandleFunc("/api/task", storage.TaskHandler)

	//Запускаем сервер
	fmt.Printf("Сервер TODO запущен! Порт %s.\n", port)
	err = http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("Ошибка при запуске сервера: %s.\n", err.Error())
	}
}
