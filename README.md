# synlabs

## Overview

**synlabs** is a web application built using Go and MongoDB, designed to manage job applications and user profiles. The application provides a RESTful API for user registration, job postings, and application management.

## Features

- User registration and authentication
- Job posting and management
- Application submission for jobs
- Profile management for users

## Technologies Used

- **Go**: The primary programming language for the backend.
- **MongoDB**: NoSQL database for storing user and job data.
- **Gin**: A web framework for building the RESTful API.
- **Docker**: For containerization and easy deployment.

## Prerequisites

Before you begin, ensure you have the following installed:

- [Go](https://golang.org/dl/)
- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Setup Instructions

Follow these steps to set up the project locally:

### 1. Clone the Repository

Clone the repository to your local machine:

```bash
git clone https://github.com/hritesh04/synlabs.git
cd synlabs
```

### 2. Create a `.env` File

Create a `.env` file in the root of the project directory. This file will store your environment variables. Here’s an example of what to include:

```plaintext
PORT=3000
DSN=mongodb://db:27017/recruitment
SECRET=your_secret_key
API_KEY=your_api_key
PARSER_URL=your_parser_url
```

Make sure to replace `your_secret_key`, `your_api_key`, and `your_parser_url` with actual values.

### 3. Build and Run with Docker

To build and run the application using Docker, execute the following command:

```bash
docker-compose -f compose-dev.yml up
```

This command will:

- Build the Go application.
- Start the MongoDB container.
- Initialize the database with the necessary collections.
- Start the backend service.

### 4. Access the Application

Once the application is running, you can access the API at:

```
http://localhost:3000
```

### 5. API Endpoints

Here are the key API endpoints available:

- **User Registration**: `POST /signup`
- **User Login**: `POST /login`
- **Upload Resume**: `POST /uploadResume`
- **Get All Jobs**: `GET /jobs`
- **Apply to Job**: `GET /jobs/apply?job_id=jod_id`
- **Create Job**: `POST /admin/job`
- **Get Job Info**: `GET /admin/job/:jobID`
- **Get all User**: `GET /admin/applicants`
- **Get User Profile**: `GET /admin/applicant/:applicantID`

### 6. Stopping the Application

To stop the application, press `CTRL + C` in the terminal where Docker Compose is running. You can also run:

```bash
docker-compose -f compose-dev.yml down -v
```

This command will stop and remove the containers and the assosiated volumes and networks.