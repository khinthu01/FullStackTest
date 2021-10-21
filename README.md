# Full Stack Software Engineer Test
This repository contains code for a simple web application that displays all orders in a database. The frontend was implemented using Vue.js and the backend was implemented using Golang. 

## Setting Up Databases
First install pymongo, pandas, psycopg2 and decouple while in the main directory using the following commands:

```
pip install pymongo
pip install pandas
pip install psycopg2
pip install decouple
```

Then set up a .env file in the parent directory and create an environment variable called PASSWORD to hold the password for your postgres account. 
Run files import_data_mongodb.py and import_data_postgres.py. Then run create_results.py. This will create a new table in your postgres database to hold the orders data in the format that will be presented. 

## Setup Instructions - Backend
