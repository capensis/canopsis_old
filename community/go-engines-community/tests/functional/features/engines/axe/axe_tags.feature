Feature: create and update alarm by main event stream
  I need to be able to create and update alarm on event

  @concurrent
  Scenario: given check event should create alarm with tags
    Given I am admin
    When I send an event and wait the end of event processing:
    """json
    {
      "event_type": "check",
      "tags": {
        "env": "prod",
        "cloud": "",
        "": "localhost"
      },
      "state": 2,
      "output": "test-output-axe-tags-1",
      "connector": "test-connector-axe-tags-1",
      "connector_name": "test-connector-name-axe-tags-1",
      "component": "test-component-axe-tags-1",
      "resource": "test-resource-axe-tags-1",
      "source_type": "resource"
    }
    """
    When I do GET /api/v4/alarms?search=test-resource-axe-tags-1
    Then the response code should be 200
    Then the response array key "data.0.tags" should contain only:
    """json
    [
      "env: prod",
      "cloud"
    ]
    """

  @concurrent
  Scenario: given check event with another tags should add new tags
    Given I am admin
    When I send an event and wait the end of event processing:
    """json
    {
      "event_type": "check",
      "tags": {
        "env": "prod",
        "cloud": "",
        "": "localhost"
      },
      "state": 2,
      "output": "test-output-axe-tags-2",
      "connector": "test-connector-axe-tags-2",
      "connector_name": "test-connector-name-axe-tags-2",
      "component": "test-component-axe-tags-2",
      "resource": "test-resource-axe-tags-2",
      "source_type": "resource"
    }
    """
    When I send an event and wait the end of event processing:
    """json
    {
      "event_type": "check",
      "tags": {
        "env": "test",
        "website": "",
        "cloud": "",
        "": ""
      },
      "state": 2,
      "output": "test-output-axe-tags-2",
      "connector": "test-connector-axe-tags-2",
      "connector_name": "test-connector-name-axe-tags-2",
      "component": "test-component-axe-tags-2",
      "resource": "test-resource-axe-tags-2",
      "source_type": "resource"
    }
    """
    When I do GET /api/v4/alarms?search=test-resource-axe-tags-2
    Then the response code should be 200
    Then the response array key "data.0.tags" should contain only:
    """json
    [
      "env: prod",
      "cloud",
      "env: test",
      "website"
    ]
    """

  @concurrent
  Scenario: given check event without tags should keep old tags
    Given I am admin
    When I send an event and wait the end of event processing:
    """json
    {
      "event_type": "check",
      "tags": {
        " env ": " prod ",
        "cloud": "  ",
        " ": "localhost"
      },
      "state": 2,
      "output": "test-output-axe-tags-3",
      "connector": "test-connector-axe-tags-3",
      "connector_name": "test-connector-name-axe-tags-3",
      "component": "test-component-axe-tags-3",
      "resource": "test-resource-axe-tags-3",
      "source_type": "resource"
    }
    """
    When I send an event and wait the end of event processing:
    """json
    {
      "event_type": "check",
      "state": 2,
      "output": "test-output-axe-tags-3",
      "connector": "test-connector-axe-tags-3",
      "connector_name": "test-connector-name-axe-tags-3",
      "component": "test-component-axe-tags-3",
      "resource": "test-resource-axe-tags-3",
      "source_type": "resource"
    }
    """
    When I do GET /api/v4/alarms?search=test-resource-axe-tags-3
    Then the response code should be 200
    Then the response array key "data.0.tags" should contain only:
    """json
    [
      "env: prod",
      "cloud"
    ]
    """

  @concurrent
  Scenario: given check event should create alarm tags
    Given I am admin
    When I send an event and wait the end of event processing:
    """json
    {
      "event_type": "check",
      "tags": {
        "test-tag-axe-tags-4-1": " prod ",
        " test-tag-axe-tags-4-2 ": " "
      },
      "state": 2,
      "output": "test-output-axe-tags-4",
      "connector": "test-connector-axe-tags-4",
      "connector_name": "test-connector-name-axe-tags-4",
      "component": "test-component-axe-tags-4",
      "resource": "test-resource-axe-tags-4",
      "source_type": "resource"
    }
    """
    When I do GET /api/v4/alarm-tags?search=test-tag-axe-tags-4 until response code is 200 and body contains:
    """json
    {
      "data": [
        {
          "value": "test-tag-axe-tags-4-1: prod"
        },
        {
          "value": "test-tag-axe-tags-4-2"
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
    When I save response sameLabelColor={{ (index .lastResponse.data 0).color }}
    When I send an event and wait the end of event processing:
    """json
    {
      "event_type": "check",
      "tags": {
        "test-tag-axe-tags-4-1": " test ",
        " test-tag-axe-tags-4-3 ": " "
      },
      "state": 2,
      "output": "test-output-axe-tags-4",
      "connector": "test-connector-axe-tags-4",
      "connector_name": "test-connector-name-axe-tags-4",
      "component": "test-component-axe-tags-4",
      "resource": "test-resource-axe-tags-4",
      "source_type": "resource"
    }
    """
    When I do GET /api/v4/alarm-tags?search=test-tag-axe-tags-4 until response code is 200 and body contains:
    """json
    {
      "data": [
        {
          "value": "test-tag-axe-tags-4-1: prod"
        },
        {
          "value": "test-tag-axe-tags-4-1: test",
          "color": "{{ .sameLabelColor }}"
        },
        {
          "value": "test-tag-axe-tags-4-2"
        },
        {
          "value": "test-tag-axe-tags-4-3"
        }
      ],
      "meta": {
        "page": 1,
        "page_count": 1,
        "per_page": 10,
        "total_count": 4
      }
    }
    """
