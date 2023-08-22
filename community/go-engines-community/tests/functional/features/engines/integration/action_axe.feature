Feature: update alarm on action
  I need to be able to update alarm on action.

  Scenario: given alarm and scenario with resolve trigger should call webhook on resolve
    Given I am admin
    When I do POST /api/v4/scenarios:
    """json
    {
      "name": "test-scenario-action-axe-1-name",
      "priority": 10040,
      "enabled": true,
      "triggers": ["resolve"],
      "actions": [
        {
          "entity_pattern": [
            [
              {
                "field": "name",
                "cond": {
                  "type": "eq",
                  "value": "test-resource-action-axe-1"
                }
              }
            ]
          ],
          "type": "webhook",
          "parameters": {
            "request": {
              "method": "POST",
              "url": "{{ .apiURL }}/api/v4/scenarios",
              "auth": {
                "username": "root",
                "password": "test"
              },
              "headers": {"Content-Type": "application/json"},
              "payload": "{\"priority\": 10041,\"name\":\"{{ `{{ .Entity.ID }}` }}\",\"enabled\":true,\"triggers\":[\"create\"],\"actions\":[{\"entity_pattern\":[[{\"field\":\"name\",\"cond\":{\"type\": \"eq\", \"value\": \"test-scenario-action-axe-1-alarm\"}}]],\"type\":\"ack\",\"drop_scenario_if_not_matched\":false,\"emit_trigger\":false}]}"
            }
          },
          "drop_scenario_if_not_matched": false,
          "emit_trigger": false
        }
      ]
    }
    """
    Then the response code should be 201
    When I wait the next periodical process
    When I send an event:
    """json
    {
      "connector" : "test-connector-action-axe-1",
      "connector_name" : "test-connector-name-action-axe-1",
      "source_type" : "resource",
      "event_type" : "check",
      "component" :  "test-component-action-axe-1",
      "resource" : "test-resource-action-axe-1",
      "state" : 2,
      "output" : "noveo alarm"
    }
    """
    When I wait the end of event processing
    When I send an event:
    """json
    {
      "connector" : "test-connector-action-axe-1",
      "connector_name" : "test-connector-name-action-axe-1",
      "source_type" : "resource",
      "event_type" : "check",
      "component" :  "test-component-action-axe-1",
      "resource" : "test-resource-action-axe-1",
      "state" : 0,
      "output" : "noveo alarm"
    }
    """
    When I wait the end of event processing
    When I send an event:
    """json
    {
      "connector" : "test-connector-action-axe-1",
      "connector_name" : "test-connector-name-action-axe-1",
      "source_type" : "resource",
      "event_type" : "resolve_close",
      "component" :  "test-component-action-axe-1",
      "resource" : "test-resource-action-axe-1",
      "output" : "noveo alarm"
    }
    """
    When I wait the end of event processing
    When I do GET /api/v4/scenarios?search=test-resource-action-axe-1
    Then the response code should be 200
    Then the response body should contain:
    """json
    {
      "data": [
        {
          "name": "test-resource-action-axe-1/test-component-action-axe-1"
        }
      ],
      "meta": {
        "total_count": 1
      }
    }
    """

  Scenario: given alarm and scenario with resolve trigger should not update alarm
    Given I am admin
    When I do POST /api/v4/scenarios:
    """json
    {
      "name": "test-scenario-action-axe-2-name",
      "priority": 10042,
      "enabled": true,
      "triggers": ["resolve"],
      "actions": [
        {
          "entity_pattern": [
            [
              {
                "field": "name",
                "cond": {
                  "type": "eq",
                  "value": "test-resource-action-axe-2"
                }
              }
            ]
          ],
          "type": "ack",
          "parameters": {
            "output": "test-output-action-axe-2"
          },
          "drop_scenario_if_not_matched": false,
          "emit_trigger": false
        }
      ]
    }
    """
    Then the response code should be 201
    When I wait the next periodical process
    When I send an event:
    """json
    {
      "connector" : "test-connector-action-axe-2",
      "connector_name" : "test-connector-name-action-axe-2",
      "source_type" : "resource",
      "event_type" : "check",
      "component" :  "test-component-action-axe-2",
      "resource" : "test-resource-action-axe-2",
      "state" : 2,
      "output" : "noveo alarm"
    }
    """
    When I wait the end of event processing
    When I send an event:
    """json
    {
      "connector" : "test-connector-action-axe-2",
      "connector_name" : "test-connector-name-action-axe-2",
      "source_type" : "resource",
      "event_type" : "check",
      "component" :  "test-component-action-axe-2",
      "resource" : "test-resource-action-axe-2",
      "state" : 0,
      "output" : "noveo alarm"
    }
    """
    When I wait the end of event processing
    When I send an event:
    """json
    {
      "connector" : "test-connector-action-axe-2",
      "connector_name" : "test-connector-name-action-axe-2",
      "source_type" : "resource",
      "event_type" : "resolve_close",
      "component" :  "test-component-action-axe-2",
      "resource" : "test-resource-action-axe-2",
      "output" : "noveo alarm"
    }
    """
    When I wait the end of event processing
    When I do GET /api/v4/alarms?search=test-resource-action-axe-2
    Then the response code should be 200
    Then the response body should contain:
    """json
    {
      "data": [
        {
          "v": {
            "connector": "test-connector-action-axe-2",
            "connector_name": "test-connector-name-action-axe-2",
            "component": "test-component-action-axe-2",
            "resource": "test-resource-action-axe-2"
          }
        }
      ],
      "meta": {
        "page": 1,
        "page_count": 1,
        "per_page": 10,
        "total_count": 1
      }
    }
    """
    When I do POST /api/v4/alarm-details:
    """json
    [
      {
        "_id": "{{ (index .lastResponse.data 0)._id }}",
        "steps": {
          "page": 1
        }
      }
    ]
    """
    Then the response code should be 207
    Then the response body should contain:
    """json
    [
      {
        "status": 200,
        "data": {
          "steps": {
            "data": [
              {
                "_t": "stateinc",
                "val": 2
              },
              {
                "_t": "statusinc",
                "val": 1
              },
              {
                "_t": "statedec",
                "val": 0
              },
              {
                "_t": "statusdec",
                "val": 0
              }
            ],
            "meta": {
              "page": 1,
              "page_count": 1,
              "per_page": 10,
              "total_count": 4
            }
          }
        }
      }
    ]
    """

  Scenario: given alarm and scenario and widget filter should filter alarms by ticket message or ticket's ticket
    Given I am admin
    When I do POST /api/v4/scenarios:
    """json
    {
      "name": "test-scenario-action-axe-3-name-1",
      "priority": 1692721198,
      "enabled": true,
      "triggers": ["create"],
      "actions": [
        {
          "entity_pattern": [
            [
              {
                "field": "name",
                "cond": {
                  "type": "eq",
                  "value": "test-resource-action-axe-3-1"
                }
              }
            ]
          ],
          "type": "webhook",
          "parameters": {
            "request": {
              "method": "POST",
              "url": "{{ .apiURL }}/api/v4/scenarios",
              "auth": {
                "username": "root",
                "password": "test"
              },
              "headers": {"Content-Type": "application/json"},
              "payload": "{\"priority\":1692721199, \"name\":\"{{ `{{ .Alarm.Value.Output }}` }}\",\"enabled\":true,\"triggers\":[\"create\"],\"actions\":[{\"entity_pattern\":[[{\"field\":\"name\",\"cond\":{\"type\": \"eq\", \"value\": \"test-scenario-action-axe-3-alarm\"}}]],\"type\":\"ack\",\"drop_scenario_if_not_matched\":false,\"emit_trigger\":false}]}"
            },
            "declare_ticket": {
              "ticket_id": "name"
            }
          },
          "drop_scenario_if_not_matched": false,
          "emit_trigger": false
        }
      ]
    }
    """
    Then the response code should be 201
    When I do POST /api/v4/scenarios:
    """json
    {
      "name": "test-scenario-action-axe-3-name-2",
      "priority": 1692721200,
      "enabled": true,
      "triggers": ["create"],
      "actions": [
        {
          "entity_pattern": [
            [
              {
                "field": "name",
                "cond": {
                  "type": "eq",
                  "value": "test-resource-action-axe-3-2"
                }
              }
            ]
          ],
          "type": "webhook",
          "parameters": {
            "request": {
              "method": "POST",
              "url": "{{ .apiURL }}/api/v4/scenarios",
              "auth": {
                "username": "root",
                "password": "test"
              },
              "headers": {"Content-Type": "application/json"},
              "payload": "{\"priority\":1692721201, \"name\":\"{{ `{{ .Alarm.Value.Output }}` }}\",\"enabled\":true,\"triggers\":[\"create\"],\"actions\":[{\"entity_pattern\":[[{\"field\":\"name\",\"cond\":{\"type\": \"eq\", \"value\": \"test-scenario-action-axe-3-alarm\"}}]],\"type\":\"ack\",\"drop_scenario_if_not_matched\":false,\"emit_trigger\":false}]}"
            },
            "declare_ticket": {
              "ticket_id": "name"
            }
          },
          "drop_scenario_if_not_matched": false,
          "emit_trigger": false
        }
      ]
    }
    """
    Then the response code should be 201
    When I do POST /api/v4/scenarios:
    """json
    {
      "name": "test-scenario-action-axe-3-name-3",
      "priority": 1692721202,
      "enabled": true,
      "triggers": ["create"],
      "actions": [
        {
          "entity_pattern": [
            [
              {
                "field": "name",
                "cond": {
                  "type": "eq",
                  "value": "test-resource-action-axe-3-3"
                }
              }
            ]
          ],
          "type": "webhook",
          "parameters": {
            "request": {
              "method": "POST",
              "url": "{{ .apiURL }}/api/v4/scenarios",
              "auth": {
                "username": "root",
                "password": "test"
              },
              "headers": {"Content-Type": "application/json"},
              "payload": "{\"priority\":1692721203, \"name\":\"{{ `{{ .Alarm.Value.Output }}` }}\",\"enabled\":true,\"triggers\":[\"create\"],\"actions\":[{\"entity_pattern\":[[{\"field\":\"name\",\"cond\":{\"type\": \"eq\", \"value\": \"test-scenario-action-axe-3-alarm\"}}]],\"type\":\"ack\",\"drop_scenario_if_not_matched\":false,\"emit_trigger\":false}]}"
            },
            "declare_ticket": {
              "ticket_id": "name"
            }
          },
          "drop_scenario_if_not_matched": false,
          "emit_trigger": false
        }
      ]
    }
    """
    Then the response code should be 201
    When I wait the next periodical process
    When I send an event:
    """json
    {
      "connector": "test-connector-action-axe-3",
      "connector_name": "test-connector-name-action-axe-3",
      "source_type": "resource",
      "event_type": "check",
      "component":  "test-component-action-axe-3",
      "resource": "test-resource-action-axe-3-1",
      "state": 2,
      "output": "test-1"
    }
    """
    When I wait the end of event processing
    When I send an event:
    """json
    {
      "connector": "test-connector-action-axe-3",
      "connector_name": "test-connector-name-action-axe-3",
      "source_type": "resource",
      "event_type": "check",
      "component":  "test-component-action-axe-3",
      "resource": "test-resource-action-axe-3-2",
      "state": 2,
      "output": "test-2"
    }
    """
    When I wait the end of event processing
    When I send an event:
    """json
    {
      "connector": "test-connector-action-axe-3",
      "connector_name": "test-connector-name-action-axe-3",
      "source_type": "resource",
      "event_type": "check",
      "component":  "test-component-action-axe-3",
      "resource": "test-resource-action-axe-3-3",
      "state": 2,
      "output": "test-3"
    }
    """
    When I wait the end of event processing
    When I do POST /api/v4/widget-filters:
    """json
    {
      "title": "test-widgetfilter-action-axe-3-1",
      "widget": "test-widget-to-alarm-get",
      "is_private": true,
      "alarm_pattern": [
        [
          {
            "field": "v.ticket.ticket",
            "cond": {
              "type": "eq",
              "value": "test-2"
            }
          }
        ]
      ]
    }
    """
    Then the response code should be 201
    When I do GET /api/v4/alarms?filters[]={{ .lastResponse._id }}&sort=asc&sort_by=v.resource
    Then the response code should be 200
    Then the response body should contain:
    """json
    {
      "data": [
        {
          "v": {
            "resource": "test-resource-action-axe-3-2"
          }
        }
      ],
      "meta": {
        "total_count": 1
      }
    }
    """
    When I do POST /api/v4/widget-filters:
    """json
    {
      "title": "test-widgetfilter-action-axe-3-2",
      "widget": "test-widget-to-alarm-get",
      "is_private": true,
      "alarm_pattern": [
        [
          {
            "field": "v.ticket.ticket",
            "cond": {
              "type": "regexp",
              "value": "test"
            }
          }
        ]
      ]
    }
    """
    Then the response code should be 201
    When I do GET /api/v4/alarms?filters[]={{ .lastResponse._id }}&sort=asc&sort_by=v.resource
    Then the response code should be 200
    Then the response body should contain:
    """json
    {
      "data": [
        {
          "v": {
            "resource": "test-resource-action-axe-3-1"
          }
        },
        {
          "v": {
            "resource": "test-resource-action-axe-3-2"
          }
        },
        {
          "v": {
            "resource": "test-resource-action-axe-3-3"
          }
        }
      ],
      "meta": {
        "total_count": 3
      }
    }
    """
    When I do POST /api/v4/widget-filters:
    """json
    {
      "title": "test-widgetfilter-action-axe-3-3",
      "widget": "test-widget-to-alarm-get",
      "is_private": true,
      "alarm_pattern": [
        [
          {
            "field": "v.ticket.m",
            "cond": {
              "type": "regexp",
              "value": "test-2"
            }
          }
        ]
      ]
    }
    """
    Then the response code should be 201
    When I do GET /api/v4/alarms?filters[]={{ .lastResponse._id }}&sort=asc&sort_by=v.resource
    Then the response code should be 200
    Then the response body should contain:
    """json
    {
      "data": [
        {
          "v": {
            "resource": "test-resource-action-axe-3-2"
          }
        }
      ],
      "meta": {
        "total_count": 1
      }
    }
    """
    When I do POST /api/v4/widget-filters:
    """json
    {
      "title": "test-widgetfilter-action-axe-3-4",
      "widget": "test-widget-to-alarm-get",
      "is_private": true,
      "alarm_pattern": [
        [
          {
            "field": "v.ticket.m",
            "cond": {
              "type": "regexp",
              "value": "test"
            }
          }
        ]
      ]
    }
    """
    Then the response code should be 201
    When I do GET /api/v4/alarms?filters[]={{ .lastResponse._id }}&sort=asc&sort_by=v.resource
    Then the response code should be 200
    Then the response body should contain:
    """json
    {
      "data": [
        {
          "v": {
            "resource": "test-resource-action-axe-3-1"
          }
        },
        {
          "v": {
            "resource": "test-resource-action-axe-3-2"
          }
        },
        {
          "v": {
            "resource": "test-resource-action-axe-3-3"
          }
        }
      ],
      "meta": {
        "total_count": 3
      }
    }
    """
