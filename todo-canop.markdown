Avertissement :
===============

Ce document est propriété Canop et toute modification externe sera refusée. Il est sur git pour la sauvegarde et la libre consultation, pas pour faire des demandes. Il reflète ce que j'ai besoin de noter plus que ce que j'ai réellement l'intention de faire.

Fait :
======

* ordonner les profondeurs dans le menu de sélection
* optimisation : les enrichissements des cartes des amis sont effectués en parallèle
* affichage PV/PVMax, faim, durée du tour et DLA des amis

En cours :
==========

* BUG : script serveur bloquant
* BUG : le centrage sur un braldun ne modifie pas la profondeur
* BUG : trop d'enrichissements. On ne semble plus détecter que la carte n'a pas changé

P1 :
====

* stockage JSON / détection changement carte : ne hasher et stocker que la vue, pas l'ensemble du message
* optimisation
* BUG : pas d'affichage pour certaines cases boisées et balisées
* stocker (mysql ?) les échoppes et autres lieux (ruines, buissons, ?) et leur connaissance par les braldûns
* faire de jolis dessins pour les routes
* PB : certaines icônes, en particulier les braldûns, sont floues sur Firefox
* dialogue/bâtiments : lien vers l'aide du bâtiment
* bouton d'activation de la session sans se déconnecter
* centrage : ne demander la profondeur au serveur que si on ne l'a pas

P2 :
====

* page d'accueil plus sympa pour le serveur braldop sur 8001
* factoriser les données des bralduns, pour ne pas répéter partout leur nom (ne mettre dans les vues, champs, échoppes, que les numéros)
* librairie des cdm
* icônes PVE/PVP devant les noms des régions
* fusionner tous les exe en go dans un unique braldop suivant le modèle d'arguments de braldopadmin

OPTM :
======

* Construction des png : c'est l'essentiel de la charge. On gagne environ 25% en modifiant image/png pour utiliser une compression plus rapide.
