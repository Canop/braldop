Avertissement :
===============

Ce document est propriété Canop et toute modification externe sera refusée. Il est sur git pour la sauvegarde et la libre consultation, pas pour faire des demandes. Il reflète ce que j'ai besoin de noter plus que ce que j'ai réellement l'intention de faire.

Fait :
======

En cours :
==========

* pb ascenseur menu dans interface scrollable & molette https://github.com/braldahim/braldahim/issues/105
* BUG extension : problème à la distribution de px (Braldahim.js ligne 96)


P1 :
====

* pb : je ne sais pas indiquer si une palissade dans le brouillard est destructible ou non

* bradmin : fonction de fusion d'images

* gestion de groupes dans l'extension
* mécanisme de partage de cartes au sein du groupe
* mécanisme de partage de vues au sein du groupe
* stocker (mysql ?) les échoppes et autres lieux (ruines, buissons, ?) et leur connaissance par les braldûns

* bug extension : on n'affiche pas les caractéristiques des terrains pour les cases sous le brouillard

* bouton pour activer la session sans se délogguer

* faire de jolis dessins pour les routes

* ordonner les profondeurs dans le menu de sélection
* essayer d'améliorer le dessin des portails
* gèrer le cas de 3 types de monstres sur une case sans braldun (j'attends que la situation se présente)
* PB : certaines icônes, en particulier les braldûns, sont floues sur Firefox
* icônes PVE/PVP devant les noms des régions
* dialogue/bâtiments : lien vers l'aide du bâtiment

P2 :
====

* utiliser webp au lieu de png lorsque ce sera intégré dans firefox (vérifier performances) [ goinstall vp8-go.googlecode.com/hg/webp n'a pas l'air prêt ]
* enlever le z inutile dans les objets palissade, champ, échoppe, etc.
* tag (optionnel) de cellule pour indiquer les coordonnées
* factoriser les données des bralduns, pour ne pas répéter partout leur nom (ne mettre dans les vues, champs, échoppes, que les numéros)
* BUG map.html : de temps en temps la carte se charge en noir et je dois faire un reload sous Chrome
* BUG map.html : mon code GO actuel n'efface pas les balises : si une balise est vue sur une case elle le reste (y compris dans la vue)
	-> mettre la date de vue dans VueEnvironnement et comparer à l'écriture dans MemMap, avec raz des routes, champs, échoppes, palissades et bosquets en cas de modif de l'environnement
* extension : système de chat de groupe (avec invitation possibles autres bralduns)

* librairie des cdm
