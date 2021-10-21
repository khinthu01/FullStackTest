# Full Stack Software Engineer Test

This repository contains code for a simple web application that displays all orders in a database. The frontend was implemented using Vue.js, the backend was implemented using Golang and the databases were populated using python scripts as such to run this project you will need to install python, pip, npm, vue.js, and go.

## Setting Up Databases

First install pymongo, pandas, psycopg2 and decouple while in the main directory using the following commands:

```
pip install pymongo
pip install pandas
pip install psycopg2
pip install decouple
```

Then set up a .env file in the parent directory and create an environment variable called PASSWORD to hold the password for your postgres account. Check that the user name, host, and ports are correct in the import_data_postgres.py file.
Run files import_data_mongodb.py and import_data_postgres.py. Then run create_results.py. This will create a new table in your postgres database to hold the orders data in the format that will be presented.

## Setup Instructions - Backend

Install the necessary modules using the following commands:

```
go get github.com/gofiber/fiber/v2
```

```
go get github.com/lib/pq
```

```
go mod tidy
```

For main.go to access the postgres database with the results table input your database name, user name, password, host and port on line 27.

```go

    const (
	    host = "localhost"
	    port = 5432
	    user = "postgres"
	    dbname = "customerorders"
	    password = "#password"
    )

```

Then run main.go to run the backend.

```
go run main.go
```

## Setup Instructions - Frontend

In a separate terminal change directory to the frontend folder and install the required dependencies using the following command:

```
npm install
```

Then run the frontend with the following command:

```
npm run serve
```
