import React, { useState } from 'react';
import TeacherDashboard from './views/TeacherDashboard';
import StudentDashboard from './views/StudentDashboard';

export default function App() {
  // Hocaya sunum yaparken ekranlar arası hızlıca geçiş yapabilmek için basit bir state
  const [view, setView] = useState('teacher');

  return (
    <div style={{ fontFamily: 'sans-serif' }}>
      {/* Üst Navigasyon Çubuğu */}
      <nav style={{ backgroundColor: '#2c3e50', padding: '15px', color: 'white', display: 'flex', gap: '20px' }}>
        <span style={{ fontWeight: 'bold', marginRight: 'auto' }}>Ölçme ve Değerlendirme Sistemi</span>
        <button 
          onClick={() => setView('teacher')} 
          style={{ backgroundColor: view === 'teacher' ? '#3498db' : 'transparent', color: 'white', border: '1px solid white', padding: '5px 10px', cursor: 'pointer' }}
        >
          Öğretmen Paneli
        </button>
        <button 
          onClick={() => setView('student')} 
          style={{ backgroundColor: view === 'student' ? '#3498db' : 'transparent', color: 'white', border: '1px solid white', padding: '5px 10px', cursor: 'pointer' }}
        >
          Öğrenci Paneli
        </button>
      </nav>

      {/* Dinamik Ekran İçeriği */}
      <div style={{ padding: '20px' }}>
        {view === 'teacher' ? <TeacherDashboard /> : <StudentDashboard />}
      </div>
    </div>
  );
}