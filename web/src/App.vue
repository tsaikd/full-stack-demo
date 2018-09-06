<template>
  <v-app>
    <v-content class="grey lighten-4">
      <router-view />
    </v-content>
  </v-app>
</template>

<script>

export default {
  name: 'App'
}
</script>

<style lang="stylus">
body
  margin 0
#app
  -webkit-font-smoothing antialiased
  -moz-osx-font-smoothing grayscale
</style>

<style lang="stylus">
.v-list.left-bar a[href]
  border-right 5px solid transparent
  &.fn-active-border-right
    border-right 5px solid #3498db !important
</style>
