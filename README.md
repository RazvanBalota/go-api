# go-api
A simple Go API service that provides a mock database for retrieving a user’s coin balance. This service returns JSON responses, including success and error handling.

## Endpoints

```bash
GET /account/coins/?username=<username>
```

## Responses:

### Success:
```json
{
  "Code": 200,
  "Balance": <int64>
}
```
### 400 Error:
```json
{
  "Code": 400,
  "Message": "Invalid username or token"
}
```
### 500 Error:
```json
{
  "Code": 500,
  "Message": "An Unexpected Error Occurred."
}
```

## Usage
- Clone the repository.

- Start the server:

```bash
go run cmd/api/main.go
```
- Test the API using Postman or Thunder Client.