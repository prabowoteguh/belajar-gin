# Go Assigment Haktiv8

## Topics

- Golang
- Back-end
- CRUD
- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [GORM](https://github.com/go-gorm/gorm)
- PostgreSQL 16
- RESTful API

## Installation

1. Clone this repository:

```
git clone https://github.com/prabowoteguh/belajar-gin.git
```

2. Go to `belajar-gin` directory:

```
cd belajar-gin
```

3. Download and install `Go` SDK `1.21.6`.

4. Build and run the application:

```
go run server.go
```

or

```
go run .
```

5. Apply migrations:

```
go run migration.go
```

## API Endpoints

You can access and try endpoint with postman here [API Postman](https://elements.getpostman.com/redirect?entityId=11831418-0db734cb-a2d5-4d29-b4ff-17ece7139939&entityType=collection)

### All Orders

- Request URL: `localhost:8000/api/orders`
- Method: `GET`
- Path: `/orders`
- Headers: `Accept:application/json, Content-Type:application/json`
- Parameters: -
- Status Code: `200`
- Response:

```json
{
  "result": [
    {
      "OrderId": 1,
      "CustomerName": "Teguh Agung",
      "Items": [
        {
          "ItemId": 1,
          "ItemCode": "123",
          "Description": "IONIC5",
          "Quantity": 1,
          "OrderId": 1,
          "CreatedAt": "0001-01-01T00:00:00Z",
          "UpdatedAt": "0001-01-01T00:00:00Z"
        }
      ],
      "OrderedAt": "2019-11-10T06:21:46+07:00",
      "CreatedAt": "2024-02-09T14:59:39.004018+07:00",
      "UpdatedAt": "2024-02-09T14:59:39.004018+07:00"
    },
    {
      "OrderId": 2,
      "CustomerName": "Teguh Agung",
      "Items": [
        {
          "ItemId": 2,
          "ItemCode": "123",
          "Description": "IONIC5",
          "Quantity": 1,
          "OrderId": 2,
          "CreatedAt": "0001-01-01T00:00:00Z",
          "UpdatedAt": "0001-01-01T00:00:00Z"
        }
      ],
      "OrderedAt": "2019-11-10T06:21:46+07:00",
      "CreatedAt": "2024-02-09T15:07:35.339938+07:00",
      "UpdatedAt": "2024-02-09T15:07:35.339938+07:00"
    },
    {
      "OrderId": 3,
      "CustomerName": "Teguh Agung",
      "Items": [
        {
          "ItemId": 3,
          "ItemCode": "123",
          "Description": "IONIC5",
          "Quantity": 1,
          "OrderId": 3,
          "CreatedAt": "2024-02-09T15:41:54.151331+07:00",
          "UpdatedAt": "2024-02-09T15:41:54.151331+07:00"
        }
      ],
      "OrderedAt": "2024-02-09T15:41:54.138938+07:00",
      "CreatedAt": "2024-02-09T15:41:54.141589+07:00",
      "UpdatedAt": "2024-02-09T15:41:54.141589+07:00"
    }
  ]
}
```

### Create Order

- Request URL: `localhost:8000/api/orders`
- Method: `POST`
- Path: `/orders`
- Headers: `Accept:application/json, Content-Type:application/json`
- Parameters:

```json
{
  "customerName": "Teguh Agung",
  "items": [
    {
      "itemCode": "123",
      "description": "IONIC5",
      "quantity": 1
    }
  ]
}
```

- Status Code: `200`
- Response:

```json
{
  "Data": {
    "OrderId": 3,
    "CustomerName": "Teguh Agung",
    "Items": [
      {
        "ItemId": 3,
        "ItemCode": "123",
        "Description": "IONIC5",
        "Quantity": 1,
        "OrderId": 3,
        "CreatedAt": "2024-02-09T15:41:54.151331+07:00",
        "UpdatedAt": "2024-02-09T15:41:54.151331+07:00"
      }
    ],
    "OrderedAt": "2024-02-09T15:41:54.138938+07:00",
    "CreatedAt": "2024-02-09T15:41:54.141589+07:00",
    "UpdatedAt": "2024-02-09T15:41:54.141589+07:00"
  },
  "status": "Data successfully added!"
}
```

### Update Order

- Request URL: `localhost:8000/api/orders/{OrderID}`
- Method: `PUT`
- Path: `/orders/{OrderID}`
- Headers: `Accept:application/json, Content-Type:application/json`
- Parameters:

```json
{
  "customerName": "Prabowo Teguh"
}
```

- Status Code: `200`
- Response:

```json
{
  "data": {
    "id": 10000011,
    "agency_name": "Test Agency Name",
    "status": "waiting",
    "is_confirmed": false,
    "is_checked": false,
    "rental_date": "01-02-2024",
    "user_name": "Test7",
    "transport_count": 7,
    "guest_count": 7,
    "admin_note": "qq",
    "note": "ww",
    "email": "777@777.test",
    "phone": "71111111111",
    "confirmed_at": null,
    "created_at": "31-01-2024 22:11:15",
    "updated_at": "31-01-2024 22:17:47"
  }
}
```

### Delete Order

- Request URL: `localhost:8000/api/orders/{OrderID}`
- Method: `DELETE`
- Path: `/orders/{OrderID}`
- Headers: `Accept:application/json, Content-Type:application/json`
- Status Code: `200`
- Response:

```json
{
  "message": "Order with id 1 Has Been Successfully Deleted",
  "success": true
}
```
