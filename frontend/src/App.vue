<template>
  <img alt="Vue logo" src="./assets/logo.png" />
  <div id="App"></div>
</template>

<script>
export default {
  name: "App",
  components: {},
  async getOrders() {
    const res = await fetch("http://localhost:8080/orderList");
    const { orders } = await res.json();

    console.log(orders);
    return { orders };
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
