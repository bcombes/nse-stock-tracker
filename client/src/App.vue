<template>
  <div>
    <section class="hero is-info">
      <div class="hero-body">
        <div class="container">
          <h1 class="title">
            NSE Market Watch
          </h1>
          <h2 class="subtitle">
            Follow the market like a boss...
          </h2>
        </div>
      </div>
      <!-- Hero footer: will stick at the bottom -->
      <div class="hero-foot">
        <nav class="tabs is-boxed level">
          <div class="container level-left">
            <ul>
              <li class="is-active"><a>All</a></li>
              <li><a>Most Traded</a></li>
              <li><a>Most Active</a></li>
              <li><a>Biggest Gainers</a></li>
              <li><a>Biggest Losers</a></li>
            </ul>
            <div class="control level-right has-icons-right">
              <input class="input" type="text" placeholder="search" v-model="query">
              <span class="icon is-small is-right">
                <i class="fas fa-search"></i>
              </span>
            </div>
          </div>
        </nav>
      </div>
    </section>
    <section class="section">
      <div class="columns is-multiline is-mobile">
        <stock-card v-for="stock in filterStocksByQuery()" :key="stock.symbol" :stock="stock" @stock-card-clicked="onCardClicked"></stock-card>
      </div>
    </section>
  </div>  
</template>

<script>
  import StockCard from './components/StockCard.vue'

  export default {
    name: 'app',
    components: { StockCard },
    methods: {
      onCardClicked() {
        alert('Hello')
      },
      filterStocksByQuery() {
        let _query = this.query;
        return this.stocks.filter(stock => stock.symbol.toLowerCase().includes(_query.toLowerCase()));
      } 
    },
    mounted() {
      axios
        .get('/pricelist')
          .then(response => {
            console.log(response)
            this.stocks = response.data.stocks
          })
          .catch(error => console.log(error))
          .finally(() => console.log('finally...'))

    },
    data() {
      return {
        stocks : [{ close:"1.3", date:"2018-06-25T00:00:00Z", high:1.3, low:1.3, movement:0, open:1.3, percent_movement:0, prev_close:1.3, symbol:"ABBEYBDS", value:1860, volume:1500},
        { close:"1.3", date:"2018-06-25T00:00:00Z", high:1.3, low:1.3, movement:0, open:1.3, percent_movement:0, prev_close:1.3, symbol:"ABBEYBDS", value:1860, volume:1500},
        { close:"1.3", date:"2018-06-25T00:00:00Z", high:1.3, low:1.3, movement:0, open:1.3, percent_movement:0, prev_close:1.3, symbol:"ABBEYBDS", value:1860, volume:1500},
        { close:"1.3", date:"2018-06-25T00:00:00Z", high:1.3, low:1.3, movement:0, open:1.3, percent_movement:0, prev_close:1.3, symbol:"ABBEYBDS", value:1860, volume:1500},
        { close:"1.3", date:"2018-06-25T00:00:00Z", high:1.3, low:1.3, movement:0, open:1.3, percent_movement:0, prev_close:1.3, symbol:"ABBEYBDS", value:1860, volume:1500},
        { close:"1.3", date:"2018-06-25T00:00:00Z", high:1.3, low:1.3, movement:0, open:1.3, percent_movement:0, prev_close:1.3, symbol:"ABBEYBDS", value:1860, volume:1500},
        { close:"1.3", date:"2018-06-25T00:00:00Z", high:1.3, low:1.3, movement:0, open:1.3, percent_movement:0, prev_close:1.3, symbol:"ABBEYBDS", value:1860, volume:1500},
        { close:"1.3", date:"2018-06-25T00:00:00Z", high:1.3, low:1.3, movement:0, open:1.3, percent_movement:0, prev_close:1.3, symbol:"ABBEYBDS", value:1860, volume:1500}],
        query: ''
      }
    }
  }
</script>

<style lang="css">
  body {
    background-color: #303030;
  }

  .hero.is-info .tabs.is-boxed li.is-active a, .hero.is-info .tabs.is-boxed li.is-active a:hover {
    background-color: #303030;
    border-color: #303030;
  }
</style>
