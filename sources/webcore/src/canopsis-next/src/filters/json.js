import { isObject } from 'lodash';

export default function (json, indents = 4) {
  try {
    if (isObject(json)) {
      return JSON.stringify(json, null, indents);
    }

    return JSON.stringify(JSON.parse(json), null, indents);
  } catch (err) {
    console.warn(err);

    return '{}';
  }
}
