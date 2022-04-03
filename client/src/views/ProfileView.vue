<script setup lang="ts">
import { capitalCase } from 'change-case';
import { NAvatar, NButton, NCard, NH2, NResult, NText } from 'naive-ui';
import { computed, ref, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';

import { users } from '@/providers/api';
import { User } from '@/types/entities';

const route = useRoute();
const router = useRouter();
const userId = computed(() => Number(route.params.userId));
const user = ref<User | undefined>();
const error = ref<{ status: import('naive-ui').ResultProps['status']; message: string } | undefined>();

watch(
  userId,
  async (userId) => {
    if (!userId || isNaN(userId)) {
      error.value = { status: '500', message: 'User ID is incorrect' };
      return;
    }

    user.value = undefined;
    error.value = undefined;

    try {
      const [data] = await users.get([userId]);
      if (data) user.value = data;
      else error.value = { status: '404', message: 'User not found' };
    } catch {
      error.value = { status: '500', message: 'Could not get data from API' };
    }
  },
  { immediate: true },
);
</script>

<template>
  <NCard style="max-width: 600px">
    <NResult v-if="error" :status="error.status" title="API error" :description="error.message">
      <template #footer>
        <NButton @click="router.push('/')">Back to home</NButton>
      </template>
    </NResult>
    <div v-else-if="user" class="row gap-20">
      <div class="column">
        <NAvatar :size="128" :src="'https://joeschmoe.io/api/v1/male/' + user.name" round color="transparent" />
      </div>
      <div class="column">
        <NH2 class="mb-5">{{ user.name }} {{ user.surname }}</NH2>
        <NText>{{ user.city }}</NText>
        <NText depth="3">{{ capitalCase(user.sex) }}, 24 years</NText>
      </div>
    </div>
  </NCard>
</template>
