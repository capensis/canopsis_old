/**
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
    name: 'UserviewRoute',
    after: ['AuthenticatedRoute', 'DataUtils', 'NotificationUtils', 'FormsUtils', 'WidgetSelectorsUtils'],
    initialize: function(container, application) {

        var AuthenticatedRoute = container.lookupFactory('route:authenticated');

        var dataUtils = container.lookupFactory('utility:data');
        var notificationUtils = container.lookupFactory('utility:notification');
        var formUtils = container.lookupFactory('utility:forms');
        var widgetSelectorsUtils = container.lookupFactory('utility:widget-selectors');

        var set = Ember.set,
            get = Ember.get,
            isNone = Ember.isNone;

        var initialLoadDone = false;

        /**
         * @class UserviewRoute
         * @extends AuthenticatedRoute
         * @constructor
         */
        var route = AuthenticatedRoute.extend({
            beforeModel: function(transition) {
                var app = this.controllerFor('application');
                app.addConcurrentLoading('userview');

                return this._super(transition);
            },

            afterModel: function(view, transition) {
                var app = this.controllerFor('application');
                app.removeConcurrentLoading('userview');

                return this._super(view, transition);
            },

            actions: {
                error: function(error, transition){
                    if (error.status === 0) {
                        console.debug('no error detected in access status');
                    } else if (error.status === 403) {
                        console.debug('go to some default route');
                    } else if (error.status === 401) {
                        console.debug('handle 401');
                    } else if (error.status === 404) {
                        this.transitionTo('/userview/view.404');
                    } else {
                        console.error(error);
                        notificationUtils.error(__('Impossible to load view.'));
                    }
                },


                /**
                 * @event toggleEditMode
                 *
                 * Toggles the edition mode for the currently opened view
                 */
                toggleEditMode: function() {
                    var applicationController = this.controllerFor('application');
                    if (get(applicationController, 'editMode') === true) {

                        console.info('Leaving edit mode');
                        set(applicationController, 'editMode', false);
                    } else {

                        // Try to rollback each widget of the view
                        var viewWidgets = widgetSelectorsUtils.children(get(this.controllerFor('userview'), 'containerwidget'));
                        for (var i = 0, l = viewWidgets.length; i < l; i++) {
                            var currentWidget = viewWidgets[i];

                            if(get(currentWidget, 'rollbackable')) {
                                currentWidget.send('rollback');
                            }
                        }

                        console.info('Entering edit mode');
                        set(applicationController, 'editMode', true);
                    }
                },

                /**
                 * @event exportCurrentView
                 *
                 * Exports the currently opened view
                 */
                exportCurrentView: function() {
                    var viewModel = this.controllerFor('userview').get('model');
                    var viewJSON = viewModel.serialize();

                    viewJSON = dataUtils.cleanJSONIds(viewJSON);

                    dataUtils.download(JSON.stringify(viewJSON), get(viewModel, 'crecord_name') + '.json', 'application/json');
                },

                /**
                 * @event toggleDevTools
                 *
                 * show or hide developper tools
                 */
                toggleDevTools: function () {
                    //FIXME don't use jquery in here, it's for views !
                    $('.debugmenu').slideToggle(200);
                },

                /**
                 * @event toggleFullscreen
                 * Toggle fullscreen and regular mode, by toggling Applicationcontroller#fullscreenMode boolean.
                 * The rest of the implementation is on handlebars templates (application and userview)
                 */
                toggleFullscreen: function() {
                    var applicationController = this.controllerFor('application');

                    var updatedFullscreenMode = !get(applicationController, 'fullscreenMode');

                    set(applicationController, 'fullscreenMode', updatedFullscreenMode);
                },

                /**
                 * @event stopLiveReporting
                 * clears the live reporting interval, ending the live reporting mode
                 */
                stopLiveReporting: function () {
                    this.setReportingInterval(undefined);
                },

                /**
                 * @event displayLiveReporting
                 *
                 * Display a pop up allowing customer to set view time
                 * parameters that will affect all widget data selection.
                 */
                displayLiveReporting: function () {

                    var userview = this;

                    var record = dataUtils.getStore().createRecord('livereporting', {
                        crecord_type: 'livereporting'
                    });

                    var recordWizard = formUtils.showNew('modelform', record, {
                        title: __('Edit live reporting')
                    });

                    recordWizard.submit.then(function(form) {

                        record = get(form, 'formContext');
                        var interval = get(record, 'dateinterval');
                        set(userview.controllerFor('application'), 'interval', interval);
                        userview.setReportingInterval(interval);
                        console.debug('setting interval to value', interval);
                    });
                },

                /**
                 * @event refresh
                 *
                 * Refreshes the currently opened view
                 */
                refresh: function() {
                    var userviewController = this.controllerFor('userview');

                    console.log('refresh', this);
                    this.refresh();
                    userviewController.send('refreshView');
                }
            },


            /**
             * @method setReportingInterval
             * @param interval
             */
            setReportingInterval: function (interval) {

                set(this.controllerFor('application'), 'interval', interval);

                var controller = this.controllerFor('userview');
                var rootWidget = get(controller, 'content.containerwidget');
                var children = widgetSelectorsUtils.children(rootWidget);


                //Set filter as void instead undefined
                if (Ember.isNone(interval)) {
                    interval = {};
                }

                for (var i = 0, l = children.length; i < l; i++) {

                    console.debug('Child widget', get(children[i], 'id'), children[i]);
                    var widgetController = get(children[i], 'controllerInstance');

                    if (!Ember.isNone(widgetController)) {
                        //each widget takes responsibility to refresh or not once update interval is called
                        widgetController.updateInterval(interval);
                    } else {
                        console.error('Unable to find controller instance for widget', get(children[i], 'id'));
                    }
                }
            },

            /**
             * @method setupController
             * @param controller
             * @param model
             */
            setupController: function(controller, model) {
                console.log('UserviewRoute setupController', arguments);

                this._super.apply(this, arguments);

                controller.setProperties({
                    'isMainView': true
                });

                if(initialLoadDone === false) {
                    initialLoadDone = true;
                }

                set(this.controllerFor('application'), 'currentViewId', get(model, 'id'));
                set(this.controllerFor('application'), 'currentViewModel', model);

                var app = this.controllerFor('application');

                if(controller.trigger) {
                    controller.trigger('refreshView');
                }
            }
        });

        application.register('route:userview', route);
    }
});
