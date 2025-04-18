# Golang AWS Secrets Manager Demo

A demonstration project showcasing AWS Secrets Manager integration with a Go application, containerized using DevPod for development.

## Architecture

- **Application**: Go HTTP server that retrieves secrets from AWS Secrets Manager
- **Containerization**: Docker with multi-stage build for minimal image size
- **Development**: DevPod for consistent development environment
- **Testing**: Go test framework with HTTP endpoint testing

## Prerequisites

- Go 1.22 or later
- Docker Desktop
- DevPod CLI
- AWS CLI with configured credentials
- Task (taskfile.dev)

## Project Structure

```
.
├── main.go              # Main application code
├── main_test.go         # Test cases
├── Dockerfile           # Multi-stage Docker build
├── devcontainer.json    # DevPod configuration
├── Taskfile.yaml        # Task automation
└── README.md           # Project documentation
```

## Setup

1. Install DevPod:

```bash
task setup:devpod
```

2. Start the development workspace:

```bash
task devpod:up
```

## Development

The project uses DevPod for development, providing a consistent environment with:

- Go toolchain
- AWS CLI
- Docker-in-Docker
- VS Code extensions

### Key Features

- **Hot Reloading**: Using Air for development
- **Testing**: Integrated test framework
- **Containerization**: Optimized Docker build
- **AWS Integration**: Secrets Manager access

## Building and Running

### Local Development

1. Start the development server:

```bash
task dev
```

2. Run tests:

```bash
task test
```

### Containerized

1. Build the Docker image:

```bash
task docker:build
```

2. Run the container:

```bash
task docker:run
```

## Testing

The project includes:

- HTTP endpoint testing
- AWS Secrets Manager integration tests
- Test coverage reporting

Run all tests:

```bash
task test:devpod
```

## Environment Variables

- `AWS_SECRET_NAME`: Name of the secret to retrieve from AWS Secrets Manager
- `DOCKER_BUILDKIT`: Enable BuildKit for faster Docker builds

## API Endpoints

### GET /secret

Retrieves a secret from AWS Secrets Manager.

**Response**:

- Success: `Secret value: <secret>`
- Error: Appropriate HTTP status code with error message

## DevPod Configuration

The development environment is configured with:

- Go 1.22
- AWS CLI
- Docker-in-Docker
- VS Code extensions for Go development

## Task Automation

The project uses Task for automation:

- `task devpod:up`: Start DevPod workspace
- `task dev`: Start development server
- `task test`: Run tests
- `task docker:build`: Build Docker image
- `task docker:run`: Run container

## Security Considerations

- AWS credentials are mounted from host machine
- Non-root user in container
- Minimal base image (scratch)
- SSL certificates included for AWS SDK
