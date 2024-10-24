# Value Indexer Service

## Overview
The Value Indexer Service is a REST API that indexes values based on a sorted file.

## Modules
The service consists of the following modules:
- **api**: API based on the echo framework.
- **provider**: The provider loads the data from the file, abstracted with interface for ease of replacement.
- **search**: The search module searches for the value in the sorted slice and returns the value if found, or nearest match in 10% tolerance. It uses binary search in order to achieve O(log n) time complexity. As an exercise, I've added support for negative values'.
- **logger**: Basic logger that utilizes stdout. Could be easily replaced by more sophisticated solution, like zap or logrus.

## Configuration
The service can be configured using environment variables defined in the `.env` file:

- `PORT`: The port on which the service will run (default: `8080`).
- `LOG_LEVEL`: The logging level (default: `INFO`).

## Usage
### Running the Application
To run the application, use:
```sh
make run
```
### Running tests
To run the tests, use:
```sh
make test
```

## API Documentation
The API consists of one endpoint:
- `GET /endpoint/{value}`: Searches for the value in the sorted slice and returns the value if found, or nearest match in 10% tolerance.

Example response:
```json
{
  "index": 67,
  "value": 6700
}
```
Closest match response:
```json
{
  "index": 68,
  "message": "closest match found",
  "value": 6800
}
```
No match response:
```json
{
  "error": "index for value not found",
  "value": "20"
}
```

## Bugs
Provided input.txt file is sorted up to the last value which is 1000000, and the previous one is 9999900. Provider validates if the file is sorted, so given that the task implies that the values are sorted, I have fixed the last value. But as a improvement sorting could be done on provider side during load.