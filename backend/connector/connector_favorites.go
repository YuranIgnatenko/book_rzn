package connector

import "database/sql"

type TableFavorites struct {
	DB *sql.DB
}

func NewTableFavorites() *TableOrders {
	return &TableOrders{}
}

func (t_orders *TableFavorites) Add() {}
func (t_orders *TableFavorites) Get() {}
