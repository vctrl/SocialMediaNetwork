import { defineStore } from 'pinia';

import { users } from '@/providers/api';
import { User } from '@/types/entities';

export const useState = defineStore('state', {
  state: () => ({
    localUser: undefined as User | undefined,
  }),

  actions: {
    async fetchLocalUser() {
      try {
        this.localUser = await users.me();
      } catch {
        this.localUser = undefined;
      }

      return this.localUser;
    },
    async logout() {
      await users.logout();
      this.localUser = undefined;
    },
  },
});
