import { SERVICE_WEATHER_STATE_COUNTERS, WEATHER_ACTIONS_TYPES } from '@/constants';

export default {
  grey: 'Gris',
  hideGrey: 'Cacher les tuiles grises',
  primaryIcon: 'Icône principale',
  secondaryIcon: 'Icône secondaire',
  massActions: 'Actions de masse',
  cannotBeApplied: 'Cette action ne peut pas être appliquée',
  actions: {
    [WEATHER_ACTIONS_TYPES.entityAck]: 'Acquitter',
    [WEATHER_ACTIONS_TYPES.entityAckRemove]: 'Supprimer l\'acquittement',
    [WEATHER_ACTIONS_TYPES.entityValidate]: 'Valider',
    [WEATHER_ACTIONS_TYPES.entityInvalidate]: 'Invalider',
    [WEATHER_ACTIONS_TYPES.entityPause]: 'Pause',
    [WEATHER_ACTIONS_TYPES.entityPlay]: 'Supprimer la pause',
    [WEATHER_ACTIONS_TYPES.entityCancel]: 'Annuler',
    [WEATHER_ACTIONS_TYPES.entityAssocTicket]: 'Associer un ticket',
    [WEATHER_ACTIONS_TYPES.entityComment]: 'Commenter l\'alarme',
    [WEATHER_ACTIONS_TYPES.executeInstruction]: 'Exécuter la consigne',
    [WEATHER_ACTIONS_TYPES.declareTicket]: 'Déclarer un incident',
  },
  iconTypes: {
    ok: 'Ok',
    minorOrMajor: 'Mineure ou Majeure',
    critical: 'Critique',
  },
  stateCounters: {
    [SERVICE_WEATHER_STATE_COUNTERS.all]: 'Nombre d\'alarmes',
    [SERVICE_WEATHER_STATE_COUNTERS.active]: 'Nombre d\'alarmes actives',
    [SERVICE_WEATHER_STATE_COUNTERS.depends]: 'Nombre de dépendances',
    [SERVICE_WEATHER_STATE_COUNTERS.ok]: 'Ok',
    [SERVICE_WEATHER_STATE_COUNTERS.minor]: 'Mineure',
    [SERVICE_WEATHER_STATE_COUNTERS.major]: 'Majeur',
    [SERVICE_WEATHER_STATE_COUNTERS.critical]: 'Critique',
    [SERVICE_WEATHER_STATE_COUNTERS.acked]: 'Acquittées',
    [SERVICE_WEATHER_STATE_COUNTERS.unacked]: 'Non acquittées',
    [SERVICE_WEATHER_STATE_COUNTERS.underPbehavior]: 'Sous PBh',
    [SERVICE_WEATHER_STATE_COUNTERS.ackedUnderPbehavior]: 'Acquittées sous PBh',
  },
  stateCountersTooltips: {
    [SERVICE_WEATHER_STATE_COUNTERS.all]: 'alarmes au total',
    [SERVICE_WEATHER_STATE_COUNTERS.active]: 'alarmes actives',
    [SERVICE_WEATHER_STATE_COUNTERS.depends]: 'dépendances',
    [SERVICE_WEATHER_STATE_COUNTERS.ok]: 'état OK',
    [SERVICE_WEATHER_STATE_COUNTERS.minor]: 'alarmes mineures',
    [SERVICE_WEATHER_STATE_COUNTERS.major]: 'alarmes majeures',
    [SERVICE_WEATHER_STATE_COUNTERS.critical]: 'alarmes critiques',
    [SERVICE_WEATHER_STATE_COUNTERS.acked]: 'alarmes acquittées',
    [SERVICE_WEATHER_STATE_COUNTERS.unacked]: 'non acquittées',
    [SERVICE_WEATHER_STATE_COUNTERS.underPbehavior]: 'sous PBh',
    [SERVICE_WEATHER_STATE_COUNTERS.ackedUnderPbehavior]: 'acquittées sous PBh',
  },
};
