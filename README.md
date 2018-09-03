# The backend

The backend is written in Go (golang), using PostgreSQL as persistance engine.

This is a brief documentation on how to get it up and running for development purposes.

## Prerequisites

1. Go is installed
    1. Download from here: https://golang.org/dl/
1. Docker is installed and docker Daemon is running
    1. Linux: Use https://get.docker.com/ - it is a convenience script that will install Docker on almost all Linux
    distros. Usage instructions can be found in the script in the top.
    1. Windows: https://docs.docker.com/docker-for-windows/install/
1. PostgreSQL is running locally (this can be done using docker/docker-compose. See below)

## Running database using docker-compose

Simply execute `docker-compose up -d`, and you will have a database running.

In order to create the correct tables, you need to *manually* execute the scripts in the `migrations` folder. One at a
time. This is not convenient and will be changed ASAP.

*NOTE: Automatic migrations was previously done automatically each time the backend started up, using the 
`go-migrate/migrate` repo. However, this introduced so many dependencies to the code that it wasn't feasable. 
Another solution is being worked out instead.*

## Running the backend

This can be done using go run.

```
go run main.go
```

This will compile and run main.go, which is the entrypoint to the backend. It will be listening for requests on
port 8081.