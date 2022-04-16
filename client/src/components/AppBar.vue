<script setup lang="ts">
import { NButton, NCard, NDropdown } from 'naive-ui';
import { DropdownMixedOption } from 'naive-ui/lib/dropdown/src/interface';
import { ref } from 'vue';
import { useRouter } from 'vue-router';

import { useState } from '@/store/state';

const router = useRouter();
const state = useState();

const options = ref<DropdownMixedOption[]>([
  { label: 'My profile', key: 'profile' },
  { label: 'Log out', key: 'logout' },
]);

const handleClick = async (key: string) => {
  switch (key) {
    case 'profile':
      router.push('users/' + state.localUser?.id);
      break;
    case 'logout':
      await state.logout();
      break;
    default:
      console.error('Unknown key in click handler');
  }
};
</script>

<template>
  <NCard id="app-bar">
    <div class="row">
      <div class="spacer"></div>
      <NDropdown :options="options" @select="handleClick">
        <NButton quaternary size="large">{{ state.localUser?.name }} {{ state.localUser?.surname }}</NButton>
      </NDropdown>
    </div>
  </NCard>
</template>

<style lang="scss">
#app-bar {
  max-width: 1280px;
  margin: 1em auto;

  .n-card__content {
    padding: 0;
  }
}

@media screen and (max-width: 1300px) {
  #app-bar {
    margin: 1em;
    width: auto;
  }
}

@media screen and (max-width: 768px) {
  #app-bar {
    margin: -1px;
    margin-bottom: 1em;
    width: auto;
    border-radius: 0;
  }
}
</style>
