package sqlite

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// FindOrCreateDB - ищет файл базы данных в папке запуска приложения.
func FindOrCreateDB(todoDB string) error {

	//Если переменная окружения не задани или пуста присвоем адрес текущего каталога
	if todoDB == "" {
		appPath, err := os.Executable()
		if err != nil {
			log.Fatal(err)
		}

		dbFile := filepath.Join(filepath.Dir(appPath), "scheduler.db")
		_, err = os.Stat(dbFile)

		if err != nil {
			if err = CreateDB(dbFile); err != nil {
				fmt.Println("Не удалось создать БД", err)
				return err
			}
		}
	}

	return nil
}
