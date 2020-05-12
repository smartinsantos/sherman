# go-auth-api
Identity service written in go

## Documentation
- Linter: https://github.com/golangci/golangci-lint
- Watcher: https://github.com/cosmtrek/air
- Migrations: https://github.com/pressly/goose
- Pack binaries with assets: https://github.com/gobuffalo/packr/tree/master/v2
- Dependency injection container: https://github.com/sarulabs/di

## Bin
- ```bin/setup```                       : Install all project dependencies
- ```bin/lint```                        : Finds lint errors in the application
- ```bin/watch```                       : Runs the server on watch mode
- ```bin/run```                         : Runs the server
- ```bin/build```                       : Build server binary
- ```bin/migration [migration-name]```  : Creates a migration
- ```bin/migrate [command]```           : Migrate the DB see cmd list
- List of commands
    - up                   Migrate the DB to the most recent version available
    - up-by-one            Migrate the DB up by 1
    - up-to [version]        Migrate the DB to a specific VERSION
    - down                 Roll back the version by 1
    - down-to [version]     Roll back to a specific VERSION
    - redo                 Re-run the latest migration
    - reset                Roll back all migrations
    - status               Dump the migration status for the current DB
    - version              Print the current version of the database


