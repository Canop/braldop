Avertissement :
===============

Ce document est propriété Canop et toute modification externe sera refusée. Il est sur git pour la sauvegarde et la libre consultation, pas pour faire des demandes. Il reflète ce que j'ai besoin de noter plus que ce que j'ai réellement l'intention de faire.

En cours :
==========

* affichage des objets mobiles (Braldûns, monstres, buissons, éléments, etc.) de la dernière vue

P1 :
====

* mettre le halo doré sur la case à l'origine du menu
* corriger l'assombrissement des vues : il disparait quand la zone vue n'intercepte pas l'écran (et de toutes manières il faut cumuler les trous, pas les assombrissements)
* corriger le README de Braldop

P2 :
====

* bug : divers algos ne sont pas adaptés (affichage par exemple, ou getBralduns) si plusieurs vues affichées se recoupent
* dans le js : table définie suivant les dimensions actuelles et la résolution (si les lieux sont visibles) permettant l'accès rapide aux dits lieux suivant x,y. Remplie en cours de dessin. Notons qu'on peut aussi envisager une sous liste des objets survolables à l'écran (vide si la résolution ne permet pas le survol)

P3 :
====

* factoriser les données des bralduns, pour ne pas répéter partout leur nom (ne mettre dans les vues, champs, échoppes, que les numéros)
* intégrer les matrices de vue à la matrice globale (une seule cellule par case) ? 

Questions :
===========

* comment trouve-t-on l'icône d'un type de monstre ?
* Charrettes : à quoi correspond le "nom_type_materiel" ? Comment trouve-t-on le propriétaire ?
