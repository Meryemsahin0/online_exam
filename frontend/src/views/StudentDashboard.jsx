import React, { useEffect, useState } from 'react';
import { apiService } from '../services/api';

export default function StudentDashboard() {
  const [exams, setExams] = useState([]);
  const studentId = "Ogr123"; // Test için sabit öğrenci ID'si

  useEffect(() => {
    loadExams();
  }, []);

  const loadExams = async () => {
    try {
      const data = await apiService.getStudentExams(studentId);
      setExams(data || []);
    } catch (error) {
      console.error("Sınavlar yüklenirken hata oluştu", error);
    }
  };

  return (
    <div style={{ padding: '20px', fontFamily: 'sans-serif' }}>
      <h2>Öğrenci Sınav Sistemi</h2>
      <p>Öğrenci Numarası: <strong>{studentId}</strong></p>
      <hr />
      
      <h3>Atanan Sınavlar Listesi</h3>
      <table border="1" cellPadding="10" style={{ width: '100%', borderCollapse: 'collapse', textAlign: 'left' }}>
        <thead>
          <tr style={{ backgroundColor: '#f2f2f2' }}>
            <th>Sınav Adı</th>
            <th>Başlangıç Tarihi</th>
            <th>Bitiş Tarihi</th>
            <th>Durum / İşlem</th>
            <th>Sınav Notu</th>
          </tr>
        </thead>
        <tbody>
          {exams.map((exam) => (
            <tr key={exam.id}>
              <td><strong>{exam.title}</strong></td>
              <td>{new Date(exam.start_time).toLocaleString()}</td>
              <td>{new Date(exam.end_time).toLocaleString()}</td>
              <td>
                {exam.is_active ? (
                  <button style={{ backgroundColor: 'green', color: 'white', padding: '5px 10px', cursor: 'pointer' }}>
                    Sınava Giriş Yap
                  </button>
                ) : (
                  <button disabled style={{ backgroundColor: '#ccc', color: '#666', padding: '5px 10px', cursor: 'not-allowed' }}>
                    Sınav Aktif Değil
                  </button>
                )}
              </td>
              <td>
                {exam.score !== null ? (
                  <span style={{ color: 'blue', fontWeight: 'bold' }}>{exam.score} / 100</span>
                ) : (
                  <span style={{ color: '#888', fontStyle: 'italic' }}>Henüz Girilmedi/Açıklanmadı</span>
                )}
              </td>
            </tr>
          ))}
          {exams.length === 0 && (
            <tr>
              <td colSpan="5" style={{ textAlign: 'center' }}>Atanmış bir sınav bulunamadı.</td>
            </tr>
          )}
        </tbody>
      </table>
    </div>
  );
}