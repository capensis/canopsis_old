Feature: modify event on event filter
  I need to be able to modify event on event filter

  @concurrent
  Scenario: given check event and drop event filter rule should update events count
    Given I am admin
    When I do POST /api/v4/eventfilter/rules:
    """json
    {
      "type": "drop",
      "description": "test-event-filter-che-event-filters-third-1-description",
      "enabled": true,
      "event_pattern": [
        [
          {
            "field": "component",
            "cond": {
              "type": "eq",
              "value": "test-component-che-event-filters-third-1"
            }
          },
          {
            "field": "event_type",
            "cond": {
              "type": "eq",
              "value": "check"
            }
          }
        ]
      ]
    }
    """
    Then the response code should be 201
    When I wait the next periodical process
    When I send an event and wait the end of event processing:
    """json
    [
      {
        "connector": "test-connector-che-event-filters-third-1",
        "connector_name": "test-connector-name-che-event-filters-third-1",
        "source_type": "resource",
        "event_type": "check",
        "component": "test-component-che-event-filters-third-1",
        "resource": "test-resource-che-event-filters-third-1-1",
        "state": 2
      },
      {
        "connector": "test-connector-che-event-filters-third-1",
        "connector_name": "test-connector-name-che-event-filters-third-1",
        "source_type": "resource",
        "event_type": "check",
        "component": "test-component-che-event-filters-third-1",
        "resource": "test-resource-che-event-filters-third-1-2",
        "state": 1
      }
    ]
    """
    When I do GET /api/v4/eventfilter/rules?search=che-event-filters-third-1 until response code is 200 and body contains:
    """json
    {
      "data": [
        {
          "description": "test-event-filter-che-event-filters-third-1-description",
          "events_count": 2
        }
      ]
    }
    """

  @concurrent
  Scenario: given check event and break event filter rule should update events count
    Given I am admin
    When I do POST /api/v4/eventfilter/rules:
    """json
    {
      "type": "break",
      "description": "test-event-filter-che-event-filters-third-2-description",
      "enabled": true,
      "event_pattern": [
        [
          {
            "field": "component",
            "cond": {
              "type": "eq",
              "value": "test-component-che-event-filters-third-2"
            }
          },
          {
            "field": "event_type",
            "cond": {
              "type": "eq",
              "value": "check"
            }
          }
        ]
      ]
    }
    """
    Then the response code should be 201
    When I wait the next periodical process
    When I send an event and wait the end of event processing:
    """json
    [
      {
        "connector": "test-connector-che-event-filters-third-2",
        "connector_name": "test-connector-name-che-event-filters-third-2",
        "source_type": "resource",
        "event_type": "check",
        "component": "test-component-che-event-filters-third-2",
        "resource": "test-resource-che-event-filters-third-2-1",
        "state": 2
      },
      {
        "connector": "test-connector-che-event-filters-third-2",
        "connector_name": "test-connector-name-che-event-filters-third-2",
        "source_type": "resource",
        "event_type": "check",
        "component": "test-component-che-event-filters-third-2",
        "resource": "test-resource-che-event-filters-third-2-2",
        "state": 1
      }
    ]
    """
    When I do GET /api/v4/eventfilter/rules?search=che-event-filters-third-2 until response code is 200 and body contains:
    """json
    {
      "data": [
        {
          "description": "test-event-filter-che-event-filters-third-2-description",
          "events_count": 2
        }
      ]
    }
    """

  @concurrent
  Scenario: given check event and enrichment event filter rule should update events count
    Given I am admin
    When I do POST /api/v4/eventfilter/rules:
    """json
    {
      "type": "enrichment",
      "description": "test-event-filter-che-event-filters-third-3-description",
      "enabled": true,
      "event_pattern": [
        [
          {
            "field": "component",
            "cond": {
              "type": "eq",
              "value": "test-component-che-event-filters-third-3"
            }
          },
          {
            "field": "event_type",
            "cond": {
              "type": "eq",
              "value": "check"
            }
          }
        ]
      ],
      "config": {
        "actions": [
          {
            "type": "set_field_from_template",
            "name": "customer",
            "description": "Customer",
            "value": "{{ `{{ .Event.ExtraInfos.customer }}` }}"
          }
        ],
        "on_success": "pass",
        "on_failure": "pass"
      }
    }
    """
    Then the response code should be 201
    When I wait the next periodical process
    When I send an event and wait the end of event processing:
    """json
    [
      {
        "connector": "test-connector-che-event-filters-third-3",
        "connector_name": "test-connector-name-che-event-filters-third-3",
        "source_type": "resource",
        "event_type": "check",
        "component": "test-component-che-event-filters-third-3",
        "resource": "test-resource-che-event-filters-third-3-1",
        "customer": "test-customer-che-event-filters-third-3-1",
        "state": 2
      },
      {
        "connector": "test-connector-che-event-filters-third-3",
        "connector_name": "test-connector-name-che-event-filters-third-3",
        "source_type": "resource",
        "event_type": "check",
        "component": "test-component-che-event-filters-third-3",
        "resource": "test-resource-che-event-filters-third-3-2",
        "customer": "test-customer-che-event-filters-third-3-2",
        "state": 1
      }
    ]
    """
    When I do GET /api/v4/eventfilter/rules?search=che-event-filters-third-3 until response code is 200 and body contains:
    """json
    {
      "data": [
        {
          "description": "test-event-filter-che-event-filters-third-3-description",
          "events_count": 2
        }
      ]
    }
    """

  @concurrent
  Scenario: given check event and enrichment event filter rule should update zero events count after update
    Given I am admin
    When I do POST /api/v4/eventfilter/rules:
    """json
    {
      "type": "enrichment",
      "description": "test-event-filter-che-event-filters-third-4-description",
      "enabled": true,
      "event_pattern": [
        [
          {
            "field": "component",
            "cond": {
              "type": "eq",
              "value": "test-component-che-event-filters-third-4"
            }
          },
          {
            "field": "event_type",
            "cond": {
              "type": "eq",
              "value": "check"
            }
          }
        ]
      ],
      "config": {
        "actions": [
          {
            "type": "set_field_from_template",
            "name": "customer",
            "description": "Customer",
            "value": "{{ `{{ .Event.ExtraInfos.customer }}` }}"
          }
        ],
        "on_success": "pass",
        "on_failure": "pass"
      }
    }
    """
    Then the response code should be 201
    Then I save response ruleId={{ .lastResponse._id }}
    When I wait the next periodical process
    When I send an event and wait the end of event processing:
    """json
    [
      {
        "connector": "test-connector-che-event-filters-third-4",
        "connector_name": "test-connector-name-che-event-filters-third-4",
        "source_type": "resource",
        "event_type": "check",
        "component": "test-component-che-event-filters-third-4",
        "resource": "test-resource-che-event-filters-third-4-1",
        "customer": "test-customer-che-event-filters-third-4-1",
        "state": 2
      },
      {
        "connector": "test-connector-che-event-filters-third-4",
        "connector_name": "test-connector-name-che-event-filters-third-4",
        "source_type": "resource",
        "event_type": "check",
        "component": "test-component-che-event-filters-third-4",
        "resource": "test-resource-che-event-filters-third-4-2",
        "customer": "test-customer-che-event-filters-third-4-2",
        "state": 1
      }
    ]
    """
    When I do GET /api/v4/eventfilter/rules?search=che-event-filters-third-4 until response code is 200 and body contains:
    """json
    {
      "data": [
        {
          "description": "test-event-filter-che-event-filters-third-4-description",
          "events_count": 2
        }
      ]
    }
    """
    When I do PUT /api/v4/eventfilter/rules/{{ .ruleId }}:
    """json
    {
      "type": "enrichment",
      "description": "test-event-filter-che-event-filters-third-4-description",
      "enabled": true,
      "event_pattern": [
        [
          {
            "field": "component",
            "cond": {
              "type": "eq",
              "value": "test-component-che-event-filters-third-4"
            }
          },
          {
            "field": "event_type",
            "cond": {
              "type": "eq",
              "value": "check"
            }
          }
        ]
      ],
      "config": {
        "actions": [
          {
            "type": "set_field_from_template",
            "name": "customer",
            "description": "Customer",
            "value": "{{ `{{ .Event.ExtraInfos.customer }}` }}"
          },
          {
            "type": "set_field_from_template",
            "name": "domain",
            "description": "Domain",
            "value": "{{ `{{ .Event.ExtraInfos.domain }}` }}"
          }
        ],
        "on_success": "pass",
        "on_failure": "pass"
      }
    }
    """
    Then the response code should be 200
    When I do GET /api/v4/eventfilter/rules?search=che-event-filters-third-4
    Then the response code should be 200
    Then the response body should contain:
    """json
    {
      "data": [
        {
          "description": "test-event-filter-che-event-filters-third-4-description",
          "events_count": 0
        }
      ]
    }
    """

  @concurrent
  Scenario: given check event and enrichment event filter rule with string slice should not update entity
    Given I am admin
    When I do POST /api/v4/eventfilter/rules:
    """json
    {
      "type": "enrichment",
      "description": "test-event-filter-che-event-filters-third-5-description",
      "enabled": true,
      "event_pattern": [
        [
          {
            "field": "component",
            "cond": {
              "type": "eq",
              "value": "test-component-che-event-filters-third-5"
            }
          },
          {
            "field": "event_type",
            "cond": {
              "type": "eq",
              "value": "check"
            }
          }
        ]
      ],
      "config": {
        "actions": [
          {
            "type": "copy_to_entity_info",
            "name": "hostgroups",
            "description": "Hostgroups",
            "value": "Event.ExtraInfos.hostgroups"
          }
        ],
        "on_success": "pass",
        "on_failure": "pass"
      }
    }
    """
    Then the response code should be 201
    When I wait the next periodical process
    When I send an event and wait the end of event processing:
    """json
    {
      "event_type": "check",
      "state": 0,
      "hostgroups": [
        "test-hostgroup-che-event-filters-third-5-3",
        "test-hostgroup-che-event-filters-third-5-1",
        "test-hostgroup-che-event-filters-third-5-2"
      ],
      "connector": "test-connector-che-event-filters-third-5",
      "connector_name": "test-connector-name-che-event-filters-third-5",
      "component": "test-component-che-event-filters-third-5",
      "resource": "test-resource-che-event-filters-third-5",
      "source_type": "resource"
    }
    """
    When I do GET /api/v4/entities?search=test-resource-che-event-filters-third-5
    Then the response code should be 200
    Then the response body should contain:
    """json
    {
      "data": [
        {
          "_id": "test-resource-che-event-filters-third-5/test-component-che-event-filters-third-5",
          "infos": {
            "hostgroups": {
              "description": "Hostgroups",
              "name": "hostgroups",
              "value": [
                "test-hostgroup-che-event-filters-third-5-1",
                "test-hostgroup-che-event-filters-third-5-2",
                "test-hostgroup-che-event-filters-third-5-3"
              ]
            }
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
    When I send an event and wait the end of event processing:
    """json
    {
      "event_type": "check",
      "state": 0,
      "hostgroups": [
        "test-hostgroup-che-event-filters-third-5-2",
        "test-hostgroup-che-event-filters-third-5-3",
        "test-hostgroup-che-event-filters-third-5-1"
      ],
      "connector": "test-connector-che-event-filters-third-5",
      "connector_name": "test-connector-name-che-event-filters-third-5",
      "component": "test-component-che-event-filters-third-5",
      "resource": "test-resource-che-event-filters-third-5",
      "source_type": "resource"
    }
    """
    When I do GET /api/v4/entities?search=test-resource-che-event-filters-third-5
    Then the response code should be 200
    Then the response body should contain:
    """json
    {
      "data": [
        {
          "_id": "test-resource-che-event-filters-third-5/test-component-che-event-filters-third-5",
          "infos": {
            "hostgroups": {
              "description": "Hostgroups",
              "name": "hostgroups",
              "value": [
                "test-hostgroup-che-event-filters-third-5-1",
                "test-hostgroup-che-event-filters-third-5-2",
                "test-hostgroup-che-event-filters-third-5-3"
              ]
            }
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

  @concurrent
  Scenario: given check event and enrichment event filter rule should copy newtags to tags in event
    Given I am admin
    When I do POST /api/v4/eventfilter/rules:
    """json
    {
      "type": "enrichment",
      "description": "test-event-filter-che-event-filters-third-6-description",
      "enabled": true,
      "event_pattern": [
        [
          {
            "field": "resource",
            "cond": {
              "type": "eq",
              "value": "test-resource-che-event-filters-third-6"
            }
          },
          {
            "field": "event_type",
            "cond": {
              "type": "eq",
              "value": "check"
            }
          }
        ]
      ],
      "config": {
        "actions": [
          {
            "type": "copy",
            "name" : "Tags",
            "description" : "Copy newtags as tags",
            "value" : "Event.ExtraInfos.newtags"
          }
        ],
        "on_success": "pass",
        "on_failure": "pass"
      }
    }
    """
    Then the response code should be 201
    Then I save response ruleId={{ .lastResponse._id }}
    When I wait the next periodical process
    When I send an event and wait the end of event processing:
    """json
    [
      {
        "connector": "test-connector-che-event-filters-third-6",
        "connector_name": "test-connector-name-che-event-filters-third-6",
        "source_type": "resource",
        "event_type": "check",
        "component": "test-component-che-event-filters-third-6",
        "resource": "test-resource-che-event-filters-third-6",
        "newtags" : {
          "RAM" : "",
          "Location": "location-che-event-filters-third-6",
          "Env": "env-che-event-filters-third-6",
          "Catégorie" : "cat-che-event-filters-third-6"
        },
        "state": 3
      }
    ]
    """
    When I do GET /api/v4/alarms?search=che-event-filters-third-6
    Then the response code should be 200
    Then the response array key "data.0.tags" should contain only:
    """json
    [
      "Catégorie: cat-che-event-filters-third-6",
      "Env: env-che-event-filters-third-6",
      "Location: location-che-event-filters-third-6",
      "RAM"
    ]
    """
