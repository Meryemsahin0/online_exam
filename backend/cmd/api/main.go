package main

import (
	"log"
	"net/http" // Standart Go HTTP paketi
	"online-exam-app/internal/config"
	deliveryHttp "online-exam-app/internal/delivery/http" // Bizim paketimize takma ad verdik
	"online-exam-app/internal/repository"
	"online-exam-app/internal/service"
)

func main() {
	// 1. Veritabanı bağlantısını kur (MongoDB v2)
	db := config.ConnectDB()

	// 2. Katmanları sırasıyla oluştur (Dependency Injection)
	questionRepo := repository.NewMongoQuestionRepository(db)
	examRepo := repository.NewMongoExamRepository(db)

	examService := service.NewExamService(questionRepo, examRepo)
	examHandler := deliveryHttp.NewExamHandler(examService) // Artık hata vermeyecek

	// 3. API Yönlendirmelerini (Routes) Tanımla
	mux := http.NewServeMux()

	// Öğretmen Endpoints
	mux.HandleFunc("/api/teacher/questions", examHandler.AddQuestion)
	mux.HandleFunc("/api/teacher/exams", examHandler.CreateExam)

	// Öğrenci Endpoints
	mux.HandleFunc("/api/student/exams", examHandler.ListExams)

	// CORS Ayarı (Frontend-Backend iletişimi için)
	corsMux := corsMiddleware(mux)

	// 4. Sunucuyu Başlat
	log.Println("Backend sunucusu 8080 portunda çalışıyor...")
	if err := http.ListenAndServe(":8080", corsMux); err != nil {
		log.Fatalf("Sunucu başlatılamadı: %v", err)
	}
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}