GOPATH=${HOME}/Go

create-base:
	mkdir -p ${GOPATH}

create-structure: create-base
	mkdir -p ${GOPATH}/src/github.com/andriolisp/hangman

copy-structure: create-structure
	cp -R ./* ${GOPATH}/src/github.com/andriolisp/hangman/
	cd ${GOPATH}/src/github.com/andriolisp/hangman/

get-dependencies: create-structure
	glide install

client: create-structure
	cd client && yarn && yarn build && cd ..

install: client
	go install

build: client
	go build -o hangman

run: build
	./hangman
