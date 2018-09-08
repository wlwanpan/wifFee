<template lang="pug">
  md-list
    md-list-item(
      @click="navigateTo('/')")
      md-icon terrain
      span.md-list-item-text Map
    md-list-item(
      @click="navigateTo('/account')")
      md-icon account_circle
      span.md-list-item-text Account
    md-list-item(
      @click="navigateTo('/api')")
      md-icon code
      span.md-list-item-text Developer Api
    md-list-item(
      @click="navigateTo('/about')")
      md-icon contact_support
      span.md-list-item-text About
</template>

<script>

export default {
  name: 'drawer-nav-list',
  methods: {
    navigateTo(route) {
      this.$emit('close-drawer')
      this.$router.push(route)
    }
  }
}
</script>

<style lang='scss' scoped>
.md-list-item {
  &:hover {
    background: #ddd;
  }
}
</style>
