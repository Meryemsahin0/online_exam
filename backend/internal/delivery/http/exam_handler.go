package http

import (
    "encoding/json"
    "net/http"
    "online-exam-app/internal/service"
    "time"
)

type ExamHandler struct {
    examService *service.ExamService
}

func NewExamHandler(s *service.ExamService) *ExamHandler {
    return &ExamHandler{examService: s}
}

// POST /api/teacher/questions
// ÖĞRETMEN: Soru havuzuna yeni soru ekler
func (h *ExamHandler) AddQuestion(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Yalnızca POST metodu destekleniyor", http.StatusMethodNotAllowed)
        return
    }

    var req struct {
        Content       string   `json:"content"`
        Options       []string `json:"options"`
        CorrectAnswer string   `json:"correct_answer"`
    }

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Geçersiz JSON formatı", http.StatusBadRequest)
        return
    }

    err := h.examService.AddQuestionToPool(r.Context(), req.Content, req.Options, req.CorrectAnswer)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"message": "Soru havuzuna başarıyla eklendi"})
}

// POST /api/teacher/exams
// ÖĞRETMEN: Havuzdaki sorulardan sınav oluşturur ve zaman atar
func (h *ExamHandler) CreateExam(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Yalnızca POST metodu destekleniyor", http.StatusMethodNotAllowed)
        return
    }

    var req struct {
        Title       string    `json:"title"`
        QuestionIDs []string  `json:"question_ids"`
        StartTime   time.Time `json:"start_time"`
        EndTime     time.Time `json:"end_time"`
    }

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Geçersiz JSON formatı", http.StatusBadRequest)
        return
    }

    err := h.examService.CreateExam(r.Context(), req.Title, req.QuestionIDs, req.StartTime, req.EndTime)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"message": "Sınav başarıyla oluşturuldu ve atandı"})
}

// GET /api/student/exams?student_id=123
// ÖĞRENCİ: Atanan tüm sınavları ve anlık durumlarını (aktiflik, not) görür
func (h *ExamHandler) ListExams(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Yalnızca GET metodu destekleniyor", http.StatusMethodNotAllowed)
        return
    }

    studentID := r.URL.Query().Get("student_id")
    if studentID == "" {
        http.Error(w, "student_id parametresi zorunludur", http.StatusBadRequest)
        return
    }

    exams, err := h.examService.ListExamsForStudent(r.Context(), studentID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(exams)
}