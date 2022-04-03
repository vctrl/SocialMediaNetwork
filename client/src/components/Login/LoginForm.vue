<script setup lang="ts">
import { AxiosError } from 'axios';
import { NButton, NForm, NFormItemRow, NIcon, NInput, useNotification } from 'naive-ui';
import { ref } from 'vue';
import { useRouter } from 'vue-router';

import Icon from '@/components/Icon';
import { users } from '@/providers/api';

const notify = useNotification();
const router = useRouter();

const payload = ref({
  login: '',
  password: '',
});

async function submit() {
  try {
    await users.login(payload.value);
    router.push('/');
  } catch (e) {
    const err = e as AxiosError;
    notify.error({ title: 'API error', description: err.response?.data?.message ?? err.message, duration: 8000 });
  }
}
</script>

<template>
  <NForm label-placement="left" label-width="auto">
    <NFormItemRow label="Username">
      <NInput v-model:value="payload.login" placeholder="Username" />
    </NFormItemRow>
    <NFormItemRow label="Password">
      <NInput v-model:value="payload.password" placeholder="Password" type="password" />
    </NFormItemRow>
    <div class="row jc-center">
      <NButton type="primary" icon-placement="right" @click="submit">
        Submit
        <template #icon>
          <NIcon><Icon>mdi mdi-login</Icon></NIcon>
        </template>
      </NButton>
    </div>
  </NForm>
</template>
