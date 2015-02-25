# -*- coding: utf-8 -*-
# --------------------------------
# Copyright (c) 2015 "Capensis" [http://www.capensis.com]
#
# This file is part of Canopsis.
#
# Canopsis is free software: you can redistribute it and/or modify
# it under the terms of the GNU Affero General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# Canopsis is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU Affero General Public License for more details.
#
# You should have received a copy of the GNU Affero General Public License
# along with Canopsis.  If not, see <http://www.gnu.org/licenses/>.
# ---------------------------------

"""Module in charge of defining topology processing in engines.

When an event occured, the related entity is retrieved and all bound
topological nodes are retrieved as well in order to execute theirs rules.

First, a topology processing is triggered when an event occured.

From this event, bound topology nodes are retrieved in order to apply node
    rules.

A typical topological task condition is an ``canopsis.task.condition.all``
composed of the ``canopsis.topology.rule.condition.new_state`` condition and
``canopsis.topology.rule.action.change_state`` action.
If this condition is checked, then other specific conditions can be applied
such as those defined in the canopsis.topology.rule.action module.
"""

from canopsis.topology.elements import Topology, TopoNode
from canopsis.topology.manager import TopologyManager
from canopsis.context.manager import Context
from canopsis.task import register_task
from canopsis.event import Event
from canopsis.check.manager import CheckManager

context = Context()
tm = TopologyManager()
_check = CheckManager()

SOURCE = 'source'
PUBLISHER = 'publisher'


@register_task
def event_processing(engine, event, manager=None, **kwargs):
    """Process input event in getting topology nodes bound to input event
    entity.

    One topology nodes are founded, executing related rules.

    :param dict event: event to process.
    :param Engine engine: engine which consumes the event.
    :param TopologyManager manager: topology manager to use.
    """

    if manager is None:
        manager = tm

    event_type = event[Event.TYPE]

    # apply processing only in case of check event
    if event_type in _check.types:
        # get source type
        source_type = event[Event.SOURCE_TYPE]
        # in case of topology node
        if source_type in [TopoNode.TYPE, Topology.TYPE]:
            # process all targets
            elt = manager.get_elts(ids=event[Topology.ID])
            if elt is not None:
                targets_by_edge = manager.get_targets(
                    ids=elt.id, add_edges=True
                )
                # get elt state in order to update all edges
                state = elt.state
                for edge_id in targets_by_edge:
                    edge, targets = targets_by_edge[edge_id]
                    # update edge state
                    edge.state = state
                    edge.save(manager=manager)
                    # process and save all targets
                    for target in targets:
                        target.process(
                            event=event, engine=engine, manager=manager,
                            **kwargs
                        )

        else:  # in case of entity event
            # get elts from entity
            entity = context.get_entity(event)
            if entity is not None:
                entity_id = context.get_entity_id(entity)
                elts = manager.get_elts(info={TopoNode.ENTITY: entity_id})
                # process all elts
                for elt in elts:
                    elt.process(
                        event=event, engine=engine, manager=manager, **kwargs
                    )

    return event
