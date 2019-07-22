const { get } = require('lodash');
const uid = require('uid');

const createRole = (browser, {
  name, description, groupId, viewId,
}) => {
  const rolesPage = browser.page.admin.roles();
  const createRoleModal = browser.page.modals.admin.createRole();
  const selectViewModal = browser.page.modals.view.selectView();

  rolesPage.clickAddButton();

  createRoleModal.verifyModalOpened()
    .setName(name)
    .setDescription(description)
    .clickSelectDefaultViewButton();

  selectViewModal.verifyModalOpened()
    .browseGroupById(groupId)
    .browseViewById(viewId)
    .verifyModalClosed();

  browser.waitForFirstXHR('account/role', 1000, () => createRoleModal.clickSubmitButton(), ({ responseData }) => {
    const response = JSON.parse(responseData);

    browser.assert.equal(response.total, 1);

    browser.globals.rolesIds.push(get(response.data, [0, '_id']));

    createRoleModal.verifyModalClosed();
  });
};

module.exports = {
  async before(browser, done) {
    await browser.maximizeWindow()
      .completed.loginAsAdmin();

    browser.globals.rolesIds = [];

    done();
  },

  after(browser, done) {
    delete browser.globals.rolesIds;

    browser.end(done);
  },

  'Create new role with data from constants': (browser) => {
    const role = {
      name: 'Test role',
      description: 'Test role description',
      groupId: '05b2e049-b3c4-4c5b-94a5-6e7ff142b28c',
      viewId: '875df4c2-027b-4549-8add-e20ed7ff7d4f',
    };

    browser.page.admin.roles()
      .navigate()
      .verifyPageElementsBefore();

    createRole(browser, role);
  },


  'Pagination on data-table': (browser) => {
    const rolesPage = browser.page.admin.roles();

    rolesPage.clickNextButton()
      .defaultPause();

    rolesPage.clickPrevButton()
      .defaultPause();

    rolesPage.selectRange(5)
      .defaultPause();
  },

  'Edit created role by data from constants': (browser) => {
    const rolesPage = browser.page.admin.roles();
    const createRoleModal = browser.page.modals.admin.createRole();
    const roleId = browser.globals.rolesIds[0];

    rolesPage.verifyPageRoleBefore(roleId)
      .clickEditButton(roleId);

    createRoleModal.verifyModalOpened()
      .clickSubmitButton()
      .verifyModalClosed();
  },

  'Delete created role': (browser) => {
    const rolesPage = browser.page.admin.roles();
    const confirmationModal = browser.page.modals.common.confirmation();
    const roleId = browser.globals.rolesIds.shift();

    rolesPage.verifyPageRoleBefore(roleId)
      .clickDeleteButton(roleId);

    confirmationModal.verifyModalOpened()
      .clickSubmitButton()
      .verifyModalClosed();
  },

  'Create several new roles with data from constants': (browser) => {
    const rolesData = [
      {
        name: `Test role ${uid()}`,
        description: 'Test role description',
        groupId: '05b2e049-b3c4-4c5b-94a5-6e7ff142b28c',
        viewId: '875df4c2-027b-4549-8add-e20ed7ff7d4f',
      }, {
        name: `Test role ${uid()}`,
        description: 'Test role description',
        groupId: '05b2e049-b3c4-4c5b-94a5-6e7ff142b28c',
        viewId: '875df4c2-027b-4549-8add-e20ed7ff7d4f',
      }, {
        name: `Test role ${uid()}`,
        description: 'Test role description',
        groupId: '05b2e049-b3c4-4c5b-94a5-6e7ff142b28c',
        viewId: '875df4c2-027b-4549-8add-e20ed7ff7d4f',
      },
    ];

    browser.page.admin.roles()
      .navigate()
      .verifyPageElementsBefore();

    rolesData.forEach((role) => {
      createRole(browser, role);
    });
  },

  'Mass delete created roles': (browser) => {
    const rolesPage = browser.page.admin.roles();
    const confirmationModal = browser.page.modals.common.confirmation();
    const { rolesIds } = browser.globals;

    rolesPage.selectRange(5)
      .defaultPause();

    rolesIds.forEach((roleId) => {
      rolesPage.verifyPageRoleBefore(roleId)
        .clickOptionCheckbox(roleId)
        .defaultPause();
    });

    rolesPage.verifyMassDeleteButton()
      .clickMassDeleteButton();

    confirmationModal.verifyModalOpened()
      .clickSubmitButton()
      .verifyModalClosed();
  },

  'Refresh button': (browser) => {
    const rolesPage = browser.page.admin.roles();

    rolesPage.waitForFirstXHR(
      'rest/default_rights/role',
      5000,
      () => rolesPage.clickRefreshButton(),
      ({ status, method, httpResponseCode }) => {
        browser.assert.equal(status, 'success');
        browser.assert.equal(method, 'GET');
        browser.assert.equal(httpResponseCode, '200');
      },
    );
  },
};
