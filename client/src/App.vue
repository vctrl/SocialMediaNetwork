<script setup lang="ts">
import { darkTheme, NConfigProvider, NNotificationProvider } from 'naive-ui';
import { RouterView } from 'vue-router';
</script>

<template>
  <NConfigProvider :theme="darkTheme">
    <NNotificationProvider>
      <div class="fill-parent column ai-center jc-center">
        <RouterView />
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
