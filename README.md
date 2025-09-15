# ULB Epreuve technique

This project include two parts
* a golang rest api server who's manage all the database interaction and make proxy for the Swagger API
* a SPA (Single Page Application) developed in React situated in the frontend sub-directory

## REST API Server

### build the sources

`go build .`

### Run the REST API Server

`go run main.go`

## SPA

#### The frontend is situated in frontend/ directory

### Install dependencies

`ǹpm install`

### build the sources

`npm run build`

### Run Node dev server

`npm run dev`

## Or you can Run both of them as a container (dont forget to install dependencies and build the sources for the SPA)

`docker-compose up -d `

## Or with Podman

`podman-compose up -d` 

