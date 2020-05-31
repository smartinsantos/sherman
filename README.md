# Sherman Starter Kit
A Go microservice implementation following Robert "Uncle Bob" Clean Architecture

### Application dependencies
- Docker: https://docs.docker.com/get-docker/ (Development)

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
    
### Bin command reference
- ```bin/init```                            : Initialize/Reset containers && database
- ```bin/up```                              : Builds and/or spins up docker containers https://docs.docker.com/compose/reference/up/  
- ```bin/down```                            : Stops docker containers https://docs.docker.com/compose/reference/down/
- ```bin/go```                              : run go commands in the app container
- ```bin/mockery```                         : generates mocks for every interface in the project under ./src/mocks
- ```bin/exec```
    - ```setup```                           : Install all project dependencies
    - ```lint```                            : Finds lint errors in the application
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

## Go Packages reference
- Linter: https://github.com/golangci/golangci-lint
- Watcher: https://github.com/cosmtrek/air
- MYSQL Driver: https://github.com/go-sql-driver/mysql
- Migrations: https://github.com/pressly/goose
- Dependency injection container: https://github.com/sarulabs/di
- Echo web framework: https://echo.labstack.com/guide
- JWT tokens: https://github.com/dgrijalva/jwt-go
- Logger: https://github.com/rs/zerolog
