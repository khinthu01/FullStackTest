<template>
  <input type="text" v-model="search" placeholder="search orders" />
  <Table :currentPageItems="currentPageItems" />
  <span>Page:</span>
  <button
    v-for="i in paginatedItems.length"
    :key="i"
    class="btn btn-secondary mr-1"
    :class="{ 'btn-success': currentPage === i }"
    @click="changePage(i)"
  >
    {{ i }}
  </button>
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
      currentPage: 1,
      itemsPerPage: 5,
      orders: [],
      search: "",
    };
  },
  async created() {
    const res = await axios.get("http://localhost:8080/orderList", {
      responseType: "json",
    });
    const orders = res.data;
    this.orders = orders;
  },
  computed: {
    filteredOrders() {
      return this.orders.filter((order) => {
        return order["OrderName"].match(this.search);
      });
    },
    paginatedItems() {
      let page = 1;
      return [].concat.apply(
        [],
        this.filteredOrders.map((item, index) =>
          index % this.itemsPerPage
            ? []
            : {
                page: page++,
                orders: this.filteredOrders.slice(
                  index,
                  index + this.itemsPerPage
                ),
              }
        )
      );
    },
    currentPageItems() {
      let currentPageItems = this.paginatedItems.find(
        (pages) => pages.page == this.currentPage
      );
      return currentPageItems ? currentPageItems.orders : [];
    },
  },
  methods: {
    changePage(pageNumber) {
      if (pageNumber != this.currentPage) {
        this.currentPage = pageNumber;
      }
    },
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
