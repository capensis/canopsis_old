import { throttle } from 'lodash';

export const widgetColumnResizingAlarmMixin = {
  props: {
    resizingColumnThrottleDelay: {
      type: Number,
      default: 10,
    },
    minColumnWidth: {
      type: Number,
      default: 40,
    },
  },
  data() {
    return {
      resizingMode: false,
      resizingColumnIndex: null,
      columnsWidthByField: {},
    };
  },
  created() {
    this.aggregatedMovementDiff = 0;
    this.throttledResizeColumnByDiff = throttle(this.resizeColumnByDiff, this.resizingColumnThrottleDelay);
  },
  beforeDestroy() {
    this.finishColumnResize();
  },
  computed: {
    tableRow() {
      return this.tableHeader?.querySelector('tr:first-of-type');
    },

    headerCells() {
      return this.tableRow?.querySelectorAll('th') ?? [];
    },

    sumOfColumnsWidth() {
      return this.calculateFullColumnsWidth(this.columnsWidthByField);
    },
  },
  methods: {
    enableResizingMode() {
      this.resizingMode = true;

      this.calculateColumnsWidths();
    },

    disableResizingMode() {
      this.resizingMode = false;
    },

    toggleResizingMode() {
      return this.resizingMode ? this.disableResizingMode() : this.enableResizingMode();
    },

    setColumnsWidth(columnsWidth) {
      this.columnsWidthByField = { ...columnsWidth };
    },

    getColumnWidthByField(field) {
      return this.columnsWidthByField[field];
    },

    getNormalizedWidth(field, newWidth) {
      return Math.max(newWidth, this.minColumnWidth);
    },

    calculateFullColumnsWidth(columnsWidth) {
      return Object.values(columnsWidth).reduce((acc, width) => acc + width, 0);
    },

    calculateElementNormalizedWidth(element, field) {
      const { width } = element.getBoundingClientRect();

      return this.getNormalizedWidth(field, width);
    },

    calculateColumnsWidths() {
      this.columnsWidthByField = [...this.headerCells].reduce((acc, headerElement) => {
        if (headerElement.dataset?.value) {
          const { value } = headerElement.dataset;

          acc[value] = this.calculateElementNormalizedWidth(headerElement, value);
        }

        return acc;
      }, {});
    },

    resizeColumnByDiff(index) {
      const diff = this.aggregatedMovementDiff;

      if (!diff) {
        return;
      }

      const resizingLeftColumn = this.headers[index].value;
      const previousLeftColumnWidth = this.getColumnWidthByField(resizingLeftColumn);
      const newLeftColumnWidth = this.getNormalizedWidth(resizingLeftColumn, previousLeftColumnWidth + diff);

      this.columnsWidthByField = {
        ...this.columnsWidthByField,
        [resizingLeftColumn]: newLeftColumnWidth,
      };

      if (newLeftColumnWidth !== previousLeftColumnWidth) {
        this.aggregatedMovementDiff = 0;
      }
    },

    handleColumnResize(event) {
      this.aggregatedMovementDiff += event.movementX;

      this.throttledResizeColumnByDiff(this.resizingColumnIndex);

      this.setHorizontalScrollbarWidth();
      this.setHorizontalScrollPosition();
    },

    startColumnResize(columnName) {
      const body = document.querySelector('body');

      this.resizingColumnIndex = this.headers.findIndex(({ value }) => value === columnName);

      if (!body) {
        return;
      }

      body.addEventListener('mousemove', this.handleColumnResize);
      body.addEventListener('mouseup', this.finishColumnResize);
      body.addEventListener('mouseleave', this.finishColumnResize);
    },

    finishColumnResize() {
      this.aggregatedMovementDiff = 0;

      const body = document.querySelector('body');

      if (!body) {
        return;
      }

      body.removeEventListener('mousemove', this.handleColumnResize);
      body.removeEventListener('mouseup', this.finishColumnResize);
      body.removeEventListener('mouseleave', this.finishColumnResize);
    },
  },
};
