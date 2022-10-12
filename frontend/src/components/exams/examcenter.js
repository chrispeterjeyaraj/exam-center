import React from "react";
import "./examcenter.css";
import { Result } from "../result/result";
export class Exams extends React.Component {
  constructor(props) {
    super(props);
    this.questions = [];
    this.timeLeft = 10;
    this.isSubmitted = false;
    this.score = 0;
    this.state = {
      questions: this.questions,
      isSubmitted: this.isSubmitted,
      timeLeft: this.timeLeft,
      isAnswered: false,
    };
    this.setCount();
  }

  componentWillMount() {
    fetch("http://localhost:1234/getQuestions", {
      method: "POST", // *GET, POST, PUT, DELETE, etc.
      mode: "cors", // no-cors, cors, *same-origin
      credentials: "include", // include, *same-origin, omit
      body: JSON.stringify({
        sessionId: localStorage.getItem("sessionId"),
      }),
      headers: {
        "Content-Type": "application/json; charset=utf-8",
      },
    })
      .then((data) => data.json())
      .then((data) => {
        console.log("questions received..", data);
        if (data.isAuth == false) {
          this.props.history.push("/");
        } else {
          this.questions = data.content;
          console.log(data.content);
          this.setState({
            questions: this.questions,
            isSubmitted: this.isSubmitted,
            timeLeft: this.timeLeft,
            isAnswered: false,
          });
        }
      });
  }

  getAnswer(event) {
    console.log("event is", event.target);
    let answer = event.target.value;
    console.log("parent id is", event.target.parentElement.parentElement.id);
    let questionId = event.target.parentElement.parentElement.id;

    for (let i = 0; i < this.state.questions.length; i++) {
      if (this.state.questions[i].questionId == questionId) {
        if (this.state.questions[i].rightAns == answer) {
          console.log("the question is currectly answered..");
          this.state.questions[i].isCurrectlyAnswered = true;
          this.score += this.state.questions[i].difficulty.weight;
          this.state.questions[i].marksAwarded =
            this.state.questions[i].difficulty.weight;
          this.state.questions[i].yourAns = event.target.value;
          this.state.questions[i].isAnswered = true;
        } else {
          console.log("the question is incorrectly answered");

          this.state.questions[i].isCurrectlyAnswered = false;
          this.score += 0;
          this.state.questions[i].marksAwarded = 0;
          this.state.questions[i].yourAns = event.target.value;
          this.state.questions[i].isAnswered = true;
        }

        this.setState({
          questions: this.questions,
          isSubmitted: this.isSubmitted,
          timeLeft: this.timeLeft,
          isAnswered: true,
        });
      }
    }
  }

  submitTest(event) {
    event.target.disabled = true;
    if (!localStorage.getItem("sessionId")) {
      this.props.history.push("/");
      this.isSubmitted = false;
      this.setState({
        questions: this.questions,
        isSubmitted: this.isSubmitted,
        timeLeft: this.timeLeft,
      });
    } else {
      this.isSubmitted = true;
      this.setState({
        questions: this.questions,
        isSubmitted: this.isSubmitted,
        timeLeft: this.timeLeft,
      });
    }
  }

  decrementTime() {
    if (this.state.timeLeft != 0) {
      this.timeLeft -= 1;
      this.setState(
        {
          questions: this.questions,
          isSubmitted: this.isSubmitted,
          timeLeft: this.timeLeft,
        },
        () => {
          this.setState({
            questions: this.questions,
            isSubmitted: this.isSubmitted,
            timeLeft: this.timeLeft,
          });
        }
      );
    } else {
      this.submitTest();
    }
  }

  setCount() {
    setInterval(this.decrementTime(), 1000);
  }

  logout() {
    localStorage.clear();
    this.props.history.push("/");
  }

  render() {
    return (
      <div className="container">
        {this.state.questions.map((question) => {
          return (
            <div
              className="question mb-2"
              id={question.questionId}
              key={question.questionId}
            >
              {question.questionId}
              <p>{question.question}</p>
              <div>
                <input
                  type="radio"
                  name={question.questionId}
                  value={question.options[0].option}
                  onClick={this.getAnswer.bind(this)}
                  disabled={question.isAnswered}
                />
                {question.options[0].option}
              </div>
              <div>
                <input
                  type="radio"
                  name={question.questionId}
                  value={question.options[1].option}
                  onClick={this.getAnswer.bind(this)}
                  disabled={question.isAnswered}
                />
                {question.options[1].option}
              </div>
              <div>
                <input
                  type="radio"
                  name={question.questionId}
                  value={question.options[2].option}
                  onClick={this.getAnswer.bind(this)}
                  disabled={question.isAnswered}
                />
                {question.options[2].option}
              </div>
              <div>
                <input
                  type="radio"
                  name={question.questionId}
                  value={question.options[3].option}
                  onClick={this.getAnswer.bind(this)}
                  disabled={question.isAnswered}
                />
                {question.options[3].option}
              </div>
            </div>
          );
        })}

        <button className="btn btn-danger" onClick={this.submitTest.bind(this)}>
          SUBMIT
        </button>

        <div>
          {this.state.isSubmitted ? (
            <Result data={this.state.questions} score={this.score} />
          ) : (
            <div className="alert alert-warning">
              Click on the Submit button to submit the test anytime
            </div>
          )}
        </div>

        <div>
          <button className="btn btn-primary" onClick={this.logout.bind(this)}>
            LOGOUT
          </button>
        </div>
      </div>
    );
  }
}
