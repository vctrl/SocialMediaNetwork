import { Injectable } from '@nestjs/common';
import { User } from '../../../src/types/entities';

const requests = new Map<User['id'], User['id'][]>();
const friends = new Map<User['id'], User['id'][]>();

function getArrayOrCreate(map: Map<User['id'], User['id'][]>, id: User['id']) {
  if (map.has(id)) return map.get(id);

  const newEntry: User['id'][] = [];
  map.set(id, newEntry);
  return newEntry;
}

@Injectable()
export class FriendsService {
  async get(userId: User['id']) {
    return getArrayOrCreate(friends, userId);
  }

  async getRequestsSent(userId: User['id']) {
    return getArrayOrCreate(requests, userId);
  }

  async getRequestsIncoming(userId: User['id']) {
    return Array.from(requests.entries())
      .filter(([_, ids]) => ids.indexOf(userId) !== -1)
      .map(([sourceUserId]) => sourceUserId);
  }

  async createRequest(sourceUserId: User['id'], targetUserId: User['id']) {
    const userRequests = getArrayOrCreate(requests, sourceUserId);
    if (userRequests.find((id) => id === targetUserId)) return;

    userRequests.push(targetUserId);
  }

  async acceptRequest(sourceUserId: User['id'], targetUserId: User['id']) {
    const userRequests = getArrayOrCreate(requests, sourceUserId);
    const index = userRequests.findIndex((id) => id === targetUserId);
    if (index === -1) return;

    await Promise.all([this.removeRequest(sourceUserId, targetUserId), this.create(sourceUserId, targetUserId)]);
  }

  async removeRequest(sourceUserId: User['id'], targetUserId: User['id']) {
    const userRequests = getArrayOrCreate(requests, sourceUserId);
    const index = userRequests.findIndex((id) => id === targetUserId);
    if (index === -1) return;

    userRequests.splice(index, 1);
  }

  async create(userId1: User['id'], userId2: User['id']) {
    const userFriends1 = getArrayOrCreate(friends, userId1);
    if (!userFriends1.find((id) => id === userId2)) userFriends1.push(userId2);

    const userFriends2 = getArrayOrCreate(friends, userId2);
    if (!userFriends2.find((id) => id === userId1)) userFriends2.push(userId1);
  }

  async remove(userId1: User['id'], userId2: User['id']) {
    const userFriends1 = getArrayOrCreate(friends, userId1);
    const index1 = userFriends1.findIndex((id) => id === userId2);
    if (index1 !== -1) userFriends1.splice(index1, 1);

    const userFriends2 = getArrayOrCreate(friends, userId2);
    const index2 = userFriends2.findIndex((id) => id === userId1);
    if (index2 !== -1) userFriends2.splice(index2, 1);
  }
}
