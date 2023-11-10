package types

type TransactionType string

var TransactionTypes = struct {
	Asset          TransactionType
	Entertainment  TransactionType
	Food           TransactionType
	Grocery        TransactionType
	Health         TransactionType
	Income         TransactionType
	Rent           TransactionType
	Subscriptions  TransactionType
	Transportation TransactionType
}{
	Asset:          "asset",
	Entertainment:  "entertainment",
	Food:           "food",
	Grocery:        "grocery",
	Health:         "health",
	Income:         "income",
	Rent:           "rent",
	Subscriptions:  "subscriptions",
	Transportation: "transportation",
}