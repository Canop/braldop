Avertissement :
===============

Ce document est propriété Canop et toute modification externe sera refusée. Il est sur git pour la sauvegarde et la libre consultation, pas pour faire des demandes. Il reflète ce que j'ai besoin de noter plus que ce que j'ai réellement l'intention de faire.

Fait :
======

* mettre en cache les images : le profiling révèle 15% du temps passé à décoder le png existant et 29% à dessiner l'ancienne image (attention mémoire ?)
* utiliser un Mutex.Lock pour éviter de backuper l'ancien png (en appelant le lock lors de l'exit) ?

En cours :
==========

* pb ascenseur menu dans interface scrollable & molette https://github.com/braldahim/braldahim/issues/105
* BUG extension : problème à la distribution de px (Braldahim.js ligne 96)


P1 :
====


* pb : je ne sais pas indiquer si une palissade dans le brouillard est destructible ou non
* code go : essayer d'utiliser le package log

* mapserver : profilage mémoire (pour un peu mieux régler le cache des images en particulier)

* bradmin : fonction de fusion d'images

* récupérer dans l'extension l'affichage des régions
* gestion de groupes dans l'extension
* mécanisme de partage de cartes au sein du groupe
* mécanisme de partage de vues au sein du groupe
* stocker (mysql ?) les échoppes autres lieux (ruines, ?) et leur connaissance par les braldûns
* optimiser gestion images png (je crois que ce qui est long est le décodage, ce qui pourrait s'optimiser par un cache)

* bouton pour activer la session sans se délogguer

* BUG : de temps en temps la carte se charge en noir et je dois faire un reload sous Chrome

* BUG : mon code GO actuel n'efface pas les balises : si une balise est vue sur une case elle le reste (y compris dans la vue)
	-> mettre la date de vue dans VueEnvironnement et comparer à l'écriture dans MemMap, avec raz des routes, champs, échoppes, palissades et bosquets en cas de modif de l'environnement
* faire de jolis dessins pour les routes

* extension : système de chat avec les bralduns en vue (avec invitations possibles d'autres bralduns)
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

