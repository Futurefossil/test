
# LockSysProject

## Description

LockSysProject is a demonstration project showing how to create a simple server-client architecture in Go, for controlling and querying a smart lock system. The server simulates a smart lock's API, capable of locking, unlocking, and reporting the lock's status. The client interacts with this API to perform operations.

## Getting Started

### Prerequisites

- Go 1.15 or higher

### Installation

Clone the repository to your local machine:

```sh
git clone https://github.com/axs7941/LockSysProject.git
cd LockSysProject
```

### Running the Server

Navigate to the server directory and start the server:

```sh
cd server
go run server.go
```

The server will start listening on port 8080 for incoming requests.

### Running the Client

Open a new terminal window, navigate to the client directory, and run the client:

```sh
cd ../client
go run main.go
```

The client will send lock, unlock, and status requests to the server and print the responses.

## API Reference

The server defines the following endpoints:

- `POST /lock`: Locks the smart lock.
- `POST /unlock`: Unlocks the smart lock.
- `GET /status`: Returns the current status of the smart lock.

## Contributing

Contributions to the LockSysProject are welcome. Please feel free to submit pull requests or create issues for bugs and feature requests.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
