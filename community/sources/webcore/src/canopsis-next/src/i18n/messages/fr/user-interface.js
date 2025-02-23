export default {
  title: 'Interface utilisateur',
  appTitle: 'Titre de l\'application',
  language: 'Langue par défaut',
  footer: 'Page d\'identification : pied de page',
  description: 'Page d\'identification : description',
  logo: 'Logo',
  infoPopupTimeout: 'Délai d\'affichage pour les popups d\'informations',
  errorPopupTimeout: 'Délai d\'affichage pour les popups d\'erreurs',
  allowChangeSeverityToInfo: 'Autorise le changement de criticité en Info',
  showHeaderOnKioskMode: 'Afficher l\'en-tête en mode kiosque',
  maxMatchedItems: 'Seuil d\'éléments avant avertissement',
  checkCountRequestTimeout: 'Délai d\'expiration de la requête',
  requiredInstructionApprove: 'Approbation des consignes requise',
  disabledTransitions: 'Désactivez certaines animations',
  disabledTransitionsTooltip: 'Cela aidera à améliorer les performances de l\'application',
  autoSuggestPbehaviorName: 'Proposer automatiquement un nom de comportement périodique',
  defaultTheme: 'Thème par défaut',
  versionDescriptionTooltip: 'Info-bulle contextuelle sur la bannière de version',
  defaultVersionDescription: `<div>
  <div>@:common.edition:<strong> {{ edition }}</strong></div>
  <div>@:common.serialName:<strong> {{ serialName }}</strong></div>
  <div>@:common.versionUpdated:<strong> {{ timestamp versionUpdated }}</strong></div>
</div>`,
  tooltips: {
    maxMatchedItems: 'Avertit l\'utilisateur lorsque le nombre d\'éléments correspondant aux modèles est supérieur à cette valeur',
    checkCountRequestTimeout: 'Définit le délai d\'expiration (en secondes) de la demande d\'éléments correspondants',
  },
};
