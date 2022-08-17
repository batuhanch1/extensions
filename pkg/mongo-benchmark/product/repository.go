package product

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProductRepository interface {
	InsertOne(product Product) error
	InsertMany(productList []Product) error
	Delete() error
	DeleteByName(name string) error
}

type productRepository struct {
	db *mongo.Database
}

func NewProductRepository(db *mongo.Database) ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) InsertOne(product Product) error {
	_, err := r.db.Collection("Product").InsertOne(context.TODO(), product)
	if err != nil {
		return err
	}
	return nil
}

func (r *productRepository) InsertMany(productList []Product) error {
	var products []interface{}

	for _, product := range productList {
		products = append(products, product)
	}

	opts := options.InsertMany().SetOrdered(false)
	_, err := r.db.Collection("Product").InsertMany(context.TODO(), products, opts)
	if err != nil {
		return err
	}
	return nil
}

func (r *productRepository) Delete() error {
	_, err := r.db.Collection("Product").DeleteMany(context.TODO(), bson.D{{}})
	return err
}
func (r *productRepository) DeleteByName(name string) error {
	_, err := r.db.Collection("Product").DeleteMany(context.TODO(), bson.D{{"Name", name}})
	return err
}
