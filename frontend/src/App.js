import React, { Component } from 'react';
import SignIn from './components/auth/signin/signin';
import SignUp from './components/auth/signup/signup';
import { Routes, Route } from 'react-router-dom'
import { Exams } from './components/exams/examcenter';
import { Result } from './components/result/result';

class App extends Component {
  render() {
    return (
      <div>
        <Routes>
          <Route path="/exams" element=<Exams /> />
          <Route path="" element=<SignIn /> />
          <Route path="/signup" element=<SignUp/> />
          <Route path="result" element=<Result /> />
        </Routes>
      </div>
    );
  }
}

export default App;
