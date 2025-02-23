import { ALARM_LIST_STEPS, ALARM_METRIC_PARAMETERS } from '@/constants';

export default {
  liveReporting: 'Set a custom date range',
  ackAuthor: 'Ack author',
  lastCommentAuthor: 'Last comment author',
  lastCommentMessage: 'Last comment message',
  metaAlarm: 'Meta alarm',
  acknowledge: 'Acknowledge',
  ackResources: 'Ack resources',
  ackResourcesQuestion: 'Do you want to ack linked resources?',
  actionsRequired: 'Actions required',
  acknowledgeAndDeclareTicket: 'Acknowledge and declare ticket',
  acknowledgeAndAssociateTicket: 'Acknowledge and associate ticket',
  otherTickets: 'Other tickets are available in the expand panel',
  noAlarmFound: 'No alarms is found according to the patterns defined',
  associateTicketResources: 'Associate ticket for resources',
  followLink: 'Follow the link "{title}"',
  hasBookmark: 'Alarm has bookmark',
  filterByBookmark: 'Filter by bookmark',
  runTest: 'Run test',
  tagFilter: 'Tag filter',
  popups: {
    exportFailed: 'Failed to export alarms list in CSV format',
    addBookmarkSuccess: 'Bookmark was added',
    addBookmarkFailed: 'Something went wrong. Bookmark wasn\'t added',
    removeBookmarkSuccess: 'Bookmark was removed',
    removeBookmarkFailed: 'Something went wrong. Bookmark wasn\'t removed',
  },
  actions: {
    titles: {
      ack: 'Ack',
      fastAck: 'Fast ack',
      ackRemove: 'Cancel ack',
      pbehavior: 'Periodical behavior',
      fastPbehaviorAdd: 'Fast periodical behavior',
      snooze: 'Snooze alarm',
      declareTicket: 'Declare ticket',
      associateTicket: 'Associate ticket',
      cancel: 'Cancel alarm',
      unCancel: 'Uncancel alarm',
      fastCancel: 'Fast cancel',
      changeState: 'Change and lock severity',
      variablesHelp: 'List of available variables',
      history: 'History',
      linkToMetaAlarm: 'Link to a meta alarm',
      removeAlarmsFromManualMetaAlarm: 'Unlink alarm from manual meta alarm',
      removeAlarmsFromAutoMetaAlarm: 'Unlink alarm from meta alarm',
      comment: 'Comment',
      exportPdf: 'Export alarm to PDF file',
      addBookmark: 'Add bookmark',
      removeBookmark: 'Remove bookmark',
    },
    iconsTitles: {
      ack: 'Ack',
      declareTicket: 'Declare ticket',
      canceled: 'Canceled',
      snooze: 'Snooze',
      pbehaviors: 'Periodic behaviors',
      grouping: 'Meta alarm',
      comment: 'Comment',
    },
    iconsFields: {
      ticketNumber: 'Ticket number',
      parents: 'Causes',
      children: 'Consequences',
    },
  },
  timeline: {
    by: 'by',
    launched: 'launched',
    junit: 'JUnit',
    groupItems: 'Group items',
    steps: {
      [ALARM_LIST_STEPS.statedec]: 'State decreased',
      [ALARM_LIST_STEPS.stateinc]: 'State increased',
      [ALARM_LIST_STEPS.changeState]: 'State changed',
      [ALARM_LIST_STEPS.stateCounter]: 'Cropped severity (since the last change of the status)',
      [ALARM_LIST_STEPS.statusdec]: 'Status changed to {status}',
      [ALARM_LIST_STEPS.statusinc]: 'Status changed to {status}',
      [ALARM_LIST_STEPS.resolve]: 'Alarm resolved',
      [ALARM_LIST_STEPS.activate]: 'Alarm is active',
      [ALARM_LIST_STEPS.snooze]: 'Alarm snoozed for {duration}',
      [ALARM_LIST_STEPS.unsnooze]: 'Alarm unsnoozed',
      [ALARM_LIST_STEPS.ack]: 'Ack',
      [ALARM_LIST_STEPS.ackRemove]: 'Cancel ack',
      [ALARM_LIST_STEPS.pbhenter]: 'Periodical behavior activated',
      [ALARM_LIST_STEPS.pbhleave]: 'Periodical behavior deactivated',
      [ALARM_LIST_STEPS.assocTicket]: 'Ticket associated',
      [ALARM_LIST_STEPS.webhookStart]: 'Webhook is launched',
      [ALARM_LIST_STEPS.webhookInProgress]: 'Webhook launched by {author} in progress...',
      [ALARM_LIST_STEPS.webhookComplete]: 'Webhook is executed successfully',
      [ALARM_LIST_STEPS.webhookFail]: 'Webhook is failed',
      [ALARM_LIST_STEPS.declareTicket]: 'Ticket is declared',
      [ALARM_LIST_STEPS.declareTicketFail]: 'Ticket declaration is failed',
      [ALARM_LIST_STEPS.declareTicketRuleInProgress]: 'Ticket declaration rule launched by {author} in progress',
      [ALARM_LIST_STEPS.declareTicketRuleComplete]: 'Ticket declaration rule is complete',
      [ALARM_LIST_STEPS.declareTicketRuleFail]: 'Ticket declaration rule is failed',
      [ALARM_LIST_STEPS.instructionStart]: 'Instruction is launched',
      [ALARM_LIST_STEPS.instructionPause]: 'Instruction paused',
      [ALARM_LIST_STEPS.instructionResume]: 'Instruction resumed',
      [ALARM_LIST_STEPS.instructionComplete]: 'Instruction is complete',
      [ALARM_LIST_STEPS.instructionAbort]: 'Instruction is aborted',
      [ALARM_LIST_STEPS.instructionFail]: 'Instruction is failed',
      [ALARM_LIST_STEPS.autoInstructionStart]: 'Instruction is launched by the system',
      [ALARM_LIST_STEPS.autoInstructionComplete]: 'Instruction is complete',
      [ALARM_LIST_STEPS.autoInstructionFail]: 'Instruction is failed',
      [ALARM_LIST_STEPS.junitTestSuiteUpdate]: 'Test suite is relaunched',
      [ALARM_LIST_STEPS.junitTestCaseUpdate]: 'Test case is relaunched',
      [ALARM_LIST_STEPS.metaalarmattach]: 'Alarm linked to meta alarm',
      [ALARM_LIST_STEPS.metaalarmdetach]: 'Alarm unlinked from meta alarm',
      [ALARM_LIST_STEPS.comment]: 'Comment added',
    },
  },
  tabs: {
    moreInfos: 'More infos',
    timeLine: 'Timeline',
    charts: 'Charts',
    alarmsChildren: 'Alarms consequences',
    trackSource: 'Track source',
    impactChain: 'Impact chain',
    entityGantt: 'Gantt chart',
    ticketsDeclared: 'Tickets declared',
    remediation: 'Remediation',
  },
  moreInfos: {
    defineATemplate: 'To define a template for this window, go to the alarms list settings',
  },
  infoPopup: 'Info popup',
  tooltips: {
    priority: 'The priority parameter is derived from the alarm severity multiplied by impact level of the entity on which the alarm is raised',
    runningManualInstructions: 'Manual instruction <strong>{title}</strong> in progress | Manual instructions <strong>{title}</strong> in progress',
    runningAutoInstructions: 'Automatic instruction <strong>{title}</strong> in progress | Automatic instructions <strong>{title}</strong> in progress',
    successfulManualInstructions: 'Manual instruction <strong>{title}</strong> is successful | Manual instructions <strong>{title}</strong> is successful',
    successfulAutoInstructions: 'Automatic instruction <strong>{title}</strong> is successful | Automatic instructions <strong>{title}</strong> is successful',
    failedManualInstructions: 'Manual instruction <strong>{title}</strong> is failed | Manual instructions <strong>{title}</strong> is failed',
    failedAutoInstructions: 'Automatic instruction <strong>{title}</strong> is failed | Automatic instructions <strong>{title}</strong> is failed',
    hasManualInstruction: 'There is a manual instruction for this type of an incident | There are a manual instructions for this type of an incident',
    resetChangeColumns: 'Reset columns ordering/resizing',
    startChangeColumns: 'Start change columns ordering/resizing',
    finishChangeColumns: 'Finish change columns ordering/resizing',
  },
  metrics: {
    [ALARM_METRIC_PARAMETERS.createdAlarms]: 'Number of created alarms',
    [ALARM_METRIC_PARAMETERS.activeAlarms]: 'Number of active alarms',
    [ALARM_METRIC_PARAMETERS.nonDisplayedAlarms]: 'Number of non-displayed alarms',
    [ALARM_METRIC_PARAMETERS.instructionAlarms]: 'Number of alarms under automatic remediation',
    [ALARM_METRIC_PARAMETERS.pbehaviorAlarms]: 'Number of alarms with PBehavior',
    [ALARM_METRIC_PARAMETERS.correlationAlarms]: 'Number of alarms with correlation',
    [ALARM_METRIC_PARAMETERS.ackAlarms]: 'Total number of acks',
    [ALARM_METRIC_PARAMETERS.ackActiveAlarms]: 'Number of active alarms with acks',
    [ALARM_METRIC_PARAMETERS.cancelAckAlarms]: 'Number of canceled acks',
    [ALARM_METRIC_PARAMETERS.ticketActiveAlarms]: 'Number of active alarms with tickets',
    [ALARM_METRIC_PARAMETERS.withoutTicketActiveAlarms]: 'Number of active alarms without tickets',
    [ALARM_METRIC_PARAMETERS.ratioCorrelation]: '% of correlated alarms',
    [ALARM_METRIC_PARAMETERS.ratioInstructions]: '% of alarms with auto remediation',
    [ALARM_METRIC_PARAMETERS.ratioTickets]: '% of alarms with tickets created',
    [ALARM_METRIC_PARAMETERS.ratioNonDisplayed]: '% of non-displayed alarms',
    [ALARM_METRIC_PARAMETERS.ratioRemediatedAlarms]: '% of manually remediated alarms',
    [ALARM_METRIC_PARAMETERS.averageAck]: 'Average time to ack alarms',
    [ALARM_METRIC_PARAMETERS.averageResolve]: 'Average time to resolve alarms',
    [ALARM_METRIC_PARAMETERS.timeToAck]: 'Time to ack alarms',
    [ALARM_METRIC_PARAMETERS.timeToResolve]: 'Time to resolve alarms',
    [ALARM_METRIC_PARAMETERS.minAck]: 'Minimum time to ack alarms',
    [ALARM_METRIC_PARAMETERS.maxAck]: 'Maximum time to ack alarms',
    [ALARM_METRIC_PARAMETERS.manualInstructionExecutedAlarms]: 'Number of manually remediated alarms',
    [ALARM_METRIC_PARAMETERS.manualInstructionAssignedAlarms]: 'Number of alarms with manual instructions',
    [ALARM_METRIC_PARAMETERS.notAckedAlarms]: 'Number of not acked alarms',
    [ALARM_METRIC_PARAMETERS.notAckedInHourAlarms]: 'Number of not acked alarms with duration 1-4h',
    [ALARM_METRIC_PARAMETERS.notAckedInFourHoursAlarms]: 'Number of not acked alarms with duration 4-24h',
    [ALARM_METRIC_PARAMETERS.notAckedInDayAlarms]: 'Number of not acked alarms older than 24h',
    [ALARM_METRIC_PARAMETERS.minResolve]: 'Min time to resolve alarms',
    [ALARM_METRIC_PARAMETERS.maxResolve]: 'Max time to resolve alarms',
  },
  fields: {
    displayName: 'Display name',
    assignedInstructions: 'Assigned instructions',
    initialOutput: 'Initial output',
    initialLongOutput: 'Initial long output',
    lastCommentInitiator: 'Last comment initiator',
    lastCommentAuthor: 'Last comment author',
    ackBy: 'Acked by',
    ackMessage: 'Acked message',
    ackInitiator: 'Acked initiator',
    stateMessage: 'State message',
    statusMessage: 'Status message',
    totalStateChanges: 'Total state changes',
    stateInitiator: 'State initiator',
    stateValue: 'State value',
    statusValue: 'Status value',
    ackAt: 'Acked at',
    stateAt: 'State changed at',
    statusAt: 'Status changed at',
    resolved: 'Resolved at',
    activationDate: 'Activation date',
    currentStateDuration: 'Current state duration',
    snoozeDuration: 'Snooze duration',
    snoozeAuthor: 'Snooze author',
    snoozeInitiator: 'Snooze initiator',
    pbhInactiveDuration: 'Pbehavior inactive duration',
    activeDuration: 'Active duration',
    eventsCount: 'Events count',
    extraDetails: 'Extra details',
    ticketAuthor: 'Ticket submitter',
    ticketId: 'Ticket id',
    ticketMessage: 'Ticket message',
    ticketInitiator: 'Ticket initiator',
    ticketCreatedAt: 'Ticket created at',
    ticketData: 'Ticket data',
    alarmInfos: 'Alarm infos',
    entityName: 'Entity name',
    entityCategoryName: 'Entity category name',
    entityType: 'Entity type',
    entityComponent: 'Entity component',
    entityConnector: 'Entity connector',
    entityImpactLevel: 'Entity impact level',
    entityKoEvents: 'Entity KO events',
    entityOkEvents: 'Entity OK events',
    entityInfos: 'Entity infos',
    entityComponentInfos: 'Entity component infos',
    entityLastPbehaviorDate: 'Entity last pbehavior date',
    openedChildren: 'Opened consequences',
    closedChildren: 'Closed consequences',
    canceledInitiator: 'Canceled initiator',
    changeState: 'Change state',
  },
};
