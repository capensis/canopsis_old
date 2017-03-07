# -*- coding: utf-8 -*-
# --------------------------------
# Copyright (c) 2017 "Capensis" [http://www.capensis.com]
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

from calendar import timegm
from datetime import datetime
from dateutil.rrule import rrulestr
from json import loads, dumps
from uuid import uuid4

from canopsis.common.utils import singleton_per_scope
from canopsis.context.manager import Context
from canopsis.configuration.configurable.decorator import (
    add_category, conf_paths
)
from canopsis.middleware.registry import MiddlewareRegistry

# Might be useful when dealing with rrules
# from dateutil.rrule import rrulestr

CONF_PATH = 'pbehavior/pbehavior.conf'
CATEGORY = 'PBEHAVIOR'


class BasePBehavior(dict):
    _FIELDS = ()
    _EDITABLE_FIELDS = ()

    def __init__(self, **kwargs):
        for key, value in kwargs.iteritems():
            if key in self._FIELDS:
                self.__dict__[key] = value

    def __repr__(self):
        return repr(self.__dict__)

    def __setitem__(self, key, item):
        if key in self._EDITABLE_FIELDS:
            self.__dict__[key] = item

    def __getitem__(self, key):
        return self._get(key)

    def __getattr__(self, item):
        return self._get(item)

    def _get(self, item):
        if item in self._FIELDS and item in self.__dict__:
            return self.__dict__[item]
        return None

    def update(self, **kwargs):
        for key, value in kwargs.iteritems():
            if key in self._EDITABLE_FIELDS:
                self.__dict__[key] = value
        return self.__dict__

    def to_dict(self):
        return self.__dict__


class PBehavior(BasePBehavior):
    NAME = 'name'
    FILTER = 'filter'
    COMMENTS = 'comments'
    TSTART = 'tstart'
    TSTOP = 'tstop'
    RRULE = 'rrule'
    ENABLED = 'enabled'
    EIDS = 'eids'
    CONNECTOR = 'connector'
    CONNECTOR_NAME = 'connector_name'
    AUTHOR = 'author'

    _FIELDS = (NAME, FILTER, COMMENTS, TSTART, TSTOP, RRULE, ENABLED, EIDS,
               CONNECTOR, CONNECTOR_NAME, AUTHOR)

    _EDITABLE_FIELDS = (NAME, FILTER, TSTART, TSTOP, RRULE, ENABLED,
                        CONNECTOR, CONNECTOR_NAME, AUTHOR)

    def __init__(self, **kwargs):
        if PBehavior.FILTER in kwargs:
            kwargs[PBehavior.FILTER] = dumps(kwargs[PBehavior.FILTER])
        super(PBehavior, self).__init__(**kwargs)

    def update(self, **kwargs):
        if PBehavior.FILTER in kwargs:
            kwargs[PBehavior.FILTER] = dumps(kwargs[PBehavior.FILTER])
        super(PBehavior, self).update(**kwargs)


class Comment(BasePBehavior):
    ID = '_id'
    AUTHOR = 'author'
    TS = 'ts'
    MESSAGE = 'message'

    _FIELDS = (ID, AUTHOR, TS, MESSAGE)
    _EDITABLE_FIELDS = (AUTHOR, MESSAGE)


@conf_paths(CONF_PATH)
@add_category(CATEGORY)
class PBehaviorManager(MiddlewareRegistry):

    PBEHAVIOR_STORAGE = 'pbehavior_storage'

    _UPDATE_FLAG = 'updatedExisting'

    @property
    def pbehavior_storage(self):
        return self[PBehaviorManager.PBEHAVIOR_STORAGE]

    @property
    def context(self):
        return singleton_per_scope(Context)

    def get(self, _id, query=None):
        """Get pbehavior by id.
        :param str id: pbehavior id
        :param dict query: filtering options
        """
        return self.pbehavior_storage.get_elements(ids=_id, query=query)

    def create(
        self,
        name, filter, author,
        tstart, tstop, rrule='',
        enabled=True, comments=None,
        connector='canopsis', connector_name='canopsis'
    ):
        """
        Method creates pbehaviro record
        :param str name: filtering options
        :param dict filter: a mongo filter that match entities from canopsis context
        :param str author: the name of the user/app that has generated the pbehavior
        :param timestamp tstart: timestamp that correspond to the start of the pbehavior
        :param timestamp tstop: timestamp that correspond to the end of the pbehavior
        :param str rrule: reccurent rule that is compliant with rrule spec
        :param bool enabled: boolean to know if pbhevior is enabled or disabled
        :param list of dict comments: a list of comments made by users
        :param str connector: a string representing the type of connector that has generated the pbehavior
        :param str connector_name:  a string representing the name of connector that has generated the pbehavior
        :return:
        """
        kw = locals()
        kw.pop('self')
        data = PBehavior(**kw)
        if not data.comments or not isinstance(data.comments, list):
            data.update(comments=[])
        else:
            for c in data.comments:
                c.update({'_id': str(uuid4())})
        result = self.pbehavior_storage.put_element(element=data.to_dict())
        return result

    def read(self, _id=None):
        """Get pbehavior or list pbehaviors.
        :param str _id: pbehavior id, _id may be equal to None
        """
        result = self.get(_id)

        return result if _id else list(result)

    def update(self, _id, **kwargs):
        """
        Update pbehavior record
        :param str _id: pbehavior id
        :param dict kwargs: values pbehavior fields
        """
        pbehavior = PBehavior(**self.get(_id))
        pbehavior.update(**kwargs)

        result = self.pbehavior_storage.put_element(
            element=pbehavior.to_dict(), _id=_id
        )

        if (PBehaviorManager._UPDATE_FLAG in result and
                result[PBehaviorManager._UPDATE_FLAG]):
            return pbehavior.to_dict()
        return None

    def delete(self, _id=None, _filter=None):
        """
        Delete pbehavior record
        :param str _id: pbehavior id
        """
        result = self.pbehavior_storage.remove_elements(
            ids=_id, _filter=_filter
        )

        return self._check_response(result)

    def create_pbehavior_comment(self, pbehavior_id, author, message):
        """
        Сreate comment for pbehavior
        :param str pbehavior_id: pbehavior id
        :param str author: author of the comment
        :param str message: text of the comment
        """
        comments = {
            Comment.ID: str(uuid4()),
            Comment.AUTHOR: author,
            Comment.TS: timegm(datetime.utcnow().timetuple()),
            Comment.MESSAGE: message
        }
        result = self.pbehavior_storage._update(
            spec={'_id': pbehavior_id},
            document={'$push': {PBehavior.COMMENTS: comments}},
            multi=False, cache=False
        )

        if (PBehaviorManager._UPDATE_FLAG in result and
                result[PBehaviorManager._UPDATE_FLAG]):
            return comments[Comment.ID]
        return None

    def update_pbehavior_comment(self, pbehavior_id, _id, **kwargs):
        """
        Update the comment record
        :param str pbehavior_id: pbehavior id
        :param str_id: comment id
        :param dict kwargs: values comment fields
        """
        pbehavior = self.get(
            pbehavior_id,
            query={PBehavior.COMMENTS: {'$elemMatch': {'_id': _id}}}
        )
        if not pbehavior:
            return None

        _comments = pbehavior[PBehavior.COMMENTS]
        if not _comments:
            return None

        comment = Comment(**_comments[0])
        comment.update(**kwargs)

        result = self.pbehavior_storage._update(
            spec={'_id': pbehavior_id, 'comments._id': _id},
            document={'$set': {'comments.$': comment.to_dict()}},
            multi=False, cache=False
        )

        if (PBehaviorManager._UPDATE_FLAG in result and
                result[PBehaviorManager._UPDATE_FLAG]):
            return comment.to_dict()
        return None

    def delete_pbehavior_comment(self, pbehavior_id, _id):
        """
        Delete comment record
        :param str pbehavior_id: pbehavior id
        :param str _id: comment id
        """
        result = self.pbehavior_storage._update(
            spec={'_id': pbehavior_id},
            document={'$pull': {PBehavior.COMMENTS: {'_id': _id}}},
            multi=False, cache=False
        )

        return self._check_response(result)

    def get_pbehaviors(self, entity_id):
        """
        Return all pbehaviors related to an entity_id, sorted by descending
        tstart.

        :param str entity_id: Id for which behaviors have to be returned

        :return: List of pbehaviors as dict, with name, tstart, tstop, rrule
          and enabled keys
        :rtype: list of dict
        """
        pbehaviors = self.pbehavior_storage.get_elements(
            query={PBehavior.EIDS: {'$in': [entity_id]}},
            sort={PBehavior.TSTART: -1}
        )
        result = [PBehavior(**pb).to_dict() for pb in pbehaviors]
        return result

    def compute_pbehaviors_filters(self):
        """
        Compute all filters and update eids attributes
        """
        pbehaviors = self.pbehavior_storage.get_elements(
            query={PBehavior.FILTER: {'$exists': True}}
        )
        for pb in pbehaviors:
            entities = self.context[Context.CTX_STORAGE].get_elements(
                query=loads(pb[PBehavior.FILTER])
            )

            pb[PBehavior.EIDS] = [e['entity_id'] for e in entities]
            self.pbehavior_storage.put_element(element=pb, _id=pb['_id'])

    def check_pbehaviors(self, entity_id, list_in, list_out):
        """
        :param str entity_id:
        :param list list_in: list of pbehavior names
        :param list list_out: list of pbehavior names
        :return: bool if the entity_id is currently in list_in arg and out list_out arg
        """
        return (self._check_pbehavior(entity_id, list_in) and
                not self._check_pbehavior(entity_id, list_out))

    def _check_pbehavior(self, entity_id, pb_names):
        """
        :param str entity_id:
        :param list pb_names: list of pbehavior names
        :return: bool if the entity_id is currently in pb_names arg
        """
        cm = singleton_per_scope(Context)
        entity = cm.get_entity_by_id(entity_id)
        event = cm.get_event(entity)

        pbehaviors = self.pbehavior_storage.get_elements(
            query={
                PBehavior.NAME: {'$in': pb_names},
                PBehavior.EIDS: {'$in': [entity_id]}
            }
        )

        names = []
        fromts = datetime.fromtimestamp
        names_append = names.append
        for pb in pbehaviors:
            tstart = fromts(pb[PBehavior.TSTART])
            tstop = fromts(pb[PBehavior.TSTOP])

            dt_list = list(
                rrulestr(pb['rrule'], dtstart=tstart).between(
                    tstart, tstop, inc=True
                )
            )

            if (len(dt_list) >= 2 and
                fromts(event['timestamp']) >= dt_list[0] and
                fromts(event['timestamp']) <= dt_list[-1]):
                names_append(pb[PBehavior.NAME])

        result = set(pb_names).isdisjoint(set(names))

        return not result

    def _check_response(self, response):
        return True if 'ok' in response and response['ok'] == 1 else False
