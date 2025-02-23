import { omit } from 'lodash';

import { generateRenderer } from '@unit/utils/vue';

import { ALARM_STATES, ALARM_STATUSES } from '@/constants';

import AlarmColumnValueStatus from '@/components/widgets/alarm/columns-formatting/alarm-column-value-status.vue';

const stubs = {
  'c-no-events-icon': true,
  'c-simple-tooltip': true,
};

describe('alarm-column-value-status', () => {
  const snapshotFactory = generateRenderer(AlarmColumnValueStatus, {
    stubs,
    attachTo: document.body,
  });

  it.each(Object.entries(ALARM_STATES))('Renders `alarm-column-value-status` with ongoing status and state: %s', (_, state) => {
    const wrapper = snapshotFactory({
      propsData: {
        alarm: {
          entity: {},
          v: {
            state: {
              val: state,
            },
            status: {
              val: ALARM_STATUSES.ongoing,
            },
          },
        },
      },
    });

    expect(wrapper).toMatchSnapshot();
  });

  it.each(
    Object.entries(omit(ALARM_STATUSES, ['ongoing', 'noEvents'])),
  )('Renders `alarm-column-value-status` with status: %s', (_, status) => {
    const wrapper = snapshotFactory({
      propsData: {
        alarm: {
          entity: {},
          v: {
            state: {
              val: ALARM_STATES.ok,
            },
            status: {
              val: status,
            },
          },
        },
      },
    });

    expect(wrapper).toMatchSnapshot();
  });

  it('Renders `alarm-column-value-status` with status: noEvents', () => {
    const wrapper = snapshotFactory({
      propsData: {
        alarm: {
          entity: {
            idle_since: 1386435600000,
          },
          v: {
            status: {
              val: ALARM_STATUSES.noEvents,
            },
          },
        },
      },
    });

    expect(wrapper).toMatchSnapshot();
  });
});
