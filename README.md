# MESS-APP
Mess-app is to manage mess bugets amoung roommates.

## Seting up
### Local environment
Setting up postgres.
```docker-compose up -d```

Installing goose
```go get bitbucket.org/liamstask/goose/cmd/goose```

To run db migration
```
source config/local.env
cd migration
goose -env local up
```
Install go-swagger
```
go get -u github.com/go-swagger/go-swagger/cmd/swagger
```

## ToDo:
 - Create architecture diagram.
 - Create swagger file for all api specification.
 - Database design for the app. Please make a doc first in docs directory.
 - Implement logger using logrus package.

# Contributions
 - Checkout branch from develop and make PR to develop.