package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/dzonib/golang-online-tavern/aggregate"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	db       *mongo.Database
	customer *mongo.Collection
}

// mongCUstomer is internal type that is used to store a customer CustomerAggregate inside it
type mongoCustomer struct {
	ID   uuid.UUID `bson:"id"`
	Name string    `bson:"name"`
}

// FromDomainModel function maps the data from the aggregate to the local internal struct
func FromDomainModel(c aggregate.Customer) mongoCustomer {
	return mongoCustomer{
		ID:   c.GetID(),
		Name: c.GetName(),
	}
}

func (m mongoCustomer) ToDomainModel() aggregate.Customer {
	c := aggregate.Customer{}

	c.SetID(m.ID)
	c.SetName(m.Name)

	return c
}

func New(ctx context.Context, connectionString string) (*MongoRepository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, err
	}

	db := client.Database("ddd")

	customers := db.Collection("customers")

	return &MongoRepository{
		db:       db,
		customer: customers,
	}, nil
}

func (mr *MongoRepository) Get(id uuid.UUID) (aggregate.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	result := mr.customer.FindOne(ctx, bson.M{"id": id})

	var c mongoCustomer
	err := result.Decode(&c)
	if err != nil {
		return aggregate.Customer{}, err
	}

	return c.ToDomainModel(), nil
}

func (mr *MongoRepository) Add(c aggregate.Customer) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	internal := FromDomainModel(c)

	_, err := mr.customer.InsertOne(ctx, internal)
	if err != nil {
		return err
	}

	return nil
}

func (mr *MongoRepository) Update(c aggregate.Customer) error {
	panic("TO IMPLEMENT")
}
