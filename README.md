# MESS-APP
Mess-app is to manage mess bugets amoung roommates.

## Seting up
Setting up postgres.
```docker-compose up -d```

Installing goose
```go get bitbucket.org/liamstask/goose/cmd/goose```

To run db migration
```
cd migration
goose -env local up
```
To generate model
```
cd internal
swagger generate model --name=[model-name] --spec=../docs/swagger.yml
eg:
swagger generate model --name=UserPayload --spec=../docs/swagger.yml
```
`model-name` is the models given in swagger file. eg: UserPayload
## ToDo:
 - Create architecture diagram.
 - Create swagger file for all api specification.
 - Database design for the app. Please make a doc first in docs directory.
 - Implement logger using logrus package.

# Contributions
 - Checkout branch from develop and make PR to develop.