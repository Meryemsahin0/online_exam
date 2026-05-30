package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func ConnectDB() *mongo.Database {
	// localhost yerine direkt 127.0.0.1 yazarak IPv6 karmaşasını çözüyoruz
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")
	
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

	log.Println("MongoDB'ye başarıyla bağlanıldı! (v2 Sürücüsü)")
	return client.Database("online_exam_db")
}