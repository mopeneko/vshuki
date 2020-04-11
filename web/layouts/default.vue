<template>
  <v-app>
    <top-menu />

    <v-content>
      <v-container fluid>
        <nuxt />
      </v-container>
    </v-content>
  </v-app>
</template>

<script lang="ts">
import Vue from 'vue'
import TopMenu from '~/components/TopMenu.vue'

export default Vue.extend({
  components: {
    TopMenu
  }
})
</script>

