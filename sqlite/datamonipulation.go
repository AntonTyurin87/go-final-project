package sqlite

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go_final_project/datawork"
	"io"
	"net/http"

	_ "modernc.org/sqlite"
)

type ErrorType struct {
	Error error `json:"error"`
}

type IDType struct {
	ID int64 `json:"id"`
}

type Storage struct {
	DB *sql.DB
}

func NewStorage(db *sql.DB) Storage {
	return Storage{DB: db}
}

// TaskHandler - обрабатывает GET, POST и DELETE запросы
func (s *Storage) TaskHandler(w http.ResponseWriter, r *http.Request) {
	//Забираем данные для валидации и возвращаем в виде структуры

	var res []byte
	var errRes ErrorType

	httpData, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Не прочитано тело запроса api/task. ", err)
	}

	data, err := datawork.TaskDataValidation(string(httpData))
	if err != nil {
		fmt.Println("Ошибка конвертации входящего значения api/task. ", err)

		errRes.Error = err
		res, err = json.Marshal(errRes)
		if err != nil {
			fmt.Println("Не удалось упаковать ошибку в JSON. ", err)
		}

	} else {

		switch r.Method {
		//Идём писать в базу
		case http.MethodPost:
			res, err = s.TaskDataWrite(data) //sqlite.TaskDataWrite(data)
			if err != nil {
				fmt.Println("Ошибка записи в БД ", err)
			}

		case http.MethodGet:

		case http.MethodDelete:
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(res)
}

// TaskDataWrite - Записывает в БД данные о внесённой задаче
func (s *Storage) TaskDataWrite(data datawork.TaskData) ([]byte, error) {

	var result []byte
	var returnData IDType

	qeryToDB := fmt.Sprintf(`
				INSERT INTO
					scheduler (date, title, comment, repeat)
						VALUES ("%s", "%s", "%s", "%s");`,
		data.Date, data.Title, data.Comment, data.Repeate)

	res, err := s.DB.Exec(qeryToDB)
	if err != nil {
		fmt.Println("Запись в БД не состоялась ", err)
		return result, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		fmt.Println("ID последней записи в БД не удалось получить ", err)
		return result, err
	}

	returnData.ID = id

	result, err = json.Marshal(returnData)
	if err != nil {
		fmt.Println("Не получилось выдать ID последней записи в виде JSON ", err)
		return result, err
	}

	//fmt.Println(string(result)) //TODO Убрать после отработки

	return result, nil
}
