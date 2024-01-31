Steps to running service : 
1. Provide postgreSQl database local, that must be set in file .env
2. run commmand: **go installs** to get all package that used in service
3. run commmand: **go run src/infra/cmd/migrate.go** to automatic create tables
4. run commmand: **go run main.go** to running services
