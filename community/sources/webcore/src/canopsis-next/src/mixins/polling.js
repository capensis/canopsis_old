/**
 * Mixin creator for polling
 *
 * @param {string} method
 * @param {string} [delayField = 'pollingDelay']
 * @returns {{
 *   data(): { timeout: null },
 *   methods: {
 *     polling(): Promise,
 *     stopPolling(): void
 *     stopPolling(): void
 *   },
 *   mounted(): void
 *   beforeDestroy(): void,
 * }}
 */
export const pollingMixinCreator = ({ method, delayField = 'pollingDelay' }) => ({
  data() {
    return {
      timeout: null,
    };
  },
  mounted() {
    this.startPolling();
  },
  beforeDestroy() {
    this.stopPolling();
  },
  methods: {
    async polling() {
      await this[method]();

      this.startPolling();
    },

    startPolling() {
      const delay = this[delayField];

      if (!delay) {
        return;
      }

      this.timeout = setTimeout(this.polling, delay);
    },

    stopPolling() {
      clearTimeout(this.timeout);

      this.timeout = null;
    },
  },
});
