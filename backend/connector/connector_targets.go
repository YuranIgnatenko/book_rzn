package connector

import (
	"backend/models"
	"database/sql"
	"fmt"
)

type ConnectorTargets struct {
}

func NewConnectorTargets() *ConnectorTargets {
	return &ConnectorTargets{}
}

func (conn *Connector) SaveTargetTargets(targethash, autor, title, price, image string) {

	db, err := sql.Open("mysql", conn.dsn())
	if err != nil {
		fmt.Printf("Error %s when opening DB\n", err)
	}
	conn.Db = db

	rows, err := conn.Db.Query(
		fmt.Sprintf(`INSERT bookrzn.Targets (target_hash,autor,title,price,image,comment) 
		VALUES ( '%v','%v','%v', '%v','%v','%v');`,
			targethash, autor, title, price, image, "comment desk")) //,
	if err != nil {
		panic(err)
	}
	//
	// обязательно иначе привысит лимит подключений и будет сбой
	defer rows.Close()

}

func (conn *Connector) GetListTargets() []models.TargetCard {
	db, err := sql.Open("mysql", conn.dsn())
	if err != nil {
		fmt.Printf("Error %s when opening DB\n", err)

	}
	conn.Db = db

	rows, err := conn.Db.Query(fmt.Sprintf(`SELECT * FROM bookrzn.Targets;`)) //,
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var targetsCard []models.TargetCard

	for rows.Next() {
		card := models.TargetCard{}
		err := rows.Scan(
			&card.Id,
			&card.TargetHash,
			&card.Autor,
			&card.Title,
			&card.Price,
			&card.Link,
			&card.Comment,
		)

		if err != nil {
			fmt.Println(err)
			continue
		}
		targetsCard = append(targetsCard, card)
	}

	return targetsCard
}
