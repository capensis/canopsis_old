import Faker from 'faker';
import { Validator } from 'vee-validate';

import { mount, shallowMount, createVueInstance } from '@unit/utils/vue';
import { FILTER_INPUT_TYPES } from '@/constants';

import CMixedField from '@/components/forms/fields/c-mixed-field.vue';

const localVue = createVueInstance();

const stubs = {
  'v-select': {
    props: ['value', 'items'],
    template: `
      <select class="v-select" :value="value" @change="$listeners.input($event.target.value)">
        <option v-for="item in items" :value="item.value" :key="item.value">
          {{ item.value }}
        </option>
      </select>
    `,
  },
  'v-text-field': {
    props: ['value'],
    template: `
      <input
        :value="value"
        class="v-text-field"
        @input="$listeners.input($event.target.value)"
      />
    `,
  },
  'v-combobox': {
    props: ['value'],
    template: `
      <input
        :value="value"
        class="v-combobox"
        @input="$listeners.input($event.target.value)"
      />
    `,
  },
};

const snapshotStubs = {
  'c-array-mixed-field': true,
  'v-select': {
    props: ['items'],
    template: `
      <div class="v-select">
        <select>
          <option v-for="(item, index) in items" :value="item.value" :key="item.value">
            <slot name="selection" :item="item" :index="index" />
            <slot name="item" :item="item" />
          </option>
        </select>
      </div>
    `,
  },
};

const factory = (options = {}) => shallowMount(CMixedField, {
  localVue,
  stubs,
  ...options,
});

describe('c-mixed-field', () => {
  it('Value set into select', () => {
    const value = Faker.datatype.number();
    const wrapper = factory({
      provide: {
        $validator: new Validator(),
      },
      propsData: {
        value,
      },
    });

    const selectElement = wrapper.find('select.v-select');

    expect(selectElement.element.value).toBe(FILTER_INPUT_TYPES.number);
  });

  it('Value changed after change the type select to string', () => {
    const value = Faker.datatype.number();
    const wrapper = factory({
      provide: {
        $validator: new Validator(),
      },
      propsData: {
        value,
      },
    });
    const selectElement = wrapper.find('select.v-select');

    selectElement.setValue(FILTER_INPUT_TYPES.string);

    const inputEvents = wrapper.emitted('input');
    expect(inputEvents).toHaveLength(1);

    const [inputEventData] = inputEvents[0];
    expect(inputEventData).toEqual(`${value}`);
  });

  it('Value changed after change the type select to number', () => {
    const number = Faker.datatype.number();
    const wrapper = factory({
      provide: {
        $validator: new Validator(),
      },
      propsData: {
        value: `${number}`,
      },
    });
    const selectElement = wrapper.find('select.v-select');

    selectElement.setValue(FILTER_INPUT_TYPES.number);

    const inputEvents = wrapper.emitted('input');
    expect(inputEvents).toHaveLength(1);

    const [inputEventData] = inputEvents[0];
    expect(inputEventData).toEqual(number);
  });

  it('Value changed after change the type select to boolean', () => {
    const wrapper = factory({
      provide: {
        $validator: new Validator(),
      },
      propsData: {
        value: Faker.datatype.string(),
      },
    });
    const selectElement = wrapper.find('select.v-select');

    selectElement.setValue(FILTER_INPUT_TYPES.boolean);

    const inputEvents = wrapper.emitted('input');
    expect(inputEvents).toHaveLength(1);

    const [inputEventData] = inputEvents[0];
    expect(inputEventData).toEqual(true);
  });

  it('Value changed after change the type select to null', () => {
    const wrapper = factory({
      provide: {
        $validator: new Validator(),
      },
      propsData: {
        value: Faker.datatype.string(),
      },
    });
    const selectElement = wrapper.find('select.v-select');

    selectElement.setValue(FILTER_INPUT_TYPES.null);

    const inputEvents = wrapper.emitted('input');
    expect(inputEvents).toHaveLength(1);

    const [inputEventData] = inputEvents[0];
    expect(inputEventData).toEqual(null);
  });

  it('Value changed after change the type select to array', () => {
    const value = Faker.datatype.string();
    const wrapper = factory({
      provide: {
        $validator: new Validator(),
      },
      propsData: {
        value,
      },
    });
    const selectElement = wrapper.find('select.v-select');

    selectElement.setValue(FILTER_INPUT_TYPES.array);

    const inputEvents = wrapper.emitted('input');
    expect(inputEvents).toHaveLength(1);

    const [inputEventData] = inputEvents[0];
    expect(inputEventData).toEqual([value]);
  });

  it('Value changed on the first type after remove selected type', async () => {
    const value = Faker.datatype.number();
    const wrapper = factory({
      provide: {
        $validator: new Validator(),
      },
      propsData: {
        value,
        types: [
          { value: FILTER_INPUT_TYPES.string },
          { value: FILTER_INPUT_TYPES.boolean },
          { value: FILTER_INPUT_TYPES.number },
        ],
      },
    });

    await wrapper.setProps({
      types: [
        { value: FILTER_INPUT_TYPES.string },
        { value: FILTER_INPUT_TYPES.boolean },
      ],
    });

    const inputEvents = wrapper.emitted('input');
    expect(inputEvents).toHaveLength(1);

    const [inputEventData] = inputEvents[0];
    expect(inputEventData).toEqual(`${value}`);
  });

  it('Value change on undefined after remove all types', async () => {
    const value = Faker.datatype.number();
    const wrapper = factory({
      provide: {
        $validator: new Validator(),
      },
      propsData: {
        value,
        types: [
          { value: FILTER_INPUT_TYPES.string },
          { value: FILTER_INPUT_TYPES.boolean },
          { value: FILTER_INPUT_TYPES.number },
        ],
      },
    });

    await wrapper.setProps({
      types: [],
    });

    const inputEvents = wrapper.emitted('input');
    expect(inputEvents).toHaveLength(1);

    const [inputEventData] = inputEvents[0];
    expect(inputEventData).toEqual(undefined);
  });

  it('Value changed to empty string after trigger the input with null value', () => {
    const wrapper = factory({
      provide: {
        $validator: new Validator(),
      },
      propsData: {
        value: 'Value',
      },
    });
    const inputElement = wrapper.find('input.v-text-field');

    inputElement.vm.$emit('input', null);

    const inputEvents = wrapper.emitted('input');
    expect(inputEvents).toHaveLength(1);

    const [inputEventData] = inputEvents[0];
    expect(inputEventData).toEqual('');
  });

  it('Value changed after trigger the input with number value', () => {
    const newNumber = Faker.datatype.number();
    const wrapper = factory({
      provide: {
        $validator: new Validator(),
      },
      propsData: {
        value: 12,
      },
    });
    const inputElement = wrapper.find('input.v-text-field');

    inputElement.setValue(`${newNumber}`);

    const inputEvents = wrapper.emitted('input');
    expect(inputEvents).toHaveLength(1);

    const [inputEventData] = inputEvents[0];
    expect(inputEventData).toEqual(newNumber);
  });

  it('Value changed after trigger the select with items', () => {
    const item = {
      text: Faker.datatype.string(),
      value: Faker.datatype.string(),
    };
    const wrapper = factory({
      provide: {
        $validator: new Validator(),
      },
      propsData: {
        value: '',
        items: [item],
      },
    });
    const comboboxInputElement = wrapper.find('input.v-combobox');

    comboboxInputElement.setValue(item.value);

    const inputEvents = wrapper.emitted('input');
    expect(inputEvents).toHaveLength(1);

    const [inputEventData] = inputEvents[0];
    expect(inputEventData).toEqual(item.value);
  });

  it('v-validate works correctly with component', async () => {
    const name = Faker.datatype.string();
    const value = Faker.datatype.string();
    const validator = new Validator();
    mount({
      inject: ['$validator'],
      components: {
        CMixedField,
      },
      props: ['name', 'value'],
      template: `
        <c-mixed-field v-validate="'required'" :name="name" :value="value" />
      `,
    }, {
      localVue,
      stubs,
      mocks: { $t: () => {} },
      provide: {
        $validator: validator,
      },
      propsData: {
        value,
        name,
      },
    });

    await validator.validateAll();

    expect(validator.fields.find({ name })).toBeTruthy();
  });

  it('Renders `c-mixed-field` with default props correctly', () => {
    const wrapper = shallowMount(CMixedField, {
      localVue,
      stubs: snapshotStubs,
      provide: {
        $validator: new Validator(),
      },
    });

    expect(wrapper.element).toMatchSnapshot();
  });

  it('Renders `c-mixed-field` with errors correctly', () => {
    const validator = new Validator();

    const wrapper = shallowMount(CMixedField, {
      localVue,
      provide: {
        $validator: validator,
      },
      stubs: snapshotStubs,
      propsData: {
        errorMessages: ['First error message', 'Second error message'],
      },
    });

    expect(wrapper.element).toMatchSnapshot();
  });

  it('Renders `c-mixed-field` with custom props correctly', () => {
    const validator = new Validator();

    const wrapper = shallowMount(CMixedField, {
      localVue,
      provide: {
        $validator: validator,
      },
      stubs: snapshotStubs,
      propsData: {
        value: 'Value',
        name: 'mixedFieldName',
        label: 'Mixed field label',
        disabled: true,
        soloInverted: true,
        flat: true,
        hideDetails: true,
        errorMessages: ['First error message', 'Second error message'],
        types: [
          { value: FILTER_INPUT_TYPES.string, text: 'Custom string' },
          { value: FILTER_INPUT_TYPES.number },
          { value: FILTER_INPUT_TYPES.boolean },
        ],
      },
    });

    expect(wrapper.element).toMatchSnapshot();
  });

  it('Renders `c-mixed-field` with string type and items correctly', () => {
    const validator = new Validator();

    const wrapper = shallowMount(CMixedField, {
      localVue,
      provide: {
        $validator: validator,
      },
      stubs: snapshotStubs,
      propsData: {
        value: false,
        name: 'mixedFieldName',
        label: 'Mixed field with boolean type',
        items: [{
          customText: 'Custom item text',
          customValue: 'Custom item value',
        }],
        itemText: 'customText',
        itemValue: 'customValue',
      },
    });

    expect(wrapper.element).toMatchSnapshot();
  });

  it('Renders `c-mixed-field` with boolean type correctly', () => {
    const validator = new Validator();

    const wrapper = shallowMount(CMixedField, {
      localVue,
      provide: {
        $validator: validator,
      },
      stubs: snapshotStubs,
      propsData: {
        value: false,
        name: 'mixedFieldName',
        label: 'Mixed field with boolean type',
      },
    });

    expect(wrapper.element).toMatchSnapshot();
  });

  it('Renders `c-mixed-field` with number type correctly', () => {
    const validator = new Validator();

    const wrapper = shallowMount(CMixedField, {
      localVue,
      provide: {
        $validator: validator,
      },
      stubs: snapshotStubs,
      propsData: {
        value: 222,
        name: 'mixedFieldName',
        label: 'Mixed field with number type',
      },
    });

    expect(wrapper.element).toMatchSnapshot();
  });

  it('Renders `c-mixed-field` with null type correctly', () => {
    const validator = new Validator();

    const wrapper = shallowMount(CMixedField, {
      localVue,
      provide: {
        $validator: validator,
      },
      stubs: snapshotStubs,
      propsData: {
        value: null,
        name: 'mixedFieldName',
        label: 'Mixed field with null type',
      },
    });

    expect(wrapper.element).toMatchSnapshot();
  });

  it('Renders `c-mixed-field` with null type correctly', () => {
    const validator = new Validator();

    const wrapper = shallowMount(CMixedField, {
      localVue,
      provide: {
        $validator: validator,
      },
      stubs: snapshotStubs,
      propsData: {
        value: [0, '1', null, false, []],
        name: 'mixedFieldName',
        label: 'Mixed field with null type',
      },
    });

    expect(wrapper.element).toMatchSnapshot();
  });
});
