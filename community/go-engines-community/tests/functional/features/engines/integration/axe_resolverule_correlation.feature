Feature: resolve meta alarm
  I need to be able to resolve meta alarm

  Scenario: given meta alarm and resolved child should auto resolve meta alarm
    Given I am admin
    When I do POST /api/v4/cat/metaalarmrules:
    """
    {
      "name": "test-metaalarmrule-axe-resolverule-correlation-1",
      "type": "attribute",
      "auto_resolve": true,
      "config": {
        "alarm_patterns": [
          {
            "v": {
              "component": "test-component-axe-resolverule-correlation-1"
            }
          }
        ]
      }
    }
    """
    Then the response code should be 201
    Then I save response metaAlarmRuleID={{ .lastResponse._id }}
    When I do POST /api/v4/resolve-rules:
    """json
    {
      "_id": "test-resolve-rule-axe-resolverule-correlation-1",
      "description": "test-resolve-rule-axe-resolverule-correlation-1-desc",
      "entity_patterns":[
        {
          "name": "test-resource-axe-resolverule-correlation-1"
        }
      ],
      "duration": {
        "seconds": 2,
        "unit": "s"
      },
      "priority": 10
    }
    """
    Then the response code should be 201
    When I wait the next periodical process
    When I send an event:
    """
    {
      "connector": "test-connector-axe-resolverule-correlation-1",
      "connector_name": "test-connector-name-axe-resolverule-correlation-1",
      "source_type": "resource",
      "event_type": "check",
      "component":  "test-component-axe-resolverule-correlation-1",
      "resource": "test-resource-axe-resolverule-correlation-1",
      "state": 2,
      "output": "test-output-axe-resolverule-correlation-1"
    }
    """
    When I wait the end of 2 events processing
    When I send an event:
    """
    {
      "connector": "test-connector-axe-resolverule-correlation-1",
      "connector_name": "test-connector-name-axe-resolverule-correlation-1",
      "source_type": "resource",
      "event_type": "check",
      "component":  "test-component-axe-resolverule-correlation-1",
      "resource": "test-resource-axe-resolverule-correlation-1",
      "state": 0,
      "output": "test-output-axe-resolverule-correlation-1"
    }
    """
    When I wait the end of 2 events processing
    When I do GET /api/v4/alarms?filter={"$and":[{"v.resolved":{"$gt":0}},{"v.meta":"{{ .metaAlarmRuleID }}"}]}&with_steps=true&correlation=true
    Then the response code should be 200
    Then the response body should contain:
    """
    {
      "data": [
        {
          "v": {
            "component": "metaalarm",
            "connector": "engine",
            "connector_name": "correlation",
            "state": {
              "val": 0
            },
            "status": {
              "val": 0
            },
            "steps": [
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
            ]
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
