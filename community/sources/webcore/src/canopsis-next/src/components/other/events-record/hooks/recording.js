import { EVENT_FILTER_PATTERN_FIELDS, MODALS } from '@/constants';

import { promisedWait } from '@/helpers/async';

import { useI18n } from '@/hooks/i18n';
import { useModals } from '@/hooks/modals';
import { useEventsRecordCurrent } from '@/hooks/store/modules/events-record-current';

/**
 * Hook to manage event recording functionality.
 *
 * @param {Function} [fetchListHandler = () => {}] - A function to fetch the list of events records.
 * @returns {Object} An object containing methods to start and stop event recording.
 */
export const useEventsRecordRecording = (fetchListHandler = () => {}) => {
  const { t } = useI18n();
  const modals = useModals();

  const { startEventsRecordCurrent, stopEventsRecordCurrent } = useEventsRecordCurrent();

  /**
   * Start recording events based on a specified event pattern.
   */
  const startRecording = () => modals.show({
    name: MODALS.applyEventFilter,
    config: {
      title: t('eventsRecord.launchEventRecording'),
      excludedAttributes: [
        { value: EVENT_FILTER_PATTERN_FIELDS.eventType },
        { value: EVENT_FILTER_PATTERN_FIELDS.state },
        { value: EVENT_FILTER_PATTERN_FIELDS.sourceType },
        { value: EVENT_FILTER_PATTERN_FIELDS.output },
        { value: EVENT_FILTER_PATTERN_FIELDS.extraInfos },
        { value: EVENT_FILTER_PATTERN_FIELDS.longOutput },
        { value: EVENT_FILTER_PATTERN_FIELDS.author },
        { value: EVENT_FILTER_PATTERN_FIELDS.initiator },
      ],
      action: eventPattern => startEventsRecordCurrent({ data: { event_pattern: eventPattern } }),
      afterSubmit: fetchListHandler,
    },
  });

  /**
   * Stop the current event recording.
   */
  const stopRecording = () => modals.show({
    name: MODALS.confirmation,
    config: {
      action: async () => {
        await stopEventsRecordCurrent();

        /**
         * We've added that to avoiding problem with async on the backend side.
         * There is 3000ms timeout on the backend side for sync
         */
        await promisedWait(3000);

        return fetchListHandler();
      },
    },
  });

  return {
    startRecording,
    stopRecording,
  };
};
