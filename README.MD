# Hangman Project

## Pre-requisites

- Go (1.11)
- Glide
- Node JS (8.1.0)
- Yarn

## API Documentation

After tha application is running the documentation can be found on the following address:
http://localhost:7100/v1/swagger

## Description

The backend of the application was made in GO using one of my boilerplates for API, the application is using internal memory database integrated with GO called Bolt

Instead of using the specific words, I add a Third Party integration getting one random word from the RandomList.com server

## Frontend

The frontend was made in React and can be viewed on the folder **client/src**, was a React Boilerplate chaged to adapt to the Hangman situation.

To run the frontend on development you should run the command **yarn start**, it will run on the port **3000**, but to integrate with the backend, the production folder is deployed into the folder **client/build** which can be accessed by the por **7000** when the backend is running

- The Frontend is on the endpoint http://localhost:7000
- The development frontend is on the endpoint http://localhost:7000

## Backend

The backend was written in go and it is separated by layers, so inside of the **server** folder you will have all the information to expose the middleware but no business rules.

The business rules and application inteligence was implemented on the **application** folder which can be used by the server layer or if it is necessary on the console layer extracting 100% of the application independently

When you run the application will expose the port **7100** and **7000** for the API and client hosting

- The API is on the endpoint http://localhost:7100/v1
- The API documentation is http://localhost:7100/v1/swagger

## Run the application

### Manually

If you want to run the application step by step, after the Pre Requisites are all installed on the Unix environemnt you need to run the following steps

```
export GOPATH=$HOME/Go
mkdir -p $GOPATH/src/github.com/andriolisp/hangman
cp -R ./* $GOPATH/src/github.com/andriolisp/hangman/
cd $GOPATH/src/github.com/andriolisp/hangman/
cd client
yarn && yarn build
cd ..
glide install && go build && ./hangman
```

### Automatically

If you want to run the application step by step, after the Pre Requisites are all installed on the Unix enviroment you need to run the following command

```
make run
```
