import React, { useState, useEffect } from 'react';
import logo from './logo.svg';
import './App.css';

function App() {
  const [data, setData] = useState<string>('');

  useEffect(() => {
    fetch('http://localhost:8080/hello') // GolangのAPIエンドポイント
      .then(response => response.json())
      .then(d => setData(d.message))
      .catch(error => console.error('Error:', error));
  }, []);

  return (
    <div>
      <h1>{data || 'Loading...'}</h1>
    </div>
  )
}

export default App;
