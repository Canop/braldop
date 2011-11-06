Avertissement :
===============

Ce document est propriété Canop et toute modification externe sera refusée. Il est sur git pour la sauvegarde et la libre consultation, pas pour faire des demandes. Il reflète ce que j'ai besoin de noter plus que ce que j'ai réellement l'intention de faire.

Fait :
======

* maquettage lancement d'action au clavier
* stocker et rétablir l'état de la grille

En cours :
==========


P1 :
====


* garder le halo or pour les objets de la carte correspondant au menu
* BUG : mon code GO actuel n'efface pas les balises : si une balise est vue sur une case elle le reste (y compris dans la vue)
	-> mettre la date de vue dans VueEnvironnement et comparer à l'écriture dans MemMap, avec raz des routes, champs, échoppes, palissades et bosquets en cas de modif de l'environnement
* faire de jolis dessins pour les routes

* serveur avec login (mdp restreint) pour la réception des vues
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

