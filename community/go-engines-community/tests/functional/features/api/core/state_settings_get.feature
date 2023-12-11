Feature: get state settings
  I need to be able to get state settings
  Only admin should be able to get state settings

  @concurrent
  Scenario: given get request and no auth user should not allow access
    When I do GET /api/v4/state-settings
    Then the response code should be 401

  @concurrent
  Scenario: given get request and auth user without permissions should not allow access
    When I am noperms
    When I do GET /api/v4/state-settings
    Then the response code should be 403

  @concurrent
  Scenario: given get list request should return ok
    When I am admin
    When I do GET /api/v4/state-settings?search=state-settings-to-get
    Then the response code should be 200
    Then the response body should contain:
    """json
    {
      "data": [
        {
          "_id": "state-settings-to-get-1",
          "method": "inherited",
          "title": "state-settings-to-get-1-title",
          "enabled": true,
          "type": "component",
          "entity_pattern": [
            [
              {
                "field": "name",
                "cond": {
                  "type": "eq",
                  "value": "test-service-state-settings-to-get-rule-pattern-1"
                }
              }
            ]
          ],
          "inherited_entity_pattern": [
            [
              {
                "field": "name",
                "cond": {
                  "type": "eq",
                  "value": "test-resource-state-settings-to-get-impacting-pattern-1"
                }
              }
            ]
          ],
          "editable": true,
          "deletable": true
        },
        {
          "_id": "state-settings-to-get-2",
          "method": "inherited",
          "title": "state-settings-to-get-2-title",
          "enabled": true,
          "type": "component",
          "entity_pattern": [
            [
              {
                "field": "name",
                "cond": {
                  "type": "eq",
                  "value": "test-service-state-settings-to-get-rule-pattern-2"
                }
              }
            ]
          ],
          "inherited_entity_pattern": [
            [
              {
                "field": "name",
                "cond": {
                  "type": "eq",
                  "value": "test-resource-state-settings-to-get-impacting-pattern-2"
                }
              }
            ]
          ],
          "editable": true,
          "deletable": true
        }
      ],
      "meta": {
        "page": 1,
        "per_page": 10,
        "page_count": 1,
        "total_count": 2
      }
    }
    """

  @concurrent
  Scenario: given get list request should return default settings on top regarding of sort
    When I am admin
    When I do GET /api/v4/state-settings?sort_by=priority&sort=asc
    Then the response code should be 200
    Then the response key "data.0._id" should be "service"
    Then the response key "data.1._id" should be "junit"
    When I do GET /api/v4/state-settings?sort_by=priority&sort=desc
    Then the response code should be 200
    Then the response key "data.0._id" should be "service"
    Then the response key "data.1._id" should be "junit"

  @concurrent
  Scenario: given get by id request should return ok
    When I am admin
    When I do GET /api/v4/state-settings/state-settings-to-get-1
    Then the response code should be 200
    Then the response body should contain:
    """json
    {
      "_id": "state-settings-to-get-1",
      "method": "inherited",
      "title": "state-settings-to-get-1-title",
      "enabled": true,
      "type": "component",
      "entity_pattern": [
        [
          {
            "field": "name",
            "cond": {
              "type": "eq",
              "value": "test-service-state-settings-to-get-rule-pattern-1"
            }
          }
        ]
      ],
      "inherited_entity_pattern": [
        [
          {
            "field": "name",
            "cond": {
              "type": "eq",
              "value": "test-resource-state-settings-to-get-impacting-pattern-1"
            }
          }
        ]
      ],
      "editable": true,
      "deletable": true
    }
    """

  @concurrent
  Scenario: given get by id request should return ok
    When I am admin
    When I do GET /api/v4/state-settings/state-settings-to-get-not found
    Then the response code should be 404
    Then the response body should contain:
    """json
    {
      "error": "Not found"
    }
    """

  @concurrent
  Scenario: given get some default settings should return ok and have deletable and editable fields
    When I am admin
    When I do GET /api/v4/state-settings/service
    Then the response code should be 200
    Then the response body should contain:
    """json
    {
      "_id": "service",
      "enabled": true,
      "editable": false,
      "deletable": false
    }
    """
    When I do GET /api/v4/state-settings/junit
    Then the response code should be 200
    Then the response body should contain:
    """json
    {
      "_id": "junit",
      "enabled": true,
      "editable": true,
      "deletable": false
    }
    """
