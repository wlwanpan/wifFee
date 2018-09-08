<template lang="pug">
  div#app
    md-app.full-page
      md-app-toolbar.md-primary
        md-button.md-icon-button(
          @click="openDrawer = !openDrawer",
          v-if="!openDrawer")
          md-icon menu
        span.md-title WifFee
        div.md-toolbar-section-end
          drawer-right-actions
      md-app-drawer(
        :md-active.sync="openDrawer",
        md-persistent="full")
        md-toolbar.md-transparent(md-elevation="0")
          span Navigation
          div.md-toolbar-section-end
            md-button.md-icon-button.md-dense(
              @click="openDrawer = !openDrawer")
              md-icon keyboard_arrow_left
        drawer-nav-list(
          @close-drawer="openDrawer = false")
      md-app-content
        router-view
</template>

<script>
import DrawerNavList from './components/DrawerNavList'
import DrawerRightActions from './components/DrawerRightActions'
import { mapGetters } from 'vuex'

export default {
  name: 'app',
  data() {
    return {
      openDrawer: false
    }
  },
  components: {
    DrawerNavList,
    DrawerRightActions
  }
}
</script>

<style lang="scss">
@import '~normalize.css';

#app {
  position: absolute;
  top: 0; bottom: 0; left: 0; right: 0;
  .md-app-content {
    padding: 0;
  }
}

.full-page {
  width: 100%; height: 100%;
}
</style>
