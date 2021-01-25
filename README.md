![Mercado Pago](https://seeklogo.com/images/M/mercado-pago-logo-CC340D0497-seeklogo.com.png)

``` sh
$ git clone git@github.com:wuilkysb/ml-test.git $GOPATH/ml-mutant-test && cd $_
```


#  Overview
API to detect mutants

# Requirements

* go v1.15
* go module
* postgres

# Build

* Install dependencies: 
```sh
$ go mod download
```

* Run test:
```sh 
$ go test ./... 
```

# Environments
#### Required environment variables

* `PORT`: port for the server
* `PG_USER`: postgres user
* `PG_PASSWORD`: postgres password
* `PG_NAME`: postgres name
* `PG_HOST`: postgres host
* `PG_PORT`: postgres port
* `PG_TIMEOUT`: postgres writer and dial timeout, example: 10s
* `PG_POOL_SIZE`: postgres pool size, example: 5
* `MIGRATIONS_COMMAND`: command for migrations (run, rollback, reset)

# Endpoints

* verify if a dna is mutant:
``` curl --location --request POST 'https://ml-mutant-w.herokuapp.com/mutant/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "dna":["ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"]
}'
```

* get stats:
``` curl --location --request GET 'https://ml-mutant-w.herokuapp.com/mutant/stats' ```