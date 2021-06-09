package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type People struct {
	People []Person `xml:"person"`
}

type Person struct {
	Id          string `xml:"id"`
	FirstName   string `xml:"first_name"`
	LastName    string `xml:"last_name"`
	Company     string `xml:"company"`
	Email       string `xml:"email"`
	IpAddress   string `xml:"ip_address"`
	PhoneNumber string `xml:"phone_number"`
}

func main() {
	xmlFile, err := os.Open("people.xml")
	if err != nil {
		fmt.Println(err)
	}

	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var people People
	var mondoData []interface{}

	xml.Unmarshal(byteValue, &people)

	for i := 0; i < len(people.People); i++ {
		mondoData = append(mondoData, people.People[i])
	}

	/* --- DB --- */

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://db"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)

	database := client.Database("exercise1db")
	peopleDatabase := database.Collection("people")

	peopleResult, err := peopleDatabase.InsertMany(ctx, mondoData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Inserted %v documents into people collection!\n", len(peopleResult.InsertedIDs))
}
