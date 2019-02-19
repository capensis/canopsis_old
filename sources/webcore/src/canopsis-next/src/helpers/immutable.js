import { get, clone, setWith, unset } from 'lodash';

/**
 * Immutable method for deep updating object field or array item
 *
 * @param {Object|Array} obj - Object will be copied and copy will be updated
 * @param {string|Array} path - Path to field or item
 * @param {*} value - New field or item value
 * @return {Object|Array}
 */
export function setIn(obj, path, value) {
  return setWith(clone(obj), path, value, clone);
}

/**
 * Immutable method for deep updating object fields
 *
 * @param {Object|Array} obj - Object will be copied and copy will be updated
 * @param {Object} pathsValuesMap - Map for paths and values ex: { 'a.b.c': 'value', 'a.b.y': 'another value' }
 * @return {Object|Array}
 */
export function setInSeveral(obj, pathsValuesMap) {
  const alreadyClonedPaths = {};
  const clonedObject = clone(obj);

  Object.keys(pathsValuesMap).forEach((path) => {
    let currentPath = '';

    setWith(clonedObject, path, pathsValuesMap[path], (value, key) => {
      currentPath += `.${key}`;

      if (alreadyClonedPaths[currentPath]) {
        return value;
      }

      alreadyClonedPaths[currentPath] = true;

      return clone(value);
    });
  });

  return clonedObject;
}

/**
 * Immutable method for deep updating object field or array item by customizer function
 *
 * @param {Object|Array} obj - Object will be copied and copy will be updated
 * @param {string|Array} path - Path to field or item
 * @param {function} [customizer=v => v] - Customizer for updated item
 * @return {Object|Array}
 */
export function setInWith(obj, path, customizer = v => v) {
  return setWith(clone(obj), path, customizer(get(obj, path)), clone);
}

/**
 * Immutable method for deep removing object field
 *
 * @param {Object|Array} obj - Object will be copied and copy will be updated
 * @param {string|Array} path - Path to field or item
 * @return {Object|Array}
 */
export function unsetIn(obj, path) {
  return unset(setIn(obj, path, get(obj, path)), path);
}

/**
 * Immutable method for deep adding new item into array
 *
 * @param {Object|Array} obj - Object will be copied and copy will be updated
 * @param {string|Array} path - Path to field or item
 * @param {*} value - Value of new array item
 * @return {Object|Array}
 */
export function addIn(obj, path, value) {
  return setIn(obj, path, [...get(obj, path, []), value]);
}

/**
 * Immutable method for deep removing item from array
 *
 * @param {Object|Array} obj - Object will be copied and copy will be updated
 * @param {string|Array} path - Path to field or item
 * @param {number} index - Index of array item
 * @return {Object|Array}
 */
export function removeIn(obj, path, index) {
  return setIn(obj, path, get(obj, path, []).filter((v, i) => i !== index));
}

export default {
  setIn,
  setInWith,
  unsetIn,
  addIn,
  removeIn,
};
