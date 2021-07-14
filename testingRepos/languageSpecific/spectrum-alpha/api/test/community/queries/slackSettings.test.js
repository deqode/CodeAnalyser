//@flow
import { request } from '../../utils';
import data from 'shared/testing/data';
import {
  SPECTRUM_COMMUNITY_ID,
  CHANNEL_MODERATOR_USER_ID,
  COMMUNITY_MODERATOR_USER_ID,
  MAX_ID,
  PREVIOUS_MEMBER_USER_ID,
} from '../../../migrations/seed/default/constants';
const channelModerator = data.users.find(
  ({ id }) => id === CHANNEL_MODERATOR_USER_ID
);
const communityModerator = data.users.find(
  ({ id }) => id === COMMUNITY_MODERATOR_USER_ID
);
const communityOwner = data.users.find(({ id }) => id === MAX_ID);
const noPermissionUser = data.users.find(
  ({ id }) => id === PREVIOUS_MEMBER_USER_ID
);

it('should not fetch slack settings if not authed', async () => {
  const query = /* GraphQL */ `
    {
      community(id: "${SPECTRUM_COMMUNITY_ID}") {
        id
        slackSettings {
          teamName
          isConnected
          hasSentInvites
          memberCount
          invitesSentAt
        }
      }
    }
  `;

  expect.assertions(1);
  const result = await request(query);

  expect(result).toMatchSnapshot();
});

it('should not fetch slack settings if no permissions', async () => {
  const query = /* GraphQL */ `
    {
      community(id: "${SPECTRUM_COMMUNITY_ID}") {
        id
        slackSettings {
          teamName
          isConnected
          hasSentInvites
          memberCount
          invitesSentAt
        }
      }
    }
  `;

  const context = { user: noPermissionUser };

  expect.assertions(1);
  const result = await request(query, { context });

  expect(result).toMatchSnapshot();
});

it('should not fetch slack settings if moderates channel', async () => {
  const query = /* GraphQL */ `
    {
      community(id: "${SPECTRUM_COMMUNITY_ID}") {
        id
        slackSettings {
          teamName
          isConnected
          hasSentInvites
          memberCount
          invitesSentAt
        }
      }
    }
  `;

  const context = { user: channelModerator };

  expect.assertions(1);
  const result = await request(query, { context });

  expect(result).toMatchSnapshot();
});

it('should fetch slack settings if moderates community', async () => {
  const query = /* GraphQL */ `
    {
      community(id: "${SPECTRUM_COMMUNITY_ID}") {
        id
        slackSettings {
          teamName
          isConnected
          hasSentInvites
          memberCount
          invitesSentAt
        }
      }
    }
  `;

  const context = { user: communityModerator };

  expect.assertions(1);
  const result = await request(query, { context });

  expect(result).toMatchSnapshot();
});

it('should fetch slack settings if owns community', async () => {
  const query = /* GraphQL */ `
    {
      community(id: "${SPECTRUM_COMMUNITY_ID}") {
        id
        slackSettings {
          teamName
          isConnected
          hasSentInvites
          memberCount
          invitesSentAt
        }
      }
    }
  `;

  const context = { user: communityOwner };

  expect.assertions(1);
  const result = await request(query, { context });

  expect(result).toMatchSnapshot();
});
