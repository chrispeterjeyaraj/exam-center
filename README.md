# Exam Center
**A simple Online Exam Application with the below features**
* Ability to create exams and questionere as an admin
* Ability to login as a user to take the exam selecting the desired questionere
* Reporting for exams taken with results and scores - admin/ user views

[![Go](https://github.com/chrispeterjeyaraj/exam-center/actions/workflows/go.yml/badge.svg)](https://github.com/chrispeterjeyaraj/exam-center/actions/workflows/go.yml)[![Node.js CI](https://github.com/chrispeterjeyaraj/exam-center/actions/workflows/node.js.yml/badge.svg)](https://github.com/chrispeterjeyaraj/exam-center/actions/workflows/node.js.yml)[![Exam Center Docker CI](https://github.com/chrispeterjeyaraj/exam-center/actions/workflows/docker-image.yml/badge.svg)](https://github.com/chrispeterjeyaraj/exam-center/actions/workflows/docker-image.yml)

[Tech Stack](#Tech) |
[Pre-requisites](#prerequisites) |
[Basic commands](#Basic-commands-to-run-the-containers) |
[Database diagram](#Database-diagram)

## Tech
- Docker for creating containers for UI, Server and DB
- React for user interface and visualization
- Go for backend APIs
- Mongo for DB

## Prerequisites
- Linux (on windows use WSL2/ others: any linux distribution)
- Docker (LTS)
- NodeJS (LTS)
- MongoDB (LTS)
- Go Compiler (1.19.2)

## Basic commands to run the application

```sh
sudo docker compose up [--build]
api1 -> http://localhost:4000 [non cors]
api2 -> http://localhost:4545 [cors] [use with web application]
mongo express -> http://localhost:8081
UI -> http://localhost:3000
```

## CI/CD Pipeline

- GO CI for API code -> Action Name -> [![Go](https://github.com/chrispeterjeyaraj/exam-center/actions/workflows/go.yml/badge.svg)](https://github.com/chrispeterjeyaraj/exam-center/actions/workflows/go.yml)

- Node CI for UI Code -> [![Node.js CI](https://github.com/chrispeterjeyaraj/exam-center/actions/workflows/node.js.yml/badge.svg)](https://github.com/chrispeterjeyaraj/exam-center/actions/workflows/node.js.yml)

- Docker CI -> [![Exam Center Docker CI](https://github.com/chrispeterjeyaraj/exam-center/actions/workflows/docker-image.yml/badge.svg)](https://github.com/chrispeterjeyaraj/exam-center/actions/workflows/docker-image.yml)


### Database diagram

![examcenter-entity-relationship drawio](https://user-images.githubusercontent.com/111428615/195823165-c463d2ce-cf32-4a29-b772-7c4157fba921.png)
