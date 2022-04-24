package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"os/exec"
	"time"
)

type friends struct {
	id      int
	name    string
	surname string
	weight  int
}

func main() {
	db, err := sql.Open("sqlite3", "friends.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
	choice := 0
	for {
		switch choice {
		case 0:
			time.Sleep(2 * time.Second)
			cmd := exec.Command("cmd", "/c", "cls")
			cmd.Stdout = os.Stdout
			cmd.Run()
			fmt.Println("Что вы хотите сделать?")
			fmt.Println("1 - Добавить друга")
			fmt.Println("2 - Удалить друга")
			fmt.Println("3 - Вывести друзей")
			fmt.Println("4 - Изменить данные друзей")
			fmt.Println("5 - Удалить все данные")
			fmt.Println("10 - Закончить работу")
			fmt.Scan(&choice)
		case 1:
			length, err := db.Query("SELECT COUNT(*) FROM friends")
			if err != nil {
				panic(err)
			}
			var count int
			for length.Next() {
				if err := length.Scan(&count); err != nil {
					panic(err)
				}
			}
			var name, surname string
			var weight int
			fmt.Println("Введите имя")
			fmt.Scan(&name)
			fmt.Println("Введите фамилию")
			fmt.Scan(&surname)
			fmt.Println("Введите вес")
			fmt.Scan(&weight)
			result, err := db.Exec("insert into friends (id, name, surname, weight) values ($1, $2, $3, $4)", count, name, surname, weight)
			if err != nil {
				panic(err)
			}
			fmt.Println(result.RowsAffected())
			choice = 0
		case 2:
			var name, surname string
			fmt.Println("Введите имя и фамилию друга, которого хотите удалить")
			fmt.Scan(&name)
			fmt.Scan(&surname)
			result, err := db.Exec("delete from friends where name = $1 and surname= $2", name, surname)
			if err != nil {
				panic(err)
			}
			fmt.Println(result.RowsAffected())
			choice = 0
		case 3:
			rows, err := db.Query("select * from friends")
			if err != nil {
				panic(err)
			}
			defer rows.Close()
			friend := []friends{}

			for rows.Next() {
				f := friends{}
				err := rows.Scan(&f.id, &f.name, &f.surname, &f.weight)
				if err != nil {
					fmt.Println(err)
					continue
				}
				friend = append(friend, f)
			}
			for _, f := range friend {
				fmt.Println(f.id, f.name, f.surname, f.weight)
			}
			choice = 0
		case 4:
			var change string
			var id int
			fmt.Println("Введите id кого хотите изменить")
			fmt.Scan(&id)
			fmt.Println("Что вы хотите изменить?")
			fmt.Scan(&change)
			if change == "weight" || change == "вес" {
				fmt.Println("Введите вес")
				var weight int
				fmt.Scan(&weight)
				result, err := db.Exec("update friends set weight = $1 where id = $2", weight, id)
				if err != nil {
					panic(err)
				}
				fmt.Println(result.RowsAffected())
			}
			if change == "name" || change == "имя" {
				var newName string
				fmt.Println("Введите новое имя")
				fmt.Scan(&newName)
				result, err := db.Exec("update friends set name = $1 where id = $2 ", newName, id)
				if err != nil {
					panic(err)
				}
				fmt.Println(result.RowsAffected())
			}
			if change == "surname" || change == "фамилию" {
				fmt.Println("Введите новую фамилию")
				var newSurname string
				fmt.Scan(&newSurname)
				result, err := db.Exec("update friends set surname = $1 where id = $2", newSurname, id)
				if err != nil {
					panic(err)
				}
				fmt.Println(result.RowsAffected())
			}
			choice = 0
		case 5:
			result, err := db.Exec("delete from friends")
			if err != nil {
				panic(err)
			}
			fmt.Println(result.RowsAffected())
			choice = 0
		case 10:
			fmt.Println("Спасибо за уделённое время")
			return
		default:
			fmt.Println("Такой команды нет, введите реализованную команду")
			choice = 0
		}
	}
}
