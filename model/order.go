package model

type Order struct {
	CustomerId string `bson:"customerid"`
	Amount     int32  `bson:"amount"`
	Product    []Product
}
type Product struct {
	ProductName string  `json:"productName"`
	Quantity    int32   `json:"quantity"`
	Price       float32 `json:"price"`
	Amount      int32   `json:"amount"`
	Description string  `json:"description"`
}
