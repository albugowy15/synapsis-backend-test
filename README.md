# Synapsis Backend Test

## Documentation

For more detailed documentation on each endpoint, including request and response formats, refer to the [Swagger documentation](https://synapsis-backend-test.fly.dev/swagger/index.html) provided by the API.

## How To Run

### Prerequisites

- Go 1.21.3 +
- Docker
- Makefile

### Steps

1. Copy app configuration

```sh
cp app.env.example app.env
```

2. Create docker network
   Before running the application, create a Docker network to ensure proper communication between containers.

```sh
make network
```

3. Pull and Run PostgreSQL Container
   Pull the PostgreSQL Docker image and run it as a container.

```sh
make postgres
```

4. Create database `synapsis_db`
   Use the following command to create the PostgreSQL database named `synapsis_db`.

```sh
make createdb
```

5. Run Database Migration
   Perform database migration to set up the schema.

```sh
make migrate_up
```

6. Seed Database
   Populate the database with initial data (if applicable).

```sh
make seed
```

7. Build the API Image
   Build the Docker image for the API.

```sh
make build_api
```

8. Run the API Container
   Launch the Docker container for the API.

```sh
make server_api
```

## Seeder Data
Inside `database/seeder` directory, you'll find database seeder files. These files contain data to populate the database. To test the protected endpoint easily, select a username from the user seeder file. You can use the password `password1for&All` for all usernames. For example:
```json
{
  "username": "angeloschmitt",
  "password": "password1for&All"
}
```
Use above credentials to Login.

## Database Schema

![Database Schema](./synapsis-backend-test.png)

With dbdiagram.io => https://dbdocs.io/kholidbughowi/synapsis-backend-test-database?view=relationships

## Deployment

This project has been deployed on the [Fly.io](https://fly.io/) platform, which offers reliable deployment services. It utilizes [Neon DB](https://neon.tech/) as its database solution to effectively manage and store data.

## Docker Image

Feel free to retrieve the Docker image from the following repository: https://hub.docker.com/r/albugowy15/synapsis-backend-test. This image is readily available for pulling and utilization within your environment.
