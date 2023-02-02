# FROM golang:alpine as builder

# # ENV GO111MODULE=on

# # Add Maintainer info
# LABEL maintainer="Anirudh Agnihotri <codewithanirudhagni@gmail.com>"

# # Install git.
# # Git is required for fetching the dependencies.
# RUN apk update && apk add --no-cache git

# # Set the current working directory inside the container 
# WORKDIR /app

# # Copy go mod and sum files 
# COPY go.mod go.sum ./

# RUN go mod tidy

# # Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
# RUN go mod download 

# # Copy the source from the current directory to the working Directory inside the container 
# COPY . .

# EXPOSE 8080

# #Command to run the executable
# CMD ["go","run","start.go"]


FROM golang:1.19-alpine

RUN apk update && apk upgrade && \
  apk add --no-cache bash git openssh

LABEL maintainer="Anirudh Agnihotri <codewithanirudhagni@gmail.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
RUN go mod tidy

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]