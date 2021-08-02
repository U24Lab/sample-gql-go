package mongoDBRepo

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/U24Lab/sample-gql-go/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TodoRepo interface {
	Save(NewTodo *model.Todo)
	GetAll() []*model.Todo
}

type database struct {
	client *mongo.Client
}

const (
	DB         = "tododb"
	COLLECTION = "TODO"
)

func New() TodoRepo {

	MONGODB := "mongodb://dev:Welcome%241@localhost:27017/?authSource=tododb&readPreference=primary&appname=mongodb-vscode%200.6.10&ssl=false"

	clientOption := options.Client().ApplyURI(MONGODB)

	clientOption = clientOption.SetMaxPoolSize(50)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	dbClient, err := mongo.Connect(ctx, clientOption)
	if err != nil {
		fmt.Println("DB Connection Error")
		log.Fatal(err)
	}
	fmt.Println("DB Connection Success!")

	return &database{
		client: dbClient,
	}
}
func (db *database) Save(todo *model.Todo) {
	collection := db.client.Database(DB).Collection(COLLECTION)
	_, err := collection.InsertOne(context.TODO(), todo)
	if err != nil {
		fmt.Println("DB  Error")
		log.Fatal(err)
	}
}

func (db *database) GetAll() []*model.Todo {
	collection := db.client.Database(DB).Collection(COLLECTION)
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println("DB  Error")
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())
	var res []*model.Todo
	for cursor.Next(context.TODO()) {
		var v *model.Todo
		err := cursor.Decode(&v)
		if err != nil {
			log.Fatal(err)
		}
		res = append(res, v)
	}
	return res

}
