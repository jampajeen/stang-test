package db

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/jampajeen/stang-test/monitor-service/model"
)

const COLLECTION_NAME = "tx_records"

type MongoDb struct {
	client *mongo.Client
	db     *mongo.Database
}

func NewMongoDb(url string, database string) (*MongoDb, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		return nil, err
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	db := client.Database(database)

	return &MongoDb{client: client, db: db}, nil
}

func (m *MongoDb) InsertTransactionRecord(doc model.TransactionRecord) (objectId string, err error) {
	db := m.db

	now := time.Now().UTC()
	doc.CreatedAt = now.String()

	collection := db.Collection(COLLECTION_NAME)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := collection.InsertOne(ctx, doc)
	if err != nil {
		log.Println(err.Error())
		return
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		objectId = oid.Hex()
	} else {
		err = errors.New("result is not ObjectID")
	}

	return
}

func (m *MongoDb) QueryTransactionRecord(id string) (records []model.TransactionRecord, err error) {
	db := m.db

	collection := db.Collection(COLLECTION_NAME)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	filter := bson.D{
		{"$or", []interface{}{
			bson.D{{"txfrom", id}},
			bson.D{{"txto", id}},
		}},
	}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var result []model.TransactionRecord
	if err = cursor.All(ctx, &result); err != nil {
		log.Println(err.Error())
		return nil, err
	}
	records = result

	return
}
