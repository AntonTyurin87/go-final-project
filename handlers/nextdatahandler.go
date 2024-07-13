package handlers

import (
	"fmt"
	"net/http"
	"time"

	"go_final_project/datawork"
)

// NextDateHandler - возвращает значение новой даты, если оно валидно.
func NextDateHandler(w http.ResponseWriter, r *http.Request) {
	now := r.FormValue("now")
	date := r.FormValue("date")
	repeat := r.FormValue("repeat")

	nowTime, err := time.Parse("20240229", now)
	if err != nil {
		fmt.Println("Ошибка конвертации входящего времени nowTime. ", err)
	}

	res, err := datawork.NextDate(nowTime, date, repeat)
	if err != nil {
		fmt.Println("Ошибка получения NextDate. ", err)
	}

	fmt.Fprint(w, res)

	//fmt.Fprintf(w, "Время сейчас %s, дата %s, повторы %s", nowTime, date, repeat)
}
