<script setup lang="ts">
import { AxiosError } from 'axios';
import {
  FormInst,
  NButton,
  NDivider,
  NForm,
  NFormItemRow,
  NIcon,
  NInput,
  NInputNumber,
  NSelect,
  useNotification,
} from 'naive-ui';
import { SelectBaseOption } from 'naive-ui/lib/select/src/interface';
import { ref } from 'vue';
import { useRouter } from 'vue-router';

import Icon from '@/components/Icon';
import { users } from '@/providers/api';
import { User } from '@/types/entities';
import { RulesFor } from '@/types/util';
import { validateAsync } from '@/use/validate';

type RegisterModel = Omit<User, 'id'> & { password: string };

const notify = useNotification();
const router = useRouter();

const trigger = ['input', 'blur', 'change'];
const sexOptions: SelectBaseOption<User['sex']>[] = [
  { label: 'Male', value: 'MALE' },
  { label: 'Female', value: 'FEMALE' },
];

const form = ref<FormInst | null>(null);
const user = ref<Partial<RegisterModel>>({});
const rules: RulesFor<RegisterModel> = {
  age: {
    validator(_, age: number) {
      if (typeof age !== 'number') return new Error('Age is required');
      if (age < 13) return new Error('Minimum age to register is 13 years');
      return true;
    },
    trigger,
  },
  city: { message: 'Must be from 5 to 50 latin characters', required: true, pattern: /^[\w\d\s.,-]{5,50}$/i, trigger },
  interests: {
    message: 'Must be from 5 to 250 latin characters',
    required: true,
    pattern: /^[\w\d\s.,-]{5,250}$/i,
    trigger,
  },
  login: { message: 'Must be from 3 to 20 latin characters', required: true, pattern: /^\w{3,20}$/i, trigger },
  name: { message: 'Must be from 2 to 20 latin characters', required: true, pattern: /^[\w\d\s.,-]{2,20}$/i, trigger },
  password: {
    validator(_, password: string) {
      if (!password) return new Error('Password is required');
      if (!password.match(/[a-z]/)) return new Error('At least 1 lowercase letter needed');
      if (!password.match(/[A-Z]/)) return new Error('At least 1 uppercase letter needed');
      if (!password.match(/\d/)) return new Error('At least 1 digit needed');
      if (password.length < 8) return new Error('Minimum length is 8 symbols');
      return true;
    },
    trigger,
  },
  sex: { message: 'Gender is required', required: true, trigger },
  surname: {
    message: 'Must be from 2 to 20 latin characters',
    required: true,
    pattern: /^[\w\d\s.,-]{2,20}$/i,
    trigger,
  },
};

async function submit() {
  if (!form.value) return;

  const errors = await validateAsync(form.value);
  if (errors?.length ?? 0 > 0)
    return notify.error({
      title: 'Error',
      description: 'The form has problems, resolve them before submitting it',
      duration: 8000,
    });

  try {
    await users.register(user.value as RegisterModel);
    router.push('/login');
  } catch (e) {
    const err = e as AxiosError;
    notify.error({ title: 'API error', description: err.response?.data?.message ?? err.message, duration: 8000 });
  }
}
</script>

<template>
  <NForm ref="form" label-placement="left" :label-width="80" :rules="rules" :model="user" :show-require-mark="false">
    <NFormItemRow label="Gender" path="sex">
      <NSelect v-model:value="user.sex" :options="sexOptions as any" />
    </NFormItemRow>
    <NFormItemRow label="First Name" path="name">
      <NInput v-model:value="user.name" placeholder="John" />
    </NFormItemRow>
    <NFormItemRow label="Last Name" path="surname">
      <NInput v-model:value="user.surname" placeholder="Doe" />
    </NFormItemRow>
    <NFormItemRow label="Location" path="city">
      <NInput v-model:value="user.city" placeholder="New York, United States" />
    </NFormItemRow>
    <NFormItemRow label="Age" path="age">
      <NInputNumber v-model:value="user.age" class="full-width" placeholder="Full years" />
    </NFormItemRow>
    <NFormItemRow label="Interests" path="interests">
      <NInput v-model:value="user.interests" placeholder="Web development, rocket-powered car soccer" type="textarea" />
    </NFormItemRow>

    <NDivider dashed />

    <NFormItemRow label="Username" path="login">
      <NInput v-model:value="user.login" placeholder="johndoe94" />
    </NFormItemRow>
    <NFormItemRow label="Password" path="password">
      <NInput v-model:value="user.password" placeholder="sup3rS3cr3t" type="password" />
    </NFormItemRow>
    <div class="row jc-center">
      <NButton type="primary" icon-placement="right" @click="submit">
        Submit
        <template #icon>
          <NIcon><Icon>mdi mdi-account-plus</Icon></NIcon>
        </template>
      </NButton>
    </div>
  </NForm>
</template>
