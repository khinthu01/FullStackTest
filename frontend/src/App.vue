<template>
  <div id="main">
    <span id="search">
      <h3>Search</h3>
      <input
        type="text"
        v-model="search"
        placeholder="search orders"
        @change="searchOrders"
        id="search-box"
      />
    </span>

    <span id="date-filter">
      <input type="date" v-model="start" />
      <input type="date" v-model="end" @change="dateFilter" />
    </span>

    <Table :currentPageItems="currentPageItems" />
    <span id="pages"
      >Page:
      <button
        v-for="i in paginatedItems.length"
        :key="i"
        @click="changePage(i)"
      >
        {{ i }}
      </button>
    </span>
  </div>
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
      start: "",
      end: "",
      url: `http://localhost:8000/orders`,
    };
  },
  async created() {
    // initial loading of all orders

    const res = await axios.get(this.url, {
      responseType: "json",
    });
    const orders = res.data;
    this.orders = orders;
  },

  computed: {
    paginatedItems() {
      // paginating orders
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
      // determines the orders that should be shown on current page
      let currentPageItems = this.paginatedItems.find(
        (pages) => pages.page == this.currentPage
      );
      return currentPageItems ? currentPageItems.orders : [];
    },
  },
  methods: {
    async searchOrders() {
      // function is called every time the input in the search bar is changed and search terms are used to send query to backend
      const searchQuery = encodeURIComponent(this.search);

      if (this.start != "" && this.end != "") {
        this.url = this.url + `&s=${searchQuery}`;
      } else {
        this.url = this.url + `?s=${searchQuery}`;
      }

      const res = await axios.get(this.url, {
        responseType: "json",
      });
      const orders = res.data;
      this.orders = orders;
    },
    async dateFilter() {
      // function is called every time the input in the search bar is changed and search terms are used to send query to backend
      const start = encodeURIComponent(this.start);
      const end = encodeURIComponent(this.end);
      if (this.search != "") {
        this.url = this.url + `&start=${start}&end=${end}`;
      } else {
        this.url = this.url + `?start=${start}&end=${end}`;
      }

      const res = await axios.get(this.url, {
        responseType: "json",
      });
      const orders = res.data;
      this.orders = orders;
    },
    changePage(pageNumber) {
      // changes the page being viewed
      if (pageNumber != this.currentPage) {
        this.currentPage = pageNumber;
      }
    },
  },
};
</script>

<style lang="scss">
@import "~@/assets/scss/vendors/bootstrap-vue/index";
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
#main {
  padding: 2em;
  padding-top: 0em;
}
#search {
  display: flex;
  margin: auto;
}

#search-box {
  margin: 0.5em;
  width: 100%;
  padding-left: 1em;
}

#date-filter {
  display: flex;
  margin-top: 0.5em;
  margin-bottom: 1em;
}

#pages {
  display: flex;
  justify-content: center;
  margin-top: 1em;
}
</style>
