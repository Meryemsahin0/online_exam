const BASE_URL = "http://localhost:8080/api";

export const apiService = {
  // ÖĞRETMEN: Soru havuzuna yeni soru ekler
  addQuestion: async (questionData) => {
    const response = await fetch(`${BASE_URL}/teacher/questions`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(questionData),
    });
    return response.json();
  },

  // ÖĞRETMEN: Yeni sınav tanımlar (Havuzdan seçilen ID'lerle)
  createExam: async (examData) => {
    const response = await fetch(`${BASE_URL}/teacher/exams`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(examData),
    });
    return response.json();
  },

  // ÖĞRENCİ: Sınav listesini, aktiflik durumunu ve notları çeker
  getStudentExams: async (studentId) => {
    const response = await fetch(`${BASE_URL}/student/exams?student_id=${studentId}`);
    return response.json();
  },
};