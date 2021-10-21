from pymongo import MongoClient
import psycopg2
from decouple import config

client = MongoClient("localhost", 27017, maxPoolSize=50)
db = client['CustomerOrders']
companies_col = db['customer_companies']

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

order_list = []


for order in order_items:
    i = 0
    while i < len(orders) and order[i][0] != order[1]:
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
    amount = int(order[3])*int(order[2])

    list_item = {
        "order_name": order_name,
        "customer_company": company_name,
        "customer_name": customer_name,
        "order_date": order_date,
        "delivered_amount": amount,
        "total_amount": amount
    }

    order_list.append(list_item)

