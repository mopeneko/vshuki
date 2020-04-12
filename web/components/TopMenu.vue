<template>
  <div>
    <v-app-bar app>
      <v-app-bar-nav-icon @click="enableDrawer" />
      <v-toolbar-title>ぶいしゅき</v-toolbar-title>
    </v-app-bar>

    <v-navigation-drawer app v-model="isDrawerEnabled">
      <v-list nav>
        <v-list-item link v-for="item in drawerItems" :key="item.title" :to="item.to">
          <v-list-item-icon>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-item-icon>

          <v-list-item-content>
            <v-list-item-title>{{ item.title }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  data() {
    return {
      isDrawerEnabled: false,
      drawerItems: [
        { icon: 'mdi-post', title: '投稿一覧', to: "/posts" }
      ]
    }
  },

  methods: {
    enableDrawer() {
      this.isDrawerEnabled = !this.isDrawerEnabled
    }
  }
})
</script>

