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
			panic(err)
		}
		return card
	}
	return card
}

func (t_targets *TableTargets) SaveParsingService(tc models.TargetCard) {
	rows, err := t_targets.DB.Query(
		fmt.Sprintf(`INSERT bookrzn.Targets (target_hash,autor,title,price,image,comment,url_source,target_type) 
		VALUES ( '%v','%v','%v', '%v','%v','%v','%v', '%v');`,
			tc.TargetHash, tc.Autor, tc.Title, tc.Price, tc.Link, "comment desk", tc.Source, tc.Tag)) //,
	if err != nil {
		panic(err)
	}
	defer rows.Close()

}

func (t_targets *TableTargets) GetListTargets() []models.TargetCard {
	rows, err := t_targets.DB.Query(`SELECT * FROM bookrzn.Targets;`)
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
			panic(err)
		}
		targetsCard = append(targetsCard, card)
	}
	return targetsCard
}

func (t_targets *TableTargets) GetListTargetsFromToken(token string) []models.TargetCard {
	fmt.Println("start get target from bd")
	rows, err := t_targets.DB.Query(fmt.Sprintf(`SELECT target_hash,count,date,id_order FROM bookrzn.Orders WHERE token = "%s";`, token))
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var targetsCard []models.TargetCard
	for rows.Next() {
		card := models.TargetCard{}
		err := rows.Scan(
			&card.TargetHash,
			&card.Count,
			&card.Date,
			&card.IdOrder,
		)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%v\n\n", card)
		fmt.Println("card", len(targetsCard))
		targetsCard = append(targetsCard, card)
	}

	resTargetsCard := make([]models.TargetCard, 0)

	for _, t_card := range targetsCard {
		rows, err = t_targets.DB.Query(fmt.Sprintf(`SELECT autor,title,image,price FROM bookrzn.Targets WHERE target_hash = "%s";`, t_card.TargetHash))
		if err != nil {
			panic(err)
		}
		defer rows.Close()
		for rows.Next() {
			card := models.TargetCard{}
			err := rows.Scan(
				&card.Autor,
				&card.Title,
				&card.Link,
				&card.Price,
			)
			if err != nil {
				panic(err)
			}
			card.TargetHash = t_card.TargetHash
			card.Count = t_card.Count
			card.Date = t_card.Date
			card.IdOrder = t_card.IdOrder
			fmt.Printf("%v\n\n", card)
			fmt.Println("card", len(targetsCard))
			resTargetsCard = append(resTargetsCard, card)
		}

	}
	return resTargetsCard
}


func (t_targets *TableTargets) GetListTargetsFromTokenHistory(token string) []models.TargetCard {
	fmt.Println("start get target from bd")
	rows, err := t_targets.DB.Query(fmt.Sprintf(`SELECT target_hash,count,date,id_order FROM bookrzn.OrdersHistory WHERE token = "%s";`, token))
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var targetsCard []models.TargetCard
	for rows.Next() {
		card := models.TargetCard{}
		err := rows.Scan(
			&card.TargetHash,
			&card.Count,
			&card.Date,
			&card.IdOrder,
		)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%v\n\n", card)
		fmt.Println("card", len(targetsCard))
		targetsCard = append(targetsCard, card)
	}

	resTargetsCard := make([]models.TargetCard, 0)

	for _, t_card := range targetsCard {
		rows, err = t_targets.DB.Query(fmt.Sprintf(`SELECT autor,title,image,price FROM bookrzn.Targets WHERE target_hash = "%s";`, t_card.TargetHash))
		if err != nil {
			panic(err)
		}
		defer rows.Close()
		for rows.Next() {
			card := models.TargetCard{}
			err := rows.Scan(
				&card.Autor,
				&card.Title,
				&card.Link,
				&card.Price,
			)
			if err != nil {
				panic(err)
			}
			card.TargetHash = t_card.TargetHash
			card.Count = t_card.Count
			card.Date = t_card.Date
			card.IdOrder = t_card.IdOrder
			fmt.Printf("%v\n\n", card)
			fmt.Println("card", len(targetsCard))
			resTargetsCard = append(resTargetsCard, card)
		}

	}
	return resTargetsCard
}
