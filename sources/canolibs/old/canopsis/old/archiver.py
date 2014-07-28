#!/usr/bin/env python
# -*- coding: utf-8 -*-
#--------------------------------
# Copyright (c) 2014 "Capensis" [http://www.capensis.com]
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

from logging import ERROR, getLogger
from time import time

from canopsis.old.storage import get_storage
from canopsis.old.account import Account
from canopsis.old.record import Record
from canopsis.old.tools import legend, uniq

legend_type = ['soft', 'hard']


class Archiver(object):
    def __init__(self, namespace, storage=None,
                 autolog=False, logging_level=ERROR):

        self.logger = getLogger('Archiver')
        self.logger.setLevel(logging_level)

        self.namespace = namespace
        self.namespace_log = namespace + '_log'

        self.autolog = autolog

        self.logger.debug("Init Archiver on %s" % namespace)

        self.account = Account(user="root", group="root")

        if not storage:
            self.logger.debug(" + Get storage")
            self.storage = get_storage(namespace=namespace,
                                       logging_level=logging_level)
        else:
            self.storage = storage

        self.collection = self.storage.get_backend(namespace)



    def beat(self):
        #Default useless values avoid crashes
        self.bagot_freq = 1
        self.bagot_time = 1
        self.furtif_time = 1
        self.restore_event = True

        self.state_config = self.storage.find({'Record_type':'state-spec'})
        if len(self.state_config) == 1:
            self.state_config = self.state_config[0]

            if 'bagot' in self.state_config:
                if 'freq' in self.state_config['bagot']:
                    self.bagot_freq = self.state_config['bagot']['freq']
                if 'time' in self.state_config['bagot']:
                    self.bagot_time = self.state_config['bagot']['time']
            if 'stealthy_time' in self.state_config:
                self.furtif_time = self.state_config['stealthy_time']
            if 'restore_event' in self.state_config:
                self.restore_event = self.state_config['restore_event']



    def check_bagot(self, event, devent):
        ts_curr = event['timestamp']
        ts_prev = devent['timestamp']
        ts_last_furtif = event['last_furtif']
        ts_last_bagot = (ts_curr - ts_last_furtif)
        freq = event['furtif_freq']

        log = 'Status is set to {} for event %s' % event['rk']

        event['last_furtif'] = event['timestamp']
        event['furtif_freq'] += 1
        event['status'] = 2

        # flowchart 8
        if (ts_last_furtif
            and ts_last_bagot <= self.bagot_time
            and freq >= self.furtif_freq):
            self.logger.debug(log.format('Bagot'))
            event['status'] = 3
        else:
            self.logger.debug(log.format('Furtif'))



    def is_furtif(self, ts_curr, ts_prev):
        ts_diff = ts_curr - ts_prev

        return (True if ts_diff <= self.furtif_time else False)



    def check_event(self, _id, event):
        changed = False
        new_event = False

        self.logger.debug(" + Event:")

        state = event['state']
        state_type = event['state_type']

        self.logger.debug("   - State:\t\t'%s'" % legend[state])
        self.logger.debug("   - State type:\t'%s'" % legend_type[state_type])

        now = int(time())

        event['timestamp'] = event.get('timestamp', now)

        try:
            # Get old record
            #record = self.storage.get(_id, account=self.account)
            change_fields = {
                'state': 1,
                'state_type': 1,
                'last_state_change': 1,
                'output': 1,
                'perf_data_array': 1,
                'timestamp': 1,
                'cancel': 1,
                'status': 1
            }

            devent = self.collection.find_one(_id, fields=change_fields)

            if not devent:
                new_event = True

            self.logger.debug(" + Check with old record:")
            log = 'Status is set to {} for event %s' % event['rk']
            old_state = devent['state']
            old_state_type = devent['state_type']
            prev_status = devent.get('status', 1)
            ts_curr = event['timestamp']
            ts_prev = devent['timestamp']
            event['last_furtif'] = devent.get('last_furtif', 0)
            event['furtif_freq'] = devent.get('furtif_freq', 0)
            event['last_state_change'] = devent.get('last_state_change', event['timestamp'])

            self.logger.debug("   - State:\t\t'%s'" % legend[old_state])
            self.logger.debug("   - State type:\t'%s'" % legend_type[old_state_type])


            if state != old_state:
                event['previous_state'] = old_state

            if state != old_state or state_type != old_state_type:
                self.logger.debug(" + Event has changed !")
                changed = True
            else:
                self.logger.debug(" + No change.")

            # Check if event is still canceled
            # flowchart 1 2 3
            # status legend:
            # 0 == Ok
            # 1 == On going
            # 2 == Stealthy
            # 3 == Bagot
            # 4 == Canceled
            if ((prev_status == 4
                 and old_state != state
                 and (self.restore_event
                      or state == 0
                      or old_state == 0))
                or prev_status != 4):

                # flowchart 4
                if (state == 0):
                    # flowchart 5 6
                    if (old_state == 0
                        or not self.is_furtif(ts_curr, ts_prev)):
                        self.logger.debug(log.format('Ok'))
                        event['status'] = 0
                        event['furtif_freq'] = 0
                    else:
                        self.check_bagot(event, devent)

                else:
                    # flowchart 7
                    if (not self.is_furtif(ts_curr, ts_prev)):
                        self.logger.debug(log.format('On going'))
                        event['status'] = 1
                        event['furtif_freq'] = 0
                    else:
                        self.check_bagot(event, devent)

            else:
                self.logger.debug(log.format('Canceled'))
                event['status'] = 4
                event['furtif_freq'] = 0


            try:
                event = self.merge_perf_data(devent, event)
            except Exception, err:
                self.logger.warning("merge_perf_data: %s" % err)

        except:
            # No old record
            self.logger.debug(" + New event")
            changed = True
            old_state = state

        if changed:
            # Tests if change is from alert to non alert
            if 'last_state_change' in event and (state == 0 or (state > 0 and old_state == 0)):
                event['previous_state_change_ts'] = event['last_state_change']
            event['last_state_change'] = event.get('timestamp', now)


        if new_event:
            self.store_new_event(_id, event)
        else:
            #Cancel management
            self.process_cancel(devent, event)
            change = {}
            # resets ack status for this alert if already set to true
            if devent.get('ack', False):
                change['ack'] = False

            for key in change_fields:
                if key in event and key in devent and devent[key] != event[key]:
                    change[key] = event[key]
                elif key in event and key not in devent:
                    change[key] = event[key]
            if change:
                self.storage.get_backend('events').update({'_id': _id}, {'$set': change})

        mid = None
        if changed and self.autolog:
            mid = self.log_event(_id, event)

        return mid

    def merge_perf_data(self, old_event, new_event):
        old_event['perf_data_array'] = old_event.get('perf_data_array', [])
        new_event['perf_data_array'] = new_event.get('perf_data_array', [])

        if new_event['perf_data_array'] != []:
            perf_data_array = old_event['perf_data_array']

            new_metrics = [ perf['metric'] for perf in new_event['perf_data_array'] ]
            old_metrics = [ perf['metric'] for perf in old_event['perf_data_array'] ]

            if new_metrics == old_metrics:
                new_event['perf_data_metrics'] = new_metrics
                return new_event

            new_event['perf_data_metrics'] = uniq(new_metrics + old_metrics)

            for new_metric in new_metrics:
                if new_metric in old_metrics:
                    perf_data_array[old_metrics.index(new_metric)] = new_event['perf_data_array'][new_metrics.index(new_metric)]
                else:
                    perf_data_array.append(new_event['perf_data_array'][new_metrics.index(new_metric)])

            new_event['perf_data_array'] = perf_data_array



        return new_event

    #Cancel and uncancel process are processed here
    def process_cancel(self, devent, event):
        #Event cancellation management

        if 'cancel' in event:

            #Saving status, in case cancel is undone
            previous_status = devent.get('status', 1)
            previous_state = devent['state']
            # Cancel value means cancel for True or undo cancel for False
            cancel = {
                'author' : event.get('author','unknown'),
                'cancel' : event['cancel'],
                'comment' : event['output'],
                'previous_status': previous_status,
                'previous_state': previous_state
            }
            # Preparing event for cancel, avoid using space and use previous output
            if 'author' in event:
                del event['author']

            event['output'] = devent.get('output', '')

            #Is it a cancel ?
            if event['cancel'] == True:
                self.logger.info("set cancel to the event")
                # If cancel is not in ok state, then it is not an alert cancellation
                event['cancel'] = cancel
                event['status'] = 4
                event['state'] = 0

            # Undo cancel ?
            elif event['cancel'] == False:
                #If event has been previously cancelled
                if 'cancel' in devent and 'cancel' in devent['cancel'] and devent['cancel']['cancel']:

                    event['cancel'] = cancel
                    #Restore previous status
                    event['status'] = devent['cancel']['previous_status']
                    event['state'] = devent['cancel']['previous_state']



            #Avoid saving useless boolean values to db
            if 'cancel' in event and isinstance(event['cancel'], bool):
                del event['cancel']



    def store_new_event(self, _id, event):
        record = Record(event)
        record.type = "event"
        record.chmod("o+r")
        record._id = _id

        self.storage.put(record, namespace=self.namespace, account=self.account)

    def log_event(self, _id, event):
        self.logger.debug("Log event '%s' in %s ..." % (_id, self.namespace_log))
        record = Record(event)
        record.type = "event"
        record.chmod("o+r")
        record.data['event_id'] = _id
        record._id = _id + '.' + str(time())

        self.storage.put(record, namespace=self.namespace_log, account=self.account)
        return record._id

    def get_logs(self, _id, start=None, stop=None):
        return self.storage.find({'event_id': _id}, namespace=self.namespace_log, account=self.account)

    def remove_all(self):
        self.logger.debug("Remove all logs and state archives")

        self.storage.drop_namespace(self.namespace)
        self.storage.drop_namespace(self.namespace_log)


