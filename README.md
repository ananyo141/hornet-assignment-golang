# Books Library API Assignment (GoLang)

This document outlines the project, detailing the structure of the codebase, the technologies used, and instructions on how to set up and run the application.

## Overview

This project is a backend server application for a book management system. It supports two types of users: Admin and Regular users, each with distinct functionalities. The application is built to handle user authentication, book management (addition and deletion of books).

## Project Structure

The project is organized as follows:

- `src/`: Contains all the source code for the application.
  - `controllers/`: Business logic for handling requests.
  - `middlewares/`: JWT authentication and authorization middleware.
  - `models/`: Data structures for users and books.
  - `routes/`: Route definitions for the application.
  - `utils/`: Login related to jwt and http responses.
  - `db/`: Operations related to database
- `.env`: Environment variables for JWT secret and other configurations.
- `main.go`: The entry point of the application.

## Features

- **Error handling** with custom error handler
- **Structured API response** for easy consumption (especially for typed languages like Dart, Rust)
- Strict **validation** and **input sanitization**
- Save books in `PostgreSQL` for **persistent book records**
- **Role based Auth** token middleware (only admin allowed to modify records)
- **Dockerized** for total isolation from host system (files remain unchanged; **image size under 13mb** via multi-stage builds)
- Clean codebase structure with **MVC architecture** with proper coding conventions according to best practices.

## Technologies Used

- **Go**: The primary programming language used to build the application.
- **GoFiber**: A high-performance HTTP framework used for building the API.
- **JWT-go**: Library for generating and validating JSON Web Tokens.
- **gORM**: Easy way to communicate with PostgreSQL.

## Setup and Running the Application

### Prerequisites

- Ensure Go is installed (version 1.22 or newer).
- Check docker is available (optional, if running via `docker compose`)

### Installation Steps

#### With `docker compose` (Recommended):

1. Go to project root and invoke

```bash
$ docker compose up -d
```

#### Manual

1. Clone the repository to your local machine:

```bash
$ git clone https://github.com/ananyo141/scalex-assignment.git
```

or copy from this repo.

2. Create .env file with `JWT_KEY` variable defined, or just copy `.env.docker` file

```bash
$ cp .env.docker .env
```

3. Build and run the application

```bash
$ go build -o main -ldflags="-s -w" ./src
$ ./main
```

4. Import `postman_collection.json` given in the project root to test the endpoints.

_NOTE:_ Make sure to setup PostgreSQL and it's .env variable while setting up
manually.

## Using Ansible

The Ansible playbook is written for an automated deploy to remote server using
docker and nginx. Just run this command on the terminal

```bash
$ ansible-playbook playbooks/deploy_golang_api.yml -i hosts
```

Make sure `ssh_key.pem` private key is present in `ansible/ssh_key.pem`.
