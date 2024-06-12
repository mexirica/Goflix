
# Goflix a microservices stream video system in Go using go-kit

This project is a comprehensive microservices system developed in Go using the go-kit framework. It encompasses several services that together create a robust and scalable platform for various functionalities such as user management, authentication, video streaming, and notifications. Below are the details of each service and the overall architecture.

## Microservices Overview

### 1. Gateway Service

The Gateway Service acts as a unified entry point for clients. It routes requests to the appropriate microservices, ensuring a seamless and secure interaction with the system.

Key Features:
- Request routing
- Load balancing
- Security enforcement

### 2. User Service

The User Service manages user information and roles, providing functionalities for user registration and management.

Key Features:
- User registration and management
- Role-based access control
- User data storage

### 3. Authentication and Authorization Service

This service handles authentication and authorization processes, supporting secure access to the system's resources.

Key Features:
- User authentication (support for OAuth2, JWT, etc.)
- Authorization mechanisms
- Token management

### 4. Video Streaming Service

The Video Streaming Service manages the streaming of video content, ensuring efficient and high-quality video delivery to users.

Key Features:
- Video upload and storage (for admin)
- Stream processing and delivery
- Video metadata management

### 5. Notification Service

The Notification Service handles the delivery of notifications to users, supporting various types of notifications such as email and SMS.

Key Features:
- Notification sending (email, SMS)
- Notification templates

## Common Functionalities

- **Development with Go**: Utilization of the go-kit framework for creating microservices.
- **Service Communication**: Use of gRPC for efficient inter-service communication.
- **Message Queues**: RabbitMQ for asynchronous communication between microservices.
- **Orchestration with Kubernetes**: For management and scalability of microservices.
- **Monitoring and Logging**: Tools for monitoring service performance and logging events and errors.
- **CI/CD**: Continuous integration and delivery pipeline using tools like GitLab CI.
- **Security**: Implementation of security best practices including authentication and authorization.
- **Cache**: Usage of Redis for caching to reduce unnecessary database calls and improve performance.

## Technologies and Tools

- **Language**: Go
- **Framework**: go-kit
- **Communication**: gRPC and RabbitMQ
- **Orchestration**: Kubernetes
- **Caching**: Redis
- **CI/CD**: GitLab CI

## :handshake: Collaborators
<table>
  <tr>
    <td align="center">
      <a href="https://github.com/mexirica">
        <img src="https://avatars.githubusercontent.com/u/67772460?v=4" width="100px;" alt="Foto de Rodrigo Mecheri no GitHub"/><br>
        <sub>
          <b>mexirica</b>
        </sub>
      </a>
    </td>
  </tr>
</table>

## :dart: Status

* In progress
```
