package links

import (
	"log"

	database "graphServer/db/mydatabase"
	"graphServer/users"
)

//Link is ...
type Link struct {
	ID      string
	Title   string
	Address string
	User    *users.User
}

// Save is ...
func (link Link) Save() int64 {
	//#3
	stmt, err := database.Db.Prepare("INSERT INTO Links(Title,Address,UserID) VALUES($1,$2,$3)")
	if err != nil {
		log.Fatal(err)
	}
	//#4
	res, err := stmt.Exec(link.Title, link.Address, link.User.ID)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Row inserted!")
	r, _ := res.RowsAffected()
	return r
}

//GetAll is..
func GetAll() []Link {
	stmt, err := database.Db.Prepare("SELECT L.id, L.title, L.address,L.UserID, L.Username FROM Links L inner join Users U on L.UserID = U.ID")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var links []Link
	var username string
	var id string
	for rows.Next() {
		var link Link
		err := rows.Scan(&link.ID, &link.Title, &link.Address, &id, &username)
		if err != nil {
			log.Fatal(err)
		}
		link.User = &users.User{
			ID:       id,
			Username: username,
		}
		links = append(links, link)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return links
}
