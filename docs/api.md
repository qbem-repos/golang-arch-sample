# API

## Routes
```go
func main() {
    r := mux.NewRouter()
    r.HandleFunc("/heathcheck", HeathCheck)
    r.HandleFunc("/your-service", HeathCheck).Methods(http.MethodGET)
    r.HandleFunc("/your-service", HeathCheck).Methods(http.MethodPOST)
}
```

## API Diagram

Visão de como os services se comunicam internamente.

```mermaid
    sequenceDiagram
   actor User
    participant Service
    participant Database

    Note left of Service: Microservice
    User-->>Service: [POST]
   Service-->> User: Response 201
    Service-->>Database: Save()
    Database-->>Service: Return of objectID | MongoDB
    User-->>Service: [GET] 
   Service-->>User: Response 200
    

```