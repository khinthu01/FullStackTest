<template>
  <input
    type="text"
    v-model="search"
    placeholder="search orders"
    @change="searchOrders"
  />
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
      datefilter: "",
    };
  },
  async created() {
    const searchQuery = encodeURIComponent(this.search);
    const url = `http://localhost:8000/orders?s=${searchQuery}`;

    const res = await axios.get(url, {
      responseType: "json",
    });
    const orders = res.data;
    this.orders = orders;
  },

  computed: {
    paginatedItems() {
      let page = 1;
      return [].concat.apply(
        [],
        this.orders.map((item, index) =>
          index % this.itemsPerPage
            ? []
            : {
                page: page++,
                orders: this.orders.slice(index, index + this.itemsPerPage),
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
    async searchOrders() {
      const searchQuery = encodeURIComponent(this.search);
      const url = `http://localhost:8000/orders?s=${searchQuery}`;

      const res = await axios.get(url, {
        responseType: "json",
      });
      const orders = res.data;
      this.orders = orders;
    },
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
