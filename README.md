# Sherman Starter Kit
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
- Docker: https://docs.docker.com/get-docker/
- Go SqlMock: https://github.com/DATA-DOG/go-sqlmock  
- JWT tokens: https://github.com/dgrijalva/jwt-go
- MYSQL Driver: https://github.com/go-sql-driver/mysql
- SQLITE3 Driver: https://github.com/mattn/go-sqlite3 
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

## License
Sherman is licensed under the MIT license. Check the [LICENSE](LICENSE) file for details.

## Author
[Sergio Martin](https://smartinsantos.github.io/)