Feature: Create a widget filter
  I need to be able to create a widget filter
  Only admin should be able to create a widget filter

  Scenario: given create request should return ok
    When I am admin
    When I do POST /api/v4/widget-filters:
    """json
    {
      "title": "test-widgetfilter-to-create-1-title",
      "widget": "test-widget-to-filter-edit",
      "is_private": true,
      "alarm_pattern": [
        [
          {
            "field": "v.component",
            "cond": {
              "type": "eq",
              "value": "test-widgetfilter-to-create-1-pattern"
            }
          }
        ]
      ],
      "entity_pattern": [
        [
          {
            "field": "name",
            "cond": {
              "type": "eq",
              "value": "test-widgetfilter-to-create-1-pattern"
            }
          }
        ]
      ],
      "pbehavior_pattern": [
        [
          {
            "field": "pbehavior_info.type",
            "cond": {
              "type": "eq",
              "value": "test-widgetfilter-to-create-1-pattern"
            }
          }
        ]
      ]
    }
    """
    Then the response code should be 201
    Then the response body should contain:
    """json
    {
      "author": "root",
      "title": "test-widgetfilter-to-create-1-title",
      "is_private": true,
      "alarm_pattern": [
        [
          {
            "field": "v.component",
            "cond": {
              "type": "eq",
              "value": "test-widgetfilter-to-create-1-pattern"
            }
          }
        ]
      ],
      "entity_pattern": [
        [
          {
            "field": "name",
            "cond": {
              "type": "eq",
              "value": "test-widgetfilter-to-create-1-pattern"
            }
          }
        ]
      ],
      "pbehavior_pattern": [
        [
          {
            "field": "pbehavior_info.type",
            "cond": {
              "type": "eq",
              "value": "test-widgetfilter-to-create-1-pattern"
            }
          }
        ]
      ]
    }
    """
    When I do GET /api/v4/widget-filters/{{ .lastResponse._id }}
    Then the response code should be 200
    Then the response body should contain:
    """json
    {
      "author": "root",
      "title": "test-widgetfilter-to-create-1-title",
      "is_private": true,
      "alarm_pattern": [
        [
          {
            "field": "v.component",
            "cond": {
              "type": "eq",
              "value": "test-widgetfilter-to-create-1-pattern"
            }
          }
        ]
      ],
      "entity_pattern": [
        [
          {
            "field": "name",
            "cond": {
              "type": "eq",
              "value": "test-widgetfilter-to-create-1-pattern"
            }
          }
        ]
      ],
      "pbehavior_pattern": [
        [
          {
            "field": "pbehavior_info.type",
            "cond": {
              "type": "eq",
              "value": "test-widgetfilter-to-create-1-pattern"
            }
          }
        ]
      ]
    }
    """

  Scenario: given create request with corporate patterns should return ok
    When I am admin
    When I do POST /api/v4/widget-filters:
    """json
    {
      "title": "test-widgetfilter-to-create-2-title",
      "widget": "test-widget-to-filter-edit",
      "is_private": false,
      "corporate_alarm_pattern": "test-pattern-to-filter-edit-1",
      "corporate_entity_pattern": "test-pattern-to-filter-edit-2",
      "corporate_pbehavior_pattern": "test-pattern-to-filter-edit-3"
    }
    """
    Then the response code should be 201
    Then the response body should contain:
    """json
    {
      "author": "root",
      "title": "test-widgetfilter-to-create-2-title",
      "is_private": false,
      "corporate_alarm_pattern": "test-pattern-to-filter-edit-1",
      "corporate_alarm_pattern_title": "test-pattern-to-filter-edit-1-title",
      "alarm_pattern": [
        [
          {
            "field": "v.component",
            "cond": {
              "type": "eq",
              "value": "test-pattern-to-filter-edit-1-pattern"
            }
          }
        ]
      ],
      "corporate_entity_pattern": "test-pattern-to-filter-edit-2",
      "corporate_entity_pattern_title": "test-pattern-to-filter-edit-2-title",
      "entity_pattern": [
        [
          {
            "field": "name",
            "cond": {
              "type": "eq",
              "value": "test-pattern-to-filter-edit-2-pattern"
            }
          }
        ]
      ],
      "corporate_pbehavior_pattern": "test-pattern-to-filter-edit-3",
      "corporate_pbehavior_pattern_title": "test-pattern-to-filter-edit-3-title",
      "pbehavior_pattern": [
        [
          {
            "field": "pbehavior_info.type",
            "cond": {
              "type": "eq",
              "value": "test-pattern-to-filter-edit-3-pattern"
            }
          }
        ]
      ]
    }
    """
    When I do GET /api/v4/widget-filters/{{ .lastResponse._id }}
    Then the response code should be 200
    Then the response body should contain:
    """json
    {
      "author": "root",
      "title": "test-widgetfilter-to-create-2-title",
      "is_private": false,
      "corporate_alarm_pattern": "test-pattern-to-filter-edit-1",
      "corporate_alarm_pattern_title": "test-pattern-to-filter-edit-1-title",
      "alarm_pattern": [
        [
          {
            "field": "v.component",
            "cond": {
              "type": "eq",
              "value": "test-pattern-to-filter-edit-1-pattern"
            }
          }
        ]
      ],
      "corporate_entity_pattern": "test-pattern-to-filter-edit-2",
      "corporate_entity_pattern_title": "test-pattern-to-filter-edit-2-title",
      "entity_pattern": [
        [
          {
            "field": "name",
            "cond": {
              "type": "eq",
              "value": "test-pattern-to-filter-edit-2-pattern"
            }
          }
        ]
      ],
      "corporate_pbehavior_pattern": "test-pattern-to-filter-edit-3",
      "corporate_pbehavior_pattern_title": "test-pattern-to-filter-edit-3-title",
      "pbehavior_pattern": [
        [
          {
            "field": "pbehavior_info.type",
            "cond": {
              "type": "eq",
              "value": "test-pattern-to-filter-edit-3-pattern"
            }
          }
        ]
      ]
    }
    """

  Scenario: given create request and no auth user should not allow access
    When I do POST /api/v4/widget-filters
    Then the response code should be 401

  Scenario: given create request and auth user without permissions should not allow access
    When I am noperms
    When I do POST /api/v4/widget-filters
    Then the response code should be 403

  Scenario: given create request and auth user without view permission should not allow access
    When I am admin
    When I do POST /api/v4/widget-filters:
    """json
    {
      "widget": "test-widget-to-filter-check-access",
      "title": "test-widgetfilter-to-create-3-title",
      "is_private": true,
      "alarm_pattern": [
        [
          {
            "field": "v.component",
            "cond": {
              "type": "eq",
              "value": "test-widgetfilter-to-create-5-pattern"
            }
          }
        ]
      ]
    }
    """
    Then the response code should be 403

  Scenario: given create request and with not exist widget should not allow access
    When I am admin
    When I do POST /api/v4/widget-filters:
    """json
    {
      "widget": "test-widget-not-exist",
      "title": "test-widgetfilter-to-create-3-title",
      "is_private": true,
      "alarm_pattern": [
        [
          {
            "field": "v.component",
            "cond": {
              "type": "eq",
              "value": "test-widgetfilter-to-create-5-pattern"
            }
          }
        ]
      ]
    }
    """
    Then the response code should be 403

  Scenario: given create request with missing fields should return bad request error
    When I am admin
    When I do POST /api/v4/widget-filters:
    """json
    {
    }
    """
    Then the response code should be 400
    Then the response body should be:
    """json
    {
      "errors": {
        "title": "Title is missing.",
        "widget": "Widget is missing.",
        "is_private": "IsPrivate is missing.",
        "alarm_pattern": "AlarmPattern is missing.",
        "corporate_alarm_pattern": "CorporateAlarmPattern is missing.",
        "entity_pattern": "EntityPattern is missing.",
        "corporate_entity_pattern": "CorporateEntityPattern is missing.",
        "pbehavior_pattern": "PbehaviorPattern is missing.",
        "corporate_pbehavior_pattern": "CorporatePbehaviorPattern is missing."
      }
    }
    """
    When I do POST /api/v4/widget-filters:
    """json
    {
      "alarm_pattern": [
        []
      ],
      "entity_pattern": [
        []
      ],
      "pbehavior_pattern": [
        []
      ]
    }
    """
    Then the response code should be 400
    Then the response body should contain:
    """json
    {
      "errors": {
        "alarm_pattern": "AlarmPattern is invalid alarm pattern.",
        "entity_pattern": "EntityPattern is invalid entity pattern.",
        "pbehavior_pattern": "PbehaviorPattern is invalid pbehavior pattern."
      }
    }
    """
    When I do POST /api/v4/widget-filters:
    """json
    {
      "alarm_pattern": [
        [
          {
            "field": "v.component",
            "cond": {
              "type": "eq",
              "value": "test-widgetfilter-to-create-3-pattern"
            }
          }
        ]
      ],
      "corporate_alarm_pattern": "test-pattern-to-filter-edit-1",
      "entity_pattern": [
        [
          {
            "field": "name",
            "cond": {
              "type": "eq",
              "value": "test-widgetfilter-to-create-3-pattern"
            }
          }
        ]
      ],
      "corporate_entity_pattern": "test-pattern-to-filter-edit-2",
      "pbehavior_pattern": [
        [
          {
            "field": "type",
            "cond": {
              "type": "eq",
              "value": "test-widgetfilter-to-create-3-pattern"
            }
          }
        ]
      ],
      "corporate_pbehavior_pattern": "test-pattern-to-filter-edit-3"
    }
    """
    Then the response code should be 400
    Then the response body should contain:
    """json
    {
      "errors": {
        "alarm_pattern": "Can't be present both AlarmPattern and CorporateAlarmPattern.",
        "entity_pattern": "Can't be present both EntityPattern and CorporateEntityPattern.",
        "pbehavior_pattern": "Can't be present both PbehaviorPattern and CorporatePbehaviorPattern."
      }
    }
    """
    When I do POST /api/v4/widget-filters:
    """json
    {
      "title": "test-widgetfilter-to-create-3-title",
      "widget": "test-widget-to-filter-edit",
      "is_private": true,
      "corporate_alarm_pattern": "test-pattern-not-exist"
    }
    """
    Then the response code should be 400
    Then the response body should contain:
    """json
    {
      "errors": {
        "corporate_alarm_pattern": "CorporateAlarmPattern doesn't exist."
      }
    }
    """
    When I do POST /api/v4/widget-filters:
    """json
    {
      "title": "test-widgetfilter-to-create-3-title",
      "widget": "test-widget-to-filter-edit",
      "is_private": true,
      "corporate_entity_pattern": "test-pattern-not-exist"
    }
    """
    Then the response code should be 400
    Then the response body should contain:
    """json
    {
      "errors": {
        "corporate_entity_pattern": "CorporateEntityPattern doesn't exist."
      }
    }
    """
    When I do POST /api/v4/widget-filters:
    """json
    {
      "title": "test-widgetfilter-to-create-3-title",
      "widget": "test-widget-to-filter-edit",
      "is_private": true,
      "corporate_pbehavior_pattern": "test-pattern-not-exist"
    }
    """
    Then the response code should be 400
    Then the response body should contain:
    """json
    {
      "errors": {
        "corporate_pbehavior_pattern": "CorporatePbehaviorPattern doesn't exist."
      }
    }
    """
