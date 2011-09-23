Avertissement :
===============

Ce document est propriété Canop et toute modification externe sera refusée. Il est sur git pour la sauvegarde et la libre consultation, pas pour faire des demandes. Il reflète ce que j'ai besoin de noter plus que ce que j'ai réellement l'intention de faire.

En cours :
==========

* affichage des objets mobiles (Braldûns, monstres, buissons, éléments, etc.) de la dernière vue
* utilisation de la taille naturelle de toutes les icônes (hors terrains) en zoom 64
* mode d'affichage "proto vue"
* montrer les boites de dialogues sur les lieux et objets

P1 :
====

* corriger l'assombrissement des vues : il disparait quand la zone vue n'intercepte pas l'écran (et de toutes manières il faut cumuler les trous, pas les assombrissements)
* corriger le README de Braldop
* optimisation des affichages de cases et de lieux : tester x et y plutôt que le rectangle (donc calculer xmin et xmax etc.) en début de dessin [sera plutôt fait via la matrice]

P2 :
====

* dans le js : table définie suivant les dimensions actuelles et la résolution (si les lieux sont visibles) permettant l'accès rapide aux dits lieux suivant x,y. Remplie en cours de dessin. Notons qu'on peut aussi envisager une sous liste des objets survolables à l'écran (vide si la résolution ne permet pas le survol)

Questions :
===========

* comment trouve-t-on l'icône d'un type de monstre ?
* Charrettes : à quoi correspond le "nom_type_materiel" ? Comment trouve-t-on le propriétaire ?
