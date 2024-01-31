Steps to running service : 
1. Provide postgreSQl database local, that must be set in file .env
  DB_HOST=localhost
  DB_PORT=5432
  DB_NAME=dating-services
  DB_USERNAME=postgres
  DB_PASSWORD=admin
  DB_SCHEMA=public
  DB_SSL_MODE=disable
  DB_DIALECT=postgres

2. run commmand: **go installs** to get all package that used in service
3. run commmand: **go run src/infra/cmd/migrate.go** to automatic create tables
4. run commmand: **go run main.go** to running services
