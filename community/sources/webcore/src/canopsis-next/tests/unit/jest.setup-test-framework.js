import { cloneDeep, isFunction, kebabCase } from 'lodash';
import { toMatchSnapshot } from 'jest-snapshot';
import registerRequireContextHook from 'babel-plugin-require-context-hook/register';
import { toMatchImageSnapshot } from 'jest-image-snapshot';
import ResizeObserver from 'resize-observer-polyfill';
import flatten from 'flat';

registerRequireContextHook();

global.ResizeObserver = ResizeObserver;
global.IntersectionObserver = jest.fn(() => ({
  observe: jest.fn(),
  unobserve: jest.fn(),
  disconnect: jest.fn(),
}));

Object.defineProperty(HTMLElement.prototype, 'innerText', {
  set(value) {
    this.textContent = value;
  },
  get() {
    return this.textContent;
  },
});

function toMatchCanvasSnapshot(canvasOrWrapper, options, ...args) {
  const canvas = canvasOrWrapper?.tagName !== 'CANVAS' && canvasOrWrapper.find
    ? canvasOrWrapper.find('canvas')?.element
    : canvasOrWrapper;

  const img = canvas.toDataURL();
  const data = img.replace(/^data:image\/(png|jpg);base64,/, '');
  const newOptions = {
    comparisonMethod: 'ssim',
    diffDirection: 'vertical',
    customDiffConfig: {
      ssim: 'fast',
    },
    failureThreshold: 0.1,
    failureThresholdType: 'percent',
    customSnapshotIdentifier: ({ currentTestName, counter }) => (
      kebabCase(`${currentTestName.replace(/(.*\sRenders\s)|(.$)/g, '')}-${counter}`)
    ),

  };

  return toMatchImageSnapshot.call(this, data, newOptions, ...args);
}

function toMatchTooltipSnapshot(wrapper) {
  const tooltip = wrapper.findTooltip();

  return toMatchSnapshot.call(this, tooltip);
}

function toMatchMenuSnapshot(wrapper) {
  const menu = wrapper.findMenu();

  return toMatchSnapshot.call(this, menu);
}

function toEmit(wrapper, event, ...data) {
  // eslint-disable-next-line no-restricted-syntax
  const emittedEvents = wrapper.emitted(event);

  if (this.isNot) {
    try {
      expect(emittedEvents).not.toBeTruthy();
    } catch (err) {
      return {
        pass: true,
        message: () => `Event '${event}' emitted`,
      };
    }
  }

  try {
    if (!data.length) {
      expect(emittedEvents).toBeTruthy();

      return { pass: true };
    }
  } catch (err) {
    return {
      pass: false,
      message: () => `Event '${event}' not emitted`,
    };
  }

  try {
    expect(emittedEvents).toHaveLength(data.length);
  } catch (err) {
    return {
      pass: false,
      message: () => [
        `Expected number of emit: ${data.length}`,
        `Received number of emit: ${emittedEvents?.length ?? 0}`,
      ].join('\n'),
    };
  }

  try {
    expect(
      cloneDeep(emittedEvents.map(events => events[0])),
    ).toEqual(data);
  } catch (err) {
    const { pass, message } = err.matcherResult ?? { pass: false, message: '' };

    return {
      pass,
      message: isFunction(message) ? message : () => message,
    };
  }

  return { pass: true };
}

function toEmitInput(wrapper, ...data) {
  return toEmit.call(this, wrapper, 'input', ...data);
}

function toHaveBeenEmit(wrapper, event) {
  return toEmit.call(this, wrapper, event);
}

function toStructureEqual(received, expected) {
  const flattenReceived = flatten(received);
  const flattenExpected = flatten(expected);

  try {
    expect(flattenReceived).toEqual(Object.keys(flattenExpected).reduce((acc, key) => {
      acc[key] = expect.any(String);

      return acc;
    }, {}));

    return { pass: true };
  } catch (err) {
    return err.matcherResult;
  }
}

function toBeDispatchedWith(received, expected) {
  try {
    expect(received).toBeCalledWith(
      expect.any(Object),
      expected,

    );

    return { pass: true };
  } catch (err) {
    const { pass, message } = err.matcherResult ?? { pass: false, message: '' };

    return {
      pass,
      message: isFunction(message) ? message : () => message,
    };
  }
}

expect.extend({
  toMatchImageSnapshot,
  toMatchCanvasSnapshot,
  toMatchTooltipSnapshot,
  toMatchMenuSnapshot,
  toEmit,
  toEmitInput,
  toHaveBeenEmit,
  toStructureEqual,
  toBeDispatchedWith,
});
