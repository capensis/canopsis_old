// http://nightwatchjs.org/gettingstarted#settings-file

/**
 * The following code helps us for requiring es6 modules
 */
process.env.VUE_CLI_BABEL_TARGET_NODE = true;
process.env.VUE_CLI_BABEL_TRANSPILE_MODULES = true;

require('@babel/register');

const path = require('path');
const deepmerge = require('deepmerge');
const loadEnv = require('@vue/cli-service/lib/util/loadEnv'); // eslint-disable-line import/no-extraneous-dependencies

/* eslint-disable import/no-extraneous-dependencies */
const seleniumServer = require('selenium-server');
/* eslint-enable import/no-extraneous-dependencies */

const localEnvPath = path.resolve(process.cwd(), '.env.local');
const baseEnvPath = path.resolve(process.cwd(), '.env');

loadEnv(localEnvPath);
loadEnv(baseEnvPath);

const sel = require('./helpers/sel');

const userOptions = JSON.parse(process.env.VUE_NIGHTWATCH_USER_OPTIONS || '{}');

const seleniumConfig = {
  start_process: true,
  server_path: seleniumServer.path,
  host: '127.0.0.1',
  port: 4444,
  cli_args: {
    'webdriver.chrome.driver': process.env.CHROME_DRIVER_PATH,
  },
};

/**
 * Put sel helper method into global object
 */
global.sel = sel;

module.exports = deepmerge({
  src_folders: [path.resolve('tests', 'e2e', 'specs')],
  output_folder: path.resolve('tests', 'e2e', 'reports'),
  custom_assertions_path: [
    path.resolve('node_modules', 'nightwatch-xhr', 'es5', 'assertions'),
    path.resolve('tests', 'e2e', 'custom-assertions'),
  ],
  custom_commands_path: [
    path.resolve('node_modules', 'nightwatch-xhr', 'es5', 'commands'),
    path.resolve('tests', 'e2e', 'custom-commands'),
  ],
  page_objects_path: [path.resolve('tests', 'e2e', 'page-objects')],
  globals_path: path.resolve('tests', 'e2e', 'globals.js'),

  selenium: seleniumConfig,

  test_settings: {
    default: {
      selenium_host: seleniumConfig.host,
      selenium_port: seleniumConfig.port,
      silent: true,

      videos: {
        fileName: 'test-result', // Required field
        nameAfterTest: true,
        format: 'mp4',
        enabled: false,
        deleteOnSuccess: true,
        path: path.resolve('tests', 'e2e', 'records'),
        resolution: '1440x900',
        fps: 15,
        input: '',
        videoCodec: 'libx264',
      },
    },

    chrome: {
      desiredCapabilities: {
        browserName: 'chrome',
        javascriptEnabled: true,
        acceptSslCerts: true,
      },
    },
  },
}, userOptions);
