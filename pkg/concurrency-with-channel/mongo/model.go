package mongo

import (
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ReadOneList []ReadOne

type ReadOne struct {
	opts   *options.FindOneOptions
	filter interface{}
}

func NewReadOne() *ReadOne {
	return &ReadOne{}
}

func (r *ReadOne) SetFindOptions(opts *options.FindOneOptions) *ReadOne {
	r.opts = opts

	return r
}

func (r *ReadOne) SetFilter(filter interface{}) *ReadOne {
	r.filter = filter

	return r
}

type InsertOneList []InsertOne

type InsertOne struct {
	opts *options.InsertOneOptions
	data interface{}
}

func NewInsertOne() *InsertOne {
	return &InsertOne{}
}

func (r *InsertOne) SetInsertOptions(opts *options.InsertOneOptions) *InsertOne {
	r.opts = opts

	return r
}

func (r *InsertOne) SetData(data interface{}) *InsertOne {
	r.data = data

	return r
}

type UpdateOneList []UpdateOne

type UpdateOne struct {
	opts   *options.UpdateOptions
	filter interface{}
	update interface{}
}

func NewUpdateOne() *UpdateOne {
	return &UpdateOne{}
}

func (r *UpdateOne) SetUpdateOptions(opts *options.UpdateOptions) *UpdateOne {
	r.opts = opts

	return r
}

func (r *UpdateOne) SetFilter(filter interface{}) *UpdateOne {
	r.filter = filter

	return r
}
func (r *UpdateOne) SetUpdate(update interface{}) *UpdateOne {
	r.update = update

	return r
}

type ReadManyList []ReadMany

type ReadMany struct {
	opts   *options.FindOptions
	filter interface{}
}

func NewReadMany() *ReadMany {
	return &ReadMany{}
}

func (r *ReadMany) SetFindOptions(opts *options.FindOptions) *ReadMany {
	r.opts = opts

	return r
}

func (r *ReadMany) SetFilter(filter interface{}) *ReadMany {
	r.filter = filter

	return r
}

type UpdateManyList []UpdateMany

type UpdateMany struct {
	opts   *options.UpdateOptions
	filter interface{}
	update interface{}
}

func NewUpdateMany() *UpdateMany {
	return &UpdateMany{}
}

func (r *UpdateMany) SetUpdateOptions(opts *options.UpdateOptions) *UpdateMany {
	r.opts = opts

	return r
}

func (r *UpdateMany) SetFilter(filter interface{}) *UpdateMany {
	r.filter = filter

	return r
}

func (r *UpdateMany) SetUpdate(update interface{}) *UpdateMany {
	r.update = update

	return r
}

type InsertManyList []InsertMany

type InsertMany struct {
	opts *options.InsertManyOptions
	data []interface{}
}

func NewInsertMany() *InsertOne {
	return &InsertOne{}
}

func (r *InsertMany) SetInsertOptions(opts *options.InsertManyOptions) *InsertMany {
	r.opts = opts

	return r
}

func (r *InsertMany) SetData(data []interface{}) *InsertMany {
	r.data = data

	return r
}

type AggregateList []Aggregate

type Aggregate struct {
	opts     *options.AggregateOptions
	pipeline interface{}
}

func NewAggregate() *Aggregate {
	return &Aggregate{}
}

func (r *Aggregate) SetAggregateOptions(opts *options.AggregateOptions) *Aggregate {
	r.opts = opts

	return r
}

func (r *Aggregate) SetPipeline(pipeline interface{}) *Aggregate {
	r.pipeline = pipeline

	return r
}
