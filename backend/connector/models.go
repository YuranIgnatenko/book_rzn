package connector

type DataRow interface{}

type DataRowTargets struct {
	// Id int
	TargetHash string
}

type DataRowFavorites struct {
	// Id int
	TargetHash string
}

type DataRowOrders struct {
	// Id int
	TargetHash string
}

type DataRowOrdersHistory struct {
	// Id int
	TargetHash string
}
