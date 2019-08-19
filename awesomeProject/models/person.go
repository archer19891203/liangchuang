package models
import (
	db "awesomeProject/databases"
	"log"
)
//定义person类型结构
type Person struct {
	Id        int    `json:"id"`
	Username string `json:"username"`
	Tel  string `json:"tel"`
}

func (p *Person) AddPerson() (id int64, err error) {
	rs, err := db.SqlDB.Exec("INSERT INTO wp_member(username, tel) VALUES (?, ?)", p.Username, p.Tel)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	return
}

func (p *Person) GetPersons() (persons []Person, err error) {
	persons = make([]Person, 0)
	rows, err := db.SqlDB.Query("SELECT id, username, tel FROM wp_member limit 10")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var person Person
		rows.Scan(&person.Id, &person.Username, &person.Tel)
		persons = append(persons, person)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}

func (p *Person) GetPerson() (person Person, err error) {
	err = db.SqlDB.QueryRow("SELECT id, username, tel FROM wp_member WHERE id=?", p.Id).Scan(
		&person.Id, &person.Username, &person.Tel,
	)
	return
}

func (p *Person) ModPerson() (ra int64, err error) {
	stmt, err := db.SqlDB.Prepare("UPDATE wp_member SET username=?, tel=? WHERE id=?")
	defer stmt.Close()
	if err != nil {
		return
	}
	rs, err := stmt.Exec(p.Username, p.Tel, p.Id)
	if err != nil {
		return
	}
	ra, err = rs.RowsAffected()
	return
}

func (p *Person) DelPerson() (ra int64, err error) {
	rs, err := db.SqlDB.Exec("DELETE FROM wp_member WHERE id=?", p.Id)
	if err != nil {
		log.Fatalln(err)
	}
	ra, err = rs.RowsAffected()
	return
}