<script setup lang="ts">
import { darkTheme, NConfigProvider, NNotificationProvider } from 'naive-ui';
import { RouterView } from 'vue-router';

import AppBar from './components/AppBar.vue';
import { useState } from './store/state';
import LoginView from './views/LoginView.vue';

const state = useState();
state.fetchLocalUser();
</script>

<template>
  <NConfigProvider :theme="darkTheme">
    <NNotificationProvider>
      <template v-if="state.localUser">
        <AppBar />
        <RouterView />
      </template>
      <div v-else class="fill-parent column ai-center jc-center">
        <LoginView />
      </div>
    </NNotificationProvider>
  </NConfigProvider>
</template>

<style lang="scss">
@import './styles/util.scss';
@import './styles/animations.scss';
body {
  background: url(@/assets/bg.png);
}
</style>
