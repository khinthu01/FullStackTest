<template>
  <div class="container">
    <table class="table table-striped table-bordered">
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
.container {
}
</style>
