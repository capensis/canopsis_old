from __future__ import unicode_literals

import re
import itertools


class HeartBeat:

    MAX_DUR_REGEXP = "^[0-9]*(s|m|h)$"
    __expected_interval_pattern = re.compile(MAX_DUR_REGEXP)
    MAPPINGS_KEY = "mappings"
    MAX_DUR_KEY = "maxduration"

    def __init__(self, mappings, maxduration):
        """Create a new heartbeat instance.

        For more information about the format see the isValid class method
        docstring.

        :param mappings: a list of dict use to map an event to a maxDuration.
        :param maxduration: a string that represent the time to wait
        before an alarm is created in the case no event link to an entity is
        received.
        """
        self.mappings = mappings
        self.max_duration = maxduration

    @staticmethod
    def validate_heartbeat_pattern(pattern):
        """
        Check if ``pattern`` is a non-empty dict and
        the all of them keys and values are strings.

        :param `dict` pattern: a Heartbeat pattern.
        :rtype: `bool`.
        """
        return bool(pattern) and isinstance(pattern, dict) and all((
            isinstance(i, basestring) for i
            in itertools.chain(*pattern.items())
        ))

    @classmethod
    def check_expected_interval(cls, expected_interval):
        return isinstance(expected_interval, str) and \
               bool(cls.__expected_interval_pattern.match(expected_interval))

    @classmethod
    def isValid(cls, heartBeat):
        """
        Check if the heartBeat given is valid.

        In order to considered valid, the `mappings` and `maxDuration`
        attributs must be valid.

        `Mappings` is a list of item.
        An `item` is an ojects with at least one key. The key and the
        associated value are both string
        `maxduration` is a string that match the follow pattern: ^[0-9]*(s|m|h)$.
        `s` means waiting XX seconds
        `m` means waiting XX minutes
        `h` means waiting XX hours

        :param dict heartBeat: a dict.
        :rtype: (bool, str).
        """
        try:
            mappings = heartBeat.mappings
        except KeyError:
            return False, "The `mappings` field is missing."

        try:
            max_duration = heartBeat.max_duration
        except KeyError:
            return False, "The `maxduration` field is missing."

        if len(mappings) == 0:
            return False, "The mappings array must not be empty."

        for mapping in mappings:
            if len(mapping) == 0:
                return False, "The mapping object must not be empty."

            for it, key in enumerate(mapping):
                if not isinstance(key, basestring):
                    return False, "{} must be a string.".format(key)

                if not isinstance(mapping[key], basestring):
                    return False, "The value of `{0}` of the mapping object"\
                        " at index {1} must be a string.".format(key, it)

        if cls.check_expected_interval(max_duration):
            return True, ""

        return False, "The maxDuration fields does not match the" \
            " regular expression {}.".format(cls.MAX_DUR_KEY)

    def to_dict(self):
        """
        Return the representation of the current instance as a dict.
        :rtype: dict
        """
        return {self.MAPPINGS_KEY: self.mappings,
                self.MAX_DUR_KEY: self.max_duration}
