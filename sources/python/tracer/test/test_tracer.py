#!/usr/bin/env python
# -*- coding: utf-8 -*-

from __future__ import unicode_literals
from unittest import TestCase, main

from canopsis.common.utils import is_mongo_successfull
from canopsis.tracer.manager import (
    Trace, TraceSetError, TraceNotFound, TracerManager
)


class Unencodable(object):
    pass


class TestTracerManager(TestCase):

    def setUp(self):
        self.manager = TracerManager('mongodb-default-testtracer://')

    def test_set_trace(self):
        r = self.manager.set_trace('test_trace', 'unittest')
        self.assertTrue(is_mongo_successfull(r))

        r = self.manager.set_trace_extra('test_trace', {'c\'est': 'extra'})
        self.assertTrue(is_mongo_successfull(r))

        r = self.manager.get_by_id('test_trace')
        self.assertTrue(isinstance(r[Trace.EXTRA], dict))

        r = self.manager.add_trace_entities('test_trace', ['test_id'])
        self.assertTrue(is_mongo_successfull(r))

        r = self.manager.get_by_id('test_trace')
        self.assertEqual(len(r[Trace.IMPACT_ENTITIES]), 1)

    def test_get_trace(self):
        r = self.manager.set_trace('test_trace_get1', 'unittest')
        self.assertTrue(is_mongo_successfull(r))

        trace = self.manager.get_by_id('test_trace_get1')

        self.assertEqual(trace[Trace.ID], 'test_trace_get1')
        self.assertEqual(trace[Trace.TRIGGERED_BY], 'unittest')
        self.assertEqual(trace[Trace.IMPACT_ENTITIES], [])
        self.assertEqual(trace[Trace.EXTRA], {})

    def test_trace_not_found(self):
        with self.assertRaises(TraceNotFound):
            self.manager.get_by_id('test_trace_doesnt_exists')

    def test_del_trace(self):
        self.manager.set_trace('test_trace_del', 'unittest')
        self.manager.get_by_id('test_trace_del')
        self.manager.del_trace('test_trace_del')
        with self.assertRaises(TraceNotFound):
            self.manager.get_by_id('test_trace_del')

    def test_set_trace_error(self):
        with self.assertRaises(TraceSetError):
            self.manager.set_trace('test_trace2', 'unittest', [Unencodable])

    def tearDown(self):
        self.manager.storage._backend.drop()


if __name__ == '__main__':
    main()
