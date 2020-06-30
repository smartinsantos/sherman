# Sherman Starter Kit
[![Go Report Card](https://goreportcard.com/badge/github.com/smartinsantos/sherman)](https://goreportcard.com/report/github.com/smartinsantos/sherman)
[![Maintainability](https://api.codeclimate.com/v1/badges/df184224058c1f4c0b9e/maintainability)](https://codeclimate.com/github/smartinsantos/sherman/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/df184224058c1f4c0b9e/test_coverage)](https://codeclimate.com/github/smartinsantos/sherman/test_coverage)

Sherman is a Golang starter kit to develop API microservices. It is designed to help kickstart a project, avoiding the boilerplate overhead.
It follows SOLID principles and attempts to follow Robert "Uncle Bob" Clean Architecture. 

## Features
- Fully "Dockerized" application.
- Endpoints for user authentication.
- JWT authentication and refresh token based session.
- Request marshaling and data validation.
- Mysql/SQLite3 Database with Migrations support.
- Application configuration thru .env file.
- Dependency injection container to handle inversion of control with ease.
- Tests
    - Interface mocks generator.
    - Mocked database query tests.
    - Coverage reports.
- Linter setup/configuration featuring golangci-lint.
- Development
    - Watcher rebuilds application on file changes
    - Pretty logs

## Getting Started
### Set up application for the first time
1. Install Docker if you haven't already https://docs.docker.com/get-docker/
2. Create an .env configuration on the root folder of the project the .env.example file provided contains all necessary configurations. 
3. Run ```bin/init```. This command will install all required dependencies, run migrations, spin up docker containers and run the application.
4. The application will run on the configured port, by default ```APP_PORT=5000```

### Stop the application
- Run ```bin/down```

### Run the application
- Run ```bin/up```

### Tests
- `bin/exec test` will automatically run all test files in the project and generate coverage files under ./coverage
- `bin/exec test [test_path]` will run tests in the provided path (no coverage will be generated) 
- When creating new tests files include the provided testing package ```src/app/testing``` like so ```import _ "[module]/src/app/testing"```. This will change the test working dir to the specified ROOT_DIR.
    
### /bin scripts reference
- ```bin/init```                            : Initialize/Reset containers && database
- ```bin/up```                              : Builds and/or spins up docker containers https://docs.docker.com/compose/reference/up/  
- ```bin/down```                            : Stops docker containers https://docs.docker.com/compose/reference/down/
- ```bin/go```                              : Run go commands in the app container (**docker container must be running)
- ```bin/mockery```                         : Generates mocks for every interface in the project under ./src/mocks
- ```bin/lint```                            : Finds lint errors in the application
- ```bin/exec```                            : Execs the following commands (**docker containers must be running)
    - ```setup```                           : Install all project dependencies
    - ```gofmt```                           : Formats .go files in /src folder
    - ```test```                            : Runs test suites
    - ```watch```                           : Runs the server on watch mode
    - ```new-migration [migration-name]```  : Creates a migration
    - ```migrate [command]```               : Migrate the DB see cmd list
        - ```up```                          : Migrate the DB to the most recent version available
        - ```up-by-one```                   : Migrate the DB up by 1
        - ```up-to [version]```             : Migrate the DB to a specific VERSION
        - ```down```                        : Roll back the version by 1
        - ```down-to [version]```           : Roll back to a specific VERSION
        - ```redo```                        : Re-run the latest migration
        - ```reset```                       : Roll back all migrations
        - ```status```                      : Dump the migration status for the current DB
        - ```version```                     : Print the current version of the database


## Project Structure
```
--bin [scripts see list of scripts]
    --cmd [scripts that run inside the containers using bin/exec]
--mocks [mock interface implementations]
--src
    --app
        --config [app configuration setup]
        --database [database related (connection, migrations, etc)]
        --registry [dependency injection container]
        --router [app router, routes/middleware setup]
        --testing [app testing package]
        --utils
          -- response [app specific http response struct]
          -- terr [app specific typed errors]
    --delivery [interface adapters layer]
    --domain [entities/aggregates layer]
    --repository [data layer]
    --service [globally available sevices (middleware, validators, etc)]
    --usecase [use case layer]
main.go [entry point]
```

## Main application dependencies
- Docker: [docs.docker.com/get-docker/](https://docs.docker.com/get-docker/)
- JWTs: [github.com/dgrijalva/jwt-go](https://github.com/dgrijalva/jwt-go)
- MYSQL Driver: [github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)
- SQLITE3 Driver: [github.com/mattn/go-sqlite3](https://github.com/mattn/go-sqlite3)
- Migrations: [github.com/pressly/goose](https://github.com/pressly/goose)
- UUIDs: [github.com/google/uuid](https://github.com/google/uuid)
- Loader .env: [github.com/joho/godotenv](https://github.com/joho/godotenv)
- Echo Web Framework: [echo.labstack.com/guide](https://echo.labstack.com/guide)
- Logger: [github.com/rs/zerolog](https://github.com/rs/zerolog)
- Dependency Injection container: [github.com/sarulabs/di](https://github.com/sarulabs/di)
- Tests: [github.com/stretchr/testify](https://github.com/stretchr/testify)
- Sql Mocks: [github.com/DATA-DOG/go-sqlmock](https://github.com/DATA-DOG/go-sqlmock)
- Mockery: [github.com/vektra/mockery](https://github.com/vektra/mockery)
- Linter: [github.com/golangci/golangci-lint](https://github.com/golangci/golangci-lint)
- Code Change Watcher: [https://github.com/cosmtrek/air](https://github.com/cosmtrek/air)

## License
Sherman is licensed under the MIT license. Check the [LICENSE](LICENSE) file for details.

## Author
[Sergio Martin](https://smartinsantos.github.io/)