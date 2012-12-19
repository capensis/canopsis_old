/*
#--------------------------------
# Copyright (c) 2011 "Capensis" [http://www.capensis.com]
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
*/
Ext.define('canopsis.controller.View', {
    extend: 'canopsis.lib.controller.ctree',

    views: ['View.TreePanel'],
    stores: ['Views', 'TreeStoreViews'],
    model: ['View'],

    logAuthor: '[controller][View]',

    init: function() {
		log.debug('Initialize ...', this.logAuthor);

		//this.formXtype = 'GroupForm'
		this.listXtype = 'ViewTreePanel';

		this.modelId = 'View';

		log.debug(' + Load treeStore ...', this.logAuthor);
		this.treeStore = Ext.data.StoreManager.lookup('TreeStoreViews');

		if (!this.treeStore.isLoading()) {
			this.treeStore.load();
		}

		this.store = Ext.data.StoreManager.lookup('Views');

		this.store.proxy.on('exception', function(proxy, response) {
			log.error('Error in request', this.logAuthor);
			var message = Ext.String.format(
				'{0}<br>{1}: {2}',
				 _('Error in request:'),
				 response.status,
				 response.statusText
			);

			global.notify.notify(_('View'), message, 'error');
		}, this);

		this.callParent(arguments);

		//binding view export pdf

    },

    bindTreeEvent: function() {
		var btns = Ext.ComponentQuery.query('#' + this.tree.id + ' [action=import]');
		for (var i = 0; i < btns.length; i++)
			btns[i].on('click', this.openFilepopup, this);

		this.tree.on('exportPdf', function(view) {
				this.getController('Reporting').launchReport(view);
			},this);

		this.tree.getView().on('getViewFile', this.getViewFile, this);

		this.tree.getView().on('OpenViewOption', this.OpenViewOption, this);
	},

	check_right_on_drop: function(node,data,overModel) {
		log.debug('checking right on record before drop', this.logAuthor);
		var ctrl = this.getController('Account');
		var src_record = data.records[0];
		var dest_record = overModel;

		//check if right to right, if not, stop event + display notif
		if (!(ctrl.check_record_right(src_record, 'w') && ctrl.check_record_right(dest_record, 'w'))) {
			global.notify.notify(_('Access denied'), _('You don\'t have the rights to modify this object'), 'error');
			return true;
		}

	},

    addLeafButton: function() {
		this.create_new_view();
	},

    addDirectoryButton: function() {
		this.create_new_directory();
	},

	itemDoubleClick: function() {
		var tree = this.tree;

		var selection = tree.getSelectionModel().getSelection()[0];
		if (selection.get('leaf')) {
			view_id = selection.get('id');
			view_name = selection.get('crecord_name');
			this.getController('Tabs').open_view({ view_id: view_id, title: view_name });
		}
	},

	_duplicateButton: function() {
		log.debug('duplicate', this.logAuthor);
		//get selected views
		var tree = this.tree;
		var selection = tree.getSelectionModel().getSelection();
		//for each selected view
		for (var i = 0; i < selection.length; i++) {
			if (selection[i].isLeaf()) {
				var view_id = 'view.' + global.account.user + '.' + global.gen_id();
				var new_record = selection[i].copy(view_id);
				Ext.data.Model.id(new_record);
				new_record.set('_id', view_id);
				new_record.set('id', view_id);
				this.add_to_home(new_record, false);
			}
		}
	},

	exportButton: function(item) {
		log.debug('Export view', this.logAuthor);
		this.getController('Schedule').scheduleWizard(item, this.tree.id);
	},

	/////////////////

	create_new_directory: function() {
		Ext.Msg.prompt(_('Directory name'), _('Please enter directory name:'), function(btn, directoryName) {
			if (btn == 'ok') {

				var record = Ext.create('canopsis.model.View');

				var directory_id = 'directory.' + global.account.user + '.' + global.gen_id();

				record.set('crecord_name', directoryName);

				record.set('_id', directory_id);
				record.set('id', directory_id);
				//need to set the empty array , otherwise treepanel send request
				//to fetch inside
				record.set('children', []);
				record.set('leaf', false);

				this.add_to_home(record, false);

			} else {
				log.debug('cancel new view', this.logAuthor);
			}
		}, this);
	},

	create_new_view: function() {
		Ext.Msg.prompt(_('View name'), _('Please enter view name:'), function(btn, viewName) {
			if (btn == 'ok') {
				// Create view
				var view_id = 'view.' + global.account.user + '.' + global.gen_id();

				//building record
				var record = Ext.create('canopsis.model.View');
				record.set('_id', view_id);
				record.set('id', view_id);
				record.set('crecord_name', viewName);
				record.set('leaf', true);

				this.add_to_home(record, true);
			} else {
				log.debug('cancel new view', this.logAuthor);
			}
		}, this);
	},

	add_to_home: function(record,open_after_put) {
		//append child
		this.open_after_put = open_after_put;
		var rootDir = 'directory.root.' + global.account.user;

		log.debug('Add view: ', this.logAuthor);
		log.debug(' + root dir: ' + rootDir, this.logAuthor);
		log.debug(' + name: ' + record.get('crecord_name'), this.logAuthor);
		log.debug(' + id: ' + record.get('_id'), this.logAuthor);

		var parentNode = this.treeStore.getNodeById(rootDir);
		var rootNode = this.treeStore.getRootNode();
		if (parentNode) {
			parentNode.appendChild(record);

			//this is a hack
			rootNode.dirty = false;

			this.treeStore.sync({
				scope: this,
				callback: function(batch) {
					var response = Ext.decode(batch.operations[0].response.responseText);

					this.request_success = true;

					var data = response.data;
					for (var i = 0; i < data.length; i++)
						if (!data[i].success)
							this.request_success = false;

					this.treeStore.load({
							scope: this,
							callback: function() {
								if (this.open_after_put == true && this.request_success) {
									var tab = this.getController('Tabs').open_view({ view_id: record.get('_id'), title: record.get('crecord_name') });
									tab.editMode();
									tab.doRedraw();

									//Hack
									Ext.getCmp('dashboardSelector').store.load();
								}
							}
					});
				}
			});


		}else {
			log.debug('Impossible to add view, root directory not found ....', this.logAuthor);
		}
	},

    checkOpen: function(view_id) {
		var maintabs = Ext.getCmp('main-tabs');
		var tab = Ext.getCmp(view_id + '.tab');

		if (tab) {
			global.notify.notify(_('Warning'), 'You must close view before renaming', 'error');
			return true;
		}else {
			return false;
		}
	},

	openFilepopup: function(file) {
		log.debug('Open file popup', this.logAuthor);
		var config = {
			_fieldLabel: _('View dump'),
			copyPasteZone: true,
			constrainTo: this.tree.id
			};
		var popup = Ext.create('canopsis.lib.view.cfile_window', config);
		popup.show();

		popup.on('save', function(file) {
			if (file.length > 0) {
				var file_type = file[0].type;
				if (file_type == '' || file_type == 'application/json') {
					this.importFile(file);
					popup.close();
				}else {
					log.debug('Wrong file type: ' + file_type, this.logAuthor);
					global.notify.notify(_('Wrong file type'), _('Please choose a correct json file'), 'info');
				}
			}

			if (file.items) {
				this.importView(file);
				popup.close();
			}

		},this);
	},

	importView: function(obj) {
		var record = Ext.create('canopsis.model.View', obj);
		record.set('_id', 'view.' + global.account.user + '.' + global.gen_id());
		record.set('leaf', true);
		this.add_to_home(record, false);
		//this.tree.store.load()
	},

	importFile: function(file) {
		log.debug('Import view file', this.logAuthor);
		var reader = new FileReader();
		reader.onload = (function(e) {
			this.importView(Ext.decode(e.target.result));
		}).bind(this);
		reader.readAsText(file[0]);
	},

	getViewFile: function(view_id) {
		log.debug('Get view file', this.logAuthor);
		window.open('/ui/view/export/' + view_id);
	},

	OpenViewOption: function(view) {
		var form = Ext.create('canopsis.view.Tabs.View_form');

		form.getForm().setValues(view.view_options);

		this.view_option_win = Ext.create('widget.window', {
			title: view.crecord_name + ' ' + _('options'),
			items: [form],
			closable: false,
			resizable: false,
			constrain: true,
			renderTo: this.tree.id,
			closeAction: 'destroy'
		}).show();

		form.on('save', function(data) {
			this.view_option_win.close();
			updateRecord('object', 'view', 'View', view._id, {view_options: data});
			this.treeStore.load();
		},this);

		form.on('close', function() {
			this.view_option_win.close();
		},this);

	}

});
