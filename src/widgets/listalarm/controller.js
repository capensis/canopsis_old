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

Ember.Application.initializer({
    name: 'ListAlarmWidget',
    after: ['TimeWindowUtils', 'DataUtils', 'WidgetFactory', 'UserconfigurationMixin', 'SchemasLoader'],
    initialize: function(container, application) {
		    var timeWindowUtils = container.lookupFactory('utility:timewindow'),
            dataUtils = container.lookupFactory('utility:data'),
			      WidgetFactory = container.lookupFactory('factory:widget'),
			      UserConfigurationMixin = container.lookupFactory('mixin:userconfiguration');

        var get = Ember.get,
            set = Ember.set,
            isNone = Ember.isNone;

        // load the viewMixin
        var listOptions = {
            mixins: [
                UserConfigurationMixin
            ]
        };

        /**
         * This widget allows to display alarms, with action possible on them.
         *
         * @memberOf canopsis.frontend.brick-listalarm
         * @mixes UserConfigurationMixin
         * @class WidgetListAlarm
         * @widget listalarm
         */
        var widget = WidgetFactory('listalarm',{

			viewMixins: [
            ],

            /**
             * Create the widget and set widget params into Ember vars
             * @method init
             */
            init: function() {
                this._super.apply(this, arguments);
				        this.fetchAlarms();
                this.valideExpression();

				        set(this, 'store', DS.Store.extend({
                    container: get(this, 'container')
                }));
                this.showParams();

                
            },

            showParams: function () {
                var controller = this;
                var params = ["popup", "title"];
                console.error("brick's parameters");
                params.forEach(function(param) {
                    console.error(param + ': ' + controller.get('model.' + param));
                });
                console.error("default_sort_column: " + controller.get('model.default_sort_column.property') + "-" + controller.get('model.default_sort_column.direction'))
            },

            /**
             * Set the reload to true in order to redraw events
             * extend the native refreshContent method
             * @method refreshContent
             */
            refreshContent: function () {
				          // Not implemented because backend too long, feature not useful for this widget
            },

            /**
             * Get the Alarms from the backend using the adapter
             * @method fetchAlarms
             */
            fetchAlarms: function() {
              var controller = this;

              var query = { tstart: 1483225200, tstop: 1583225200 };

              var adapter = dataUtils.getEmberApplicationSingleton().__container__.lookup('adapter:alerts');
            	adapter.findQuery('alerts', query).then(function (result) {
                    // onfullfillment
					          var alerts = get(result, 'data');
                    controller.setAlarmsForShow(alerts[0]['alarms']);

                    console.log('alerts::', alerts);
              }, function (reason) {
                    // onrejection
                    // console.error('ERROR in the adapter: ', reason);
              });
            },

            /**
             * Get the Alarms from the backend using the adapter
             * @method valideExpression
             */
            valideExpression: function () {
              var controller = this;

              var query = {
                expression: 'a=1'
              };

              var adapter = dataUtils.getEmberApplicationSingleton().__container__.lookup('adapter:alertexpression');
            	adapter.findQuery('alertexpression', query).then(function (result) {
                    // onfullfillment
					          var result = get(result, 'data');
                    console.error('alertexpression result', result);
              }, function (reason) {
                    // onrejection
                    console.error('ERROR in the adapter: ', reason);
              });
            },

            setAlarmsForShow: function (alarms) {
              var fields = ['component', 'resource', 'connector', 'connector_name'];
              var controller = this;
              var alarmsArr = alarms.map(function(alarm) {
                var newAlarm = {};
                fields.forEach(function(field) {
                  newAlarm[field] = alarm['v'][field]
                })
                return newAlarm;
              });

              this.set('alarms', alarmsArr);
              this.set('fields', fields);
            }

        }, listOptions);

        application.register('widget:listalarm', widget);
    }
});
