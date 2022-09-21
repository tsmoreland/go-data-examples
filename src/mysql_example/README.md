# MYSQL Example code in Go


## Database Migrations

The migrations are handled using golang-migrate which can be installed using 

- windows: ```scoop install migrate``` 
- mac: ```brew install golang-migrate```
- linux: run the following commands
    1. curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | sudo bash
    2. sudo apt-get update
    3. sudo apt-get install migrate

Create initial migration using:

```migrate create -ext sql -dir db/migrations -seq init_schema```

where init_schema is the schema name (or database name) and ```db/migrations``` is where the migrations will be stored.  This creates 2 SQL files one for applying the migration (up) and one for removing it (down)


### Applying Migrations

Run the following command:

```migrate -path db/migration -database "mysql://user:pass@tcp(hostname:port)/schema" -verbose up```

