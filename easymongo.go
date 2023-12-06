package easymongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	mClient     *mongo.Client
	mCollection *mongo.Collection
}

var singletonInstance *MongoInstance

func New(uri string, db string, collection string) *MongoInstance {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPIOptions).SetMaxPoolSize(200)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println(err)
		return &MongoInstance{}
	}
	return &MongoInstance{
		mClient:     client,
		mCollection: client.Database(db).Collection(collection),
	}
}

func NewStatic(uri string, db string, collection string) *MongoInstance {
	if singletonInstance == nil {
		singletonInstance = New(uri, db, collection)
	}
	return singletonInstance
}

func (m *MongoInstance) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	m.mClient.Disconnect(ctx)
}

func (m *MongoInstance) InsertOne(document interface{}) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	return m.mCollection.InsertOne(ctx, document)
}

func (m *MongoInstance) InsertMany(documents []interface{}) (*mongo.InsertManyResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	return m.mCollection.InsertMany(ctx, documents)
}

func (m *MongoInstance) DeleteOne(filter bson.D) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	return m.mCollection.DeleteOne(ctx, filter)
}

func (m *MongoInstance) UpdateOne(filter bson.D, update bson.D) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	options := options.Update().SetUpsert(true)
	return m.mCollection.UpdateOne(ctx, filter, update, options)
}

func (m *MongoInstance) FindOne(filter bson.D) *mongo.SingleResult {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	return m.mCollection.FindOne(ctx, filter)
}
