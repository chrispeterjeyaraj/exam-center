# OnlineExam
Online Exam App using React and Node

This is a full stack Mern Application using REST APIs, with React and Node.

## How to run ?

1. clone this repository 
2. go inside the backend folder and hit npm install
3. Similarly go inside the onlinetest folder and hit npm install

## Flow
The first page of the app is the landing page of logging in the student. A sessionChecker middleware is made which checks the cookie of the incoming request and provides proper sessions management.

Proper error and use case handling has been done to prevent any abnormal behaviour and maintain robustness.

If logged in successfully, the sessionId is saved in the localStorage to check and authenticate the user everytime he tries to visit an API. Then he enters a test where he gets 13 questions and answers them by submitting the test.

All the questions are dynamic as they straightaway come from the databse.

The App is made with performance in mind , in terms of connectivity and internet, as minimal use of backend queries is made.
Maximum work is done at the frontend to provide a better experiance.

The users submits the test and can see his report of correct and incorrect answers with the marks awarded for each answer according to the difficulty and tags(stored and maintainded in the database).

he can also see the aread which he needs to work on.

## Dependencies used
1. React-dom
2. React Router
3. Mongoose
4. Expressjs
5. Express sessions
6. SessionStore for saving the sessions in databse
7. MongoDB as the database
8. nodemon
9. bodyparser
10. mlab
