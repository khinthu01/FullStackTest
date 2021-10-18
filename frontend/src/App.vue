<template>
  <img alt="Vue logo" src="./assets/logo.png" />
  <Table :orders="orders" />
</template>

<script>
import axios from "axios";
import Table from "./components/Table";

export default {
  name: "App",
  components: {
    Table,
  },
  data() {
    return {
      orders: [],
    };
  },
  async created() {
    const res = await axios.get("http://localhost:8080/orderList", {
      responseType: "json",
    });
    const orders = res.data;
    this.orders = orders;
  },
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
