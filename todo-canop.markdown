Avertissement :
===============

Ce document est propriété Canop et toute modification externe sera refusée. Il est sur git pour la sauvegarde et la libre consultation, pas pour faire des demandes. Il reflète ce que j'ai besoin de noter plus que ce que j'ai réellement l'intention de faire.

Fait :
======


En cours :
==========

* matériel
* BUG : braldun KO pas dessinés ? Bug dans le script de vue ? -> oui, attendre de revoir des KO...

P1 :
====

* dialogue/bâtiments : lien vers l'aide du bâtiment
* tabac, lingot, aliment, potion

* deux monstres du même type ou de types différents : les afficher l'un à côté de l'autre

* affichage communauté dans la bulle et le menu
* petits-pas pour marquer les cases accessibles en marche  (sur méthode spécifique pour l'intégration)
* BUG : souvent des non initialisations d'images outline
* corriger le README de Braldop
* gestion de la profondeur (une couche par profondeur, le goto fait basculer la profondeur affichée)

* BUG : mon code GO actuel n'efface pas les balises : si une balise est vue sur une case elle le reste (y compris dans la vue)
	-> mettre la date de vue dans VueEnvironnement et comparer à l'écriture dans MemMap, avec raz des routes, champs, échoppes, palissades et bosquets en cas de modif de l'environnement

P2 :
====

* option pour le quadrillage
* tag (optionnel) de cellule pour indiquer les coordonnées
* extension braldop : déclencher les appels du script de vue (ou transmettre directement la vue ?) lors des mouvements [suivant planning intégration complète]
* bug : divers algos ne sont pas adaptés (affichage par exemple, ou getBralduns) si plusieurs vues affichées se recoupent
* points de distinction, gredins et redresseurs

P3 :
====

* factoriser les données des bralduns, pour ne pas répéter partout leur nom (ne mettre dans les vues, champs, échoppes, que les numéros)

Questions :
===========

