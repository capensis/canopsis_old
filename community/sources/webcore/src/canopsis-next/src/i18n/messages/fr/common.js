import {
  ALARM_STATES,
  ALARM_STATUSES,
  PATTERN_FIELD_TYPES,
  PATTERN_OPERATORS,
  TRIGGERS_TYPES,
} from '@/constants';

export default {
  ok: 'Ok',
  undefined: 'Non défini',
  entity: 'Entité | Entités',
  alarm: 'Alarme | Alarmes',
  service: 'Service',
  widget: 'Widget',
  addWidget: 'Ajouter un widget',
  addTab: 'Ajouter un onglet',
  shareLink: 'Créer un lien de partage',
  addPbehavior: 'Ajouter un comportement périodique',
  refresh: 'Rafraîchir',
  toggleEditView: 'Activer/Désactiver le mode édition',
  toggleEditViewSubtitle: 'Si vous souhaitez enregistrer les positions des widgets, vous devez désactiver le mode édition',
  name: 'Nom',
  namePrefix: 'Préfixe du nom',
  description: 'Description',
  author: 'Auteur',
  submit: 'Soumettre',
  cancel: 'Annuler',
  continue: 'Continuer',
  stop: 'Fin',
  options: 'Options',
  type: 'Type',
  quitEditing: 'Quitter le mode édition',
  enabled: 'Activé(e)',
  disabled: 'Désactivé(e)',
  login: 'Connexion',
  yes: 'Oui',
  no: 'Non',
  default: 'Défaut',
  confirmation: 'Êtes-vous sûr(e) ?',
  parameter: 'Paramètre | Paramètres',
  by: 'Par',
  date: 'Date',
  comment: 'Commentaire | Commentaires',
  end: 'Fin',
  start: 'Début',
  message: 'Message',
  preview: 'Aperçu',
  recursive: 'Récursif',
  select: 'Sélectionner',
  states: 'Criticités',
  state: 'Criticité',
  sla: 'SLA',
  authors: 'Auteurs',
  stat: 'Statistiques',
  trend: 'Tendance',
  user: 'Utilisateur | Utilisateurs',
  role: 'Rôle | Rôles',
  import: 'Importer',
  export: 'Exporter',
  profile: 'Profil',
  username: 'Identifiant utilisateur',
  password: 'Mot de passe',
  authKey: 'Clé d\'authentification',
  widgetId: 'Identifiant du widget',
  connect: 'Connexion',
  optional: 'Optionnel',
  logout: 'Se déconnecter',
  title: 'Titre',
  save: 'Sauvegarder',
  label: 'Label',
  field: 'Champ',
  value: 'Valeur',
  limit: 'Limite',
  add: 'Ajouter',
  create: 'Créer',
  delete: 'Supprimer',
  show: 'Afficher',
  hide: 'Cacher',
  edit: 'Éditer',
  duplicate: 'Dupliquer',
  play: 'Lecture',
  copyLink: 'Copier le lien',
  parse: 'Analyser',
  home: 'Accueil',
  step: 'Étape',
  paginationItems: 'Affiche {first} à {last} sur {total} Entrées',
  apply: 'Appliquer',
  from: 'Depuis',
  to: 'Vers',
  tags: 'Tags',
  actionsLabel: 'Actions',
  noResults: 'Pas de résultats',
  result: 'Résultat',
  exploitation: 'Exploitation',
  administration: 'Administration',
  forbidden: 'Accès refusé',
  notFound: 'Introuvable',
  search: 'Recherche',
  filters: 'Filtres',
  filter: 'Filtre',
  emptyObject: 'Objet vide',
  startDate: 'Date de début',
  endDate: 'Date de fin',
  link: 'Lien | Liens',
  stack: 'Pile',
  edition: 'Édition',
  icon: 'Icône | Icônes',
  fullscreen: 'Plein écran',
  interval: 'Période',
  status: 'Statut',
  unit: 'Unité',
  delay: 'Délai',
  begin: 'Début',
  timezone: 'Fuseau horaire',
  reason: 'Raison',
  or: 'OU',
  and: 'ET',
  priority: 'Priorité',
  clear: 'Nettoyer',
  deleteAll: 'Tout supprimer',
  payload: 'Payload',
  note: 'Note',
  output: 'Message',
  created: 'Date de création',
  updated: 'Date de dernière modification',
  expired: 'Date d\'expiration',
  accessed: 'Consulté à',
  lastEventDate: 'Date du dernier événement',
  lastPbehaviorDate: 'Date du dernier comportement périodique',
  activated: 'Activé',
  pattern: 'Modèle | Modèles',
  correlation: 'Corrélation',
  periods: 'Périodes',
  range: 'Plage',
  duration: 'Durée',
  previous: 'Précédent',
  next: 'Suivant',
  eventPatterns: 'Modèles des événements',
  alarmPatterns: 'Modèles des alarmes',
  entityPatterns: 'Modèles des entités',
  pbehaviorPatterns: 'Modèles des comportements périodiques',
  totalEntityPatterns: 'Modèles des entités (Total)',
  serviceWeatherPatterns: 'Modèles des météos de service',
  addFilter: 'Ajouter un filtre',
  id: 'Identifiant',
  reset: 'Réinitialiser',
  selectColor: 'Sélectionner la couleur',
  disableDuringPeriods: 'Désactiver pendant les pauses',
  retryDelay: 'Délai entre les tentatives',
  retryUnit: 'Unité des tentatives',
  retryCount: 'Nombre de tentatives après échec',
  ticket: 'Ticket | Tickets',
  method: 'Méthode',
  url: 'URL',
  category: 'Catégorie',
  infos: 'Infos',
  impactLevel: 'Niveau d\'impact',
  impactState: 'État d\'impact',
  loadMore: 'Charger plus',
  download: 'Télécharger',
  initiator: 'Initiateur',
  percent: 'Pourcentage | Pourcentages',
  number: 'Nombre | Nombres',
  tests: 'Tests',
  total: 'Total',
  error: 'Erreur | Erreurs',
  failures: 'Échecs',
  skipped: 'Ignoré',
  current: 'Actuel',
  average: 'Moyenne',
  information: 'Information | Informations',
  file: 'Fichier',
  group: 'Groupe | Groupes',
  view: 'Vue | Vues',
  tab: 'Onglet | Onglets',
  access: 'Accès',
  communication: 'Communication | Communications',
  general: 'Général',
  notification: 'Notification | Notifications',
  dismiss: 'Rejeter',
  approve: 'Approuver',
  summary: 'Résumé',
  recurrence: 'Récurrence',
  statistics: 'Statistiques',
  action: 'Action | Actions',
  minimal: 'Minimal',
  optimal: 'Optimal',
  graph: 'Graphique | Graphiques',
  systemStatus: 'État du système',
  downloadAsPng: 'Télécharger en PNG',
  rating: 'Evaluation | Evaluations',
  sampling: 'Échantillonnage',
  parametersToDisplay: '{count} paramètres à afficher',
  uptime: 'Uptime',
  maintenance: 'Maintenance',
  downtime: 'Downtime',
  totalActiveTime: 'Temps d\'activité total',
  inactiveTime: 'Temps d\'inactivité',
  toTheTop: 'Vers le haut',
  time: 'Temps',
  lastModifiedOn: 'Dernière modification le',
  lastModifiedBy: 'Dernière modification par',
  exportAsCsv: 'Exporter en csv',
  exportToPdf: 'Exporter au format PDF',
  exportToJson: 'Exporter au format JSON',
  exportFieldToJson: 'Exporter {field} au format JSON',
  exportFieldToPdfOrJson: 'Exporter {field} au format PDF ou JSON',
  copyFieldPath: 'Copier le chemin {field}',
  copyFieldPathOrExportFieldToPdf: 'Copiez le chemin du {field} ou exportez le {field} au format PDF',
  copyToClipboard: 'Copier dans le presse-papier',
  copyPathToClipboard: 'Copier le chemin dans le presse-papiers',
  criteria: 'Critères',
  ratingSettings: 'Paramètres d\'évaluation',
  pbehavior: 'Comportement périodique | Comportements périodiques',
  activePbehavior: 'Activer le comportement périodique | Activer les comportements périodiques',
  searchBy: 'Recherché par',
  dictionary: 'Dictionnaire',
  condition: 'Condition | Conditions',
  template: 'Template',
  pbehaviorList: 'Lister les comportements périodiques',
  canceled: 'Annulé',
  snooze: 'Snooze',
  snoozed: 'En attente',
  impact: 'Impact | Impacts',
  depend: 'Depend | Depends',
  componentInfo: 'Information du composant | Informations du composant',
  connector: 'Type de connecteur',
  connectorName: 'Nom du connecteur',
  component: 'Composant',
  resource: 'Ressource',
  ack: 'Acquittement',
  acked: 'Acquitté',
  extraInfo: 'Extra info | Extra infos',
  custom: 'Personnalisé',
  eventType: 'Type d\'événement',
  sourceType: 'Type de Source',
  cycleDependency: 'Dépendance au cycle',
  checkPattern: 'Vérification du modèle',
  checkFilter: 'Vérifier le filtre',
  itemFound: '{count} élément trouvé | {count} éléments trouvés',
  canonicalType: 'Type canonique',
  map: 'Cartographie | Cartographies',
  instructions: 'Consignes',
  playlist: 'Liste de lecture | Listes de lecture',
  ctrlZoom: 'Utilisez ctrl + molette de la souris pour zoomer',
  calendar: 'Calendrier',
  tag: 'Tag | Tags',
  sharedTokens: 'Jetons partagés',
  notAvailable: 'Indisponible',
  addMore: 'Ajouter plus',
  more: 'plus',
  all: 'Tous',
  attribute: 'Attribut',
  timeTaken: 'Temps passé',
  enginesMetrics: 'Métriques des moteurs',
  failed: 'Échoué',
  close: 'Fermer',
  alarmId: 'Identifiant de l\'alarme',
  entityId: 'ID d\'entité',
  longOutput: 'Sortie longue',
  timestamp: 'Horodatage',
  countOfMax: '{count} sur {total}',
  trigger: 'Déclencheur | Déclencheurs',
  column: 'Colonne | Colonnes',
  countOfTotal: '{count} sur {total}',
  deprecatedTrigger: 'Ce déclencheur n\'est plus pris en charge',
  initialLongOutput: 'Sortie initiale longue',
  totalStateChanges: 'Nombre de changements d\'état',
  noReceivedEvents: 'Aucun événement reçu pendant {duration} par certaines dépendances',
  frequencyLimit: 'Nombre d\'oscillations',
  clearSearch: 'Ne plus appliquer cette recherche',
  noData: 'Aucune donnée',
  noColumns: 'Veuillez sélectionner au moins une colonne',
  theme: 'Thème | Thèmes',
  systemName: 'Nom du système',
  emitTrigger: 'Émettre un déclencheur',
  header: 'En-tête | En-têtes',
  headerKey: 'Clé d\'en-tête',
  headerValue: 'Valeur d\'en-tête',
  rule: 'Règle | Règles',
  copyValue: 'Copier la valeur',
  copyProperty: 'Copier la propriété',
  copyPropertyPath: 'Copier le chemin de la propriété',
  hidden: 'Caché',
  numberField: 'Champ numérique',
  chart: 'Graphique | Graphiques',
  currentDate: 'Date actuelle',
  chooseFile: 'Choisir le fichier',
  seeAlarms: 'Voir les alarmes',
  seeEntities: 'Voir les entités',
  new: 'Nouvelle',
  regexp: 'Expression régulière',
  archive: 'Archive',
  convertToCustomColumn: 'Convertir la colonne en personnalisé',
  event: 'Événement | Événements',
  showMore: 'Afficher plus ({current} sur {total})',
  availability: 'Disponibilité',
  testQuery: 'Tester la requête',
  webhookStatus: 'Statut du webhook',
  webhookComplete: 'Webhook terminé',
  noResponse: 'Pas de réponse',
  loadingItems: 'Chargement d\'éléments...',
  lastComment: 'Dernier commentaire',
  serialName: 'Nom de série',
  local: 'Local',
  server: 'Serveur',
  pressEnterToApply: 'Appuyez sur <kbd>Entrée</kbd> pour appliquer',
  width: 'Largeur',
  variableTypes: {
    string: 'Chaîne de caractères',
    number: 'Nombre',
    boolean: 'Booléen',
    null: 'Nul',
    array: 'Tableau',
    object: 'Object',
  },
  mixedField: {
    types: {
      [PATTERN_FIELD_TYPES.string]: '@:common.variableTypes.string',
      [PATTERN_FIELD_TYPES.number]: '@:common.variableTypes.number',
      [PATTERN_FIELD_TYPES.boolean]: '@:common.variableTypes.boolean',
      [PATTERN_FIELD_TYPES.null]: '@:common.variableTypes.null',
      [PATTERN_FIELD_TYPES.stringArray]: '@:common.variableTypes.array',
    },
  },
  saveChanges: 'Sauvegarder',
  ordinals: {
    first: 'Premier',
    second: 'Second',
    third: 'Troisième',
    fourth: 'Quatrième',
    fifth: 'Cinquième',
  },
  times: {
    millisecond: 'milliseconde | millisecondes',
    second: 'seconde | secondes',
    minute: 'minute | minutes',
    hour: 'heure | heures',
    day: 'jour | jours',
    week: 'semaine | semaines',
    month: 'mois | mois',
    year: 'année | années',
  },
  timeFrequencies: {
    secondly: 'Par seconde',
    minutely: 'Par minute',
    hourly: 'Par heure',
    daily: 'Quotidien',
    weekly: 'Hebdomadiare',
    monthly: 'Mensuel',
    yearly: 'Annuel',
  },
  weekDays: {
    monday: 'Lundi',
    tuesday: 'Mardi',
    wednesday: 'Mercredi',
    thursday: 'Jeudi',
    friday: 'Vendredi',
    saturday: 'Samedi',
    sunday: 'Dimanche',
  },
  shortWeekDays: {
    monday: 'lun.',
    tuesday: 'mar.',
    wednesday: 'mer.',
    thursday: 'jeu.',
    friday: 'ven.',
    saturday: 'sam.',
    sunday: 'dim.',
  },
  months: {
    january: 'Janvier',
    february: 'Février',
    march: 'Mars',
    april: 'Avril',
    may: 'Mai',
    june: 'Juin',
    july: 'Juillet',
    august: 'Août',
    september: 'Septembre',
    october: 'Octobre',
    november: 'Novembre',
    december: 'Décembre',
  },
  shortMonths: {
    january: 'janv.',
    february: 'févr.',
    march: 'mars',
    april: 'avr.',
    may: 'mai',
    june: 'juin',
    july: 'juil.',
    august: 'août',
    september: 'sep.',
    october: 'oct',
    november: 'nov.',
    december: 'déc.',
  },
  stateTypes: {
    [ALARM_STATES.ok]: 'Ok',
    [ALARM_STATES.minor]: 'Mineur',
    [ALARM_STATES.major]: 'Majeur',
    [ALARM_STATES.critical]: 'Critique',
  },
  statusTypes: {
    [ALARM_STATUSES.closed]: 'Fermée',
    [ALARM_STATUSES.ongoing]: 'En cours',
    [ALARM_STATUSES.flapping]: 'Bagot',
    [ALARM_STATUSES.stealthy]: 'Furtive',
    [ALARM_STATUSES.cancelled]: 'Annulée',
    [ALARM_STATUSES.noEvents]: 'Pas d\'événements',
  },
  operators: {
    [PATTERN_OPERATORS.equal]: 'Égal',
    [PATTERN_OPERATORS.contains]: 'Contient',
    [PATTERN_OPERATORS.notEqual]: 'n\'est pas égal',
    [PATTERN_OPERATORS.notContains]: 'Ne contient pas',
    [PATTERN_OPERATORS.beginsWith]: 'Commence par',
    [PATTERN_OPERATORS.notBeginWith]: 'Ne commence pas par',
    [PATTERN_OPERATORS.endsWith]: 'Se termine par',
    [PATTERN_OPERATORS.notEndWith]: 'Ne se termine pas par',
    [PATTERN_OPERATORS.exist]: 'Existe',
    [PATTERN_OPERATORS.notExist]: 'N\'existe pas',

    [PATTERN_OPERATORS.hasEvery]: 'Contient chaque',
    [PATTERN_OPERATORS.hasOneOf]: 'Contient l\'un des',
    [PATTERN_OPERATORS.isOneOf]: 'Est l\'un de',
    [PATTERN_OPERATORS.hasNot]: 'Ne contient pas',
    [PATTERN_OPERATORS.isNotOneOf]: 'N\'est pas l\'un des',
    [PATTERN_OPERATORS.isEmpty]: 'Est vide',
    [PATTERN_OPERATORS.isNotEmpty]: 'N\'est pas vide',

    [PATTERN_OPERATORS.higher]: 'Plus haut que',
    [PATTERN_OPERATORS.lower]: 'Plus bas que',

    [PATTERN_OPERATORS.longer]: 'Supérieur à',
    [PATTERN_OPERATORS.shorter]: 'Inférieur à',

    [PATTERN_OPERATORS.ticketAssociated]: 'Un ticket est associé',
    [PATTERN_OPERATORS.ticketNotAssociated]: 'Un ticket n\'est pas associé',

    [PATTERN_OPERATORS.canceled]: 'Annulé',
    [PATTERN_OPERATORS.notCanceled]: 'Non annulé',

    [PATTERN_OPERATORS.snoozed]: 'En attente',
    [PATTERN_OPERATORS.notSnoozed]: 'Non mis en attente',

    [PATTERN_OPERATORS.acked]: 'Acquitté',
    [PATTERN_OPERATORS.notAcked]: 'Non acquitté',

    [PATTERN_OPERATORS.isGrey]: 'Tuile grise',
    [PATTERN_OPERATORS.isNotGrey]: 'Tuile non grise',

    [PATTERN_OPERATORS.with]: 'Avec',
    [PATTERN_OPERATORS.without]: 'Sans',

    [PATTERN_OPERATORS.activated]: 'Activé',
    [PATTERN_OPERATORS.inactive]: 'Inactif',

    [PATTERN_OPERATORS.isMetaAlarm]: 'Est-ce une méta-alarme',
    [PATTERN_OPERATORS.isNotMetaAlarm]: 'Ce n\'est pas une méta-alarme',
    [PATTERN_OPERATORS.ruleIs]: 'La règle est',

    [PATTERN_OPERATORS.regexp]: 'Expression régulière',
  },
  triggers: {
    [TRIGGERS_TYPES.create]: {
      text: 'Création d\'alarme',
    },
    [TRIGGERS_TYPES.statedec]: {
      text: 'Diminution de la criticité',
    },
    [TRIGGERS_TYPES.changestate]: {
      text: 'Changement et verrouillage de la criticité',
    },
    [TRIGGERS_TYPES.stateinc]: {
      text: 'Augmentation de la criticité',
    },
    [TRIGGERS_TYPES.changestatus]: {
      text: 'Changement de statut (flapping, bagot, ...)',
    },
    [TRIGGERS_TYPES.ack]: {
      text: 'Acquittement d\'une alarme',
    },
    [TRIGGERS_TYPES.ackremove]: {
      text: 'Suppression de l\'acquittement d\'une alarme',
    },
    [TRIGGERS_TYPES.cancel]: {
      text: 'Annulation d\'une alarme',
    },
    [TRIGGERS_TYPES.uncancel]: {
      text: 'Annulation de l\'annulation d\'une alarme',
      helpText: 'L\'annulation ne peut se faire que par un événement posté sur l\'API',
    },
    [TRIGGERS_TYPES.comment]: {
      text: 'Commentaire sur une alarme',
    },
    [TRIGGERS_TYPES.declareticket]: {
      text: 'Déclaration de ticket depuis l\'interface graphique',
    },
    [TRIGGERS_TYPES.declareticketwebhook]: {
      text: 'Déclaration de ticket depuis un webhook',
    },
    [TRIGGERS_TYPES.assocticket]: {
      text: 'Association de ticket sur une alarme',
    },
    [TRIGGERS_TYPES.snooze]: {
      text: 'Mise en veille d\'une alarme',
    },
    [TRIGGERS_TYPES.unsnooze]: {
      text: 'Sortie de veille d\'une alarme',
    },
    [TRIGGERS_TYPES.resolve]: {
      text: 'Résolution d\'une alarme',
    },
    [TRIGGERS_TYPES.activate]: {
      text: 'Activation d\'une alarme',
    },
    [TRIGGERS_TYPES.pbhenter]: {
      text: 'Comportement périodique démarré',
    },
    [TRIGGERS_TYPES.pbhleave]: {
      text: 'Comportement périodique terminé',
    },
    [TRIGGERS_TYPES.instructionfail]: {
      text: 'Consigne manuelle en erreur',
    },
    [TRIGGERS_TYPES.autoinstructionfail]: {
      text: 'Consigne automatique en erreur',
    },
    [TRIGGERS_TYPES.instructionjobfail]: {
      text: 'Job de remédiation en erreur',
    },
    [TRIGGERS_TYPES.instructionjobcomplete]: {
      text: 'Job de remédiation terminé',
    },
    [TRIGGERS_TYPES.instructioncomplete]: {
      text: 'Consigne manuelle terminée',
    },
    [TRIGGERS_TYPES.autoinstructioncomplete]: {
      text: 'Consigne automatique terminée',
    },
    [TRIGGERS_TYPES.autoinstructionresultok]: {
      text: 'L\'alarme est en état OK après toutes les instructions automatiques',
    },
    [TRIGGERS_TYPES.autoinstructionresultfail]: {
      text: 'L\'alarme n\'est pas dans l\'état OK après toutes les instructions automatiques',
    },
    [TRIGGERS_TYPES.eventscount]: {
      text: 'L\'alarme a été reçue pour un certain nombre d\'événements',
      selectedText: 'L\'alarme a été reçue {additionalValue} événements',
      additionalFieldLabel: 'Nombre d\'événements',
    },
  },
  request: {
    timeout: 'Temps libre',
    timeoutSettings: 'Paramètres de délai d\'attente',
    repeatRequest: 'Répéter la demande',
    skipVerify: 'Ne pas vérifier les certificats HTTPS',
    headersHelpText: 'Sélectionnez la clé et la valeur de l\'en-tête ou saisissez-les manuellement',
    emptyHeaders: 'Aucun en-tête ajouté pour le moment',
    urlHelp: '<p>Les variables accessibles sont : <strong>.Alarm</strong>, <strong>.Entity</strong> et <strong>.Children</strong></p>'
      + '<i>Quelques exemples :</i>'
      + '<pre>"https://exampleurl.com?resource={{ .Alarm.Value.Resource }}"</pre>'
      + '<pre>"https://exampleurl.com?entity_id={{ .Entity.ID }}"</pre>'
      + '<pre>"https://exampleurl.com?children_count={{ len .Children }}"</pre>'
      + '<pre>"https://exampleurl.com?children={{ range .Children }}{{ .ID }}{{ end }}"</pre>',
  },

  fileSelector: {
    dragAndDrop: {
      label: 'Glisser-déposer pour télécharger',
      labelAction: 'ou parcourir',
    },
    fileTypes: {
      svg: 'fichier SVG',
    },
  },
};
