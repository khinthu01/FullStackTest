import psycopg2
from decouple import config


password = config('PASSWORD')
connection = psycopg2.connect(host="localhost", port=5432, dbname='postgres', user='postgres', password=password)

connection.autocommit = True
db_cursor = connection.cursor()

db_sql = '''CREATE database customerOrders'''

db_cursor.execute(db_sql)

db_cursor.close()
connection.close()

connection = psycopg2.connect(host="localhost", port=5432, dbname='customerorders', user='postgres', password=password)

db_cursor = connection.cursor()

create_tbl_commands = (
    """
    CREATE TABLE deliveries (
        id INTEGER PRIMARY KEY,
        order_item_id INTEGER NOT NULL,
        delivered_quantity INTEGER NOT NULL
    )
    """,
    """
    CREATE TABLE order_items (
        id INTEGER PRIMARY KEY,
        order_id INTEGER NOT NULL,
        price_per_unit FLOAT(10),
        quantity INTEGER NOT NULL,
        product VARCHAR(255) NOT NULL
    )
    """,
    """
    CREATE TABLE orders (
        id INTEGER PRIMARY KEY,
        created_at TIMESTAMP NOT NULL,
        order_name VARCHAR(10),
        customer_id VARCHAR(10)
    )
    """
)

for command in create_tbl_commands:
    db_cursor.execute(command)


file1 = "test_data (2)/test_data (2)/Test task - Postgres - deliveries.csv"
with open(file1, 'r') as file:
    next(file)
    db_cursor.copy_from(file, "deliveries", null="", columns=("id", "order_item_id", "delivered_quantity"), sep=",")

file2 = "test_data (2)/test_data (2)/Test task - Postgres - order_items.csv"
with open(file2, 'r') as file:
    next(file)
    db_cursor.copy_from(file, "order_items", null="", columns=("id", "order_id", "price_per_unit", "quantity", "product"), sep=",")

file3 = "test_data (2)/test_data (2)/Test task - Postgres - orders.csv"
with open(file3, 'r') as file:
    next(file)
    db_cursor.copy_from(file, "orders", null="", columns=("id", "created_at",  "order_name", "customer_id"), sep=",")


db_cursor.close()

connection.commit()
connection.close()