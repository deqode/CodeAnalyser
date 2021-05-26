// @flow
import type { GraphQLContext } from '../../';
import { isAdmin } from '../../utils/permissions';
import { getCount, getGrowth } from '../../models/utils';

export default async (_: any, __: any, { user }: GraphQLContext) => {
  if (!isAdmin(user.id)) return null;

  return {
    count: await getCount('directMessageThreads'),
    weeklyGrowth: await getGrowth(
      'directMessageThreads',
      'weekly',
      'createdAt'
    ),
    monthlyGrowth: await getGrowth(
      'directMessageThreads',
      'monthly',
      'createdAt'
    ),
    quarterlyGrowth: await getGrowth(
      'directMessageThreads',
      'quarterly',
      'createdAt'
    ),
  };
};
