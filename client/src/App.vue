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
                <li v-for="(filter, index) in stockFilters" :class="{ 'is-active': index === activeFilter, 'has-text-weight-semibold': index === activeFilter  }"><a @click="activeFilter = index"> {{ filter }}</a></li>
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
      <!--<div class="columns is-multiline is-mobile"> -->
        <paginate name="filteredStocks" :list="filteredStocks" tag="div" class="columns is-multiline is-mobile" :per="20">
          <stock-card v-for="stock in paginated('filteredStocks')" :key="stock.symbol" :stock="stock" @stock-card-clicked="onCardClicked"></stock-card> 
        </paginate>
      <!--</div>-->
      <div class="container has-text-centered has-text-weight-bold">
        <paginate-links for="filteredStocks" :show-step-links="true" :step-links="{ next: '>>', prev: '<<' }" ></paginate-links>
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
    },
    computed: {
      filteredStocks: function() {
              //filterStocks() {
        let _query = this.query;
        let queriedStocks = [];
        if(_query != null && _query  != '') {
          queriedStocks =  this.stocks.filter(stock => stock.symbol.toLowerCase().includes(_query.toLowerCase()));  
        } else {
          queriedStocks = this.stocks;
        }
        let filteredStocks = [];
        switch(this.activeFilter) {
          case 1: //Most traded sort by volume
            filteredStocks = queriedStocks.sort((a,b) => b.volume - a.volume)
            break;
          case 2: //Most active sort by percentage movement
            filteredStocks = queriedStocks
              .filter(stock => stock.movement > 0)
              .sort((a,b) => b.movement - a.movement)
            break;
          case 3: //Most active sort by percentage movement
            filteredStocks = queriedStocks
              .filter(stock => stock.movement < 0)
              .sort((a,b) => a.movement - b.movement)
            break;  
          default:
            filteredStocks = queriedStocks;
            break;  
        }
        return filteredStocks;
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
        query: '',
        stockFilters: ['All', 'Most Traded', 'Biggest Gainers', 'Biggest Losers'],
        activeFilter: 0,
        paginate: ['filteredStocks']
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
  ul.paginate-links {
    list-style-type: none;
    padding: 0;
  }

  ul.paginate-links li {
    display: inline-block;
    margin: 0 10px;
  }

  ul.paginate-links > li.active {
    background-color: #505050;
  }

  ul.paginate-links > li.active > a {
    color: #f0f0f0;
  }
</style>
