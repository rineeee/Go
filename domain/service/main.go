package main

import (
	"context"
	"sample/domain/model"
	
	//"sample/domain/model/user"
	//"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See")
	}
	// context는 하나의 맥락
	// context.Background()와 context.TODO()는 empty context
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	
	coll := client.Database("test").Collection("user")

	/*
	name := "kim harin"
	
	//bson.D : 하나의 BSON 도큐먼트. 순서가 중요한 경우 사용
	//bson.M : 순서가 없는 map 형태. 순서를 유지하지 않는 다는 점을 빼면 D와 같음
	//bson.A : 하나의 BSON array 형태
	//bson.E : D 타입 내부에서 사용하는 하나의 엘리먼트
	var result bson.M
	err = coll.FindOne(context.TODO(), bson.D{{"name", name}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the name %s\n", name)
		return
	}
	if err != nil {
		panic(err)
	}

	//fmt.Print("%s\n", result)
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonData))
   */


	userarray := [2]model.User{}
	userarray[0].Name = "park"
	userarray[0].Age = 14

	userarray[1].Name = "kang"
	userarray[1].Age = 15

	users := []interface{}{}
	users= append(users, userarray[0],userarray[1])

	result, err := coll.InsertMany(context.TODO(), users)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
	

	
}
