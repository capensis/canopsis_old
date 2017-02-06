Ember.Application.initializer({
    name: 'component-alarmtd',
    initialize: function(container, application) {
        var get = Ember.get,
            set = Ember.set,
            isNone = Ember.isNone;

        /**
         * This is the eventcategories component for the widget calendar
         *
         * @class eventcategories component
         * @memberOf canopsis.frontend.brick-calendar
         */
        var component = Ember.Component.extend({
            tagName: 'td',
            renderers: ['state_val', 'state_t', 'status_val', 'ack'],

            init: function() {
                this._super();

                set(this, 'alarm', get(this, 'alarm'));
                set(this, 'field', get(this, 'field'));                
              },

            click: function () {
                this.sendAction('action', this.get('alarm'), this.get('field'));
            },

            value: function() {
                var alarm = get(this, 'alarm');
                var field = get(this, 'field');
                return alarm[field.getValue];
            }.property('alarm', 'field'),

            hasRenderer: function () {
                return this.get('renderers').includes(this.get('field.name').replace(/\./g, "_"))
            }.property('alarm'),

            renderer: function () {
                return this.get('field.name').replace(/\./g, "_")
            }.property('alarm.name')
        });

        application.register('component:component-alarmtd', component);
    }
});