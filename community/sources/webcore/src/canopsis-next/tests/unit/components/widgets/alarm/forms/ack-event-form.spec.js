import Faker from 'faker';
import flushPromises from 'flush-promises';

import { createVueInstance, generateShallowRenderer, generateRenderer } from '@unit/utils/vue';
import { createInputStub, createCheckboxInputStub } from '@unit/stubs/input';
import { createActivatorElementStub } from '@unit/stubs/vuetify';

import AckEventForm from '@/components/widgets/alarm/forms/ack-event-form.vue';

const localVue = createVueInstance();

const stubs = {
  'c-description-field': true,
  'v-text-field': createInputStub('v-text-field'),
  'v-checkbox': createCheckboxInputStub('v-checkbox'),
  'v-tooltip': createActivatorElementStub('v-tooltip'),
};

const snapshotStubs = {
  'c-description-field': true,
};

const selectTextField = wrapper => wrapper.find('.v-text-field');
const selectDescriptionField = wrapper => wrapper.find('c-description-field-stub');
const selectCheckboxField = wrapper => wrapper.find('.v-checkbox');

describe('ack-event-form', () => {
  const factory = generateShallowRenderer(AckEventForm, {
    localVue,
    stubs,
    attachTo: document.body,
  });
  const snapshotFactory = generateRenderer(AckEventForm, {
    localVue,
    stubs: snapshotStubs,
    attachTo: document.body,
  });

  test('Ticket changed after trigger text field', () => {
    const form = {
      ticket: Faker.datatype.string(),
      output: Faker.datatype.string(),
      ack_resources: Faker.datatype.boolean(),
    };
    const wrapper = factory({
      propsData: {
        form,
      },
    });

    const ticket = Faker.datatype.string();

    const textField = selectTextField(wrapper);

    textField.setValue(ticket);

    expect(wrapper).toEmit('input', { ...form, ticket });
  });

  test('Output changed after trigger description field', () => {
    const form = {
      ticket: Faker.datatype.string(),
      output: Faker.datatype.string(),
      ack_resources: Faker.datatype.boolean(),
    };
    const wrapper = factory({
      propsData: {
        form,
      },
    });

    const output = Faker.datatype.string();

    const descriptionField = selectDescriptionField(wrapper);

    descriptionField.vm.$emit('input', output);

    expect(wrapper).toEmit('input', { ...form, output });
  });

  test('Ack resources changed after trigger checkbox field', async () => {
    const form = {
      ticket: Faker.datatype.string(),
      output: Faker.datatype.string(),
      ack_resources: Faker.datatype.boolean(),
    };
    const wrapper = factory({
      propsData: {
        form,
      },
    });

    await flushPromises();

    const ackResources = !form.ack_resources;

    selectCheckboxField(wrapper).vm.$emit('change', ackResources);

    expect(wrapper).toEmit('input', { ...form, ack_resources: ackResources });
  });

  test('Renders `ack-event-form` with default props', () => {
    const wrapper = snapshotFactory();

    expect(wrapper.element).toMatchSnapshot();
  });

  test('Renders `ack-event-form` with custom props', () => {
    const wrapper = snapshotFactory({
      propsData: {
        form: {
          ticket: 'ticket',
          output: 'output',
          ack_resources: true,
        },
        isNoteRequired: true,
      },
    });

    expect(wrapper.element).toMatchSnapshot();
  });
});
