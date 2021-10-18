<template>
  <div class="container">
    <table>
      <thead>
        <th>Order name</th>
        <th>Customer Company</th>
        <th>Customer name</th>
        <th>Order date</th>
        <th>Delivered Amount</th>
        <th>Total Amount</th>
      </thead>
      <tbody>
        <tr v-for="order in orders" :key="order['OrderName']">
          <td>{{ order["OrderName"] }}</td>
          <td>{{ order["CustomerCompany"] }}</td>
          <td>{{ order["CustomerName"] }}</td>
          <td>{{ order["OrderDate"] }}</td>
          <td>{{ order["DeliveredAmount"] }}</td>
          <td>{{ order["TotalAmount"] }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  name: "Table",
  props: { orders: Array },
};
</script>

<style>
table {
  font-family: Arial, Helvetica, sans-serif;
  border-collapse: collapse;
  width: 100%;
}
td,
th {
  border: 1px solid #ddd;
  padding: 8px;
}
tr:hover {
  background-color: #ddd;
}
</style>
