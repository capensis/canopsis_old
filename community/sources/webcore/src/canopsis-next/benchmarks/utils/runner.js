const BenchmarkLauncher = require('./BenchmarkLauncher');
const ConsoleReporter = require('./reporters/ConsoleReporter');
const FileReporter = require('./reporters/FileReporter');

const benchmarkLauncher = new BenchmarkLauncher();

const runBenchmarks = async (options) => {
  const { jsonName } = options;

  await benchmarkLauncher.run(options);

  benchmarkLauncher.addReporter(
    new ConsoleReporter(),
    new FileReporter({ name: jsonName }),
  );

  benchmarkLauncher.report();
};

module.exports = {
  benchmark: benchmarkLauncher.benchmark,
  runBenchmarks,
};
