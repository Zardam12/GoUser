package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//User Model
type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name,omitempty"`
	Dob         string             `bson:"dob,omitempty"`
	Address     string             `bson:"address,omitempty"`
	Description string             `bson:"description,omitempty"`
	CreateAt    time.Time          `bson:"createAt,omitempty"`
}

// Monogo Atlas Address
const connectionString = "mongodb+srv://Zardam:kakashisorrow@cluster0.grhn5fj.mongodb.net/?retryWrites=true&w=majority"
const dbName = "DbZardam"
const colName = "users"

var collection *mongo.Collection

// making connection
func init() {
	clientOption := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database(dbName).Collection(colName) //creating dataBase and collection
	fmt.Println("connection successfull")
}

func addOneUser(user User) {
	inserted, err := collection.InsertOne(context.Background(), user) // inserting user in mongodb
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("inserted with id:", inserted.InsertedID)
}

// deleteUser
func deleteUser(name string) {
	fmt.Println(name)
	count, err := collection.DeleteMany(context.Background(), bson.M{"name": name})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*count)
}
func getAlluser() {
	cur, err := collection.Find(context.Background(), bson.D{{}}) // getting all the user
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		var user bson.M
		if err := cur.Decode(&user); err != nil {
			log.Fatal(err)
		}
		fmt.Println(user)
	}
}

func main() {
	fmt.Print("enter the no of users: ")
	m := input()
	n, err := strconv.Atoi(m)
	if err != nil {

	}
	// inserting users
	for i := 0; i < n; i++ {
		u := inputUser()
		addOneUser(u) // adding user to database
		fmt.Println()
	}

	//delete user
	fmt.Print("do you want to delete user:(yes/no) ")
	if flag := input(); flag == "yes" {
		fmt.Println("enter the user name which you want to delete: ")
		name := input()
		deleteUser(name)
	}

	// Query user
	getAlluser()
}

func input() string {
	sc := bufio.NewReader(os.Stdin)
	sc.Reset(os.Stdin)
	val, err := sc.ReadString('\n')
	if err != nil {
		log.Fatal()
	}
	return val[:len(val)-2]
}

func inputUser() User {
	fmt.Print("Enter the name: ")
	name := input()
	fmt.Print("Enter the dob: ")
	dob := input()
	fmt.Print("Enter the address: ")
	address := input()
	fmt.Print("Enter the description: ")
	description := input()
	return User{
		Name:        name,
		Dob:         dob,
		Address:     address,
		Description: description,
		CreateAt:    time.Now(),
	}
}
