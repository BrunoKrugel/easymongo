// Package easymongo provides a simple and easy-to-use wrapper for MongoDB operations in Go.
package easymongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoInstance represents a MongoDB connection instance.
type MongoInstance struct {
	mClient     *mongo.Client
	mCollection *mongo.Collection
}

var singletonInstance *MongoInstance

// New creates and returns a new instance of MongoInstance.
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

// NewStatic returns a singleton instance of MongoInstance.
func NewStatic(uri string, db string, collection string) *MongoInstance {
	if singletonInstance == nil {
		singletonInstance = New(uri, db, collection)
	}
	return singletonInstance
}

// Close disconnects the MongoDB client.
func (m *MongoInstance) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	m.mClient.Disconnect(ctx)
}

// InsertOne inserts a single document into the MongoDB collection.
func (m *MongoInstance) InsertOne(document interface{}) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	return m.mCollection.InsertOne(ctx, document)
}

// InsertMany inserts multiple documents into the MongoDB collection.
func (m *MongoInstance) InsertMany(documents []interface{}) (*mongo.InsertManyResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	return m.mCollection.InsertMany(ctx, documents)
}

// DeleteOne deletes a single document from the MongoDB collection based on the specified filter.
func (m *MongoInstance) DeleteOne(filter bson.D) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	return m.mCollection.DeleteOne(ctx, filter)
}

// UpdateOne updates a single document in the MongoDB collection based on the specified filter and update.
func (m *MongoInstance) UpdateOne(filter bson.D, update bson.D) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	options := options.Update().SetUpsert(true)
	return m.mCollection.UpdateOne(ctx, filter, update, options)
}

// FindOne retrieves a single document from the MongoDB collection based on the specified filter.
func (m *MongoInstance) FindOne(filter bson.D) *mongo.SingleResult {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	return m.mCollection.FindOne(ctx, filter)
}
