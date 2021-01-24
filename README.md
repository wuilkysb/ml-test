![RappiOne](https://one.rappi.com/assets/images/landing/logo.svg)

``` sh
$ git clone git@bitbucket.org:rappinc/one-catalog-management.git $GOPATH/src/bitbucket.org/rappinc/one-catalog-management && cd $_
```


#  Overview
Management for RappiOne catalog history

# Requirements

* go v1.13
* go module
* postgres
* redis

# Build

* Install dependencies: 
```sh
$ go mod download
```

* Generate Swagger doc:
```sh
$ swag init
```

* Run test:
```sh 
$ go test ./... 
```

# Environments
#### Required environment variables

* `SERVER_HOST`: host for the server
* `SERVER_PORT`: port for the server
* `PG_USER`: postgres user
* `PG_PASSWORD`: postgres password
* `PG_NAME`: postgres name
* `PG_HOST`: postgres host
* `PG_PORT`: postgres port
* `PG_TIMEOUT`: postgres writer and dial timeout, example: 10s
* `PG_TIMEOUT`: postgres timeout to support failover
* `PG_POOL`: postgres pool size
* `RUN_MIGRATIONS`: true or false string to run migrations
* `ROLLBACK_MIGRATIONS`: true or false string to rollback migrations
* `RESET_MIGRATIONS`: true or false string to reset all migrations
* `POSTFIX`: rappi ansible postfix
* `SIGNALFX_SERVICE_NAME`: signal service name
* `SIGNALFX_ENDPOINT_URL`: signal service url
* `X_APPLICATION_ID`: rappi ansible x_application_id

# Contributors

* Moisés Morillo
* Wuilkys Becerra
* Alirio Gutiérrez
* Piero Pinzon

# License

This project is property of Rappi
