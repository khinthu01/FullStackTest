from pymongo import MongoClient
import pandas as pd


class MongoDB():
    def __init__(self, dBName):
        self.dBName = dBName

        self.client = MongoClient("localhost", 27017, maxPoolSize=50)
        self.DB = self.client[self.dBName]


    def insertData(self, filepath, collectionName):
        df = pd.read_csv(filepath)
        data = df.to_dict('records')

        collection = self.DB[collectionName]
        collection.insert_many(data, ordered=False)
        print("All the data has been exported to Mongo DB server.")

if __name__ == "__main__":
    mongodb = MongoDB(dBName = 'CustomerOrders')
    mongodb.insertData('test_data (2)/test_data (2)/Test task - Mongo - customer_companies.csv', collectionName= 'customer_companies')
    mongodb.insertData('test_data (2)/test_data (2)/Test task - Mongo - customers.csv', collectionName= 'customers')
    mongodb.insertData('test_data (2)/test_data (2)/Test task - Orders.csv', collectionName='orders')
    