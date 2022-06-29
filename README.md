# book-rest-api
REST API developed using Go, Gin, GORM and MySQL that allows a user to create, read, update and delete books from a library

## Requirements
Go v1.18+  
MySQL v8.0+

## Installation
After cloning the repository, in the root directory run the following command:
```go
go get .
```

Go to the `db.go` file inside the data folder and replace the connection information with the settings of your MySQL instance

In the root directory run the command:
```go
go run main.go
```

## How to Use

`GET /api/books` Returns a list of books  

**Response:**
```
[
    {
        "ID": 1,
        "CreatedAt": "2022-06-22T15:26:50.406-03:00",
        "UpdatedAt": "2022-06-22T15:26:50.406-03:00",
        "DeletedAt": null,
        "title": "Shoe Dog",
        "author": "Phil Knight",
        "quantity": 5
    },
    {
        "ID": 2,
        "CreatedAt": "2022-06-22T19:35:55.662-03:00",
        "UpdatedAt": "2022-06-22T19:35:55.662-03:00",
        "DeletedAt": null,
        "title": "The 48 Laws of Power",
        "author": "Robert Greene",
        "quantity": 7
    }
]
```

`GET /api/book/{id}` Returns the book with the specified id  

**Response:**
```
{
    "ID": 1,
    "CreatedAt": "2022-06-22T15:26:50.406-03:00",
    "UpdatedAt": "2022-06-22T15:26:50.406-03:00",
    "DeletedAt": null,
    "title": "Shoe Dog",
    "author": "Phil Knight",
    "quantity": 5
}
```

`POST /api/book` Creates a new book and returns the created book  

**Request body:**
```
{
    "title": "Think and Grow Rich",
    "author": "Napoleon Hill",
    "quantity": 9
}
```

**Response:**
```
{
    "ID": 3,
    "CreatedAt": "2022-06-29T10:12:40.013-03:00",
    "UpdatedAt": "2022-06-29T10:12:40.013-03:00",
    "DeletedAt": null,
    "title": "Think and Grow Rich",
    "author": "Napoleon Hill",
    "quantity": 9
}
```

`PUT /api/book/{id}` Updates the book with the specified id  

**Request body:**
```
{
    "title": "Rich Dad Poor Dad",
    "author": "Robert Kiyosaki",
    "quantity": 9
}
```

**Response:**
```
204 No Content
```

`DELETE /api/book/{id}` Deletes the book with the specified id  

**Response:**
```
204 No Content
```
