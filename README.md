This example microservice for screenshot elements website

Endpoints:

`[POST] localhost:8080` - Create screenshot element
```json
{
    "element": "[alt=Go]",
    "name": "test122",
    "url": "https://pkg.go.dev/encoding/base64"
}
```

`[GET] localhost:8080` - Get all screenshots data  
`[GET] localhost:8080/image/:id` - Get image