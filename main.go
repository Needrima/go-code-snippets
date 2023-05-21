package main

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Comment struct {
	ID primitive.ObjectID `bson:"_id"`
	Name string `bson:"name"`
	Comment string `bson:"comment"`
}

func main() {
	// loading env variables
	if err := godotenv.Load("app.env"); err != nil {
		log.Fatal("loading environmental variables:", err)
	}

	bgCtx := context.Background()
	// database connection
	ctx, cancle := context.WithTimeout(bgCtx, time.Second * 10)
	defer cancle()

	clientOption := options.Client().ApplyURI(os.Getenv("DB_CONN_STRING"))
	client, err := mongo.Connect(ctx, clientOption)
	if err != nil {
		log.Fatal("connecting to database:", err)
	}
	collection := client.Database("my-comments-app").Collection("comments")

	// getting html files as templates
	tpl := template.Must(template.ParseGlob("./html/*.html"))

	// routing
	router := chi.NewRouter()
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		cur, err := collection.Find(bgCtx, bson.M{})
		if err != nil {
			log.Println("getting comments:", err)
			http.Error(w, "something went wrong", 500)
			return
		}

		comments := []Comment{}
		if err := cur.All(bgCtx, &comments); err != nil {
			log.Println("serializing comments:", err)
			http.Error(w, "something went wrong", 500)
			return
		}

		tpl.ExecuteTemplate(w, "index.html", comments)
	})

	router.Post("/", func(w http.ResponseWriter, r *http.Request) {
		newComment := Comment{
			ID: primitive.NewObjectID(),
			Name: r.FormValue("name"),
			Comment: r.FormValue("comment"),
		}

		if _, err := collection.InsertOne(bgCtx, newComment); err != nil {
			log.Println("inserting comment into database:", err)
			http.Error(w, "something went wrong", 500)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// starting server
	log.Println("serving on port:", port)
	http.ListenAndServe(":"+port, router)
}
