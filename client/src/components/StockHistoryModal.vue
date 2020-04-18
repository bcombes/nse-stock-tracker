<template>
	<div v-show="show" class="modal is-active">
	  <div class="modal-background"></div>
	  <div class="modal-card">
	    <header class="modal-card-head">
	      <p class="modal-card-title">{{ symbol }}</p>
	      <button class="delete" aria-label="close" @click="closeModal"></button>
	    </header>
	    <section class="modal-card-body">
	    	<div class="control">
			  <div class="select" >
			    <select v-model="selected_chart" >
			      <option disabled value="">Please select one</option>
			      <option value='time-series'>Time-series</option>
			      <option value='candle-stick'>Candle-stick</option>
			    </select>
			  </div>
			</div>
	    	<div id="plot"></div>
	    </section>
	    <footer class="modal-card-foot">
	      <button class="button" @click="closeModal">Close</button>
	    </footer>
	  </div>
	</div>
</template>

<script>
	export default {
		props: ['symbol'], 
		data() {
			return {
				selected_chart: 'time-series',
				show: false,
				data: {},
				traces: [],
				layout: {},
				axes: { xmin: 0, xmax: 0, ymin: 0,ymax: 0 }
			}	
		},
		watch: {
			selected_chart: function(value) {
				console.log('selected chart updated')
				this.updatePlot();
			},
			symbol: function(symbol) {
				//console.log("symbol changed..." + symbol)
				axios
			        .get('/stockhistory/' + symbol)
			          .then(response => {
			            console.log(response)
			            this.data = response.data;
			            this.updatePlot();
			            //this.traces = response.data.stocks

			          })
			          .catch(error => console.log(error))
			          .finally(() => console.log('finally... 2'))	

			}
		}, 
		computed: {
			time_series_layout: function() { 
				return {
				  title: this.symbol.charAt(0) + this.symbol.slice(1).toLowerCase() + ' Stock History',
				  xaxis: {
				    autorange: true,
				    range: [ this.xmin, this.xmax],
				    rangeselector: {buttons: [
				        {
				          count: 1,
				          label: '1m',
				          step: 'month',
				          stepmode: 'backward'
				        },
				      	{
				          count: 3,
				          label: '3m',
				          step: 'month',
				          stepmode: 'backward'
				        },
				        {
				          count: 6,
				          label: '6m',
				          step: 'month',
				          stepmode: 'backward'
				        },
				    	{
				          count: 12,
				          label: '1yr',
				          step: 'month',
				          stepmode: 'backward'
				        },
				        {step: 'all'}
				      ]},
				    rangeslider: {range: [this.xmin, this.xmax]},
				    type: 'date'
				  },
				  yaxis: {
				    autorange: true,
				    range: [this.ymin, this.ymax],
				    type: 'linear'
				  }
				}
			},
			candle_stick_layout: function() { 
				return {
				  dragmode: 'zoom', 
				  margin: {
				    r: 10, 
				    t: 25, 
				    b: 40, 
				    l: 60
				  }, 
				  showlegend: false, 
				  xaxis: {
				    autorange: true, 
				    domain: [0, 1], 
				    range: [this.xmin, this.xmax], 
				    rangeselector: {buttons: [
				        {
				          count: 1,
				          label: '1m',
				          step: 'month',
				          stepmode: 'backward'
				        },
				      	{
				          count: 3,
				          label: '3m',
				          step: 'month',
				          stepmode: 'backward'
				        },
				        {
				          count: 6,
				          label: '6m',
				          step: 'month',
				          stepmode: 'backward'
				        },
				    	{
				          count: 12,
				          label: '1yr',
				          step: 'month',
				          stepmode: 'backward'
				        },
				        {step: 'all'}
				      ]},
				    rangeslider: {range: [this.xmin, this.xmax]}, 
				    title: 'Date', 
				    type: 'date'
				  }, 
				  yaxis: {
				    autorange: true, 
				    domain: [0, 1], 
				    range: [this.ymin, this.ymax], 
				    type: 'linear'
				  }
				}
			}		
		},
		methods: {
			closeModal() {
				this.show = false;
				this.traces = [];
			},
			updatePlot() {
				this.traces = [];
				let data = this.data;
				this.axes.xmin = data.dates[0];
	            this.axes.xmax = data.dates[data.dates.length - 1];
	          	this.axes.ymin = Math.min(...data.closes);
	            this.axes.ymax = Math.max(...data.closes)

	            if(this.selected_chart === 'time-series') {
		           	let time_series_trace = {
					  type: "scatter",
					  mode: "lines",
					  name: this.symbol,
					  x: data.dates,
					  y: data.closes,
					  line: {color: '#17BECF'}
					}
					this.traces.push(time_series_trace)
					this.layout = this.time_series_layout;
	            } else {
	            	let candle_stick_trace = {
					  x: data.dates, 
					  close: data.closes, 
					  decreasing: {line: {color: '#7F7F7F'}}, 
					  high: data.highs, 
					  increasing: {line: {color: '#17BECF'}}, 
					  line: {color: 'rgba(31,119,180,1)'}, 
					  low: data.lows,
					  open: data.opens, 
					  type: 'candlestick', 
					  xaxis: 'x', 
					  yaxis: 'y'
					}	

	            	this.traces.push(candle_stick_trace)
					this.layout = this.candle_stick_layout;
	            }

	            this.show = true;
	            Plotly.newPlot('plot', this.traces, this.layout);
			}
		}
	}
</script>

<style>
</style>