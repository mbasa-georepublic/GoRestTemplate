# GoRestTemplate

Sample GoLang Application that can accept REST requests using GIN that stores
data into a SQLite table using GORM. This can be easily modified to use a 
PostgreSQL database instead. 


This sample application is based on this [GIN + GORM tutorial](https://blog.logrocket.com/rest-api-golang-gin-gorm/)



### To Install Gin and Gorm

```
go get github.com/gin-gonic/gin gorm.io/gorm
```

then

```
go mod download
```

### To Build the application

```
go clean;go build
```

### To Change Default Port and other Environmental Settings

```
PORT=8081 ./webservice
```
