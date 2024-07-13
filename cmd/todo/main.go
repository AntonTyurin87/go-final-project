package main

import (
	"fmt"

	"go_final_project/handlers"
	"go_final_project/sqlite"
	"net/http"
	"os"
)

func main() {

	port := os.Getenv("TODO_PORT")
	todoDB := os.Getenv("TODO_DBFILE")

	if port == "" {
		port = ":7540"
	}

	if todoDB == "" {
		todoDB = "./"
	}

	//Проверить существует ли файл БД. Если его нет, то создать БД.
	err := sqlite.FindOrCreateDB(todoDB)
	if err != nil {
		fmt.Println("Ошибка с базой данных ", err)
	}

	//Запускаем Web интерфейс
	http.Handle("/", http.FileServer(http.Dir("./web")))

	http.HandleFunc("/api/nextdate", handlers.NextDateHandler)

	//Запускаем сервер
	fmt.Printf("Сервер TODO запущен! Порт %s.", port)
	err = http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("Ошибка при запуске сервера: %s", err.Error())
	}
}

/*
http.HandlerFunc("/", func(w http.ResponseWriter, r *http.Request)) {\
	}

		// запускаем сервер
	fmt.Printf("Сервер TODO запущен! Порт %s.", port)
	http.ListenAndServe(port, nil)

func startWeb(w http.ResponseWriter, r *http.Request) {

	http.FileServer(http.Dir("./web"))
}

*/
