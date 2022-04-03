<script setup lang="ts">
import { NCard, NTabPane, NTabs } from 'naive-ui';

import LoginForm from '@/components/Login/LoginForm.vue';
import RegisterForm from '@/components/Login/RegisterForm.vue';
</script>

<template>
  <NCard style="max-width: 500px">
    <NTabs size="large" animated>
      <NTabPane name="login" tab="Log in">
        <LoginForm />
      </NTabPane>
      <NTabPane name="register" tab="Register">
        <RegisterForm />
      </NTabPane>
    </NTabs>
  </NCard>
</template>
