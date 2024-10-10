package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DbCOnnerction *sql.DB

type Person struct {
	Name string
	Age  int
}

func main() {
	DbConnection, _ := sql.Open("sqlite3", "./example.sql")
	defer DbConnection.Close()
	cmd := `CREATE TABLE IF NOT EXISTS person(
	name STRING,
	age INT)`
	_, err := DbConnection.Exec(cmd)
	if err != nil {
		log.Fatal(err)
	}

	// cmd = `INSERT INTO person (name, age) VALUES (?, ?)`
	// _, err = DbConnection.Exec(cmd, "Nancy", 20)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// cmd = "UPDATE person SET age = ? WHERE name = ?"
	// _, err = DbConnection.Exec(cmd, 25, "Mike")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// マルチプルセレクト
	fmt.Println("マルチプルセレクト")
	cmd = "SELECT * FROM person"
	rows, err := DbConnection.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()
	var pp []Person

	for rows.Next() {
		var p Person
		// 1行ずつ取得してエラーハンドリング
		err := rows.Scan(&p.Name, &p.Age)
		if err != nil {
			log.Fatalln(err)
		}
		pp = append(pp, p)
	}
	// rowsをまとめてエラーハンドリング
	err = rows.Err()
	if err != nil {
		log.Fatalln(err)
	}

	for _, p := range pp {
		fmt.Println(p)
	}

	// シングルセレクト
	fmt.Println("シングルセレクト")
	cmd = "SELECT * FROM person where age = ?"
	row := DbConnection.QueryRow(cmd, 25)
	var p Person
	err = row.Scan(&p.Name, &p.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No row")
		} else {
			log.Println(err)
		}
	}
	fmt.Println(p.Name, p.Age)

	/*
		result
		マルチプルセレクト
		{Nancy 20}
		{Mike 25}
		{Nancy 20}
		{Nancy 20}
		シングルセレクト
		Mike 25
	*/

	cmd = "DELETE FROM person WHERE name = ?"
	_, err = DbConnection.Exec(cmd, "Nancy")
	if err != nil {
		log.Fatalln(err)
	}
	/*
		result
		sqlite> select * from person;
		Nancy|20
		Mike|25
		Nancy|20

		sqlite> select * from person;
		Mike|25
	*/

	// SQLインジェクションをユーザーが出来てしまう例
	/*
		sqlite> select * from person;
		Mike|25
		Mr.X|100
	*/
	tableName := "person; INSERT INTO person (name, age) VALUES ('Mr.X', 100);"
	cmd = fmt.Sprintf("SELECT * FROM %s", tableName)
	rows, _ = DbConnection.Query(cmd)
	defer rows.Close()

	var pp2 []Person
	for rows.Next() {
		var p Person
		err := rows.Scan(&p.Name, &p.Age)
		if err != nil {
			log.Fatalln(err)
		}
		pp2 = append(pp2, p)
	}
	err = rows.Err()
	if err != nil {
		log.Fatalln(err)
	}

	for _, p := range pp2 {
		fmt.Println(p)
	}
}
