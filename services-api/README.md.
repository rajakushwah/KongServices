# Services API


## Overview


The Services API is a RESTful API built in Go Programming Language. It provides CRUD operations for managing a collection of services and their versions. This document outlines the design considerations, assumptions, and trade-offs made during the development of this API.


## Table of Contents


*   [Project Structure](#project-structure)
*   [Design Considerations](#design-considerations)
*   [Assumptions](#assumptions)
*   [Trade-offs](#trade-offs)
*   [Endpoints](#endpoints)
*   [How to Run the Application](#how-to-run-the-application)



## Project Structure



*   `main.go`: Entry point of the application. Initializes the router and starts the server.
*   `handlers/handlers.go`: Contains the handler functions for all API endpoints.
*   `models/models.go`: Defines the data structures (Service, ServiceVersion).
*   `utils/utils.go`: Contains utility functions for handling HTTP responses.
*   `go.mod`: Go module file for managing dependencies.
*   `README.md`: Project documentation.


## Design Considerations


1.  **RESTful API Design:**


 *   The API follows RESTful principles, using standard HTTP methods (GET, POST, PUT, DELETE) to perform CRUD operations on services.
 *   Endpoints are designed to be intuitive and resource-based (e.g., `/api/services`, `/api/services/{id}`).
2.  **Modularity:**


 *   The code is structured into multiple packages (`handlers`, `models`, `utils`) to promote separation of concerns and maintainability.
 *   Each package has a specific responsibility, making it easier to understand, test, and modify the code.
3.  **Error Handling:**


 *   The API uses consistent error responses, with appropriate HTTP status codes and JSON error messages.
 *   The `utils.RespondWithError` function provides a centralized way to send error responses.
4.  **Data Storage:**


 *   For simplicity, the API uses in-memory storage (slices) to store services and their versions.
 *   In a production environment, this would be replaced with a persistent database (e.g., PostgreSQL, MySQL).



## Assumptions


1.  **In-Memory Storage:**


 *   The API assumes that the data volume is small enough to be stored in memory. This is suitable for development and testing but not for production.

2.  **Basic Input Validation:**


 *   The API performs basic input validation (e.g., checking for invalid JSON payloads) but does not include comprehensive validation of all input parameters.



## Trade-offs


1.  **In-Memory vs. Persistent Storage:**


 *   *Trade-off:* Using in-memory storage simplifies development but limits scalability and persistence.
 *   *Rationale:* In-memory storage was chosen for its simplicity and ease of setup. For a production environment, a database would be necessary.
2.  **`gorilla/mux` vs. Standard `net/http`:**


 *   *Trade-off:* Using `gorilla/mux` adds an external dependency but provides more advanced routing features.
 *   *Rationale:* `gorilla/mux` was chosen for its flexibility and ease of use in defining complex routes and extracting parameters from URLs.

3.  **Limited Input Validation vs. Comprehensive Validation:**


 *   *Trade-off:* Performing limited input validation simplifies development but increases the risk of data integrity issues.
 *   *Rationale:* Basic input validation was implemented to prevent common errors. More comprehensive validation could be added in the future.
4.  **Simple Sorting vs. Advanced Sorting:**


 *   *Trade-off:* Only sorting by name is implemented for simplicity, while advanced sorting options are omitted.
 *   *Rationale:* Sorting by name covers basic needs, while more complex sorting can be implemented as needed using database queries in production.


## Endpoints


| Method | Endpoint                      | Description                                           |
| :----- | :---------------------------- | :---------------------------------------------------- |
| GET  | `/api/services`               | Retrieve all services with filtering, sorting, pagination. |
| POST | `/api/services`               | Create a new service.                                 |
| GET  | `/api/services/{id}`          | Retrieve a single service by ID.                        |
| PUT  | `/api/services/{id}`          | Update an existing service.                             |
| DELETE | `/api/services/{id}`          | Delete a service.                                     |
| GET  | `/api/services/{id}/versions` | Retrieve service versions by ServiceID.               |
| GET  | `/api/services/name/{name}`   | Retrieve a service by name.                             |
| GET  | `api/services?filter=Priority Services`   | Filter a service by name.                             |
| GET  | `api/services?sort_by=name&sort_order=desc`   | sort a service by name.                             |
| GET  | `api/services?page=2&limit=3&filter=service&sort_by=name&sort_order=asec`   | All Filters.                                                    |


## How to Run the Application


1.  **Prerequisites:**


 *   Go installed (version 1.17 or higher)
2.  **Clone the repository:**


 ```sh
 git clone [repository-url]
 cd services-api
 go mod tidy
go run main.go

3.  Test Endpoints with provided postman collection
