import React, { useState } from 'react';
import { apiService } from '../services/api';

export default function TeacherDashboard() {
  // Soru Havuzu Form State
  const [question, setQuestion] = useState({ content: '', options: ['', '', '', ''], correct_answer: '' });
  // Sınav Oluşturma Form State
  const [exam, setExam] = useState({ title: '', question_ids: '', start_time: '', end_time: '' });

  const handleAddQuestion = async (e) => {
    e.preventDefault();
    await apiService.addQuestion(question);
    alert("Soru havuzuna eklendi!");
    setQuestion({ content: '', options: ['', '', '', ''], correct_answer: '' });
  };

  const handleCreateExam = async (e) => {
    e.preventDefault();
    // Virgülle ayrılan ID'leri array'e çeviriyoruz
    const idsArray = exam.question_ids.split(',').map(id => id.trim());
    
    // Tarihleri ISO string formatına (Go'nun time.Time formatına) çeviriyoruz
    const examData = {
      title: exam.title,
      question_ids: idsArray,
      start_time: new Date(exam.start_time).toISOString(),
      end_time: new Date(exam.end_time).toISOString()
    };

    await apiService.createExam(examData);
    alert("Sınav oluşturuldu ve öğrencilere atandı!");
    setExam({ title: '', question_ids: '', start_time: '', end_time: '' });
  };

  return (
    <div style={{ padding: '20px', fontFamily: 'sans-serif' }}>
      <h2>Öğretmen Yönetim Paneli</h2>
      <hr />
      
      {/* 1. SORU HAVUZU FORMU */}
      <h3>1. Soru Havuzuna Soru Ekle</h3>
      <form onSubmit={handleAddQuestion} style={{ display: 'flex', flexDirection: 'column', gap: '10px', maxWidth: '400px' }}>
        <input type="text" placeholder="Soru Metni" value={question.content} onChange={e => setQuestion({...question, content: e.target.value})} required />
        {question.options.map((opt, i) => (
          <input key={i} type="text" placeholder={`Şık ${i+1}`} value={opt} onChange={e => {
            const newOpts = [...question.options];
            newOpts[i] = e.target.value;
            setQuestion({...question, options: newOpts});
          }} required />
        ))}
        <input type="text" placeholder="Doğru Cevap (Örn: Şık 1'deki metin)" value={question.correct_answer} onChange={e => setQuestion({...question, correct_answer: e.target.value})} required />
        <button type="submit">Havuza Soru Kaydet</button>
      </form>

      <br /><hr /><br />

      {/* 2. SINAV OLUŞTURMA FORMU */}
      <h3>2. Yeni Sınav Ata</h3>
      <form onSubmit={handleCreateExam} style={{ display: 'flex', flexDirection: 'column', gap: '10px', maxWidth: '400px' }}>
        <input type="text" placeholder="Sınav Başlığı" value={exam.title} onChange={e => setExam({...exam, title: e.target.value})} required />
        <input type="text" placeholder="Soru ID'leri (Virgülle ayırın: id1, id2)" value={exam.question_ids} onChange={e => setExam({...exam, question_ids: e.target.value})} required />
        <label>Başlangıç Zamanı:</label>
        <input type="datetime-local" value={exam.start_time} onChange={e => setExam({...exam, start_time: e.target.value})} required />
        <label>Bitiş Zamanı:</label>
        <input type="datetime-local" value={exam.end_time} onChange={e => setExam({...exam, end_time: e.target.value})} required />
        <button type="submit" style={{ backgroundColor: 'green', color: 'white' }}>Sınavı Yayınla</button>
      </form>
    </div>
  );
}