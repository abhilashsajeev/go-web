# go-web
Learning app for golang


### Steps to get this work

Install *docker* and *docker-compose*
go to taskmanager folder

```
docker-compose build
```

then

```
docker-compose up
```

API will run at port 9000
goto `http://localhost:9000/` to see 404 message

use postman to generate user `http://localhost:9000/users/register` POST with request body 
```javascript
{
  data: {
    firstname: "yourname"
    lastname: "lastname"
    email : "a@a.com"
    password: "123456"
  }
}
```


use http://localhost:9000/users/login POST to login with request body

```javascript
{
  data: {
    email: "a@a.com",
    password: "123456"
  }
}
```
