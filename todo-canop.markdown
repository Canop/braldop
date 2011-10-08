Avertissement :
===============

Ce document est propriété Canop et toute modification externe sera refusée. Il est sur git pour la sauvegarde et la libre consultation, pas pour faire des demandes. Il reflète ce que j'ai besoin de noter plus que ce que j'ai réellement l'intention de faire.

Fait :
======

En cours :
==========


P1 :
====

* extension : ajouter un lien vers la carte de groupe
* définition d'une API pour le blabla
* BUG : quand un truc est listé dans deux vues il apparait deux fois dans les bulles -> une seule matrice pour toutes les vues ? vérif à la compil ?
* icônes PVE/PVP devant les noms des régions
* problème : temps de chargement excessif. Les images ne sont-elles pas cachées ? profiler.
* matériel
* dialogue/bâtiments : lien vers l'aide du bâtiment
* BUG : souvent des non initialisations d'images outline
* BUG : mon code GO actuel n'efface pas les balises : si une balise est vue sur une case elle le reste (y compris dans la vue)
	-> mettre la date de vue dans VueEnvironnement et comparer à l'écriture dans MemMap, avec raz des routes, champs, échoppes, palissades et bosquets en cas de modif de l'environnement
* points de distinction, gredins et redresseurs

P2 :
====

* structure générique plus propre côté js pour décrire l'accès aux images, les documentations associées
* tag (optionnel) de cellule pour indiquer les coordonnées
* extension braldop : déclencher les appels du script de vue (ou transmettre directement la vue ?) lors des mouvements [suivant planning intégration complète]
* bug : divers algos ne sont pas adaptés (affichage par exemple, ou getBralduns) si plusieurs vues affichées se recoupent
* factoriser les données des bralduns, pour ne pas répéter partout leur nom (ne mettre dans les vues, champs, échoppes, que les numéros)

