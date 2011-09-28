Avertissement :
===============

Ce document est propriété Canop et toute modification externe sera refusée. Il est sur git pour la sauvegarde et la libre consultation, pas pour faire des demandes. Il reflète ce que j'ai besoin de noter plus que ce que j'ai réellement l'intention de faire.

En cours :
==========


P1 :
====

* BUG : braldun KO pas dessinés ?
* affichage des runes et buissons
* affichage des palissades et bosquets
* dialogue/bâtiments : lien vers l'aide du bâtiment

* affichage match de soule
* petits-pas pour marquer les cases accessibles en marche  (sur méthode spécifique pour l'intégration)
* BUG : peau au pluriel ne prend pas un s...
* BUG : souvent des non initialisations d'images outline
* corriger le README de Braldop
* gestion de la profondeur (une couche par profondeur, le goto fait basculer la profondeur affichée)
* corriger le pluriel de peau

* BUG : je pense que mon code GO actuel n'efface pas les balises : si une balise est vue sur une case elle le reste (y compris dans la vue)
	-> mettre la date de vue dans VueEnvironnement et comparer à l'écriture dans MemMap, avec raz des routes, champs, échoppes et bosquets en cas de modif de l'environnement ?

P2 :
====

* extension braldop : déclencher les appels du script de vue (ou transmettre directement la vue ?) lors des mouvements [suivant planning intégration complète]
* bug : divers algos ne sont pas adaptés (affichage par exemple, ou getBralduns) si plusieurs vues affichées se recoupent
* points de distinction, gredins et redresseurs

P3 :
====

* factoriser les données des bralduns, pour ne pas répéter partout leur nom (ne mettre dans les vues, champs, échoppes, que les numéros)

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
