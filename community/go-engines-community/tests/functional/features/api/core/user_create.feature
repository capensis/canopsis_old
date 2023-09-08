Feature: Create a user
  I need to be able to create a user
  Only admin should be able to create a user

  @concurrent
  Scenario: given create request should return ok
    When I am admin
    When I do POST /api/v4/users:
    """json
    {
      "name": "test-user-to-create-1-name",
      "firstname": "test-user-to-create-1-firstname",
      "lastname": "test-user-to-create-1-lastname",
      "email": "test-user-to-create-1-email@canopsis.net",
      "roles": [
        "test-role-to-user-edit-2",
        "test-role-to-user-edit-1"
      ],
      "ui_language": "fr",
      "ui_theme": "canopsis",
      "ui_groups_navigation_type": "top-bar",
      "enable": true,
      "defaultview": "test-view-to-edit-user",
      "password": "test-password"
    }
    """
    Then the response code should be 201
    Then the response body should contain:
    """json
    {
      "_id": "test-user-to-create-1-name",
      "defaultview": {
        "_id": "test-view-to-edit-user",
        "title": "test-view-to-edit-user-title"
      },
      "email": "test-user-to-create-1-email@canopsis.net",
      "enable": true,
      "external_id": "",
      "firstname": "test-user-to-create-1-firstname",
      "lastname": "test-user-to-create-1-lastname",
      "name": "test-user-to-create-1-name",
      "roles": [
        {
          "_id": "test-role-to-user-edit-2",
          "name": "test-role-to-user-edit-2",
          "defaultview": {
            "_id": "test-view-to-edit-user",
            "title": "test-view-to-edit-user-title"
          }
        },
        {
          "_id": "test-role-to-user-edit-1",
          "name": "test-role-to-user-edit-1",
          "defaultview": {
            "_id": "test-view-to-edit-user",
            "title": "test-view-to-edit-user-title"
          }
        }
      ],
      "source": "",
      "ui_groups_navigation_type": "top-bar",
      "ui_language": "fr",
      "ui_theme": {
        "name": "Canopsis",
        "colors": {
          "main": {
            "primary": "#2fab63",
            "secondary": "#2b3e4f",
            "accent": "#82b1ff",
            "error": "#ff5252",
            "info": "#2196f3",
            "success": "#4caf50",
            "warning": "#fb8c00",
            "background": "#ffffff",
            "active_color": "#000",
            "font_size": 2
          },
          "table": {
            "background": "#fff",
            "row_color": "#fff",
            "hover_row_color": "#eee"
          },
          "state": {
            "ok": "#00a65a",
            "minor": "#fcdc00",
            "major": "#ff9900",
            "critical": "#f56954"
          }
        }
      }
    }
    """
    When I do GET /api/v4/users/{{ .lastResponse._id}}
    Then the response code should be 200
    Then the response body should contain:
    """json
    {
      "_id": "test-user-to-create-1-name",
      "defaultview": {
        "_id": "test-view-to-edit-user",
        "title": "test-view-to-edit-user-title"
      },
      "email": "test-user-to-create-1-email@canopsis.net",
      "enable": true,
      "external_id": "",
      "firstname": "test-user-to-create-1-firstname",
      "lastname": "test-user-to-create-1-lastname",
      "name": "test-user-to-create-1-name",
      "roles": [
        {
          "_id": "test-role-to-user-edit-2",
          "name": "test-role-to-user-edit-2",
          "defaultview": {
            "_id": "test-view-to-edit-user",
            "title": "test-view-to-edit-user-title"
          }
        },
        {
          "_id": "test-role-to-user-edit-1",
          "name": "test-role-to-user-edit-1",
          "defaultview": {
            "_id": "test-view-to-edit-user",
            "title": "test-view-to-edit-user-title"
          }
        }
      ],
      "source": "",
      "ui_groups_navigation_type": "top-bar",
      "ui_language": "fr",
      "ui_theme": {
        "name": "Canopsis",
        "colors": {
          "main": {
            "primary": "#2fab63",
            "secondary": "#2b3e4f",
            "accent": "#82b1ff",
            "error": "#ff5252",
            "info": "#2196f3",
            "success": "#4caf50",
            "warning": "#fb8c00",
            "background": "#ffffff",
            "active_color": "#000",
            "font_size": 2
          },
          "table": {
            "background": "#fff",
            "row_color": "#fff",
            "hover_row_color": "#eee"
          },
          "state": {
            "ok": "#00a65a",
            "minor": "#fcdc00",
            "major": "#ff9900",
            "critical": "#f56954"
          }
        }
      }
    }
    """

  @concurrent
  Scenario: given create request should auth new user by base auth
    When I am admin
    When I do POST /api/v4/users:
    """json
    {
      "name": "test-user-to-create-3-name",
      "firstname": "test-user-to-create-3-firstname",
      "lastname": "test-user-to-create-3-lastname",
      "email": "test-user-to-create-3-email@canopsis.net",
      "roles": [
        "test-role-to-user-edit-1"
      ],
      "ui_language": "fr",
      "ui_theme": "canopsis",
      "ui_groups_navigation_type": "top-bar",
      "enable": true,
      "defaultview": "test-view-to-edit-user",
      "password": "test-password"
    }
    """
    Then the response code should be 201
    When I am authenticated with username "test-user-to-create-3-name" and password "test-password"
    When I do GET /api/v4/account/me
    Then the response code should be 200
    Then the response body should contain:
    """json
    {
      "_id": "test-user-to-create-3-name",
      "name": "test-user-to-create-3-name"
    }
    """

  @concurrent
  Scenario: given create request should auth new user by password
    When I am admin
    When I do POST /api/v4/users:
    """json
    {
      "name": "test-user-to-create-4-name",
      "firstname": "test-user-to-create-4-firstname",
      "lastname": "test-user-to-create-4-lastname",
      "email": "test-user-to-create-4-email@canopsis.net",
      "roles": [
        "test-role-to-user-edit-1"
      ],
      "ui_language": "fr",
      "ui_theme": "canopsis",
      "ui_groups_navigation_type": "top-bar",
      "enable": true,
      "defaultview": "test-view-to-edit-user",
      "password": "test-password"
    }
    """
    Then the response code should be 201
    When I do POST /api/v4/login:
    """json
    {
      "username": "test-user-to-create-4-name",
      "password": "test-password"
    }
    """
    When I set header Authorization=Bearer {{ .lastResponse.access_token }}
    When I do GET /api/v4/account/me
    Then the response code should be 200
    Then the response body should contain:
    """json
    {
      "_id": "test-user-to-create-4-name",
      "name": "test-user-to-create-4-name"
    }
    """

  @concurrent
  Scenario: given create request and no auth user should not allow access
    When I do POST /api/v4/users
    Then the response code should be 401

  @concurrent
  Scenario: given create request and auth user by api key without permissions should not allow access
    When I am noperms
    When I do POST /api/v4/users
    Then the response code should be 403

  @concurrent
  Scenario: given invalid create request should return errors
    When I am admin
    When I do POST /api/v4/users:
    """json
    {
      "roles": [
        "test-role-to-user-edit-1",
        "not-exist"
      ],
      "defaultview": "not-exist"
    }
    """
    Then the response code should be 400
    Then the response body should be:
    """json
    {
      "errors": {
        "defaultview": "DefaultView doesn't exist.",
        "email": "Email is missing.",
        "enable": "IsEnabled is missing.",
        "name": "Name is missing.",
        "password": "Password is missing.",
        "roles": "Roles doesn't exist."
      }
    }
    """

  @concurrent
  Scenario: given create request with already exists name should return error
    When I am admin
    When I do POST /api/v4/users:
    """json
    {
      "name": "test-user-to-check-unique-name-name"
    }
    """
    Then the response code should be 400
    Then the response body should contain:
    """json
    {
      "errors": {
        "name": "Name already exists."
      }
    }
    """

  @concurrent
  Scenario: given create request with already exists id should return error
    When I am admin
    When I do POST /api/v4/users:
    """json
    {
      "name": "test-user-to-check-unique-name"
    }
    """
    Then the response code should be 400
    Then the response body should contain:
    """json
    {
      "errors": {
        "name": "Name already exists."
      }
    }
    """

  @concurrent
  Scenario: given create request with invalid password should return error
    When I am admin
    When I do POST /api/v4/users:
    """json
    {
      "password": "1"
    }
    """
    Then the response code should be 400
    Then the response body should contain:
    """json
    {
      "errors": {
        "password": "Password should be 8 or more."
      }
    }
    """

  @concurrent
  Scenario: given create request should create user with source and external_id
    When I am admin
    When I do POST /api/v4/users:
    """json
    {
      "name": "test-user-to-create-5-name",
      "firstname": "test-user-to-create-5-firstname",
      "lastname": "test-user-to-create-5-lastname",
      "email": "test-user-to-create-5-email@canopsis.net",
      "roles": [
        "test-role-to-user-edit-1"
      ],
      "ui_language": "fr",
      "ui_groups_navigation_type": "top-bar",
      "enable": true,
      "defaultview": "test-view-to-edit-user",
      "source": "saml",
      "external_id": "saml_id"
    }
    """
    Then the response code should be 201
    When I do GET /api/v4/users/{{ .lastResponse._id}}
    Then the response code should be 200
    Then the response body should contain:
    """json
    {
      "_id": "test-user-to-create-5-name",
      "defaultview": {
        "_id": "test-view-to-edit-user",
        "title": "test-view-to-edit-user-title"
      },
      "email": "test-user-to-create-5-email@canopsis.net",
      "enable": true,
      "firstname": "test-user-to-create-5-firstname",
      "lastname": "test-user-to-create-5-lastname",
      "name": "test-user-to-create-5-name",
      "roles": [
        {
          "_id": "test-role-to-user-edit-1",
          "name": "test-role-to-user-edit-1",
          "defaultview": {
            "_id": "test-view-to-edit-user",
            "title": "test-view-to-edit-user-title"
          }
        }
      ],
      "ui_groups_navigation_type": "top-bar",
      "ui_language": "fr",
      "source": "saml",
      "external_id": "saml_id"
    }
    """

  @concurrent
  Scenario: given create request when only source exists should return error
    When I am admin
    When I do POST /api/v4/users:
    """json
    {
      "name": "test-user-to-create-6-name",
      "firstname": "test-user-to-create-6-firstname",
      "lastname": "test-user-to-create-6-lastname",
      "email": "test-user-to-create-6-email@canopsis.net",
      "roles": [
        "test-role-to-user-edit-1"
      ],
      "ui_language": "fr",
      "ui_groups_navigation_type": "top-bar",
      "enable": true,
      "defaultview": "test-view-to-edit-user",
      "source": "saml"
    }
    """
    Then the response code should be 400
    Then the response body should contain:
    """json
    {
      "errors": {
        "external_id": "ExternalID is required when Source is present."
      }
    }
    """

  @concurrent
  Scenario: given create request when only external_id exists should return error
    When I am admin
    When I do POST /api/v4/users:
    """json
    {
      "name": "test-user-to-create-6-name",
      "firstname": "test-user-to-create-6-firstname",
      "lastname": "test-user-to-create-6-lastname",
      "email": "test-user-to-create-6-email@canopsis.net",
      "roles": [
        "test-role-to-user-edit-1"
      ],
      "ui_language": "fr",
      "ui_groups_navigation_type": "top-bar",
      "enable": true,
      "defaultview": "test-view-to-edit-user",
      "external_id": "saml_id"
    }
    """
    Then the response code should be 400
    Then the response body should contain:
    """json
    {
      "errors": {
        "source": "Source is required when ExternalID is present."
      }
    }
    """

  @concurrent
  Scenario: given create request with wrong source should return error
    When I am admin
    When I do POST /api/v4/users:
    """json
    {
      "name": "test-user-to-create-6-name",
      "firstname": "test-user-to-create-6-firstname",
      "lastname": "test-user-to-create-6-lastname",
      "email": "test-user-to-create-6-email@canopsis.net",
      "roles": [
        "test-role-to-user-edit-1"
      ],
      "ui_language": "fr",
      "ui_groups_navigation_type": "top-bar",
      "enable": true,
      "defaultview": "test-view-to-edit-user",
      "source": "some",
      "external_id": "saml_id"
    }
    """
    Then the response code should be 400
    Then the response body should contain:
    """json
    {
      "errors": {
        "source": "Source must be one of [ldap cas saml] or empty."
      }
    }
    """

  @concurrent
  Scenario: given create request with source and password should return error
    When I am admin
    When I do POST /api/v4/users:
    """json
    {
      "name": "test-user-to-create-6-name",
      "firstname": "test-user-to-create-6-firstname",
      "lastname": "test-user-to-create-6-lastname",
      "email": "test-user-to-create-6-email@canopsis.net",
      "roles": [
        "test-role-to-user-edit-1"
      ],
      "ui_language": "fr",
      "ui_groups_navigation_type": "top-bar",
      "enable": true,
      "defaultview": "test-view-to-edit-user",
      "password": "test-password",
      "source": "some",
      "external_id": "saml_id"
    }
    """
    Then the response code should be 400
    Then the response body should contain:
    """json
    {
      "errors": {
        "source": "Can't be present both Source and Password."
      }
    }
    """

  @concurrent
  Scenario: given create request without source and without password should return error
    When I am admin
    When I do POST /api/v4/users:
    """json
    {
      "name": "test-user-to-create-6-name",
      "firstname": "test-user-to-create-6-firstname",
      "lastname": "test-user-to-create-6-lastname",
      "email": "test-user-to-create-6-email@canopsis.net",
      "roles": [
        "test-role-to-user-edit-1"
      ],
      "ui_language": "fr",
      "ui_groups_navigation_type": "top-bar",
      "enable": true,
      "defaultview": "test-view-to-edit-user"
    }
    """
    Then the response code should be 400
    Then the response body should contain:
    """json
    {
      "errors": {
        "password": "Password is missing."
      }
    }
    """

  @concurrent
  Scenario: given create request with invalid ui_theme should return error
    When I am admin
    When I do POST /api/v4/users:
    """json
    {
      "name": "test-user-to-create-6-name",
      "firstname": "test-user-to-create-6-firstname",
      "lastname": "test-user-to-create-6-lastname",
      "email": "test-user-to-create-6-email@canopsis.net",
      "roles": [
        "test-role-to-user-edit-1"
      ],
      "ui_theme": "not found",
      "ui_language": "fr",
      "ui_groups_navigation_type": "top-bar",
      "enable": true,
      "defaultview": "test-view-to-edit-user"
    }
    """
    Then the response code should be 400
    Then the response body should contain:
    """json
    {
      "errors": {
        "ui_theme": "UITheme doesn't exist."
      }
    }
    """
    
  @concurrent
  Scenario: given create request with empty ui_theme should return default theme
    When I am admin
    When I do POST /api/v4/users:
    """json
    {
      "name": "test-user-to-create-7-name",
      "firstname": "test-user-to-create-7-firstname",
      "lastname": "test-user-to-create-7-lastname",
      "email": "test-user-to-create-7-email@canopsis.net",
      "roles": [
        "test-role-to-user-edit-2",
        "test-role-to-user-edit-1"
      ],
      "ui_language": "fr",
      "ui_groups_navigation_type": "top-bar",
      "enable": true,
      "defaultview": "test-view-to-edit-user",
      "password": "test-password"
    }
    """
    Then the response code should be 201
    Then the response body should contain:
    """json
    {
      "ui_theme": {
        "name": "Canopsis",
        "colors": {
          "main": {
            "primary": "#2fab63",
            "secondary": "#2b3e4f",
            "accent": "#82b1ff",
            "error": "#ff5252",
            "info": "#2196f3",
            "success": "#4caf50",
            "warning": "#fb8c00",
            "background": "#ffffff",
            "active_color": "#000",
            "font_size": 2
          },
          "table": {
            "background": "#fff",
            "row_color": "#fff",
            "hover_row_color": "#eee"
          },
          "state": {
            "ok": "#00a65a",
            "minor": "#fcdc00",
            "major": "#ff9900",
            "critical": "#f56954"
          }
        }
      }
    }
    """

  @concurrent
  Scenario: given create request with custom ui_theme should return ok
    When I am admin
    When I do POST /api/v4/users:
    """json
    {
      "name": "test-user-to-create-8-name",
      "firstname": "test-user-to-create-8-firstname",
      "lastname": "test-user-to-create-8-lastname",
      "email": "test-user-to-create-8-email@canopsis.net",
      "roles": [
        "test-role-to-user-edit-2",
        "test-role-to-user-edit-1"
      ],
      "ui_language": "fr",
      "ui_theme": "test_theme_to_pick_1",
      "ui_groups_navigation_type": "top-bar",
      "enable": true,
      "defaultview": "test-view-to-edit-user",
      "password": "test-password"
    }
    """
    Then the response code should be 201
    Then the response body should contain:
    """json
    {
      "ui_theme": {
        "name": "test_theme_to_pick_1",
        "colors": {
          "main": {
            "primary": "#AAAAAA",
            "secondary": "#AAAAAA",
            "accent": "#AAAAAA",
            "error": "#AAAAAA",
            "info": "#AAAAAA",
            "success": "#AAAAAA",
            "warning": "#AAAAAA",
            "background": "#AAAAAA",
            "active_color": "#AAAAAA",
            "font_size": 2
          },
          "table": {
            "background": "#AAAAAA",
            "row_color": "#AAAAAA",
            "shift_row_color": "#AAAAAA",
            "hover_row_color": "#AAAAAA"
          },
          "state": {
            "ok": "#AAAAAA",
            "minor": "#AAAAAA",
            "major": "#AAAAAA",
            "critical": "#AAAAAA"
          }
        }
      }
    }
    """

  @concurrent
  Scenario: given user create request and delete color theme, picked theme should be replaced to default
    When I am admin
    When I do POST /api/v4/users:
    """json
    {
      "name": "test-user-to-create-9-name",
      "firstname": "test-user-to-create-9-firstname",
      "lastname": "test-user-to-create-9-lastname",
      "email": "test-user-to-create-9-email@canopsis.net",
      "roles": [
        "test-role-to-user-edit-2",
        "test-role-to-user-edit-1"
      ],
      "ui_language": "fr",
      "ui_theme": "test_theme_to_pick_4",
      "ui_groups_navigation_type": "top-bar",
      "enable": true,
      "defaultview": "test-view-to-edit-user",
      "password": "test-password"
    }
    """
    Then the response code should be 201
    Then the response body should contain:
    """json
    {
      "ui_theme": {
        "name": "test_theme_to_pick_4",
        "colors": {
          "main": {
            "primary": "#AAAAAA",
            "secondary": "#AAAAAA",
            "accent": "#AAAAAA",
            "error": "#AAAAAA",
            "info": "#AAAAAA",
            "success": "#AAAAAA",
            "warning": "#AAAAAA",
            "background": "#AAAAAA",
            "active_color": "#AAAAAA",
            "font_size": 2
          },
          "table": {
            "background": "#AAAAAA",
            "row_color": "#AAAAAA",
            "shift_row_color": "#AAAAAA",
            "hover_row_color": "#AAAAAA"
          },
          "state": {
            "ok": "#AAAAAA",
            "minor": "#AAAAAA",
            "major": "#AAAAAA",
            "critical": "#AAAAAA"
          }
        }
      }
    }
    """
    When I do DELETE /api/v4/color-themes/test_theme_to_pick_4
    Then the response code should be 204
    When I do GET /api/v4/users/test-user-to-create-9-name
    Then the response code should be 200
    Then the response body should contain:
    """json
    {
      "ui_theme": {
        "name": "Canopsis",
        "colors": {
          "main": {
            "primary": "#2fab63",
            "secondary": "#2b3e4f",
            "accent": "#82b1ff",
            "error": "#ff5252",
            "info": "#2196f3",
            "success": "#4caf50",
            "warning": "#fb8c00",
            "background": "#ffffff",
            "active_color": "#000",
            "font_size": 2
          },
          "table": {
            "background": "#fff",
            "row_color": "#fff",
            "hover_row_color": "#eee"
          },
          "state": {
            "ok": "#00a65a",
            "minor": "#fcdc00",
            "major": "#ff9900",
            "critical": "#f56954"
          }
        }
      }
    }
    """
