Avertissement :
===============

Ce document est propriété Canop et toute modification externe sera refusée. Il est sur git pour la sauvegarde et la libre consultation, pas pour faire des demandes. Il reflète ce que j'ai besoin de noter plus que ce que j'ai réellement l'intention de faire.

En cours :
==========

* dialogue/bâtiments : lien vers l'aide du bâtiment
* version 4 des scripts de vue
* affichage des objets mobiles (Braldûns, monstres, buissons, éléments, etc.) de la dernière vue

P1 :
====

* BUG : peau au pluriel ne prend pas un s...
* BUG : souvent des non initialisations d'images outline
* lecture et affichage monstres (faut déjà que j'en voie...)
* corriger l'assombrissement des vues : il disparait quand la zone vue n'intercepte pas l'écran (et de toutes manières il faut cumuler les trous, pas les assombrissements)
* corriger le README de Braldop

P2 :
====

* extension braldop : déclencher les appels du script de vue (ou transmettre directement la vue ?) lors des mouvements
* bug : divers algos ne sont pas adaptés (affichage par exemple, ou getBralduns) si plusieurs vues affichées se recoupent

P3 :
====

* factoriser les données des bralduns, pour ne pas répéter partout leur nom (ne mettre dans les vues, champs, échoppes, que les numéros)
* intégrer les matrices de vue à la matrice globale (une seule cellule par case) ? 

Questions :
===========


Note :
======

Boule : 

	Sur Braldahim, les cases sont ordonnées telles que : 
	 - Haut gauche : Braldûns, monstres
	 - Haut droit : lieu ou échoppe ou champ
	 - Bas gauche : cadavre, castars et runes, buisson, Braldûns KO. Les 5 images peuvent être affichés en même temps
	 - Bas droit : les charrettes + éléments (matériels, aliments, potions, équipements, munitions, minerais bruts, lingôt, parties plantes brutes, parties plantes préparées, grains, ingrédients, tabac)

	Généralement, ce qu'il y a en bas à gauche correspond aux combats : runes / castars / Braldûns KO / cadavre. Le buisson est là, mais c'est parce qu'il n'y avait pas de place ailleurs.

	Sur Braldop, les castars sont partis en bas à droite. Est-ce possible de les mettre en bas à gauche ? 
