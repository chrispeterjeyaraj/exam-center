# Exam Center
**A simple Online Exam Application with the below features**
* Ability to create exams and questionere as an admin
* Ability to login as a user to take the exam selecting the desired questionere
* Reporting for exams taken with results and scores - admin/ user views

[Tech](#Tech) |
[Basic commands](#Basic commands to run the containers) |
[Database diagram](#Database diagram)

## Tech
- Docker for creating containers for UI, Server and DB
- React for user interface and visualization
- Go for backend APIs
- Mongo for DB

## Basic commands to run the containers

### FrontEnd application
```sh
cd frontend
sudo docker compose -f docker-compose.dev.yml up
in browser -> http://localhost:3000
```

### BackEnd application
```sh
cd backend
sudo docker compose up [--build]
api -> http://localhost:4000
mongo express -> http://localhost:8081
```
### Database diagram

![examcenter-entity-relationship drawio](https://user-images.githubusercontent.com/111428615/195823165-c463d2ce-cf32-4a29-b772-7c4157fba921.png)
