# Philippine Standard Geographic Code (PSGC) REST API

[![Star](https://img.shields.io/github/stars/ej-agas/ph-locations.svg?style=flat-square)](https://github.com/ej-agas/ph-locations/stargazers) [![License](https://img.shields.io/github/license/ej-agas/ph-locations.svg?style=flat-square)](https://github.com/ej-agas/ph-locations/blob/main/LICENSE) [![Release](https://img.shields.io/github/release/ej-agas/ph-location.svg?style=flat-square)](https://github.com/ej-agas/ph-locations/releases) 

## 📖 API Documentation
The API Documentation can be found by accessing the url
```
http://localhost:6945/docs/index.html
```
**Note:** The host and port will change when you set a custom host and port in the `build/app.env` file.

## 💻 Run locally

### 📝 Prerequisites
1. [git](https://git-scm.com/) Installed in your system
2. [Docker](https://www.docker.com/) Installed in your system
3. [golang-migrate/migrate](https://github.com/golang-migrate/migrate) Installed in your system
5. [ej-agas/psgc-publication-parser](https://github.com/ej-agas/psgc-publication-parser) Installed in your system
5. [GNU Make](https://www.gnu.org/software/make) (Optional)

### 📝 Steps
Clone the repository
```shell
git clone git@github.com:ej-agas/ph-locations.git && cd ph-locations
```

Copy `.env.example` and `docker-compose.yaml.dist`
```shell
cp build/app.env.example build/app.env
cp build/docker-compose.yaml.dist build/docker-compose.yaml
```
(Optional) Replace the default password of the PostgreSQL user in the `build/app.env` file you've just created.
```.dotenv
POSTGRES_PASSWORD=SET_A_SECURE_PASSWORD
```
Run make docker
```shell
make docker
```
If you don't have GNU make installed in your system,  run the following command
```shell
docker build -t ph-locations:latest . -f build/Dockerfile.dev && docker compose -f build/docker-compose.yaml up -d --build
```
Try to access the root path
```shell
curl localhost:6945
```
If you get the JSON below, then everything went well
```json
{
  "status": 200,
  "message": "Philippine Standard Geographic Code (PSGC) REST API"
}
```

Enter the `ph_locations_db` container by running the command:
```shell
docker exec -it ph_locations_db bash
```
Inside the container, enter the PostgreSQL shell by running the command:
```shell
psql --user ph_locations_user --password -d postgres
```
***Note:*** The shell will prompt for a password, use the `POSTGRES_PASSWORD` value in your `build/app.env` file

Once inside, execute a `CREATE DATABASE` SQL query.
```postgresql
CREATE DATABASE ph_locations_db;
```
Exit the PostgreSQL shell by running exit and also for the `ph_locations_db` container
```shell
exit
```

Run the migrations using [golang-migrate/migrate](https://github.com/golang-migrate/migrate)
```shell
migrate -source file://migrations -database postgresql://ph_locations_user:YOUR_DATABASE_PASSWORD@localhost:5432/ph_locations_db?sslmode=disable up
```
***Note:*** Change the `YOUR_DATABASE_PASSWORD` string in the command to the `POSTGRES_PASSWORD` value in your `build/app.env` file.

Run the PSGC publication parser 
```shell
psgc-pub-parser parse publication_datafile.xlsx  --host localhost --port 5432 --db ph_locations_db --user ph_locations_user --password YOUR_DATABASE_PASSWORD
```
***Note:*** Change the `YOUR_DATABASE_PASSWORD` flag value in the command to the `POSTGRES_PASSWORD` value in your `build/app.env` file.
