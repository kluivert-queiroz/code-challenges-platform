# Code Challenge Platform (WIP)

A platform for running and testing code challenges in various programming languages. The system uses Docker containers to safely execute user-submitted code and verify the results against expected outputs.

## Next steps
- Add support to multiple languages
- Add database and a way to create challenges from UI

## Overview

This platform consists of two main components:
- A web-based UI for submitting code solutions
- A Go backend server that handles code execution and verification

The system works by:
1. Receiving code submissions through the UI
2. Creating isolated Docker containers for code execution
3. Running the submitted code with test inputs
4. Verifying the output against expected results
5. Returning the execution results to the user

## Prerequisites

- Docker
- Docker Compose

## Running the Project

1. Start the services using Docker Compose:
```bash
docker compose up
```

The platform will be available at `http://localhost:8080` (UI) and `http://localhost:3000` (API server).

## Architecture

### Frontend
- Web-based UI for submitting code solutions
- Supports multiple programming languages (only ts for now)
- Real-time feedback on code execution results

### Backend
- Go server that manages code execution
- Creates a isolated Docker container for running submissions
- Handles input/output validation
- Can support multiple programming languages through Docker images


## Security

- All code execution happens in isolated Docker containers
- Resource limits and timeout are enforced to prevent abuse
- No network access is allowed from within the containers

