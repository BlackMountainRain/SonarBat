# Sonar Bat

[![codecov](https://codecov.io/gh/BlackMountainRain/SonarBat/graph/badge.svg?token=XUMKY9D690)](https://codecov.io/gh/BlackMountainRain/SonarBat)

🦇 Sonar Bat is a large-scale network detection system designed to help users configure and execute network-related tasks on specific endpoints or hosts. It features a user-friendly web interface (webui) for task configuration and is built using a microservice architecture with Go as the primary programming language.

## Features

- **Web-based Task Configuration**: Configure network detection tasks through an intuitive webui.
- **API**: Expose an API(mainly task configuration) for programmatic access and integration with other systems(In-House or Third-Party).
- **Host Matching**: Dispatch tasks to specific endpoints based on host characteristics and requirements.
- **Microservice Architecture**: The service is composed of independent microservices for scalability and maintainability.
- **Efficient and Concurrent**: Leverages Go's performance and concurrency capabilities for efficient task execution.
- **Massive Task Support**: Sonar Bat can handle and process a large volume of tasks simultaneously, making it suitable for large-scale network operations.

## Installation
(WIP)

## Usage
(WIP)

## Development
```bash
# Create a new proto in api
kratos proto add api/task/v1/task.proto

# generate the client codes based on the proto
make api

# generate the server codes based on the proto
kratos proto server api/task/v1/task.proto -t internal/task/service
```

### Create a new table
```bash
go run -mod=mod entgo.io/ent/cmd/ent new Task
```
Edit the schema in `ent/schema/task.go` and run the following command to generate the table
```bash
go run -mod=mod entgo.io/ent/cmd/ent generate ./ent
```

## Contributing

Contributions are welcome! Please follow the guidelines outlined in the [CONTRIBUTING.md](CONTRIBUTING.md) file.

## License

This project is licensed under the [MIT License](LICENSE).

## Contact

For any questions or inquiries, please contact us :)