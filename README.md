# REST-API-BookCatalog-User-Gin

This is an simple REST API. This API build with Go Language with gin framework.
The purpose of this API is to learn to make a REST API.

## Prerequisite

- Go version 1.16.5
- Run command go mod tidy to install dependency

## Project structure

```
REST-API-BookCatalog
└── config          # load any config at .env.example
└── entity          # seems like model
└── handler         # seems like controller
└── transport       # standarization request from client and response data
|   ├── request
|   ├── response
└── usecase         # manage apps rule/flow
└── repository      # manage any parameter from usecase into query in database
└── server          # routing directory
main.go
```

## Schemas Database

Create a database with the schema below before running the project.

### Tabel 'tbl_users'

| Column   | Type         | Constraint  |
| -------- | ------------ | ----------- |
| id       | serial       | PRIMARY KEY |
| username | varchar(100) | -           |
| email    | varchar(100) | -           |

**Query**

```
CREATE TABLE tbl_users(
    user_id INT AUTO_INCREMENT,
    username VARCHAR(100),
    email VARCHAR(100),
    PRIMARY KEY(user_id)
);
```

### Tabel 'tbl_books'

| Column       | Type         | Constraint  |
| ------------ | ------------ | ----------- |
| book_id      | serial       | PRIMARY KEY |
| book_name    | varchar(100) | -           |
| book_creator | varchar(100) | -           |

**Query**

```
CREATE TABLE tbl_books(
    book_id INT AUTO_INCREMENT,
    book_name VARCHAR(100),
    book_creator
    VARCHAR(100),
    PRIMARY KEY(user_id)
);
```

## How to run

- Clone this repository
- Please run `cp .env.example .env` at your terminal
- Add your environment variable in file `.env`
- Run go run main.go
- Run page localhost:8080/{endpoint} in your postman

## API Documentation

import file [postman_collection.json](./simple-restapi-go.postman_collection.json) in Postman.
