# GoRestTemplate

Sample GoLang Application that can accept REST requests using GIN that stores
data into a SQLite table using GORM

This sample application is based on this [tutorial](https://blog.logrocket.com/rest-api-golang-gin-gorm/)


### To Install Gin and Gorm

```
go get github.com/gin-gonic/gin gorm.io/gorm
```

then

```
go mod download
```

### To Change Default Port

```
PORT=8081 ./webservice
```
