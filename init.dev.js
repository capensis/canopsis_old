/*
 * Copyright (c) 2015 "Capensis" [http://www.capensis.com]
 *
 * This file is part of Canopsis.
 *
 * Canopsis is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * Canopsis is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with Canopsis. If not, see <http://www.gnu.org/licenses/>.
 */

 require.config({
    paths: {
        'listalarm': 'canopsis/brick-listalarm/src/widgets/listalarm/listalarm',

    }
});

 define([
    'canopsis/brick-listalarm/src/adapters/alertexpression',
    'canopsis/brick-listalarm/src/adapters/alerts',
    'canopsis/brick-listalarm/src/serializers/alertexpression',
    'canopsis/brick-listalarm/src/serializers/alerts',
    'canopsis/brick-listalarm/src/widgets/listalarm/controller',
    'ehbs!listalarm'
], function () {
    
});

