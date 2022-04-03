<script setup lang="ts">
import { NButton, NForm, NFormItemRow, NIcon, NInput } from 'naive-ui';

import Icon from '@/components/Icon';
</script>

<template>
  <NForm label-placement="left" label-width="auto">
    <NFormItemRow label="Username">
      <NInput placeholder="Username" />
    </NFormItemRow>
    <NFormItemRow label="Password">
      <NInput placeholder="Password" type="password" />
    </NFormItemRow>
    <div class="row jc-center">
      <NButton type="primary" icon-placement="right">
        Submit
        <template #icon>
          <NIcon><Icon>mdi mdi-login</Icon></NIcon>
        </template>
      </NButton>
    </div>
  </NForm>
</template>
