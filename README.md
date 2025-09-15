# ULB Epreuve technique

## structure

This project include two parts
* a golang rest api server who's manage all the database interaction and make proxy for the Swagger API
* a SPA (Single Page Application) developed in React situated in the frontend sub-directory

## Summary

I didnt have the time to finish all the parts of the test
but i think than do some of all will give you better insights about my skills

* Part 1: my rest api return data from the 3 tables (no crud)
* Part 2: i finished the sub part a (le bulletin in BulletinHandler) but not he part b
* Part 3: You can look at the lessons and inscriptions/bulletins on The SPA

## REST API Server

### build the sources

`go build .`

### Run the REST API Server

`go run main.go`

The REST API Server will be available on http://localhost:8000/api/v1/

#### The routes are

http://localhost:8000/api/v1/liste_inscriptions

http://localhost:8000/api/v1/liste_cours

http://localhost:8000/api/v1/liste_notes

http://localhost:8000/api/v1/etudiants/{matricule}/annee/{annee}/bulletin


## SPA

The frontend is situated in frontend/ directory

### Install dependencies

`ǹpm install`

### build the sources

`npm run build`

### Run Node dev server

`npm run dev`

You can now access the SPA on http://localhost:8000 (via the Go endpoint)

or on http://localhost:5173 (the node test server)

#### SPA routes are 

http://localhost:5173/cours

http://localhost:5173/inscriptions

http://localhost:5173/bulletin/{matricule}/{annee}


## docker or podman container 

It will run it into a container (it was initially setup to add an idp service but did'nt have time to integrate it)

`docker-compose up -d ` or `podman-compose up -d`
http://localhost:8000

[!NOTE]
dont forget to install dependencies and build the sources for the SPA before running containers

[!NOTE]
About AI: I used LeChat from Mistral to generate some more refined templates for the SPA

