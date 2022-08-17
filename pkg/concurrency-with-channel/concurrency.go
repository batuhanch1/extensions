package concurrency_with_channel

import (
	"fmt"
	mongo2 "go.mongodb.org/mongo-driver/mongo"
	"mongoDBBenchmark/pkg/concurrency-with-channel/httpClient"
	"mongoDBBenchmark/pkg/concurrency-with-channel/mongo"
)

func catchError(chanResult chan<- *ChanResult) {
	if r := recover(); r != nil {
		err, ok := r.(error)
		if !ok {
			err = fmt.Errorf("%v", r)
		}
		chanResult <- newChanResult().setError(err)
	}
}
func receiveResultFromChannel(chanResult chan *ChanResult, channelLimit int) (results Results, err error) {
	for i := 0; i < channelLimit; i++ {
		result := <-chanResult
		if result.Err != nil {
			return nil, result.Err
		}

		results = append(results, result.Result)
	}
	return results, nil
}

func HttpGET(getList httpClient.GetList) (results Results, err error) {
	var (
		getResultChan = make(chan *ChanResult)
		channelCount  = 0
	)

	client := httpClient.NewHttpClient()

	for _, get := range getList {
		var isChannelLimitOver bool
		channelCount++

		go func(outGet httpClient.Get, out chan<- *ChanResult) {
			defer catchError(out)
			response, responseErr := client.GET(outGet)
			out <- newChanResult().setResult(response).setError(responseErr)
		}(get, getResultChan)

		if isChannelLimitOver, err = results.CheckChannelLimit(channelLimit, channelCount, getResultChan); err != nil {
			return nil, err
		}

		if isChannelLimitOver {
			channelCount = 0
		}
	}
	if _, err = results.CheckChannelLimit(0, channelCount, getResultChan); err != nil {
		return nil, err
	}
	return results, nil
}
func HttpPOST(postList httpClient.PostList) (results Results, err error) {
	var (
		getResultChan = make(chan *ChanResult)
		channelCount  = 0
	)

	client := httpClient.NewHttpClient()

	for _, post := range postList {
		var isChannelLimitOver bool
		channelCount++

		go func(outPost httpClient.Post, out chan<- *ChanResult) {
			defer catchError(out)
			response, responseErr := client.POST(outPost)
			out <- newChanResult().setResult(response).setError(responseErr)
		}(post, getResultChan)

		if isChannelLimitOver, err = results.CheckChannelLimit(channelLimit, channelCount, getResultChan); err != nil {
			return nil, err
		}

		if isChannelLimitOver {
			channelCount = 0
		}
	}

	if _, err = results.CheckChannelLimit(0, channelCount, getResultChan); err != nil {
		return nil, err
	}
	return results, nil
}

func MongoRead(db *mongo2.Database, collectionName string, readList mongo.ReadOneList) (results Results, err error) {
	var (
		getResultChan = make(chan *ChanResult)
		channelCount  = 0
	)

	mongoRepository := mongo.NewMongoRepository(db)

	for _, read := range readList {
		var isChannelLimitOver bool
		channelCount++

		go func(outRead mongo.ReadOne, out chan<- *ChanResult) {

			defer catchError(out)
			response, responseErr := mongoRepository.ReadOne(collectionName, outRead)
			out <- newChanResult().setResult(response).setError(responseErr)

		}(read, getResultChan)

		if isChannelLimitOver, err = results.CheckChannelLimit(channelLimit, channelCount, getResultChan); err != nil {
			return nil, err
		}

		if isChannelLimitOver {
			channelCount = 0
		}
	}
	if _, err = results.CheckChannelLimit(0, channelCount, getResultChan); err != nil {
		return nil, err
	}
	return results, nil
}
func MongoUpdate(db *mongo2.Database, collectionName string, updateList mongo.UpdateOneList) (err error) {
	var (
		getResultChan = make(chan *ChanResult)
		channelCount  = 0
		results       Results
	)

	mongoRepository := mongo.NewMongoRepository(db)

	for _, update := range updateList {
		var isChannelLimitOver bool
		channelCount++

		go func(outUpdate mongo.UpdateOne, out chan<- *ChanResult) {

			defer catchError(out)
			responseErr := mongoRepository.UpdateOne(collectionName, outUpdate)
			out <- newChanResult().setError(responseErr)

		}(update, getResultChan)

		if isChannelLimitOver, err = results.CheckChannelLimit(channelLimit, channelCount, getResultChan); err != nil {
			return
		}

		if isChannelLimitOver {
			channelCount = 0
		}
	}

	if _, err = results.CheckChannelLimit(0, channelCount, getResultChan); err != nil {
		return
	}
	return
}
func MongoInsert(db *mongo2.Database, collectionName string, insertList mongo.InsertOneList) (err error) {
	var (
		getResultChan = make(chan *ChanResult)
		channelCount  = 0
		results       Results
	)

	mongoRepository := mongo.NewMongoRepository(db)

	for _, insert := range insertList {
		var isChannelLimitOver bool
		channelCount++

		go func(outInsert mongo.InsertOne, out chan<- *ChanResult) {

			defer catchError(out)
			responseErr := mongoRepository.InsertOne(collectionName, outInsert)
			out <- newChanResult().setError(responseErr)

		}(insert, getResultChan)

		if isChannelLimitOver, err = results.CheckChannelLimit(channelLimit, channelCount, getResultChan); err != nil {
			return
		}

		if isChannelLimitOver {
			channelCount = 0
		}
	}

	if _, err = results.CheckChannelLimit(0, channelCount, getResultChan); err != nil {
		return
	}
	return
}

func MongoReadMany(db *mongo2.Database, collectionName string, readManyList mongo.ReadManyList) (results Results, err error) {
	var (
		getResultChan = make(chan *ChanResult)
		channelCount  = 0
	)

	mongoRepository := mongo.NewMongoRepository(db)

	for _, readMany := range readManyList {
		var isChannelLimitOver bool
		channelCount++

		go func(outRead mongo.ReadMany, out chan<- *ChanResult) {

			defer catchError(out)
			response, responseErr := mongoRepository.ReadMany(collectionName, outRead)
			out <- newChanResult().setResult(response).setError(responseErr)

		}(readMany, getResultChan)

		if isChannelLimitOver, err = results.CheckChannelLimit(channelLimit, channelCount, getResultChan); err != nil {
			return nil, err
		}

		if isChannelLimitOver {
			channelCount = 0
		}
	}
	if _, err = results.CheckChannelLimit(0, channelCount, getResultChan); err != nil {
		return nil, err
	}
	return results, nil
}
func MongoUpdateMany(db *mongo2.Database, collectionName string, updateManyList mongo.UpdateManyList) (err error) {
	var (
		getResultChan = make(chan *ChanResult)
		channelCount  = 0
		results       Results
	)

	mongoRepository := mongo.NewMongoRepository(db)

	for _, updateMany := range updateManyList {
		var isChannelLimitOver bool
		channelCount++

		go func(outUpdate mongo.UpdateMany, out chan<- *ChanResult) {

			defer catchError(out)
			responseErr := mongoRepository.UpdateMany(collectionName, outUpdate)
			out <- newChanResult().setError(responseErr)

		}(updateMany, getResultChan)

		if isChannelLimitOver, err = results.CheckChannelLimit(channelLimit, channelCount, getResultChan); err != nil {
			return
		}

		if isChannelLimitOver {
			channelCount = 0
		}
	}

	if _, err = results.CheckChannelLimit(0, channelCount, getResultChan); err != nil {
		return
	}
	return
}
func MongoInsertMany(db *mongo2.Database, collectionName string, insertManyList mongo.InsertManyList) (results Results, err error) {
	var (
		getResultChan = make(chan *ChanResult)
		channelCount  = 0
	)

	mongoRepository := mongo.NewMongoRepository(db)

	for _, insertMany := range insertManyList {
		var isChannelLimitOver bool
		channelCount++

		go func(outRead mongo.InsertMany, out chan<- *ChanResult) {

			defer catchError(out)
			responseErr := mongoRepository.InsertMany(collectionName, outRead)
			out <- newChanResult().setError(responseErr)

		}(insertMany, getResultChan)

		if isChannelLimitOver, err = results.CheckChannelLimit(channelLimit, channelCount, getResultChan); err != nil {
			return nil, err
		}

		if isChannelLimitOver {
			channelCount = 0
		}
	}
	if _, err = results.CheckChannelLimit(0, channelCount, getResultChan); err != nil {
		return nil, err
	}
	return results, nil
}

func MongoAggregate(db *mongo2.Database, collectionName string, aggregateList mongo.AggregateList) (results Results, err error) {
	var (
		getResultChan = make(chan *ChanResult)
		channelCount  = 0
	)

	mongoRepository := mongo.NewMongoRepository(db)

	for _, aggregate := range aggregateList {
		var isChannelLimitOver bool
		channelCount++

		go func(outRead mongo.Aggregate, out chan<- *ChanResult) {

			defer catchError(out)
			response, responseErr := mongoRepository.Aggregate(collectionName, aggregate)
			out <- newChanResult().setResult(response).setError(responseErr)

		}(aggregate, getResultChan)

		if isChannelLimitOver, err = results.CheckChannelLimit(channelLimit, channelCount, getResultChan); err != nil {
			return nil, err
		}

		if isChannelLimitOver {
			channelCount = 0
		}
	}
	if _, err = results.CheckChannelLimit(0, channelCount, getResultChan); err != nil {
		return nil, err
	}
	return results, nil
}
