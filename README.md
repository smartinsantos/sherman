# Sherman Starter Kit
Sherman is a Golang starter kit to develop API microservices. It is designed to help kickstart a project, avoiding the boilerplate overhead.

It follows SOLID principles, inspired by Robert "Uncle Bob" Clean Architecture. 

## Features
- Fully "Dockerized" application.
- Endpoints for user authentication.
- JWT authentication and refresh token based session.
- Request marshaling and data validation.
- Mysql Database with Migrations support.
- Application configuration thru .env file.
- Dependency injection container to handle inversion of control with ease.
- Tests
    - Interface mocks generator.
    - Mocked database query tests.
    - Coverage reports.
    - Complete test coverage.
- Linter setup/configuration featuring golangci-lint.
- Development
    - Watcher rebuilds application on file changes
    - Pretty logs

## Application dependencies
- Docker: https://docs.docker.com/get-docker/
- Go SqlMock: https://github.com/DATA-DOG/go-sqlmock  
- JWT tokens: https://github.com/dgrijalva/jwt-go
- MYSQL Driver: https://github.com/go-sql-driver/mysql
- UUIDs: github.com/go-sql-driver/mysql
- DotEnv: github.com/joho/godotenv
- Echo web framework: https://echo.labstack.com/guide
- Logger: https://github.com/rs/zerolog
- Dependency injection container: https://github.com/sarulabs/di
- Testify: https://github.com/stretchr/testify
- Linter: https://github.com/golangci/golangci-lint
- Watcher: https://github.com/cosmtrek/air
- Migrations: https://github.com/pressly/goose
- Mockery: https://github.com/vektra/mockery

## Getting Started
### Set up application for the first time
- run ```bin/init```

### Run the application
- run ```bin/up```

### Connect from a mysql client to the database
```
DB_NAME=[your-database-name]
DB_USER=[your-user-name]
DB_PASS=[your-user-password]
DB_HOST=localhost
DB_PORT=5001 * Mapped host 5001:3306 container
```   

### Tests
- `bin/exec test` will automatically run all test files in the project and generate coverage files under ./coverage
- `bin/exec test [test_path]` will run tests in the provided path (no coverage will be generated) 
- when creating new tests files include testing package on every test file. this will change the test working dir to the specified ROOT_DIR ```import _ "[module]/src/app/testing"```
    
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
    - ```lint```                            : Finds lint errors in the application
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
