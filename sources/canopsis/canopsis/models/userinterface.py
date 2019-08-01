#!/usr/bin/env python
# -*- coding: utf-8 -*-

"""
UserInterface object.
"""

from __future__ import unicode_literals


class UserInterface(object):

    """
    Representation of an user interface element.
    """

    # Keys as seen in db
    _ID = '_id'
    APP_TITLE = 'app_title'
    FOOTER = 'footer'
    LOGIN_PAGE_DESCRIPTION = 'login_page_description'
    LOGO = 'logo'
    LANGUAGE = 'language'

    def __init__(self, _id, app_title=None, footer=None, login_page_description=None, logo=None, language=None, *args, **kwargs):

        self._id = _id
        self.app_title = app_title
        self.footer = footer
        self.login_page_description = login_page_description
        self.logo = logo
        self.language = language

        if args not in [(), None] or kwargs not in [{}, None]:
            print('Ignored values on creation: {} // {}'.format(args, kwargs))

    def __str__(self):
        return '{}'.format(self._id)

    def __repr__(self):
        return '<UserInterface {}>'.format(self.__str__())

    def to_dict(self):
        """
        Give a dict representation of the object.

        :rtype: dict
        """
        dictionnary = {
            self.APP_TITLE: self.app_title,
            self.FOOTER: self.footer,
            self.LOGIN_PAGE_DESCRIPTION: self.login_page_description,
            self.LOGO: self.logo,
            self.LANGUAGE: self.language
        }

        return dictionnary
