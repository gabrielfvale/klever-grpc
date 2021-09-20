# Klever gRPC

## Technical Challenge

The Technical Challenge consists of creating an **API** with **Golang** using **gRPC** with stream pipes that exposes an Upvote service endpoints.

The API will provide the user an interface to upvote or downvote a known list of the main Cryptocurrencies (Bitcoin, ethereum, litecoin, etc..).

Technical requirements:

- Keep the code in Github
- The API must have a read, insert, delete and update interfaces.
- The API must guarantee the typing of user inputs. If an input is expected as a string, it can only be received as a string.
- The API must contain unit test of methods it uses

You can choose the database but the structs used with it should support Marshal/Unmarshal with bson, json and struct

Extra:

- Deliver the whole solution running in some free cloud service
- The API have a method that stream a live update of the current sum of the votes from a given Cryptocurrency

## Usage

### Setting your environment variables

Copy and edit the `.env.example` to `.env`. Set the `MONGO_URI` variable to your MongoDB uri.

### Using MongoDB with Docker

Run the following command to get the Mongo instance running.

```console
docker-compose -f mongo/docker-compose.yml up -d
```

### Running the server

```console
go run cmd/grpc/server/main.go
```

### Running the client

```console
go run cmd/grpc/client/main.go
```

### Client usage

```console
./client [command] [flags]
```

### Available commands

### `client help`

Show command help.

### `client create` (alias `client add`)

Creates a crypto record, with the required flags `--symbol` and `--name`

### `client read [symbol]`

Reads a crypto record, with the symbol positional argument.

### `client update`

Updates a crypto record, with the required flag `--symbol` and optional flags `--name`, `--upvotes` and `--downvotes`

### `client list`

Lists the crypto records in the database.

### `client upvote [symbol]`

Upvotes a crypto record, with the symbol positional argument.

### `client downvote [symbol]`

Downvotes a crypto record, with the symbol positional argument.

### `client subscribe [symbol]`

Subscribes to a crypto record for real time votes reading.

**Note**: requires a replicaSet as the databse.
