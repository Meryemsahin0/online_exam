package config

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func ConnectDB() *mongo.Database {
	// Eğer canlı ortamda (Render) MONGO_URI tanımlıysa onu al, yoksa lokalde çalışması için 127.0.0.1 kullan
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://127.0.0.1:27017"
	}

	clientOptions := options.Client().ApplyURI(mongoURI)
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(clientOptions)
	if err != nil {
		log.Fatalf("MongoDB bağlantı hatası: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("MongoDB Ping hatası: %v", err)
	}

	log.Println("MongoDB'ye başarıyla bağlanıldı!")
	return client.Database("online_exam_db")
}