# Hello there! Welcome to _Stocks excercise_
This is a project whose purpose is to showcase the integration of different technologies to solve a specific set of requirements.

## Requirements
- Connect to external API
- Handle errors
- Store data in database
- Send data to frontend
- Data search & filter options for user

## Technologies
### Frontend
- Vue 3 (with Pinia)
- Typescript
- TailwindCSS

### Backend & DB
- Go (with Gin)
- CockroachDB

## Project structure
### Frontend
For simplicity, I followed the standard App -> View -> Component arquitecture of Vue.
### Backend
I used Gin because it includes things like CORS, routing, and JSON binding "out of the box". It's also pretty flexible so I could implemment a "Django like" architecture to improve readability.

## Instalation & How to run...
### Prerequisites
- Make sure you have a node version >= 20. The latest LTS version available at the time of developing this project was 22.19.0
    - You can download node here: [Node.js](https://nodejs.org/es)
- Install [Go](https://go.dev/)
### How to run
```bash
#Start go server
go run cmd/server/main.go

#start Vue app
npm run dev

#MAKE SURE YOU ARE IN DE CORRECT DIRECTORY.