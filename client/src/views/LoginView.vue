<script setup lang="ts">
import { NCard, NTabPane, NTabs } from 'naive-ui';
import { ref, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';

import LoginForm from '@/components/Login/LoginForm.vue';
import RegisterForm from '@/components/Login/RegisterForm.vue';

const route = useRoute();
const router = useRouter();

const tab = ref(route.params.tab as string);

watch(route, () => (tab.value = route.params.tab as string));
</script>

<template>
  <NCard style="max-width: 500px" content-style="padding-top: 0" class="login-card">
    <NTabs
      v-model:value="tab"
      size="large"
      animated
      type="line"
      :tabs-padding="24"
      pane-style="padding-top: 24px"
      @update:value="router.push('/' + tab)"
    >
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
