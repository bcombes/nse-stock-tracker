<template>
     <div :class="trend_indicator" class="has-text-light column is-2 is-inline-block-mobile" @click="onCardClicked">
      <p class="title is-6 has-text-light has-text-weight-bold"> {{ stock.symbol }} </p>
      <p class="subtitle is-7 has-text-light has-text-weight-bold"> &#x20a6;{{ stock.close | toTwoDecimalPlaces }}</p>  <!-- price --> 
      <span class="icon">
        <i class="fas fa-angle-double-up"></i>
      </span>
      <p class="change is-size-7" > {{ stock.movement }}</p>
      <p class="percent is-size-7" >{{ stock.percent_movement }}%</p>  <!-- percentage change -->  
      <p class="volume is-size-7">{{ stock.volume | addCommaSeparator }}</p> <!-- volume --> 
    </div>
</template>

<script>
	export default {
		props: ['stock'],
		methods: {
      		onCardClicked() {
      			this.$emit('stock-card-clicked', this.stock.symbol)
      		}
    	},
    	filters: {
    		toTwoDecimalPlaces(value) {
    			return Number(value).toFixed(2)
    		},
    		addCommaSeparator(value) {
    			return Number(value).toLocaleString()
    		}
    	},
    	data() {
    		return {
    			trend_indicator:''
    		}
    	},
    	mounted() {
    		if(this.stock.movement > 0) {
    			this.trend_indicator = "has-background-primary";
    		} else if(this.stock.movement < 0) {
    			this.trend_indicator = "has-background-danger";
    		} else {
    			this.trend_indicator = "has-background-info";
    		}
    	}
	}
</script>

<style lang="css">
	.volume {
	  display: inline;
	  float: left;
	  clear: left;
	}
	.percent, .change {
	  float:right;
	}
	.percent {
	  clear: both;
	}

	.column {
	  font-weight: bold;
	  margin: 5px 5px;
	}
	.icon {
	  float: left;
	}
	/*
	.box:not(:last-child) {
  		margin-bottom: 0;
  	}*/
</style>