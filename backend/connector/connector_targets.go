package connector

import (
	"backend/models"
	"fmt"
)

func (conn *Connector) SaveParsingService(tc models.TargetCard) {
	// fmt.Println(`(autor,title,price,image,comment,url_source,target_type)`)
	// fmt.Println(tc.Autor, tc.Title, tc.Price, tc.Link, "comment desk", tc.Source, tc.Tag)
	// fmt.Println()

	rows, err := conn.Db.Query(
		fmt.Sprintf(`INSERT bookrzn.Targets (target_hash,autor,title,price,image,comment,url_source,target_type) 
		VALUES ( '%v','%v','%v', '%v','%v','%v','%v', '%v');`,
			tc.TargetHash, tc.Autor, tc.Title, tc.Price, tc.Link, "comment desk", tc.Source, tc.Tag)) //,
	if err != nil {
		panic(err)
	}
	//
	// обязательно иначе привысит лимит подключений и будет сбой
	defer rows.Close()

}

func (conn *Connector) GetListTargets() []models.TargetCard {

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
			&card.Source,
			&card.Tag,
		)

		if err != nil {
			fmt.Println(err)
			continue
		}
		targetsCard = append(targetsCard, card)
	}

	return targetsCard
}
