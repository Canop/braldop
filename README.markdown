Présentation :
==============

[Braldop](http://canop.org/braldop) est un complément au jeu [Braldahim](http://www.braldahim.com).

Il s'agit tout à la fois :

* d'une librairie cartographique et visuelle, [Goût-Gueule](http://forum.braldahim.com/viewtopic.php?f=30&t=1223), [maintenant intégrée](http://forum.braldahim.com/viewtopic.php?f=30&t=1236) dans le jeu lui-même et disponible pour tous ceux qui souhaitent augmenter leur interface tactique de groupe
* d'un système tout prêt de partage de vues et cartes, pour les joueurs solo ou les groupes
* d'une extension pour le navigateur Chrome, intégrant entre autres la cartographie à l'interface standard du jeu

Le projet est géré par cano.petrole@gmail.com. N'hésitez pas à appeler au secours si nécessaire.



Installation du système cartographique :
========================================

Cas 1 :
-------

Vous disposez d'une machine linux visible sur internet et sur laquelle vous pouvez installer des logiciels.

* vous installez le [go](http://golang.org) (branche *weekly*)
* vous récupérez le présent machin via git
* vous configurez *scripts/config.sh*
* vous lancez *scripts/update.sh* ou *scripts/update-vue.sh*
* si ça marche, vous le mettez dans un cron, avec exécution régulière
* sinon, ben vous me contactez en irc/gtalk/fofo...

Cas 2 :
-------

Vous disposez d'une part d'une machine linux sur laquelle vous pouvez installer des choses et d'autre part d'un hébergement de fichiers sur internet (on trouve ça gratuitement).

Dans ce cas vous procédez comme dans le cas 1 sauf que vous configurez dans *config.sh* comme répertoire de déploiement le répertoire de votre serveur.

Installation de l'extension Braldop :
=====================================

Visitez avec votre navigateur Chrome le site [Braldop](http://canop.org/braldop) et cliquez sur le bouton d'installation. Voilà.
