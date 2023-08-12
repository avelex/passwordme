import React from 'react';
import './App.css';
import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom';
import PasswordGenerationPage from './Generation';
import Sidebar from './ Sidebar';


function Passwords() {
  return <h2>Passwords Page Soon...</h2>;
}

function App() {
  return (
    <Router>
      <div className="App">
        <Sidebar />
        <div className="separator"></div>
        <main className="content">
          <Routes>
            <Route path="/" element={<PasswordGenerationPage />}></Route>
            <Route path="/passwords" element={<Passwords />}></Route>
          </Routes>
        </main>
      </div>
    </Router>
  );
}

export default App;
