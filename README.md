# Qbem Architecture with Golang | Sample

This is an simples sample

## TODO's

- [ ] Template de doc
  - [ ] API
  - [ ] Queue/Kafka/NatsIO ...N
  - [ ] Docker
  - [ ] Docker Compose
  - [ ] CI/CD Drone yml
  - [ ] Plugins VScode
  - [ ] Deploy
- [ ] Arch
  - [ ] API
  - [ ] Professional
    - [ ] Entities
      - [ ] main.go 
      - [ ] main_test.go 
    - [ ] Repos
      - [ ] ProfessionalRepo
    - [ ] UseCases
  - [ ] Schedules
    - [ ] Entities
    - [ ] Repos
    - [ ] UseCases
- [ ] Sandbox to Front End | Camilo
- [ ] Broadcast | ???


## TODO: Visão geral do micro serviço
```mermaid
    sequenceDiagram
    participant Service
    participant Broker
    participant Database

    Note left of Service: Microservice
    Service-->>Database: Save()
    Service->>+Broker: Send an message
    Note right of Broker: Topics and Queues
```


## TODO: Service
```mermaid
    sequenceDiagram
    participant Service
    participant API_EXTERNA
    participant Database

    Note left of Service: Microservice
    Service-->>API_EXTERNA: Request data of the Professional
    API_EXTERNA-->>Service: Response data of the Professional
    Service->>+API_EXTERNA: Request data of the Shedule
    API_EXTERNA-->>Service: Response data of the Shedule
    Service->>+Broker: Send an message
    Note right of Broker: Topics and Queues | PUB/SUB or SAGA???
    

```

## Reference
```mermaid
    sequenceDiagram
    participant DeviceGSM
    participant Server
    participant Mobile
    Note left of DeviceGSM: DeviceGSM IoT
    DeviceGSM-->>Server: Socket Connection
    DeviceGSM->>+Server: Send []bytes
    Server-->>-Mobile: Send data
    Server-->>Mobile: Send data []bytes 
    Mobile->>Server: [GET] Request >> Send device_ID 
    Note right of Mobile: App Mobile solicitando alguma info.
    Server-->>DeviceGSM: Send command
```