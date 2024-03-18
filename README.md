<!-- markdownlint-configure-file { "MD024": false } -->
# StoriCard

## Description

[Tech_Challenge_-_Software_Engineer](docs/Tech_Challenge_-_Software_Engineer.pdf)

## Endpoints

### Process File

#### Request

```curl
curl --location 'http://localhost:3000/api/v1/transaction/process' \
--header 'Content-Type: application/json' \
--data-raw '{
    "fileName": "file-example.csv",
    "email": "joaltoroc@jatc.co"
}'
```

#### Response

```json
{
    "status": 200,
    "message": "Success",
    "data": {
        "executionID": "17aab49b-f2dd-4ae5-aa68-7b31becff7cb"
    }
}
```

### Get All Execution

#### Request

```curl
curl --location 'http://localhost:3000/api/v1/transaction'
```

#### Response

```json
{
    "status": 200,
    "message": "Success",
    "data": [
        {
            "ID": 1,
            "ExecutionID": "17aab49b-f2dd-4ae5-aa68-7b31becff7cb",
            "FileID": 1,
            "TypeTransaction": "credit",
            "Date": "2024-07-15T00:00:00Z",
            "Value": 60.5,
            "CreatedAt": "2024-03-18T16:03:28.421Z"
        },
        {
            "ID": 2,
            "ExecutionID": "17aab49b-f2dd-4ae5-aa68-7b31becff7cb",
            "FileID": 2,
            "TypeTransaction": "debit",
            "Date": "2024-07-28T00:00:00Z",
            "Value": -10.3,
            "CreatedAt": "2024-03-18T16:03:28.421Z"
        },
        {
            "ID": 3,
            "ExecutionID": "17aab49b-f2dd-4ae5-aa68-7b31becff7cb",
            "FileID": 3,
            "TypeTransaction": "debit",
            "Date": "2024-08-02T00:00:00Z",
            "Value": -20.46,
            "CreatedAt": "2024-03-18T16:03:28.421Z"
        },
        {
            "ID": 4,
            "ExecutionID": "17aab49b-f2dd-4ae5-aa68-7b31becff7cb",
            "FileID": 4,
            "TypeTransaction": "credit",
            "Date": "2024-08-13T00:00:00Z",
            "Value": 10,
            "CreatedAt": "2024-03-18T16:03:28.421Z"
        }
    ]
}
```

### Get By Execution ID

#### Request

```curl
curl --location 'http://localhost:3000/api/v1/transaction/17aab49b-f2dd-4ae5-aa68-7b31becff7cb'
```

#### Response

```json
{
    "status": 200,
    "message": "Success",
    "data": [
        {
            "ID": 1,
            "ExecutionID": "17aab49b-f2dd-4ae5-aa68-7b31becff7cb",
            "FileID": 1,
            "TypeTransaction": "credit",
            "Date": "2024-07-15T00:00:00Z",
            "Value": 60.5,
            "CreatedAt": "2024-03-18T16:03:28.421Z"
        },
        {
            "ID": 2,
            "ExecutionID": "17aab49b-f2dd-4ae5-aa68-7b31becff7cb",
            "FileID": 2,
            "TypeTransaction": "debit",
            "Date": "2024-07-28T00:00:00Z",
            "Value": -10.3,
            "CreatedAt": "2024-03-18T16:03:28.421Z"
        },
        {
            "ID": 3,
            "ExecutionID": "17aab49b-f2dd-4ae5-aa68-7b31becff7cb",
            "FileID": 3,
            "TypeTransaction": "debit",
            "Date": "2024-08-02T00:00:00Z",
            "Value": -20.46,
            "CreatedAt": "2024-03-18T16:03:28.421Z"
        },
        {
            "ID": 4,
            "ExecutionID": "17aab49b-f2dd-4ae5-aa68-7b31becff7cb",
            "FileID": 4,
            "TypeTransaction": "credit",
            "Date": "2024-08-13T00:00:00Z",
            "Value": 10,
            "CreatedAt": "2024-03-18T16:03:28.421Z"
        }
    ]
}
```

## Email

![email](docs/email.png)

## Based on

[Golang Clean Architecture](https://github.com/DoWithLogic/golang-clean-architecture)

## Golang Clean Architecture

This is an example of implementation of Clean Architecture with S.O.L.I.D Principles in Go (Golang) projects.

Rule of Clean Architecture by Uncle Bob

- Independent of Frameworks. The architecture does not depend on the existence of some library of feature laden software. This allows you to use such frameworks as tools, rather than having to cram your system into their limited constraints.
- Testable. The business rules can be tested without the UI, Database, Web Server, or any other external element.
- Independent of UI. The UI can change easily, without changing the rest of the system. A Web UI could be replaced with a console UI, for example, without changing the business rules.
- Independent of Database. You can swap out Oracle or SQL Server, for Mongo, BigTable, CouchDB, or something else. Your business rules are not bound to the database.
- Independent of any external agency. In fact your business rules simply don’t know anything at all about the outside world.

More at [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2011/11/22/Clean-Architecture.html)

## 🔍 What S.O.L.I.D Principle is?

SOLID is an acronym representing five essential design principles for writing maintainable and scalable software. These principles were introduced by Robert C. Martin (Uncle Bob) and have become fundamental guidelines for good software design.

### SOLID Principles

#### Single Responsibility Principle (SRP)

The SRP states that a class should have only one reason to change, meaning it should have a single, well-defined responsibility. This principle encourages the separation of concerns, making code more modular and easier to maintain.

#### Open/Closed Principle (OCP)

The OCP emphasizes that software entities (classes, modules, functions) should be open for extension but closed for modification. To achieve this, use abstractions (e.g., interfaces, abstract classes) and dependency injection to allow adding new functionality without altering existing code.

#### Liskov Substitution Principle (LSP)

The LSP states that objects of a derived class should be substitutable for objects of the base class without affecting the correctness of the program. In other words, derived classes must adhere to the contract defined by their base classes.

#### Interface Segregation Principle (ISP)

The ISP suggests that clients should not be forced to depend on interfaces they do not use. Create smaller, more focused interfaces rather than large, monolithic ones. This avoids unnecessary dependencies and promotes flexibility.

#### Dependency Inversion Principle (DIP)

The DIP promotes high-level modules (e.g., use cases) to depend on abstractions (e.g., interfaces) rather than concrete implementations. This inversion of dependencies allows for flexibility and testability by injecting dependencies from external sources.

### Benefits of SOLID Principles

- Improved code maintainability and readability.
- Easier collaboration among developers on large projects.
- Reduced code duplication.
- Better testability and test coverage.
- Increased adaptability to changing requirements.

## 🔥 Layers of Domain

- [Controller / Handler](https://github.com/joaltoroc/storicard/tree/main/internal/transaction/handler/v1)
- [Data Transfer Object (DTO)](https://github.com/joaltoroc/storicard/tree/main/internal/transaction/dtos)
- [Usecase](https://github.com/joaltoroc/storicard/tree/main/internal/transaction/usecase)
- [Entity](https://github.com/joaltoroc/storicard/tree/main/internal/transaction/entities)
- [Repository](https://github.com/joaltoroc/storicard/tree/main/internal/transaction/repository)

## The diagram

![golang clean architecture](docs/clean-architecture.png)

## 🏗️ How To Run

setup environment and running on local

```bash
make run    # Start the database, run migrations, and start the application locally
```

## ✨ References

Golang:

- [Go Documentation](https://golang.org/doc/)
- [Go-Standards](https://github.com/golang-standards/project-layout)
- [Go For Industrial Programming](https://peter.bourgon.org/go-for-industrial-programming/)
- [Uber Go Style Guide](https://github.com/uber-go/guide)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [typical-rest-server](https://github.com/typical-go/typical-rest-server/)
- [go-clean-arch](https://github.com/bxcodec/go-clean-arch)
- [a-clean-way-to-implement-database-transaction-in-golang](https://dev.to/techschoolguru/a-clean-way-to-implement-database-transaction-in-golang-2ba)

Architecture

- [SOLID Principle](https://en.wikipedia.org/wiki/SOLID)

Framework

- [Echo framework](https://echo.labstack.com/)
