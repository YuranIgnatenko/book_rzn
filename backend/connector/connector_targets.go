package connector

import (
	"backend/models"
	"database/sql"
	"fmt"
)

type TableTargets struct {
	DB *sql.DB
}

func NewTableTargets() *TableOrders {
	return &TableOrders{}
}

// FROM bookrzn.Targets TO ...
func (t_targets *TableTargets) GetTargetsCardsFromHash(target_hash string) models.TargetCard {
	rows, err := t_targets.DB.Query(fmt.Sprintf(`SELECT * FROM bookrzn.Targets WHERE target_hash = '%s';`, target_hash)) //,
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	card := models.TargetCard{}

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
			&card.Source,
			&card.Tag,
		)

		if err != nil {
			continue
		}
		return card
	}
	return card
}

func (conn *Connector) SaveParsingService(tc models.TargetCard) {
	rows, err := conn.Db.Query(
		fmt.Sprintf(`INSERT bookrzn.Targets (target_hash,autor,title,price,image,comment,url_source,target_type) 
		VALUES ( '%v','%v','%v', '%v','%v','%v','%v', '%v');`,
			tc.TargetHash, tc.Autor, tc.Title, tc.Price, tc.Link, "comment desk", tc.Source, tc.Tag)) //,
	if err != nil {
		panic(err)
	}
	defer rows.Close()

}

func (conn *Connector) GetListTargets() []models.TargetCard {
	rows, err := conn.Db.Query(`SELECT * FROM bookrzn.Targets;`)
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
			&card.Source,
			&card.Tag,
		)
		if err != nil {
			continue
		}
		targetsCard = append(targetsCard, card)
	}
	return targetsCard
}

func (conn *Connector) GetListTargetsFromToken(token string) []models.TargetCard {
	fmt.Println("start get target from bd")
	rows, err := conn.Db.Query(fmt.Sprintf(`SELECT * FROM bookrzn.Orders WHERE token='%s';`, token))
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
			&card.Source,
			&card.Tag,
		)
		if err != nil {
			continue
		}
		fmt.Printf("%v\n\n", card)
		fmt.Println("card", len(targetsCard))
		targetsCard = append(targetsCard, card)
	}
	return targetsCard
}
