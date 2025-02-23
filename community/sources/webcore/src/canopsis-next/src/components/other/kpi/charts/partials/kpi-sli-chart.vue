<template>
  <bar-chart
    :datasets="datasets"
    :options="sliChartOptions"
    :dark="$system.dark"
  >
    <template #actions="{ chart }">
      <chart-export-actions
        :downloading="downloading"
        :chart="chart"
        class="mt-4"
        v-on="$listeners"
      />
    </template>
  </bar-chart>
</template>

<script>
import { COLORS } from '@/config';
import {
  DATETIME_FORMATS,
  KPI_SLI_GRAPH_BAR_PERCENTAGE,
  SAMPLINGS,
  KPI_SLI_GRAPH_DATA_TYPE,
  MAX_TIME_VALUE_BY_SAMPLING,
  TIME_UNITS_BY_SAMPLING,
} from '@/constants';

import { colorToRgba } from '@/helpers/color';
import { convertDurationToString, fromSeconds } from '@/helpers/date/duration';
import { getDateLabelBySampling } from '@/helpers/entities/metric/list';

import BarChart from '@/components/common/chart/bar-chart.vue';
import ChartExportActions from '@/components/common/chart/chart-export-actions.vue';

export default {
  inject: ['$system'],
  components: { ChartExportActions, BarChart },
  props: {
    metrics: {
      type: Array,
      default: () => [],
    },
    dataType: {
      type: String,
      default: KPI_SLI_GRAPH_DATA_TYPE.percent,
    },
    sampling: {
      type: String,
      default: SAMPLINGS.day,
    },
    responsive: {
      type: Boolean,
      default: false,
    },
    downloading: {
      type: Boolean,
      default: false,
    },
    minDate: {
      type: Number,
      required: false,
    },
  },
  computed: {
    maxValueBySampling() {
      return MAX_TIME_VALUE_BY_SAMPLING[this.sampling];
    },

    maxValueByType() {
      if (this.dataType === KPI_SLI_GRAPH_DATA_TYPE.percent) {
        return 100;
      }

      return this.maxValueBySampling;
    },

    samplingUnit() {
      return TIME_UNITS_BY_SAMPLING[this.sampling];
    },

    unit() {
      if (this.dataType === KPI_SLI_GRAPH_DATA_TYPE.percent) {
        return '%';
      }

      return this.samplingUnit;
    },

    datasets() {
      const { downtime, maintenance, uptime } = this.metrics.reduce((acc, metric) => {
        const x = metric.timestamp * 1000;

        acc.downtime.push({ x, y: this.convertValueBySamplingUnit(metric.downtime) });
        acc.maintenance.push({ x, y: this.convertValueBySamplingUnit(metric.maintenance) });
        acc.uptime.push({ x, y: this.convertValueBySamplingUnit(metric.uptime) });

        return acc;
      }, {
        downtime: [],
        maintenance: [],
        uptime: [],
      });

      return [{
        backgroundColor: colorToRgba(COLORS.kpi.uptime),
        barPercentage: KPI_SLI_GRAPH_BAR_PERCENTAGE,
        label: this.$t('common.uptime'),
        order: 1,
        data: uptime,
      }, {
        backgroundColor: colorToRgba(COLORS.kpi.downtime),
        barPercentage: KPI_SLI_GRAPH_BAR_PERCENTAGE,
        label: this.$t('common.downtime'),
        order: 2,
        data: downtime,
      }, {
        backgroundColor: colorToRgba(COLORS.kpi.maintenance),
        barPercentage: KPI_SLI_GRAPH_BAR_PERCENTAGE,
        label: this.$t('common.maintenance'),
        order: 3,
        data: maintenance,
      }];
    },

    sliChartOptions() {
      return {
        animation: false,
        responsive: this.responsive,
        scales: {
          x: {
            type: 'time',
            stacked: true,
            ticks: {
              min: this.minDate * 1000,
              max: Date.now(),
              source: 'data',
              font: {
                size: 11,
                family: 'Arial, sans-serif',
              },
              callback: this.getChartTimeTickLabel,
            },
          },
          y: {
            stacked: true,
            max: this.maxValueByType,
            beginAtZero: true,
            ticks: {
              font: {
                size: 11,
                family: 'Arial, sans-serif',
              },
              callback: this.getChartValueTickLabel,
            },
          },
        },
        interaction: {
          intersect: false,
          mode: 'index',
        },
        plugins: {
          legend: {
            position: 'top',
            align: 'end',
            labels: {
              boxWidth: 20,
              boxHeight: 20,
            },
          },
          tooltip: {
            itemSort: (a, b) => b.dataset.order - a.dataset.order,
            callbacks: {
              title: this.getChartTooltipTitle,
              label: this.getChartTooltipLabel,
            },
          },
        },
      };
    },
  },
  methods: {
    getChartValueTickLabel(value) {
      return `${value}${this.unit}`;
    },

    getChartTimeTickLabel(_, index, data) {
      const { value } = data[index] ?? {};

      return getDateLabelBySampling(value, this.sampling).split('\n');
    },

    getChartTooltipTitle(data) {
      const [dataset] = data;
      const { x: timestamp } = dataset.raw;

      return getDateLabelBySampling(timestamp, this.sampling);
    },

    getChartTooltipLabel(tooltip) {
      const { raw, dataset } = tooltip;

      const label = dataset.label.toLowerCase();

      if (this.dataType === KPI_SLI_GRAPH_DATA_TYPE.percent) {
        return `${raw.y}${this.unit} ${label}`;
      }

      const duration = convertDurationToString(
        raw.y,
        DATETIME_FORMATS.refreshFieldFormat,
        this.samplingUnit,
      );

      return `${duration} ${label}`;
    },

    convertValueBySamplingUnit(value) {
      if (this.dataType === KPI_SLI_GRAPH_DATA_TYPE.percent) {
        return value;
      }

      return fromSeconds(value, this.samplingUnit);
    },
  },
};
</script>
