<script setup lang="ts">
import { NButton, NDivider, NForm, NFormItemRow, NIcon, NInput, NInputNumber } from 'naive-ui';
import { computed } from 'vue';

import Icon from '@/components/Icon';
import { Gender, User } from '@/types/entities';
import { allValid, useFieldProps, useValidated, ValidationResult } from '@/use/validate';

const age = useValidated(18, (age): ValidationResult => {
  if (age < 13) return { isValid: false, message: 'Users below 13 are not allowed to register' };
  if (age > 100) return { isValid: false, message: 'You are joking, right?' };
  return { isValid: true };
});
const city = useValidated(
  '',
  (city): ValidationResult => (city.length < 3 ? { isValid: false, message: 'Too short' } : { isValid: true }),
);
const interests = useValidated('', (interests): ValidationResult => ({ isValid: true }));
const password = useValidated('', (password): ValidationResult => ({ isValid: true }));
const login = useValidated('', (login): ValidationResult => ({ isValid: true }));
const name = useValidated('', (name): ValidationResult => ({ isValid: true }));
const sex = useValidated('MALE' as Gender, (sex): ValidationResult => ({ isValid: true }));
const surname = useValidated('', (surname): ValidationResult => ({ isValid: true }));

const fieldProps = useFieldProps({ age, city, interests, login, name, password, sex, surname });
const canSave = allValid([age, city, interests, login, name, password, sex, surname]);

const user = computed<Omit<User, 'id'> & { password: string }>(() => ({
  age: age.value,
  city: city.value,
  interests: interests.value,
  login: login.value,
  name: name.value,
  password: password.value,
  sex: sex.value,
  surname: surname.value,
}));
</script>

<template>
  <NForm label-placement="left" :label-width="80">
    <NFormItemRow
      label="First Name"
      :feedback="fieldProps.name.feedback.value"
      :validation-status="fieldProps.name.status.value"
    >
      <NInput v-model:value="name" placeholder="First Name" />
    </NFormItemRow>
    <NFormItemRow
      label="Last Name"
      :feedback="fieldProps.surname.feedback.value"
      :validation-status="fieldProps.surname.status.value"
    >
      <NInput v-model:value="surname" placeholder="Last Name" />
    </NFormItemRow>
    <NFormItemRow
      label="Location"
      :feedback="fieldProps.city.feedback.value"
      :validation-status="fieldProps.city.status.value"
    >
      <NInput v-model:value="city" placeholder="Location" />
    </NFormItemRow>
    <NFormItemRow
      label="Age"
      :feedback="fieldProps.age.feedback.value"
      :validation-status="fieldProps.age.status.value"
    >
      <NInputNumber v-model:value="age" class="full-width" />
    </NFormItemRow>
    <NFormItemRow
      label="Interests"
      :feedback="fieldProps.interests.feedback.value"
      :validation-status="fieldProps.interests.status.value"
    >
      <NInput v-model:value="interests" placeholder="Interests" type="textarea" />
    </NFormItemRow>

    <NDivider dashed />

    <NFormItemRow
      label="Username"
      :feedback="fieldProps.login.feedback.value"
      :validation-status="fieldProps.login.status.value"
    >
      <NInput v-model:value="login" placeholder="Username" />
    </NFormItemRow>
    <NFormItemRow
      label="Password"
      :feedback="fieldProps.password.feedback.value"
      :validation-status="fieldProps.password.status.value"
    >
      <NInput v-model:value="password" placeholder="Password" type="password" />
    </NFormItemRow>
    <div class="row jc-center">
      <NButton type="primary" icon-placement="right" :disabled="!canSave">
        Submit
        <template #icon>
          <NIcon><Icon>mdi mdi-account-plus</Icon></NIcon>
        </template>
      </NButton>
    </div>
  </NForm>
</template>
