package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PeopleResource struct{}

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

func (rs PeopleResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", rs.List)

	r.Route("/{id}", func(r chi.Router) {
		r.Use(PersonCtx)
		r.Get("/", rs.Get)
	})

	return r
}

func (rs PeopleResource) List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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
	peopleCollection := database.Collection("people")

	cursor, err := peopleCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	var people []bson.M
	if err = cursor.All(ctx, &people); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(people)
}

func PersonCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "id", chi.URLParam(r, "id"))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (rs PeopleResource) Get(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id").(string)
	w.Header().Set("Content-Type", "application/json")
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
	peopleCollection := database.Collection("people")

	var person bson.M
	if err = peopleCollection.FindOne(ctx, bson.M{"id": id}).Decode(&person); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(person)
}
