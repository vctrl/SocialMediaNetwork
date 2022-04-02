import { Axios } from 'axios';

import { User } from '@/types/entities';

const api = new Axios({
  baseURL: '',
});

export default {
  users: {
    register: (data: User & { password: string }) => api.post<User>('/users/register', data),
    login: (data: { login: User['login']; password: string }) => api.post<User>('/users/login', data),
    get: (ids: User['id'][]) => api.get<User[]>(`/users/${ids.join(',')}`),
    edit: (id: User['id'], data: Omit<Partial<User>, 'id'>) => api.put<User>(`/users/${id}`, data),
    remove: (id: User['id']) => api.delete<void>(`/users/${id}`),
  },
  friends: {
    get: () => api.get<User['id'][]>(`/friends`),
    remove: (id: User['id']) => api.delete<User['id'][]>(`/friends/requests/${id}`),
    getIncomingRequests: () => api.get<User['id'][]>('/friends/requests'),
    getSentRequests: () => api.get<User['id'][]>('/friends/sent_requests'),
    sendRequest: (id: User['id']) => api.post(`/friends/${id}`),
    acceptRequest: (id: User['id']) => api.post(`/friends/${id}/accept`),
  },
};
