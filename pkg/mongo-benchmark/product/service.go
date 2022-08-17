package product

import (
	"mongoDBBenchmark/pkg/trace"
)

const TestOneRowRecord = 1000
const TestTwoRowRecord = 10000

//const TestThreeRowRecord = 100000

var TestScenarioList = []int{
	TestOneRowRecord, TestTwoRowRecord, //TestThreeRowRecord,
}

type ProductService interface {
	InsertOneWithEmptyTable() (err error)
	InsertManyWithEmptyTable() (err error)
	InsertOneWithFullCollection() (err error)
	InsertManyWithFullCollection() (err error)
	InsertOneMillionRecord() (err error)
}

type productService struct {
	ProductRepository ProductRepository
}

func NewProductService(productRepository ProductRepository) ProductService {
	return &productService{productRepository}
}

func (r *productService) InsertOneWithEmptyTable() (err error) {
	defer trace.Exit(trace.Enter())

	for _, testScenario := range TestScenarioList {

		productList := createMockData(testScenario)

		for _, product := range productList {
			if err = r.ProductRepository.InsertOne(product); err != nil {
				return err
			}
		}

		if err = r.ProductRepository.Delete(); err != nil {
			return err
		}
	}
	return err
}
func (r *productService) InsertManyWithEmptyTable() (err error) {
	defer trace.Exit(trace.Enter())

	for _, testScenario := range TestScenarioList {

		productList := createMockData(testScenario)

		if err = r.ProductRepository.InsertMany(productList); err != nil {
			return err
		}

		if err = r.ProductRepository.Delete(); err != nil {
			return err
		}
	}
	return err
}

func (r *productService) InsertOneWithFullCollection() (err error) {
	defer trace.Exit(trace.Enter())

	for _, testScenario := range TestScenarioList {

		productList := createMockData(testScenario)

		for _, product := range productList {
			if err = r.ProductRepository.InsertOne(product); err != nil {
				return err
			}
		}

		if err = r.ProductRepository.DeleteByName("ProductName"); err != nil {
			return err
		}
	}
	return err
}
func (r *productService) InsertManyWithFullCollection() (err error) {
	defer trace.Exit(trace.Enter())

	for _, testScenario := range TestScenarioList {

		productList := createMockData(testScenario)

		if err = r.ProductRepository.InsertMany(productList); err != nil {
			return err
		}

		if err = r.ProductRepository.DeleteByName("ProductName"); err != nil {
			return err
		}
	}

	return err
}

func (r *productService) InsertOneMillionRecord() (err error) {
	defer trace.Exit(trace.Enter())

	var productList ProductList

	for i := 0; i < 1000000; i++ {
		productList = append(productList, NewProductWithName("onemillion"))
	}

	if err = r.ProductRepository.InsertMany(productList); err != nil {
		return err
	}
	return err
}

func createMockData(testScenario int) ProductList {
	defer trace.Exit(trace.Enter())
	var productList ProductList

	for i := 0; i < testScenario; i++ {
		productList = append(productList, NewProduct())
	}

	return productList
}
