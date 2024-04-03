import { EVENT_FILTER_ENRICHMENT_ACTIONS_TYPES, EVENT_FILTER_FAILURE_TYPES, EVENT_FILTER_TYPES } from '@/constants';

export default {
  externalData: 'Données externes',
  actionsRequired: 'Veuillez ajouter au moins une action',
  configRequired: 'Aucune configuration définie. Veuillez ajouter au moins un paramètre de configuration',
  idHelp: 'Si ce champ n\'est pas renseigné, un identifiant unique sera généré automatiquement à la création de la règle',
  editPattern: 'Éditer le modèle',
  advanced: 'Avancée',
  addAField: 'Ajouter un champ',
  simpleEditor: 'Éditeur simple',
  field: 'Champ',
  value: 'Valeur',
  advancedEditor: 'Éditeur avancé',
  comparisonRules: 'Règles de comparaison',
  editActions: 'Actions',
  addAction: 'Ajouter une action',
  editAction: 'Éditer une action',
  actions: 'Actions',
  onSuccess: 'En cas de succès',
  onFailure: 'En cas d\'échec',
  configuration: 'Configuration',
  resource: 'ID de ressource ou modèle',
  component: 'ID de composant ou modèle',
  connector: 'ID ou modèle de connecteur',
  connectorName: 'Nom ou modèle de connecteur',
  duringPeriod: 'Appliqué pendant cette période uniquement',
  enrichmentOptions: 'Options d\'enrichissement',
  changeEntityOptions: 'Modifier les options d\'entité',
  eventsFilteredSinceLastUpdate: 'Evénements filtrés depuis la dernière mise à jour',
  errorsSinceLastUpdate: 'Erreurs depuis la dernière mise à jour',
  markAsRead: 'Marquer comme lu',
  filterByType: 'Filtrer par type',
  copyEventToClipboard: 'Copier l\'événement dans le presse papier',
  event: 'Evénement',
  eventCopied: 'Evénement copié dans le presse papier',
  syntaxIsValid: 'La syntaxe est valide',
  types: {
    [EVENT_FILTER_TYPES.drop]: 'Suppression',
    [EVENT_FILTER_TYPES.break]: 'Break',
    [EVENT_FILTER_TYPES.enrichment]: 'Enrichissement',
    [EVENT_FILTER_TYPES.changeEntity]: 'Changement d\'entité',
  },
  failureTypes: {
    [EVENT_FILTER_FAILURE_TYPES.invalidPattern]: 'Pattern invalide',
    [EVENT_FILTER_FAILURE_TYPES.invalidTemplate]: 'Template invalide',
    [EVENT_FILTER_FAILURE_TYPES.externalDataMongo]: 'Mongo',
    [EVENT_FILTER_FAILURE_TYPES.externalDataApi]: 'API externe',
    [EVENT_FILTER_FAILURE_TYPES.other]: 'Autre',
  },
  tooltips: {
    addValueRuleField: 'Ajouter une règle',
    editValueRuleField: 'Éditer la règle',
    addObjectRuleField: 'Ajouter un groupe de règles',
    editObjectRuleField: 'Éditer le groupe de règles',
    removeRuleField: 'Supprimer le groupe/la règle',
  },
  validation: {
    incorrectRegexOnSetTagsValue: 'Valeur non valide : la valeur de l\'action set_tags doit contenir une expression régulière pour extraire les groupes <name> et <value>',
  },
  actionsTypes: {
    [EVENT_FILTER_ENRICHMENT_ACTIONS_TYPES.copy]: {
      text: 'Copier une valeur d\'un champ d\'événement à un autre',
      message: 'Cette action permet de copier la valeur ou une paire clé+valeur d\'un contrôle dans un événement.',
      description: 'Les paramètres de l\'action sont :\n- valeur : le nom du champ dont la valeur doit être copiée. Il peut s\'agir d\'un champ d\'événement, d\'un sous-groupe d\'une expression régulière ou d\'une donnée externe.\n- description (facultatif) : la description.\n- nom : le nom du champ événement dans lequel la valeur doit être copiée.',
    },
    [EVENT_FILTER_ENRICHMENT_ACTIONS_TYPES.copyToEntityInfo]: {
      text: 'Copier une valeur d\'un champ d\'un événement vers une information d\'une entité',
      message: 'Cette action est utilisée pour copier la valeur du champ d\'un événement dans le champ d\'une entité.',
      description: 'Les paramètres de l\'action sont :\n- description (optionnel) : la description.\n- nom : le nom du champ d\'une entité.\n- valeur : le nom du champ dont la valeur doit être copiée. Il peut s\'agir d\'un champ d\'événement, d\'un sous-groupe d\'une expression régulière ou d\'une donnée externe.',
    },
    [EVENT_FILTER_ENRICHMENT_ACTIONS_TYPES.setEntityInfo]: {
      text: 'Définir une information d\'une entité sur une constante',
      message: 'Cette action permet de définir les informations dynamiques d\'une entité correspondant à l\'événement.',
      description: 'Les paramètres de l\'action sont :\n- description (optionnel) : la description.\n- nom : le nom du champ.\n- valeur : la valeur d\'un champ.',
    },
    [EVENT_FILTER_ENRICHMENT_ACTIONS_TYPES.setEntityInfoFromTemplate]: {
      text: 'Définir une chaîne d\'informations sur une entité à l\'aide d\'un modèle',
      message: 'Cette action permet de modifier les informations dynamiques d\'une entité correspondant à l\'événement.',
      description: 'Les paramètres de l\'action sont :\n- description (optionnel) : la description\n- nom : le nom du champ.\n- valeur : le modèle utilisé pour déterminer la valeur de la donnée.\nDes modèles {{.Event.NomDuChamp}}, des expressions régulières ou des données externes peuvent être utilisés.',
    },
    [EVENT_FILTER_ENRICHMENT_ACTIONS_TYPES.setField]: {
      text: 'Définir un champ d\'un événement sur une constante',
      message: 'Cette action peut être utilisée pour modifier un champ de l\'événement.',
      description: 'Les paramètres de l\'action sont :\n- description (optionnel) : la description\n- nom : le nom du champ.\n- valeur : la nouvelle valeur du champ.',
    },
    [EVENT_FILTER_ENRICHMENT_ACTIONS_TYPES.setFieldFromTemplate]: {
      text: 'Définir un champ de chaîne d\'un événement à l\'aide d\'un modèle',
      message: 'Cette action vous permet de modifier un champ d\'événement à partir d\'un modèle.',
      description: 'Les paramètres de l\'action sont :\n- description (optionnel) : la description\n- nom : le nom du champ.\n- valeur : le modèle utilisé pour déterminer la valeur du champ.\n  Des modèles {{.Event.NomDuChamp}}, des expressions régulières ou des données externes peuvent être utilisés.',
    },
    [EVENT_FILTER_ENRICHMENT_ACTIONS_TYPES.setTags]: {
      text: 'Définir les balises d\'un champ à l\'aide d\'une correspondance d\'expression rationnelle',
      message: 'Cette action peut être utilisée pour définir des balises provenant d\'autres événements filtrés à l\'aide d\'une correspondance d\'expression rationnelle.',
      description: 'TDes ags au format « Nom : Valeur » peuvent être définis à l\'aide de cette action.\nLe filtre d\'événement avec un champ filtré par correspondance d\'expression régulière avec les variables de nom et de valeur doit être appliqué en premier.\nExemples d\'expressions régulières :\n<ul><li><code>(?P&lt;value&gt;[a-zA-Z]+)\\\\s+(?P&lt;name&gt;[a-zA-Z]+);</code> - pour les chaînes qui contient des tableaux de <code>value name;</code> divisé par <code>;</code></li><li><code>(?P&lt;name&gt;[a-zA-Z]+)\\ \\s+(?P&lt;value&gt;[a-zA-Z]+);</code> - pour les chaînes qui contiennent des tableaux de <code>name value;</code> divisé par <code>;</code>< /li><li><code>(?P&lt;name&gt;[a-zA-Z]+):\\s+(?P&lt;value&gt;[a-zA-Z]+);</code> - pour les chaînes qui contient des tableaux de <code>name: value;</code> divisé par <code>;</code></li></ul>Les paramètres de l\'action sont :\n- description (facultatif) : la description.\n- valeur : le champ du filtre d\'événement à partir duquel le nom et la valeur sont récupérés.',
    },
    [EVENT_FILTER_ENRICHMENT_ACTIONS_TYPES.setTagsFromTemplate]: {
      text: 'Définir les balises d\'un champ à l\'aide d\'un modèle',
      message: 'Cette action peut être utilisée pour définir des balises provenant d\'autres champs d\'événement à l\'aide d\'un modèle.',
      description: 'Des balises au format « Nom : Valeur » peuvent être définies à l\'aide de cette action.\nLes paramètres de l\'action sont :\n- description (facultatif) : la description.\n- name : le nom du groupe de tags\n- valeur : le modèle utilisé pour déterminer la valeur de la donnée.\nDes modèles {{.Event.NomDuChamp}}, des expressions régulières ou des données externes peuvent être utilisés.',
    },
  },
};
