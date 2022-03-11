import { mount, shallowMount, createVueInstance } from '@unit/utils/vue';

import { createSelectInputStub } from '@unit/stubs/input';
import { SORT_ORDERS } from '@/constants';

import DefaultSortColumn from '@/components/side-bars/settings/fields/common/default-sort-column.vue';

const localVue = createVueInstance();

const stubs = {
  'v-select': createSelectInputStub('v-select'),
};

const factory = (options = {}) => shallowMount(DefaultSortColumn, {
  localVue,
  stubs,

  ...options,
});

const snapshotFactory = (options = {}) => mount(DefaultSortColumn, {
  localVue,

  parentComponent: {
    provide: {
      list: {
        register: jest.fn(),
        unregister: jest.fn(),
      },
      listClick: jest.fn(),
    },
  },

  ...options,
});

const selectColumnSelectField = wrapper => wrapper.findAll('select.v-select').at(0);
const selectOrderSelectField = wrapper => wrapper.findAll('select.v-select').at(1);

describe('default-sort-column', () => {
  const columns = [
    { value: 'first', label: 'First' },
    { value: 'second', label: 'Second' },
  ];

  it('Column changed after trigger select field with columns', () => {
    const [firstColumn, secondColumn] = columns;
    const wrapper = factory({
      propsData: {
        value: {
          column: firstColumn.value,
          order: SORT_ORDERS.desc,
        },
        columns,
      },
    });

    const columnField = selectColumnSelectField(wrapper);

    columnField.setValue(secondColumn.value);

    const inputEvents = wrapper.emitted('input');

    expect(inputEvents).toHaveLength(1);

    const [eventData] = inputEvents[0];
    expect(eventData).toEqual({ column: secondColumn.value, order: SORT_ORDERS.desc });
  });

  it('Order changed after trigger select field with orders', () => {
    const [firstColumn] = columns;
    const wrapper = factory({
      propsData: {
        value: {
          column: firstColumn.value,
          order: SORT_ORDERS.desc,
        },
        columns,
      },
    });

    const orderField = selectOrderSelectField(wrapper);

    orderField.setValue(SORT_ORDERS.asc);

    const inputEvents = wrapper.emitted('input');

    expect(inputEvents).toHaveLength(1);

    const [eventData] = inputEvents[0];
    expect(eventData).toEqual({ column: firstColumn.value, order: SORT_ORDERS.asc });
  });

  it('Renders `default-sort-column` with default props', () => {
    const wrapper = snapshotFactory();

    const menuContents = wrapper.findAllMenus();

    expect(wrapper.element).toMatchSnapshot();

    menuContents.wrappers.forEach(({ element }) => {
      expect(element).toMatchSnapshot();
    });
  });

  it('Renders `default-sort-column` with custom props', () => {
    const wrapper = snapshotFactory({
      propsData: {
        value: {
          column: 'createdAt',
          order: SORT_ORDERS.desc,
        },
        columns: [{
          label: 'Created at',
          value: 'createdAt',
        }],
        columnsLabel: 'Columns label',
      },
    });

    const menuContents = wrapper.findAllMenus();

    expect(wrapper.element).toMatchSnapshot();

    menuContents.wrappers.forEach(({ element }) => {
      expect(element).toMatchSnapshot();
    });
  });
});
