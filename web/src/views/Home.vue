<i18n>
en-US:
  heading: 'Coin History'
  loading: 'Loading ...'
</i18n>

<template>
  <v-container fluid>
    <v-layout column align-center>
      <div v-if="loading" v-t="'loading'" />
      <canvas ref="chart" />
    </v-layout>
  </v-container>
</template>

<script>
import Backend from '@/api/backend/index.js'
import Chart from 'chart.js'

export default {
  metaInfo () {
    return {
      title: this.$t('heading')
    }
  },
  created () {
    return this.fetch()
  },
  data () {
    return {
      loading: false,
      chart: null,
      symbol: 'BTC',
      historyData: []
    }
  },
  computed: {
    chartData () {
      return this.historyData.map((data) => {
        return {
          x: new Date(data.timestamp),
          y: data.price
        }
      })
    }
  },
  methods: {
    fetch () {
      this.loading = true
      return Backend.GetCoinHistory({ symbol: this.symbol })
        .then((list) => {
          this.historyData = list
          this.draw()
          this.loading = false
        })
        .catch((err) => {
          this.loading = false
          throw err
        })
    },

    cleanChart () {
      if (this.chart) {
        this.chart.destroy()
        this.chart = null
      }
    },

    draw () {
      this.cleanChart()
      const ctx = this.$refs['chart'].getContext('2d')
      this.chart = new Chart(ctx, {
        type: 'line',
        data: {
          datasets: [
            {
              label: this.symbol,
              backgroundColor: 'rgba(124, 179, 66, .5)',
              borderColor: 'rgb(124, 179, 66)',
              data: this.chartData,
              fill: true
            }
          ]
        },
        options: {
          scales: {
            xAxes: [
              {
                type: 'time'
              }
            ]
          }
        }
      })
    }
  },
  beforeDestroy () {
    this.cleanChart()
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
