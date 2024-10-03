import { useState, useEffect } from 'react';
import './App.css';

interface feedDataItem {
  title: string;
  description: string;
  link: string;
}

interface feedData {
  title: string;
  description: string;
  feedDataItems: feedDataItem[];
}

function App() {
  const [data, setData] = useState<feedData>();

  useEffect(() => {
    fetch('http://localhost:8080/get_feed') // GolangのAPIエンドポイント
      .then(response => response.json())
      .then((d: feedData) => setData(d))
      .catch(error => console.error('Error:', error));
  }, []);

  return (
    <div>
      <h1>{data?.title || 'Loading...'}</h1>
      <h1>{data?.description || 'Loading...'}</h1>
      <div>
        {(data?.feedDataItems ?? []).map((item, index) => (
          <div key={index}>
            <h2>{index}: {item.title}</h2>
            <p>{item.description}</p>
            <a href={item.link}>{item.link}</a>
          </div>
        ))}
      </div>
    </div>
  );
}

export default App;
