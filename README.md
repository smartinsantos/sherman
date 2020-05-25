# Sherman Starter Kit
A Go microservice implementation following Robert "Uncle Bob" Clean Architecture

# Application dependencies
- Golang: https://golang.org/doc/install
- Docker: https://docs.docker.com/get-docker/

# First time setup
- run ```bin/setup```

# To run the application
- Terminal [ 1 ]
    - run ```bin/up```
- Terminal [ 2 ]
    - if first time setup run ```bin/migrate up```
    - run ```bin/watch```
    
## Bin command reference
- ```bin/up```                              : Builds && spins up docker containers
    - List of commands https://docs.docker.com/compose/reference/up/  
- ```bin/down```                            : Stops docker containers
    - List of commands https://docs.docker.com/compose/reference/down/
- ```bin/setup```                           : Install all project dependencies
- ```bin/lint```                            : Finds lint errors in the application
- ```bin/watch```                           : Runs the server on watch mode
- ```bin/run```                             : Runs the server
- ```bin/new-migration [migration-name]```  : Creates a migration
- ```bin/migrate [command]```               : Migrate the DB see cmd list
- ```bin/prod```                            : Setup/Build/Run application for production
    - List of commands
        - up                   Migrate the DB to the most recent version available
        - up-by-one            Migrate the DB up by 1
        - up-to [version]      Migrate the DB to a specific VERSION
        - down                 Roll back the version by 1
        - down-to [version]    Roll back to a specific VERSION
        - redo                 Re-run the latest migration
        - reset                Roll back all migrations
        - status               Dump the migration status for the current DB
        - version              Print the current version of the database

## Go Packages reference
- Linter: https://github.com/golangci/golangci-lint
- Watcher: https://github.com/cosmtrek/air
- MYSQL Driver: https://github.com/go-sql-driver/mysql
- Migrations: https://github.com/pressly/goose
- Dependency injection container: https://github.com/sarulabs/di
- Echo web framework: https://echo.labstack.com/guide
- JWT tokens: https://github.com/dgrijalva/jwt-go
- Logger: https://github.com/rs/zerolog
