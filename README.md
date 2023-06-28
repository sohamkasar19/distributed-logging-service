
# Distributed Logging Service

A high-performance distributed logging service built with Go and MongoDB. It provides a RESTful API that supports CRUD operations for log entries.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)

## Features

- Built with Go and utilizes GoLang's powerful concurrency features.
- RESTful API that supports Create, Retrieve, Update, Delete (CRUD) operations for log entries.
- Uses MongoDB for persistent storage.
- Log aggregation and filtering feature using Go routines and channels.
- Dockerized for easy setup and distribution.

## Installation

This project requires Go and MongoDB to be installed on your machine.

1. Clone the repository
git clone https://github.com/<Your Github Username>/Distributed-Logging-Service.git

2. Navigate to the project directory
cd Distributed-Logging-Service


3. Install the dependencies
go mod download

4. Build the docker image
docker build -t go-logging-service .


5. Run the docker container
docker run -p 8080:8080 go-logging-service

## Usage

### Endpoints:

- POST `/logs`: Create a new log entry
- GET `/logs/{id}`: Retrieve a specific log entry
- PUT `/logs/{id}`: Update a specific log entry
- DELETE `/logs/{id}`: Delete a specific log entry

### Sample Payload:

{
    "id": "1234",
    "source": "source1",
    "message": "This is a test log message",
    "timestamp": "2023-06-30T10:30:00Z"
}



## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
