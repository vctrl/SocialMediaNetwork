import axios from 'axios';

import { User } from '@/types/entities';

const api = axios.create({
  baseURL: '/api',
});

export const users = {
  register: async (data: Omit<User, 'id'> & { password: string }) =>
    (await api.post<User>('users/register', data)).data,
  login: async (data: { login: User['login']; password: string }) => (await api.post<User>('users/login', data)).data,
  logout: async () => (await api.post('users/logout')) as void,
  me: async () => (await api.get<User>('users/me')).data,
  get: async (ids: User['id'][]) => (await api.get<User[]>(`users/${ids.join(',')}`)).data,
  edit: async (id: User['id'], data: Omit<Partial<User>, 'id'>) => (await api.put<User>(`users/${id}`, data)).data,
  remove: async (id: User['id']) => (await api.delete<void>(`users/${id}`)).data,
};

export const friends = {
  get: async () => (await api.get<User['id'][]>(`friends`)).data,
  remove: async (id: User['id']) => (await api.delete<User['id'][]>(`friends/requests/${id}`)).data,
  getIncomingRequests: async () => (await api.get<User['id'][]>('friends/requests')).data,
  getSentRequests: async () => (await api.get<User['id'][]>('friends/sent_requests')).data,
  sendRequest: async (id: User['id']) => (await api.post(`friends/${id}`)).data,
  acceptRequest: async (id: User['id']) => (await api.post(`friends/${id}/accept`)).data,
};
