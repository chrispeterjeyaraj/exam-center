import React, { Component } from 'react';
import Login from './components/login/login';
import { Routes, Route } from 'react-router-dom'
import { Exams } from './components/exams/examcenter';
import { Result } from './components/result/result';

class App extends Component {
  render() {
    return (
      <div>
        <Routes>
          <Route path="/exams" element=<Exams /> />
          <Route path="" element=<Login/> />
          <Route path="result" element=<Result /> />
        </Routes>
      </div>
    );
  }
}

export default App;
