package datawork

import (
	"fmt"
	"time"
)

// DateToDB - приводит дату к формату, который воспринимает БД
func DateToDB(date string) (string, error) {

	var dateToDb string

	dateTime, err := time.Parse("20060102", date)
	if err != nil {
		fmt.Println("Строковые данные даты не корректны. ", err)
		return dateToDb, err
	}

	dateToDb = fmt.Sprint(dateTime.Format("2006-01-02"))

	return dateToDb, nil
}

// DateFromDB - приводит дату к формату, который воспринимает фронтенд
func DateFromDB(date string) (string, error) {

	var deteFromDB string

	dateTime, err := time.Parse("2006-01-02", date[:10])
	if err != nil {
		fmt.Println("Строковые данные даты из БД не корректны. ", err)
		return deteFromDB, err
	}

	deteFromDB = fmt.Sprint(dateTime.Format("20060102"))

	return deteFromDB, nil
}
