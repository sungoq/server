# sungoq
Queue engine made with Go

# API
## HTTP

### Create a Topic
```
POST /topics HTTP/1.1
Content-Type: application/json
Content-Length: 22

{
    "name": "chat"
}
```
### Get All Topics
```
GET /topics HTTP/1.1
```

### Delete a Topic
```
DELETE /topics?name=notification HTTP/1.1
```

### Publish a Message
```
POST /publish HTTP/1.1
Content-Type: application/json
Content-Length: 44

{
    "topic": "chat",
    "message": "hi"
}
```
