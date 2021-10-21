from pymongo import MongoClient
import psycopg2
from decouple import config

client = MongoClient("localhost", 27017, maxPoolSize=50)
db = client['CustomerOrders']
companies_col = db['customer_companies'].find({})

companies = []

for company in companies_col:
    companies.append(company)

customers_col = db['customers'].find({})

customers = []

for customer in customers_col:
    customers.append(customer)


password = config('PASSWORD')
connection = psycopg2.connect(host="localhost", port=5432, dbname='customerorders', user='postgres', password=password)

connection.autocommit = True
db_cursor = connection.cursor()

orders_sql = """SELECT * from public.orders"""

db_cursor.execute(orders_sql)
orders = db_cursor.fetchall()


order_items_sql = """SELECT * from public.order_items"""

db_cursor.execute(order_items_sql)
order_items = db_cursor.fetchall()

#order_list = []

create_table = (
    """
    CREATE TABLE orderlist (
        order_item_id INTEGER PRIMARY KEY,
        order_name VARCHAR(10),
        customer_company VARCHAR(30),
        customer_name VARCHAR(30),
        order_date TIMESTAMP NOT NULL,
        delivered_amount FLOAT(6),
        total_amount FLOAT(6)
    )
    """)

db_cursor.execute(create_table)

for order in order_items:
    i = 0
    while i < len(orders) and orders[i][0] != order[1]:
        i += 1
    
    order_name = orders[i][2]
    order_date = orders[i][1]
    customer_id = orders[i][3]

    i = 0
    while i < len(customers) and customers[i]['user_id'] != customer_id:
        i += 1
    
    customer_name = customers[i]['name']
    company_id = customers[i]['company_id']

    i = 0 
    while i < len(companies) and companies[i]['company_id'] != company_id:
        i += 1

    company_name = companies[i]['company_name']

    if order[3] == None or order[2] == None:
        amount = 0
    else:
        amount = int(order[3])*float(order[2])

    insert_command = (
        """INSERT INTO orderlist(
            order_item_id,
            order_name,
            customer_company,
            customer_name,
            order_date,
            delivered_amount,
            total_amount
        ) VALUES (
            %s, %s, %s, %s, %s, %s, %s
        )
        """)

    values = (
        order[0],
        order_name,
        company_name,
        customer_name,
        order_date,
        amount,
        amount
    )

    db_cursor.execute(insert_command, values)


db_cursor.close()

connection.commit()
connection.close()



