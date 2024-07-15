package datawork

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func NextDate(now time.Time, date string, repeat string) (string, error) {

	var result string
	var nextDate time.Time

	//Парсит время для старта повторений
	dateStart, err := time.Parse("20060102", date)
	if err != nil {
		fmt.Println("Не верные данные date. ", err)
		return result, err
	}

	// Обрабатывает данные для повторений
	switch string(repeat[0]) {
	//Для года
	case "y":
		if yearParser(repeat) == nil {
			nextDate = dateStart.AddDate(1, 0, 0)
			for now.After(nextDate) {
				nextDate = nextDate.AddDate(1, 0, 0)
			}
			result = nextDate.String()
		} else {
			return result, yearParser(repeat)
		}
	//Для дней
	case "d":
		days, err := dayParser(repeat)
		if err != nil {
			return result, err
		}

		nextDate = dateStart.AddDate(0, 0, days)
		for now.After(nextDate) {
			nextDate = nextDate.AddDate(0, 0, days)
		}
		result = nextDate.String()

		//Для недель
	case "w":
		//TODO Описать логику поиска ближайшего дня недели.
	}

	return result, nil
}

// yearParser - проверяет корректность данных для повторения каждый год
func yearParser(repeat string) error {
	if repeat == "y" {
		return nil
	}
	err1 := errors.New("значение для повторений не корректно")
	return err1
}

// dayParser - проверяет корректность данных для повторения каждые несколько дней
func dayParser(repeat string) (int, error) {
	var days int

	dayData := strings.Split(repeat, " ")

	if len(dayData) < 2 {
		err2 := errors.New("не указан интервал в дняхй")
		fmt.Println(err2)
		return days, err2
	}

	dayCount, err := strconv.Atoi(dayData[1])
	if err != nil {
		fmt.Println("Неверный формат дней для повторений. ", err)
		return days, err
	}

	if dayCount > 400 {
		err3 := errors.New("превышен максимально допустимый интервал")
		fmt.Println(err3)
		return days, err3
	}

	days = dayCount

	return days, nil
}

// - проверяет корректность данных для назначения повторений на день недели
func weekParser(repeat string) ([]time.Weekday, error) {

	weekMap := map[int]time.Weekday{1: time.Monday, 2: time.Tuesday, 3: time.Wednesday, 4: time.Thursday, 5: time.Friday, 6: time.Saturday, 7: time.Sunday}
	var weekDay []time.Weekday
	var dayNumber int
	var err error

	weekDays := strings.Split(repeat, " ")

	//Проверка на наличие дня недели
	if len(weekDays) < 2 {
		err4 := errors.New("не верный день недели")
		fmt.Println(err4)
		return weekDay, err4
	}

	//Проверка на наличие одного дня недели
	if len(weekDays[1]) == 1 {
		dayNumber, err = strconv.Atoi(weekDays[1])

		if err != nil {
			fmt.Println("не верное значение дня недели", err)
			return weekDay, err
		}

		if 0 >= dayNumber || dayNumber >= 8 {
			err5 := errors.New("не верный день недели")
			fmt.Println(err5)
			return weekDay, err5
		}

		weekDay = append(weekDay, weekMap[dayNumber])

		//Если дней не один
	} else {

		for _, value := range strings.Split(weekDays[1], ",") {
			day, err := strconv.Atoi(value)

			if err != nil {
				fmt.Println("не верное значение дня недели", err)
				return weekDay, err
			}

			if 0 >= day || day >= 8 {
				err6 := errors.New("не верный день недели")
				fmt.Println(err6)
				return weekDay, err6
			}

			weekDay = append(weekDay, weekMap[day])
		}
	}

	return weekDay, nil
}

/*

	nextDate = dateStart
	for now > nextDate {
		nextDate = dateStart.AddDate(1, 0, 0)
	}

	func repeatParser(now time.Time, repeat string) (time.Duration, error) {

	var err error

}
*/
