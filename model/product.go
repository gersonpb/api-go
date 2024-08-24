package model

type Product struct {
	ID int 				`json:"id"`
	Name string			`json:"name"`
	Quantidade int16	`json:"amount"`
	Price float64		`json:"price"`
}