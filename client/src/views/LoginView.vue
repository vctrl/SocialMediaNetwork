<script setup lang="ts">
import { NCard, NTabPane, NTabs } from 'naive-ui';

import LoginForm from '@/components/Login/LoginForm.vue';
import RegisterForm from '@/components/Login/RegisterForm.vue';
</script>

<template>
  <NCard style="max-width: 500px" content-style="padding-top: 0" class="login-card">
    <NTabs size="large" animated type="line" :tabs-padding="24" pane-style="padding-top: 24px">
      <NTabPane name="login" tab="Log in">
        <LoginForm />
      </NTabPane>
      <NTabPane name="register" tab="Register">
        <RegisterForm />
      </NTabPane>
    </NTabs>
  </NCard>
</template>

<style lang="scss">
.login-card .n-tabs-nav--line-type.n-tabs-nav {
  margin: 0 -24px;
}
</style>
