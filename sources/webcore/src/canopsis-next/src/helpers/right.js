import { get } from 'lodash';

export function checkUserAccess(user, rightId, rightMask) {
  const checksum = get(user, ['rights', rightId, 'checksum'], 0);

  return (checksum & rightMask) === rightMask;
}
