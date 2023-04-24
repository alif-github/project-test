
# Project Test API (With JWT Authentication)

The Project Test API is contains some API for get data job position from companies. This Project use stacks such HTTP Rest API and JWT Authentication also for access main API.


## Tech Stack

**Server:** Golang, JWT Authentication, gorilla mux HTTP


## Run Locally

#### 1. Local environment
- Make sure postgresql has installed, else you can download https://www.postgresql.org/download/
- Create new connection on db, you can use pgAdmin or access bin folder postgresql. more on https://www.enterprisedb.com/postgres-tutorials/connecting-postgresql-using-psql-and-pgadmin
- Create new folder for save source repository
- Clone the project
```bash
  git clone https://github.com/alif-github/project-test.git
```
- Go to the project directory
```bash
  cd project-test
```
- Edit location config development on file .env to source folder
- Run server, as the default environment will be "development"
```bash
  go run main.go development
```
- Server will be running on http://localhost:9078

#### 2. Docker environment
- Clone the project
```bash
  git clone https://github.com/alif-github/project-test.git
```
- Go to the project directory
```bash
  cd project-test
```
- Compose docker compose file inside
```bash
  docker compose -f ./project-compose.yaml up
```
- Server will be running on http://localhost:8080

## API Reference

#### 1. Register User

```http
  POST /v1/project/auth/register
```
##### Request Body JSON (Example)
```json
{
    "full_name": "Alif Yudha",
    "username": "alif15",
    "password": "Alif123@"
}
```
##### Response Body JSON (Example)
```json
{
    "success": true,
    "request_id": "838b6dbc-0b0e-44e0-bc94-d792859b6a5a",
    "message": "Sukses Register User Baru",
    "data": null
}
```

#### 2. Login
```http
  POST /v1/project/auth/login
```
##### Request Body JSON (Example)
```json
{
    "username": "alif15",
    "password": "Alif123@"
}
```
##### Response Body JSON (Example)
```json
{
    "success": true,
    "request_id": "0c287f28-0d76-47c9-b065-5ece28b4bf68",
    "message": "Login Berhasil",
    "data": null
}
```

#### 3. Logout
```http
  GET /v1/project/auth/logout
```

#### 4. Get Detail Data
```http
  GET /v1/project/api/recruitment/position/${ID}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |

#### 5. Get List Data
```http
  GET /v1/project/api/recruitment/position
```

| Query Param   | Example Value | Description                       |
| :----------   | :------------ | :-------------------------------- |
| `page`        | `1`           | Pagination queue from 1 to many   |
| `description` | `golang`      | Key words for search position     |
| `location`    | `New york`    | Key words for search location     |
| `full_time`   | `false`       | Type of full time position        |