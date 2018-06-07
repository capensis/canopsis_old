export default {
  common: {
    submit: 'Soumettre',
    login: 'Connexion',
    username: 'Nom d\'utilisateur',
    password: 'Mot de passe',
    title: 'Titre',
    save: 'Sauvegarder',
    label: 'Label',
    value: 'Valeur',
    add: 'Ajouter',
    delete: 'Supprimer',
    edit: 'Éditer',
    parse: 'Compiler',
    home: 'Accueil',
    entries: 'entrées',
    showing: 'afficher',
    to: 'à',
    of: 'de',
    actionsLabel: 'Actions',
    noResults: 'Pas de résultats',
    actions: {
      close: 'Fermer',
      acknowledge: 'Acquitter',
      acknowledgeAndReport: 'Acquitter et signaler un incident',
      saveChanges: 'Sauvegarder',
      reportIncident: 'Signaler un incident',
    },
    times: {
      second: 'seconde | secondes',
      minute: 'minute | minutes',
      hour: 'heure | heures',
      day: 'jour | jours',
      week: 'semaine | semaines',
      month: 'mois | mois',
      year: 'année | années',
    },
  },
  alarmList: {
    actions: {
      ack: 'ack',
      fastAck: 'Ack rapide',
      ackRemove: 'Annuler ack',
      pbehavior: 'Comportement périodique',
      snooze: 'Snooze',
      pbehaviorList: 'Lister comportements pédioriques',
      declareTicket: 'Signaler ticket',
      associateTicket: 'Associer ticket',
      cancel: 'Annuler alarme',
      changeState: 'Changer criticité',
    },
  },
  alarmListSettings: {
    alarmListSettings: 'Paramètres du bac à alarmes',
    widgetTitle: 'Titre du widget',
    columnName: 'Nom de la colonne',
    defaultSortColumn: 'Colonne de tri par défaut',
    columnNames: 'Nom des colonnes',
    periodicRefresh: 'Rafraichissement périodique',
    defaultNumberOfElementsPerPage: 'Nombre d\'élement/page par défaut',
    elementsPerPage: 'Élements par page',
    filterOnOpenResolved: 'Filtre sur Open/Resolved',
    open: 'Ouverte',
    resolved: 'Résolue',
    filters: 'Filtres',
    selectAFilter: 'Sélectionner un filtre',
    infoPopup: 'Info popup',
  },
  modals: {
    createAckEvent: {
      title: 'Ajouter un événement: ack',
      tooltips: {
        ackResources: 'Voulez-vous acquitter les ressources liées ?',
      },
      fields: {
        ticket: 'Numéro du ticket',
        output: 'Note',
        ackResources: 'Ack ressources',
      },
    },
    createSnoozeEvent: {
      title: 'Ajouter un événement: snooze',
      fields: {
        duration: 'Durée',
      },
    },
    createCancelEvent: {
      title: 'Ajouter un événement: Annuler',
      fields: {
        output: 'Note',
      },
    },
    createChangeStateEvent: {
      title: 'Ajouter un événement: Changer état',
      states: {
        ok: 'Info',
        minor: 'Mineur',
        major: 'Majeur',
        critical: 'Critique',
      },
      fields: {
        output: 'Note',
      },
    },
    createPbehavior: {
      title: 'Ajoute run comportement périodique à ces éléments ?',
      fields: {
        name: 'Nom',
        start: 'Début',
        stop: 'Fin',
        reason: 'Raison',
        type: 'Type',
        rRuleQuestion: 'Ajouter une rrule à ce comportement périodique ?',
      },
    },
    createAckRemove: {
      title: 'Ajouter un événement: ackremove',
    },
    createDeclareTicket: {
      title: 'Ajouter un événement: declareticket',
    },
    createAssociateTicket: {
      title: 'Ajouter un événement: associateticket',
      fields: {
        ticket: 'Numéro du ticket',
      },
    },
  },
  tables: {
    alarmGeneral: {
      title: 'Generale',
      author: 'Auteur',
      connector: 'Connecteur',
      component: 'Composant',
      resource: 'Ressource',
      output: 'Output',
      lastUpdateDate: 'Date de dernière modification',
      creationDate: 'Date de création',
      duration: 'Durée',
      state: 'État',
      status: 'Status',
    },
    /**
     * This object for pbehavior fields from database
     */
    pbehaviorList: {
      name: 'Nom',
      author: 'Auteur',
      connector: 'Connecteur',
      connector_name: 'Nom du connecteur',
      enabled: 'Actif',
      tstart: 'Démarre',
      tstop: 'Finis',
      type_: 'Type',
      reason: 'Raison',
      rrule: 'Rrule',
    },
  },
  rRule: {
    advancedHint: 'Séparer les nombres par une virgule',
    textLabel: 'Rrule',
    stringLabel: 'Résumé',
    tabs: {
      simple: 'Simple',
      advanced: 'Avancé',
    },
    errors: {
      main: 'La Rrule choisis n\'est pas valide. Nous vous recommandons de la modifier avant de sauvegarder',
    },
    fields: {
      freq: 'Fréquence',
      until: 'Jusqu\'à',
      byweekday: 'Par jour de la semaine',
      count: 'Répéter',
      interval: 'Interval',
      wkst: 'Semaine de début',
      bymonth: 'Par mois',
      bysetpos: {
        label: 'Par position',
        tooltip: 'Si renseigné, doit être un ou plusieurs nombres entiers, positifs ou négatifs. Chaque entier correspondra à la n-ième occurence de la règle dans l\'interval de fréquence. Par exemple, une \'bysetpos\' de -1 combinée à une fréquence mensuelle, et une \'byweekday\' de (Lundi, Mardi, Mercredi, Jeudi, Vendredi), va nous donner le dernier jour travaillé de chaque mois',
      },
      bymonthday: {
        label: 'Par jour du mois',
        tooltip: 'Si renseigné, doit être un ou plusieurs nombres entiers, correspondant aux jours du mois auquel appliquer la récurrence.',
      },
      byyearday: {
        label: 'Par jour de l\'année',
        tooltip: 'Si renseigné, doit être un ou plusieurs nombres entiers, correspondant aux jours de l\'année auquel appliquer la récurrence.',
      },
      byweekno: {
        label: 'Par semaine n°',
        tooltip: 'Si renseigné, doit être un ou plusieurs nombres entiers, correspondant aux numéros de semaine a laquelle appliquer la récurrence. Les numéros de semaines sont ceux de ISO8601, la première semaine de l\'année étant celle contenant au moins 4 jours de cette année',
      },
      byhour: {
        label: 'Par heure',
        tooltip: 'Si renseigné, doit être un ou plusieurs nombres entiers, correspondant aux heures auquel appliquer la récurrence.',
      },
      byminute: {
        label: 'Par minute',
        tooltip: 'Si renseigné, doit être un ou plusieurs nombres entiers, correspondant aux minutes auquel appliquer la récurrence.',
      },
      bysecond: {
        label: 'Par seconde',
        tooltip: 'Si renseigné, doit être un ou plusieurs nombres entiers, correspondant aux secondes auquel appliquer la récurrence.',
      },
    },
  },
  errors: {
    default: 'Une erreur s\'est produite...',
  },
  mFilterEditor: {
    tabs: {
      visualEditor: 'Éditeur visuel',
      advancedEditor: 'Éditeur avancé',
      results: 'Résultats',
    },
    buttons: {
      addRule: 'Ajouter une règle',
      addGroup: 'Ajouter un groupe',
      deleteGroup: 'Supprimer un groupe',
    },
    resultsTableHeaders: {
      connector: 'Connecteur',
      connectorName: 'Nom du connecteur',
      component: 'Composant',
      resource: 'Ressource',
    },
  },
};
