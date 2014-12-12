#!/usr/bin/env python
# -*- coding: utf-8 -*-
# --------------------------------
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

from unittest import TestCase, main

from canopsis.task import (
    TASK_PATH, TASK_PARAMS, TaskError,
    get_task_with_params, run_task,
    get_task, register_tasks, unregister_tasks, register_task,
    __TASK_PATHS as TASK_PATHS,
)

from canopsis.common.utils import path


def test_condition_true(event, ctx, **kwargs):
    return True


def test_condition_false(event, ctx, **kwargs):
    return False


MESSAGE = 'message'


def test_exception(event, ctx, **kwargs):
    raise Exception()

COUNT = 'count'


def test_condition_count(event, ctx, **kwargs):
    ctx[COUNT] = 0
    return True


def test_action(event, ctx, **kwargs):
    ctx[COUNT] += 1
    return ctx[COUNT]


def test_wrong_params():
    pass


class TaskUnregisterTest(TestCase):

    def setUp(self):
        # clear dico and add only self names
        TASK_PATHS.clear()
        self.names = 'a', 'b'
        for name in self.names:
            register_tasks(**{name: None})

    def test_unregister(self):
        for name in self.names:
            unregister_tasks(name)
            self.assertNotIn(name, TASK_PATHS)

    def test_unregister_all(self):
        unregister_tasks(*self.names)
        self.assertFalse(TASK_PATHS)

    def test_unregister_clear(self):
        unregister_tasks()
        self.assertFalse(TASK_PATHS)


class TaskRegistrationTest(TestCase):

    def setUp(self):
        # clean task paths
        TASK_PATHS.clear()
        self.tasks = {'a': None, 'b': None}
        register_tasks(**self.tasks)

    def test_register(self):
        """
        Check for registered task in registered tasks
        """
        for task in self.tasks:
            self.assertIn(task, TASK_PATHS)

    def test_register_raise(self):

        self.assertRaises(TaskError, register_tasks, **self.tasks)

    def test_register_force(self):

        register_tasks(force=True, **self.tasks)


class GetTaskTest(TestCase):

    def setUp(self):
        # clean all task paths
        TASK_PATHS.clear()

    def test_get_unregisteredtask(self):

        getTaskTest = path(GetTaskTest)
        task = get_task(getTaskTest)
        self.assertIs(task, GetTaskTest)

    def test_get_registeredtask(self):
        task_path = 'a'
        register_tasks(**{task_path: GetTaskTest})
        task = get_task(path=task_path)
        self.assertIs(task, GetTaskTest)


class TaskRegistrationDecoratorTest(TestCase):

    def setUp(self):
        TASK_PATHS.clear()

    def test_register_without_parameters(self):

        @register_task
        def register():
            pass
        self.assertIn('register', TASK_PATHS)

    def test_register(self):

        @register_task()
        def register():
            pass
        self.assertIn('register', TASK_PATHS)

    def test_registername(self):

        name = 'toto'

        @register_task(name)
        def register():
            pass
        self.assertIn(name, TASK_PATHS)

    def test_raise(self):
        name = 'toto'
        register_tasks(**{name: None})

        error = False
        try:
            @register_task
            def toto():
                pass
        except Exception:
            error = True
        self.assertTrue(error)


class GetTaskWithParamsTest(TestCase):

    def setUp(self):

        self.wrong_function = 'test.test'

        self.existing_function = 'canopsis.rule.get_task_with_params'

    def test_none_task_from_str(self):

        task_conf = self.wrong_function

        self.assertRaises(TaskError, get_task_with_params, task_conf=task_conf)

    def test_none_task_from_dict(self):

        task_conf = {TASK_PATH: self.wrong_function}

        self.assertRaises(TaskError, get_task_with_params, task_conf=task_conf)

    def test_none_task_from_dict_with_task_name(self):

        task_name = 'test'

        task_conf = {task_name: self.wrong_function}

        self.assertRaises(
            TaskError,
            get_task_with_params,
            task_conf=task_conf, task_name=task_name)

    def test_none_task_from_dict_with_task_name_and_dict(self):

        task_name = 'test'

        task_conf = {task_name: {TASK_PATH: self.wrong_function}}

        self.assertRaises(
            TaskError,
            get_task_with_params,
            task_conf=task_conf, task_name=task_name)

    def test_task_from_str(self):

        task_conf = self.existing_function

        task, params = get_task_with_params(task_conf=task_conf)

        self.assertEqual((task, params), (get_task_with_params, {}))

    def test_task_from_dict(self):

        task_conf = {TASK_PATH: self.existing_function}

        task, params = get_task_with_params(task_conf=task_conf)

        self.assertEqual((task, params), (get_task_with_params, {}))

    def test_task_from_dict_with_task_name(self):

        task_name = 'test'

        task_conf = {task_name: self.existing_function}

        task, params = get_task_with_params(
            task_conf=task_conf, task_name=task_name)

        self.assertEqual((task, params), (get_task_with_params, {}))

    def test_task_from_dict_with_task_name_and_dict(self):

        task_name = 'test'

        task_conf = {task_name: {TASK_PATH: self.existing_function}}

        task, params = get_task_with_params(
            task_conf=task_conf, task_name=task_name)

        self.assertEqual((task, params), (get_task_with_params, {}))

    def test_task_from_dict_with_params(self):

        param = 1

        task_conf = {
            TASK_PATH: self.existing_function,
            TASK_PARAMS: param}

        task, params = get_task_with_params(task_conf=task_conf)

        self.assertEqual((task, params), (get_task_with_params, param))

    def test_task_from_dict_with_name_and_params(self):

        param = 1

        task_name = 'test'

        task_conf = {
            task_name:
            {
                TASK_PATH: self.existing_function,
                TASK_PARAMS: param
            }
        }

        task, params = get_task_with_params(
            task_conf=task_conf, task_name=task_name)

        self.assertEqual((task, params), (get_task_with_params, param))

    def test_cache(self):

        task_conf = self.existing_function

        task_not_cached_0, _ = get_task_with_params(
            task_conf=task_conf, cached=False)

        task_not_cached_1, _ = get_task_with_params(
            task_conf=task_conf, cached=False)

        self.assertTrue(task_not_cached_0 is task_not_cached_1)

        task_cached_0, _ = get_task_with_params(task_conf=task_conf)

        task_cached_1, _ = get_task_with_params(task_conf=task_conf)

        self.assertTrue(task_cached_0 is task_cached_1)


if __name__ == '__main__':
    main()
