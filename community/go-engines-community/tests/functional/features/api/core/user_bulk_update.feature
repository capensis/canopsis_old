Feature: Bulk update users
  I need to be able to bulk update users
  Only admin should be able to bulk update users

  Scenario: given bulk update request and no auth user should not allow access
    When I do PUT /api/v4/bulk/users
    Then the response code should be 401

  Scenario: given bulk update and auth user by api key without permissions should not allow access
    When I am noperms
    When I do PUT /api/v4/bulk/users
    Then the response code should be 403

  Scenario: given bulk update request should return multi status and should be handled independently
    When I am admin
    Then I do PUT /api/v4/bulk/users:
    """json
    [
      {
        "_id": "test-user-to-bulk-update-1",
        "name": "test-user-to-bulk-update-1-updated",
        "firstname": "test-user-to-bulk-update-1-firstname-updated",
        "lastname": "test-user-to-bulk-update-1-lastname-updated",
        "email": "test-user-to-bulk-update-1-email-updated@canopsis.net",
        "role": "test-role-to-edit-user",
        "ui_language": "fr",
        "ui_groups_navigation_type": "top-bar",
        "enable": true,
        "password": "test-password-updated",
        "defaultview": "test-view-to-edit-user",
        "ui_tours": {
          "test-tour-to-bulk-update-user-1": true
        }
      },
      {
        "_id": "test-user-to-bulk-update-1",
        "name": "test-user-to-bulk-update-1-updated",
        "firstname": "test-user-to-bulk-update-1-firstname-updated-twice",
        "lastname": "test-user-to-bulk-update-1-lastname-updated",
        "email": "test-user-to-bulk-update-1-email-updated@canopsis.net",
        "role": "test-role-to-edit-user",
        "ui_language": "fr",
        "ui_groups_navigation_type": "top-bar",
        "enable": true,
        "password": "test-password-updated",
        "defaultview": "test-view-to-edit-user",
        "ui_tours": {
          "test-tour-to-bulk-update-user-1": true
        }
      },
      {
        "role": "not-exist",
        "defaultview": "not-exist"
      },
      {
        "name": "test-user-to-check-unique-name-name"
      },
      {
        "name": "test-user-to-check-unique-name"
      },
      [],
      {
        "_id": "test-user-to-bulk-update-2",
        "name": "test-user-to-bulk-update-2-updated",
        "firstname": "test-user-to-bulk-update-2-firstname-updated",
        "lastname": "test-user-to-bulk-update-2-lastname-updated",
        "email": "test-user-to-bulk-update-2-email-updated@canopsis.net",
        "role": "test-role-to-edit-user",
        "ui_language": "fr",
        "ui_groups_navigation_type": "top-bar",
        "password": "test-password-updated",
        "enable": true,
        "defaultview": "test-view-to-edit-user",
        "ui_tours": {
          "test-tour-to-bulk-update-user-2": true
        }
      }
    ]
    """
    Then the response code should be 207
    Then the response body should contain:
    """json
    [
      {
        "id": "test-user-to-bulk-update-1",
        "status": 200,
        "item": {
          "_id": "test-user-to-bulk-update-1",
          "name": "test-user-to-bulk-update-1-updated",
          "firstname": "test-user-to-bulk-update-1-firstname-updated",
          "lastname": "test-user-to-bulk-update-1-lastname-updated",
          "email": "test-user-to-bulk-update-1-email-updated@canopsis.net",
          "role": "test-role-to-edit-user",
          "ui_language": "fr",
          "ui_groups_navigation_type": "top-bar",
          "enable": true,
          "password": "test-password-updated",
          "defaultview": "test-view-to-edit-user",
          "ui_tours": {
            "test-tour-to-bulk-update-user-1": true
          }
        }
      },
      {
        "id": "test-user-to-bulk-update-1",
        "status": 200,
        "item": {
          "_id": "test-user-to-bulk-update-1",
          "name": "test-user-to-bulk-update-1-updated",
          "firstname": "test-user-to-bulk-update-1-firstname-updated-twice",
          "lastname": "test-user-to-bulk-update-1-lastname-updated",
          "email": "test-user-to-bulk-update-1-email-updated@canopsis.net",
          "role": "test-role-to-edit-user",
          "ui_language": "fr",
          "ui_groups_navigation_type": "top-bar",
          "enable": true,
          "password": "test-password-updated",
          "defaultview": "test-view-to-edit-user",
          "ui_tours": {
            "test-tour-to-bulk-update-user-1": true
          }
        }
      },
      {
        "status": 400,
        "item": {
          "role": "not-exist",
          "defaultview": "not-exist"
        },
        "errors": {
          "_id": "ID is missing.",
          "defaultview": "DefaultView doesn't exist.",
          "email": "Email is missing.",
          "enable": "IsEnabled is missing.",
          "name": "Name is missing.",
          "password": "Password is missing.",
          "role": "Role doesn't exist."
        }
      },
      {
        "status": 400,
        "item": {
          "name": "test-user-to-check-unique-name-name"
        },
        "errors": {
          "_id": "ID is missing.",
          "name": "Name already exists."
        }
      },
      {
        "status": 400,
        "item": {
          "name": "test-user-to-check-unique-name"
        },
        "errors": {
          "_id": "ID is missing.",
          "name": "Name already exists."
        }
      },
      {
        "status": 400,
        "item": [],
        "error": "value doesn't contain object; it contains array"
      },
      {
        "id": "test-user-to-bulk-update-2",
        "status": 200,
        "item": {
          "_id": "test-user-to-bulk-update-2",
          "name": "test-user-to-bulk-update-2-updated",
          "firstname": "test-user-to-bulk-update-2-firstname-updated",
          "lastname": "test-user-to-bulk-update-2-lastname-updated",
          "email": "test-user-to-bulk-update-2-email-updated@canopsis.net",
          "role": "test-role-to-edit-user",
          "ui_language": "fr",
          "ui_groups_navigation_type": "top-bar",
          "password": "test-password-updated",
          "enable": true,
          "defaultview": "test-view-to-edit-user",
          "ui_tours": {
            "test-tour-to-bulk-update-user-2": true
          }
        }
      }
    ]
    """
    When I do GET /api/v4/users?search=test-user-to-bulk-update
    Then the response code should be 200
    Then the response body should contain:
    """json
    {
      "data": [
        {
          "_id": "test-user-to-bulk-update-1",
          "authkey": "5ez4e3jj-7e1e-5c2g-0e91-e079f72o6424",
          "defaultview": {
            "_id": "test-view-to-edit-user",
            "title": "test-view-to-edit-user-title"
          },
          "email": "test-user-to-bulk-update-1-email-updated@canopsis.net",
          "enable": true,
          "external_id": "",
          "firstname": "test-user-to-bulk-update-1-firstname-updated-twice",
          "lastname": "test-user-to-bulk-update-1-lastname-updated",
          "name": "test-user-to-bulk-update-1-updated",
          "role": {
            "_id": "test-role-to-edit-user",
            "name": "test-role-to-edit-user",
            "defaultview": {
              "_id": "test-view-to-edit-user",
              "title": "test-view-to-edit-user-title"
            }
          },
          "source": "",
          "ui_groups_navigation_type": "top-bar",
          "ui_language": "fr",
          "ui_tours": {
            "test-tour-to-bulk-update-user-1": true
          }
        },
        {
          "_id": "test-user-to-bulk-update-2",
          "authkey": "5ez4e3jj-7e1e-5c2g-0e91-e079f72o6424",
          "defaultview": {
            "_id": "test-view-to-edit-user",
            "title": "test-view-to-edit-user-title"
          },
          "email": "test-user-to-bulk-update-2-email-updated@canopsis.net",
          "enable": true,
          "external_id": "",
          "firstname": "test-user-to-bulk-update-2-firstname-updated",
          "lastname": "test-user-to-bulk-update-2-lastname-updated",
          "name": "test-user-to-bulk-update-2-updated",
          "role": {
            "_id": "test-role-to-edit-user",
            "name": "test-role-to-edit-user",
            "defaultview": {
              "_id": "test-view-to-edit-user",
              "title": "test-view-to-edit-user-title"
            }
          },
          "source": "",
          "ui_groups_navigation_type": "top-bar",
          "ui_language": "fr",
          "ui_tours": {
            "test-tour-to-bulk-update-user-2": true
          }
        }
      ],
      "meta": {
        "page": 1,
        "page_count": 1,
        "per_page": 10,
        "total_count": 2
      }
    }
    """
    When I am authenticated with username "test-user-to-bulk-update-1" and password "test-password-updated"
    When I do GET /api/v4/account/me
    Then the response code should be 200
    Then the response body should contain:
    """json
    {
      "_id": "test-user-to-bulk-update-1",
      "name": "test-user-to-bulk-update-1-updated"
    }
    """
    When I am authenticated with username "test-user-to-bulk-update-2" and password "test-password-updated"
    When I do GET /api/v4/account/me
    Then the response code should be 200
    Then the response body should contain:
    """json
    {
      "_id": "test-user-to-bulk-update-2",
      "name": "test-user-to-bulk-update-2-updated"
    }
    """
