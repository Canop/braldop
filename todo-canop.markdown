Avertissement :
===============

Ce document est propriété Canop et toute modification externe sera refusée. Il est sur git pour la sauvegarde et la libre consultation, pas pour faire des demandes. Il reflète ce que j'ai besoin de noter plus que ce que j'ai réellement l'intention de faire.

En cours :
==========

* affichage des objets mobiles (Braldûns, etc.) de la dernière vue

P1 :
====

* extension chrome pour le décochage de l'activation au login
* utilisation de la taille naturelle de toutes les icônes (hors terrains) en zoom 64
* optimisation des affichages de cases et de lieux : tester x et y plutôt que le rectangle (donc calculer xmin et xmax etc.) en début de dessin

P2 :
====

* dans le js : table définie suivant les dimensions actuelles et la résolution (si les lieux sont visibles) permettant l'accès rapide aux dits lieux suivant x,y. Remplie en cours de dessin. Notons qu'on peut aussi envisager une sous liste des objets survolables à l'écran (vide si la résolution ne permet pas le survol)
