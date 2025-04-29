import React from 'react';
import logo from './logo.svg';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import './App.css';
import './Downloads'
import DownloadPage from './Downloads';

function App() {
  return (
    <Router>
      <div>
        <Route path="/" element={<DownloadPage />} />
      </div>
    </Router>
  );
}

export default App;
