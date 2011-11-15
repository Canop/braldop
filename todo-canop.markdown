Avertissement :
===============

Ce document est propriété Canop et toute modification externe sera refusée. Il est sur git pour la sauvegarde et la libre consultation, pas pour faire des demandes. Il reflète ce que j'ai besoin de noter plus que ce que j'ai réellement l'intention de faire.

Fait :
======

* structure de profiling du code go

En cours :
==========

* pb ascenseur menu dans interface scrollable & molette https://github.com/braldahim/braldahim/issues/105

P1 :
====

* pb de flash dans initBraldopFetch https://github.com/braldahim/braldahim/blob/master/braldahim-static/public/js/braldop/main.js

* ne pas transmettre les env en json mais les lire dans la carte png (en comportement optionnel pour la compatibilité Braldahim)
* serveur avec login (mdp restreint) pour la réception des vues

* BUG : mon code GO actuel n'efface pas les balises : si une balise est vue sur une case elle le reste (y compris dans la vue)
	-> mettre la date de vue dans VueEnvironnement et comparer à l'écriture dans MemMap, avec raz des routes, champs, échoppes, palissades et bosquets en cas de modif de l'environnement
* faire de jolis dessins pour les routes

* envoi des vues (json) depuis l'extension braldop
* ajout d'un onglet pour la carte grace à l'extension

* extension : système de chat avec les bralduns en vue (avec invitations possibles d'autres bralduns)
* ordonner les profondeurs dans le menu de sélection
* protection par mot de passe et interface de sélection de la carte
* intégrer les nouveaux dessins (carottes, au moins)
* essayer d'améliorer le dessin des portails
* essayer coffeescript
* gèrer le cas de 3 types de monstres sur une case sans braldun (j'attends que la situation se présente)
* PB : certaines icônes, en particulier les braldûns, sont floues sur Firefox
* icônes PVE/PVP devant les noms des régions
* dialogue/bâtiments : lien vers l'aide du bâtiment

P2 :
====

* enlever le z inutile dans les objets palissade, champ, échoppe, etc.
* tag (optionnel) de cellule pour indiquer les coordonnées
* factoriser les données des bralduns, pour ne pas répéter partout leur nom (ne mettre dans les vues, champs, échoppes, que les numéros)
* compression du json (utile ? si mod_deflate sur le serveur, ça sera compressé génériquement)

