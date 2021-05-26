// @flow
const debug = require('debug')('athena:should-get-email');
import { getUsersSettings } from '../models/usersSettings';
import { getUserById } from 'shared/db/queries/user';

const getEmailStatus = (
  userId: string,
  notificationType: string
): Promise<boolean> => {
  debug(`check email status for user#${userId}`);
  return Promise.all([getUsersSettings(userId), getUserById(userId)])
    .then(([userSettings, user]) => {
      if (
        !userSettings ||
        !userSettings.notifications.types[notificationType] ||
        userSettings.notifications.types[notificationType].email === false
      ) {
        debug(`user#${userId} disabled email notifications`);
        return false;
      }

      if (user.isOnline) {
        debug(`user#${userId} is online, not sending email`);
        return false;
      }

      debug(`user#${userId} is not online and has email notifications enabled`);
      return true;
    })
    .catch(err => {
      console.error('❌ Error in job:\n');
      console.error(err);
      return false;
    });
};

export default getEmailStatus;
