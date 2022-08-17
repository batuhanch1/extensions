package product

import (
	"github.com/google/uuid"
	"math"
)

type ProductList []Product

type Product struct {
	Rate       float64     `bson:"Rate,omitempty"`
	ID         string      `bson:"ID,omitempty"`
	Name       string      `bson:"Name,omitempty"`
	MerchantId string      `bson:"MerchantId,omitempty"`
	Properties *Properties `bson:"Properties,omitempty"`
	Price      *Price      `bson:"Price,omitempty"`
	Stock      *Stock      `bson:"Stock,omitempty"`
}

func NewProduct() Product {
	return Product{
		Rate:       5,
		ID:         uuid.New().String(),
		Name:       "ProductName",
		MerchantId: uuid.New().String(),
		Properties: NewProperties(),
		Price:      NewPrice(),
		Stock:      NewStock(),
	}
}
func NewProductWithName(name string) Product {
	return Product{
		Rate:       5,
		ID:         uuid.New().String(),
		Name:       "name",
		MerchantId: uuid.New().String(),
		Properties: NewProperties(),
		Price:      NewPrice(),
		Stock:      NewStock(),
	}
}

type Properties struct {
	Renk  string
	Beden string
}

func NewProperties() *Properties {
	return &Properties{
		Renk:  "Renk",
		Beden: "Beden",
	}
}

type Price struct {
	SalePrice     float64
	ListPrice     float64
	DiscountPrice float64
}

func NewPrice() *Price {
	return &Price{
		SalePrice:     math.MaxFloat64,
		ListPrice:     math.MaxFloat64,
		DiscountPrice: math.MaxFloat64,
	}
}

type Stock struct {
	ReelStockQuantity    int
	VirtualStockQuantity int
}

func NewStock() *Stock {
	return &Stock{
		ReelStockQuantity:    math.MaxInt,
		VirtualStockQuantity: math.MaxInt,
	}
}
