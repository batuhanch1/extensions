package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	concurrency_with_channel "mongoDBBenchmark/pkg/concurrency-with-channel"
	"mongoDBBenchmark/pkg/concurrency-with-channel/httpClient"
	mongo2 "mongoDBBenchmark/pkg/concurrency-with-channel/mongo"
	"mongoDBBenchmark/pkg/mongo-benchmark/product"
	"mongoDBBenchmark/pkg/trace"
	"time"
)

func main() {
	defer trace.Exit(trace.Enter())

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return
	}
	defer client.Disconnect(ctx)
	mongoDb := client.Database("Benchmark")
	//productRepository := product.NewProductRepository(mongoDb)
	//productService := product.NewProductService(productRepository)

	//HttpGetTestExample()
	//HttpPostTestExample()
	MongoInsertTestExample(mongoDb, "Product")
	MongoUpdateTestExample(mongoDb, "Product")
	//MongoReadTestExample(mongoDb, "Product")
	MongoReadManyTestExample(mongoDb, "Product")
	//EmptyTableTestExample(productService)
	//FullTableTestExample(productService)
}

func EmptyTableTestExample(productService product.ProductService) {
	defer trace.Exit(trace.Enter())
	var err error

	for i := 1; i <= 1; i++ {
		err = productService.InsertOneWithEmptyTable()
		if err != nil {
			break
		}
	}
	for i := 1; i <= 1; i++ {
		err = productService.InsertManyWithEmptyTable()
		if err != nil {
			break
		}
	}
}

func FullTableTestExample(productService product.ProductService) {
	defer trace.Exit(trace.Enter())

	var err error

	if err = productService.InsertOneMillionRecord(); err != nil {
		return
	}

	for i := 1; i <= 1; i++ {
		if err = productService.InsertOneWithFullCollection(); err != nil {
			return
		}
	}
	for i := 1; i <= 1; i++ {
		if err = productService.InsertManyWithFullCollection(); err != nil {
			return
		}
	}
}

func HttpGetTestExample() {
	var (
		getList httpClient.GetList
		results concurrency_with_channel.Results
		err     error
	)

	for i := 0; i < 201; i++ {
		get := httpClient.NewGet().SetPath("").SetHost("").SetQueryString("")
		getList = append(getList, *get)
	}

	if results, err = concurrency_with_channel.HttpGET(getList); err != nil {
		return
	}

	for _, result := range results {
		//your model
		var model = new(interface{})
		if err = json.Unmarshal(result, model); err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	return
}

func HttpPostTestExample() {
	var (
		postList httpClient.PostList
		results  concurrency_with_channel.Results
		err      error
	)
	for i := 0; i < 201; i++ {
		post := httpClient.NewPost().SetPath("").SetHost("").SetRequest(nil)
		postList = append(postList, *post)
	}

	if results, err = concurrency_with_channel.HttpPOST(postList); err != nil {
		return
	}
	for _, result := range results {
		//your model
		var model = new(interface{})
		if err = json.Unmarshal(result, model); err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	return
}

func MongoReadTestExample(db *mongo.Database, collectionName string) {
	var (
		readList mongo2.ReadOneList
		results  concurrency_with_channel.Results
		err      error
	)

	for i := 0; i < 201; i++ {
		read := mongo2.NewReadOne().SetFilter(bson.D{{"Name", "ProductName"}}).SetFindOptions(nil)
		readList = append(readList, *read)
	}

	if results, err = concurrency_with_channel.MongoRead(db, collectionName, readList); err != nil {
		return
	}

	for _, result := range results {
		//your model
		var product = new(product.Product)
		if err = json.Unmarshal(result, product); err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Printf("Id: %v\n", product.ID)
	}
	return
}

func MongoInsertTestExample(db *mongo.Database, collectionName string) {
	var insertList mongo2.InsertOneList

	//create insert one list
	for i := 0; i < 201; i++ {
		var product = product.NewProduct()
		insert := mongo2.NewInsertOne().SetData(product).SetInsertOptions(nil)
		insertList = append(insertList, *insert)
	}

	if err := concurrency_with_channel.MongoInsert(db, collectionName, insertList); err != nil {
		fmt.Println(err.Error())
	}

	return
}

func MongoUpdateTestExample(db *mongo.Database, collectionName string) {
	var updateList mongo2.UpdateOneList

	//create update one list
	for i := 0; i < 201; i++ {
		filterQuery := bson.D{{"Rate", 5.0}}
		updateQuery := bson.D{{"$set", bson.D{{"Name", fmt.Sprintf("%s_%v", "ProductName", i)}}}}
		update := mongo2.NewUpdateOne().SetFilter(filterQuery).SetUpdate(updateQuery).SetUpdateOptions(nil)
		updateList = append(updateList, *update)
	}

	if err := concurrency_with_channel.MongoUpdate(db, collectionName, updateList); err != nil {
		fmt.Println(err.Error())
	}

	return
}

func MongoReadManyTestExample(db *mongo.Database, collectionName string) {
	var (
		readList mongo2.ReadManyList
		results  concurrency_with_channel.Results
		err      error
	)

	for i := 0; i < 201; i++ {
		read := mongo2.NewReadMany().SetFilter(bson.D{{"Name", "ProductName"}}).SetFindOptions(nil)
		readList = append(readList, *read)
	}

	if results, err = concurrency_with_channel.MongoReadMany(db, collectionName, readList); err != nil {
		return
	}

	for _, result := range results {
		//your model
		var productList = new(product.ProductList)
		if err = json.Unmarshal(result, productList); err != nil {
			fmt.Println(err.Error())
			return
		}

		for _, product := range *productList {
			fmt.Printf("Id: %v\n", product.ID)
		}
	}
	return
}
