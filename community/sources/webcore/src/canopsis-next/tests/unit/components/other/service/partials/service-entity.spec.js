import { generateRenderer, generateShallowRenderer } from '@unit/utils/vue';
import { createAuthModule, createMockedStoreModules } from '@unit/utils/store';
import { mockModals } from '@unit/utils/mock-hooks';

import { ENTITY_TYPES, MODALS, PBEHAVIOR_TYPE_TYPES, USER_PERMISSIONS } from '@/constants';

import ServiceEntity from '@/components/other/service/partials/service-entity.vue';

const stubs = {
  'v-expansion-panel-content': {
    template: '<div><slot name="header"/><slot/></div>',
  },
  'service-entity-header': true,
  'service-entity-info-tab': true,
  'service-entity-tree-of-dependencies-tab': true,
  'pbehaviors-simple-list': true,
};

const snapshotStubs = {
  'service-entity-header': true,
  'service-entity-info-tab': true,
  'service-entity-tree-of-dependencies-tab': true,
  'pbehaviors-simple-list': true,
};

const selectServiceEntityHeader = wrapper => wrapper.find('service-entity-header-stub');
const selectServiceEntityInfoTab = wrapper => wrapper.find('service-entity-info-tab-stub');
const selectTabItems = wrapper => wrapper.findAll('.v-tab');
const selectTabItemByIndex = (wrapper, index) => selectTabItems(wrapper).at(index);

describe('service-entity', () => {
  const $modals = mockModals();
  const entity = {
    _id: 'service-id',
    alarm_id: 'service-alarm-id',
    source_type: ENTITY_TYPES.service,
    pbehaviors: [{
      type: {
        type: PBEHAVIOR_TYPE_TYPES.pause,
      },
    }],
    ack: null,
  };

  const updateSelected = jest.fn();
  const removeUnavailable = jest.fn();

  const { authModule, currentUserPermissionsById } = createAuthModule();

  const store = createMockedStoreModules([authModule]);

  const snapshotFactory = generateRenderer(ServiceEntity, {
    store,

    stubs: snapshotStubs,
    listeners: {
      'update:selected': updateSelected,
      'remove:unavailable': removeUnavailable,
    },
    propsData: {
      entity,
    },
  });

  const factory = generateShallowRenderer(ServiceEntity, {
    store,

    stubs,
    mocks: { $modals },
    listeners: {
      'update:selected': updateSelected,
      'remove:unavailable': removeUnavailable,
    },
    propsData: {
      entity,
    },
  });

  test('Selected update emitted after trigger entity header', async () => {
    const wrapper = factory();

    const header = selectServiceEntityHeader(wrapper);

    await header.triggerCustomEvent('update:selected');

    expect(updateSelected).toHaveBeenCalled();
  });

  test('Remove unavailable emitted after trigger entity header', async () => {
    const wrapper = factory();

    const header = selectServiceEntityHeader(wrapper);

    await header.triggerCustomEvent('remove:unavailable');

    expect(removeUnavailable).toHaveBeenCalled();
  });

  test('Instruction executed after trigger entity info tab', async () => {
    const wrapper = factory();

    const info = selectServiceEntityInfoTab(wrapper);
    const assignedInstruction = {};

    await info.triggerCustomEvent('execute', assignedInstruction);

    expect($modals.show).toBeCalledWith(
      {
        name: MODALS.executeRemediationSimpleInstruction,
        config: {
          assignedInstruction,
          alarmId: entity.alarm_id,
          onClose: expect.any(Function),
          onComplete: expect.any(Function),
        },
      },
    );

    const [modalArguments] = $modals.show.mock.calls[0];

    modalArguments.config.onClose();

    expect(wrapper).toHaveBeenEmit('refresh');
  });

  test('Renders `service-entity` with default props', async () => {
    const wrapper = snapshotFactory();

    await wrapper.openAllExpansionPanels();

    expect(wrapper).toMatchSnapshot();
  });

  test('Renders `service-entity` with default props and declare ticket permission', async () => {
    currentUserPermissionsById.mockReturnValueOnce(({
      [USER_PERMISSIONS.business.serviceWeather.actions.entityDeclareTicket]: {
        actions: [],
      },
    }));
    const wrapper = snapshotFactory();

    await wrapper.openAllExpansionPanels();

    expect(wrapper).toMatchSnapshot();
  });

  test('Renders `service-entity` with custom props', async () => {
    const wrapper = snapshotFactory({
      propsData: {
        entity: {
          _id: 'service-id',
          source_type: ENTITY_TYPES.component,
          pbehaviors: [],
          alarm_id: 'alarm-id',
        },
        selected: true,
        lastActionUnavailable: true,
        entityNameField: 'custom_name',
        widgetParameters: {
          param: 'param',
        },
      },
    });

    await wrapper.openAllExpansionPanels();

    expect(wrapper).toMatchSnapshot();
  });

  test('Renders `service-entity` with tree of deps tab', async () => {
    const wrapper = snapshotFactory({});

    await wrapper.openAllExpansionPanels();

    await selectTabItemByIndex(wrapper, 1).trigger('click');

    expect(wrapper).toMatchSnapshot();
  });

  test('Renders `service-entity` with pbehaviors tab', async () => {
    currentUserPermissionsById.mockReturnValueOnce(({
      [USER_PERMISSIONS.business.serviceWeather.actions.entityManagePbehaviors]: {
        actions: [],
      },
    }));

    const wrapper = snapshotFactory({
      store: createMockedStoreModules([authModule]),
    });

    await wrapper.openAllExpansionPanels();

    await selectTabItemByIndex(wrapper, 1).trigger('click');

    expect(wrapper).toMatchSnapshot();
  });

  test('Renders `service-entity` after opened flag updated', async () => {
    const wrapper = snapshotFactory();

    await wrapper.setData({ opened: [true] });

    expect(wrapper).toMatchSnapshot();
  });
});
