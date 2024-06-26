# 2 - Web server

Web servers need something else apart from the ability to handle requests and responses. They need to be able to access
data and perform operations on it in a way that is efficient and easy to maintain. This is where the architecture of the
web service comes into play.

## Objective

This project is focused on improving the architecture of the web service to allow more complex scenarios and better
handling of the data.

To achieve this, the project will address the following points:

- Implementing a nodes repository to handle the access to the nodes data.
- Injecting the repository into the service.
- Implementing tests to validate the endpoint behaviour.
- Organizing the project structure to better separate the concerns, trying to follow the good practices of the Go
  community.

## Results

### Dependency injection and project structure

All the functionality of the previous project is maintained, but the code is now organized in a different way.

I have been looking for a "standard" way to organize the code in Go, but I couldn't find a great consensus.
Some examples try to adapt the Domain-Driven Design (DDD) structure from Java to Go. Those examples consider the
packages as units of functionality instead of a tool to organize the code ([go-ddd](https://www.citerus.se/go-ddd/)),
but personally, I have found them quite unclear.

That is why I have decided to create a structure that follows the general structure of the Go projects, but it is
aligned with the approach I typically use for Java projects.

The project is organized as follows:

```
├── cmd                         Entry points of the applications
│   └── api
│       ├── bootstrap           Bootstrapping of the application and its dependencies
│       └── main.go             Main function
│
└── internal                    Internal packages, not shared with other applications
        ├── nodes               Context of the nodes
        │   ├── domain          Domain objects, repositories
        │   └── application     Application services         
        └── platform            Infrastructure code
            ├── server          Web server implementation
            └── storage         Data storage implementation
```

In my opinion, the priority when organizing a project of these characteristics is to make it as maintainable
as possible, so I have tried to create a structure that is easy to understand and navigate.
I think focusing on shareable and reusable packages in the internal directory can lead to the opposite effect, so I did
not take it into account when organizing the project.

I have considered the platform package as a place to put the infrastructure code (although I would normally place it in
a infrastructure directory, next to domain and application), such as the web server and the storage implementations.
This is a common practice in Go projects.

### Testing

I have implemented tests for the service and the repository. I have explored the functionality of the testing package
implementing unit tests, table-driven tests, analyzing the coverage and writing a benchmark for the service.

#### Coverage commands

```shell
go test ./... -coverpkg=./... -coverprofile=test_coverage
go tool cover -html=test_coverage
```

![Html test coverage result](./test_coverage.png)

## Next steps

There are a number of topics I want to explore in the future:

- Sql implementation of the repository.
- Implementing a cache layer.
- Implementing a logging layer.
- Configuration management.
- Deployment of the application.
- Usage of mocks in the tests.
- Events and messaging.

