<template>
  <div>
    <v-app-bar app>
      <v-app-bar-nav-icon @click="enableDrawer" />
      <v-toolbar-title>ぶいしゅき</v-toolbar-title>
    </v-app-bar>

    <v-navigation-drawer app v-model="isDrawerEnabled">
    </v-navigation-drawer>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  data() {
    return {
      isDrawerEnabled: false
    }
  },

  methods: {
    enableDrawer() {
      this.isDrawerEnabled = !this.isDrawerEnabled
    }
  }
})
</script>

