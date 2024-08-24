# A Discord-like Platform for Learning and Collaboration

Welcome to StudyBud-Go, a web application inspired by Discord, designed to enhance learning and facilitate information exchange in group settings. This Go-based version builds on the foundation laid by [Denis Ivy's original Django project](https://github.com/divanov11/StudyBud), offering a robust platform for creating virtual spaces to discuss topics, share knowledge, and collaborate effectively.

## Features

- **Account Creation and Registration:** 
  - Users can sign up and create an account.
  - Secure login allows users to access their accounts.

- **Room Creation:** 
  - Create rooms on various topics to encourage conversation and study.

- **Room Participation:** 
  - Join existing rooms and engage in discussions.

- **Messaging:** 
  - Send and receive messages within rooms.

- **Recent Activities Log:** 
  - Stay informed about ongoing activities and conversations.

- **Searching:** 
  - Search for topics or discussions to join and participate in.

## Technologies

- **Go:** Programming language for core development.
- **go-chi:** Router for handling HTTP requests.
- **Redis:** Session management and authentication.
- **PostgreSQL:** Database for data storage.
- **GORM:** ORM library for database interactions.
- **zerolog:** Logging library for efficient logging.

## Architecture

StudyBud-Go utilizes Clean Architecture principles for a scalable and maintainable codebase, structured into:
- **Entities:** Core business logic and data models.
- **Use Cases:** Application-specific business rules.
- **Interface Adapters:** Interfaces for interacting with external systems.
- **Frameworks and Drivers:** Go’s standard library and third-party frameworks.


## Getting Started

### Prerequisites
- Docker and Docker Compose installed

### Setup and Installation

1. **Clone the Repository:**
   ```bash
   git clone https://github.com/ElyarSadig/StudyBud-Go.git
   ```
2. **Navigate to the Project Directory:**
   ```bash
   cd StudyBud-Go
   ```

3. **Setup Environment Variables:**
   Create a `.env` file in the project root directory with the following contents:
   ```env
   DB_USER=admin
   DB_PASSWORD=password
   REDIS_PASSWORD=password
   REDIS_USERNAME=root
   SESSION_PRIVATE_KEY=password12345678
   ```

4. **Start the Services:**
   Use Docker Compose to build and start the application and its dependencies:
   ```bash
   docker-compose up -d
   ```

   This command will build and start the application along with its dependencies (PostgreSQL and Redis) in detached mode.

5. **Run the Application:**
   The Dockerfile specifies the `CMD` to run the application with the `config-stage.yaml` configuration file, migration, and seeding by default. If you need to run different commands, modify the Dockerfile or override the `CMD` at runtime.

6. **Access the Application:**
   - The application will be available at `http://localhost:8080`.

### Configuration
- **Dockerfile Configuration:**
  - The `Dockerfile` is set up to copy the application binary and configuration files into the Docker image.
  - Configuration is handled through `config-stage.yaml`, which is copied into the `/configs` directory inside the container.
  - To modify configurations, update `config-stage.yaml` or build with different configurations.

- **Environment Variables:** 
  - Make sure your `.env` file is properly configured with the necessary environment variables.

### Health Check
- **Application Health Check:** 
  The application exposes a health check endpoint at `http://localhost:8080/health`. Docker Compose uses this endpoint to verify that the application is running correctly.


## Acknowledgments

- Based on the original StudyBud project developed with Django by [Denis Ivy](https://github.com/divanov11/StudyBud).
- Thanks to the Go community for their valuable tools and libraries.

## Contributing

Contributions are welcome! Just submit a pull request :)