#!/usr/bin/env python
# -*- coding: utf-8 -*-
# --------------------------------
# Copyright (c) 2016 "Capensis" [http://www.capensis.com]
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

from unittest import TestCase

from canopsis.alerts import AlarmField
from canopsis.alerts.manager import Alerts
from canopsis.check import Check
from canopsis.middleware.core import Middleware


class BaseTest(TestCase):
    def setUp(self):
        self.alarm_storage = Middleware.get_middleware_by_uri(
            'storage-periodical-testalarm://'
        )
        self.config_storage = Middleware.get_middleware_by_uri(
            'storage-default-testconfig://'
        )
        self.filter_storage = Middleware.get_middleware_by_uri(
            'storage-default-testalarmfilter://'
        )

        self.manager = Alerts()
        self.manager[Alerts.ALARM_STORAGE] = self.alarm_storage
        self.manager[Alerts.CONFIG_STORAGE] = self.config_storage
        self.manager[Alerts.FILTER_STORAGE] = self.filter_storage

        self.config_storage.put_element(
            element={
                '_id': 'test_config',
                'crecord_type': 'statusmanagement',
                'bagot_time': 3600,
                'bagot_freq': 10,
                'stealthy_time': 300,
                'stealthy_show': 600,
                'restore_event': True,
                'auto_snooze': False,
                'snooze_default_time': 300,
            },
            _id='test_config'
        )

    def tearDown(self):
        """Teardown"""
        self.alarm_storage.remove_elements()
        self.filter_storage.remove_elements()

    def gen_fake_alarm(self, moment):
        """
        Generate a fake alarm/value.
        """
        alarm_id = '/fake/alarm/id'
        alarm = self.manager.make_alarm(
            alarm_id,
            {
                'connector': 'fake-connector',
                'connector_name': 'fake-connector-name',
                'component': 'c',
                'timestamp': moment
            }
        )

        value = alarm[self.manager[Alerts.ALARM_STORAGE].VALUE]
        value[AlarmField.state.value] = {
            't': moment,
            'val': Check.MINOR
        }
        value[AlarmField.steps.value] = [
            {
                '_t': 'stateinc',
                't': moment,
                'a': 'fake-author',
                'm': 'fake-message',
                'val': Check.MINOR
            }
        ]

        return alarm, value
