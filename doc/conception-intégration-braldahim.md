Glossaire/Modèle
================

* Case : correspond à l'un des points (x,y, z), x et y variant actuellement respectivement dans [-800, 800] et [-500, 500], les y positifs au nord. Z pouvant aller jusqu'en -13
* Environnement : il y en a exactement un par case :
    - caverne
    - gazon
    - marais
    - plaine
    - montagne

* Fond : il y en a exactement un par case (il y a au total moins de 255 fonds possibles, images dispo sur https://github.com/braldahim/braldahim/tree/master/braldahim-static/public/images/vue/environnement). Il correspond à l'environnement éventuellement modifié par :
    - une balise (= affichage de l'environnement + la balise en haut à gauche)
    - un portail, un portail + balise
    - une palissade, une palissade + balise
    - un tunnel (dans les cavernes uniquement, creusé par les mineurs), un tunnel + une balise
    - une crevasse (dans les cavernes uniquement), une crevasse + une balise (cela peut arriver techniquement)
    - un bosquet de type hêtres ou chenes ou peupliers ou erables, bosquet + balise
    - une route / des pavés
    - une eau de type mer ou peuprofonde ou profonde ou lac, eau + balise (==> change rarement, mais certaines portions de carte peuvent être modifiés pour ajouter ou supprimer une eau)
    - inconnu : pour les limites des mines. Il n'y a des mines en z < 0que lorsqu'au niveau 0, c'est de la montagne.


* Lieu : il y en a 0 ou 1 par case. On distingue les lieux liés à une ville et les autres.
* Lieu public : lieu connu de tous (lieux de villes) et définis dans le fichier lieux_villes.csv
* Lieu privé : lieu qui doit être découvert (vu) pour être connu. La connaissance que l'on a du lieu date de la dernière fois que l'on a vu la case qui le contient. Ces "lieux" peuvent être des échoppes, des champs.
* Objet : les objets dont l'existence ou la position sont éphémères : braldûns, gibiers, monstres, charettes, buissons, ingrédients, etc. Leur nombre par case n'est pas limité. On considère que les objets hors vue sont inconnus et ne doivent pas être affichés.

Aux objets et lieux peuvent être associées des informations qui dépassent leur type : nom, propriétaire, quantité, etc.

Il n'y a rien d'autre dans la vue ou sur la carte.

Chaque case peut-être :

* inconnue
* connue mais non visible (sous le "brouillard de guerre")
* visible
 
Le choix a priori est que les fonds et lieux connus mais non visibles soient affichés (avec leurs lieux informations associées) dans l'état dans lequel ils étaient la dernière fois qu'ils ont été vus par le braldûn.

Formats de stockage et d'envoi au navigateur
============================================

Les fonds
---------
 Je propose de stocker les fonds sous la forme d'une image (encodée en gif) de 1600x1000 pixels. Chaque "couleur" correspondrait à l'un des types de fond possible. Il s'agit de la carte complète non accessible aux joueurs.
 La connaissance qu'a un braldûn des fonds serait une image similaire, à base transparente, pour laquelle seuls les pixels correspondant à des cases vues seraient coloriés. Le processus de modification est trivial puisqu'il consiste à copier sur cette image le rectangle de la carte complète correspondant à la zone vue.
 Cette image serait envoyée au navigateur via une url de la forme *fond_carte-idBraldun-mdpresreint.gif?nbMoves=xxx*, ceci pour protéger l'image, la rendre disponible aux interfaces externes et permettre au navigateur de la cacher tant que le braldun n'a pas bougé (nbMoves).
 Le stockage de cette image sur le serveur peut être assurée de plusieurs façons (bd, disque, etc.) le choix dépendant au final sans doute de la technique utilisée pour l'édition.

 Cette image doit pouvoir être affichée telle quelle. On définira donc la palette de telle sorte que chaque couleur soit la moyenne (HSV?) des couleurs de l'illustration du terrain utilisée en 64x64.


Les lieux publics
-----------------
 Ils peuvent être envoyés au navigateur dans un fichier json constant dont l'url ne changerait qu'en cas de modification du contenu.
 
Les lieux privés
----------------
 Deux techniques sont possibles pour l'envoi :
 - une liste unique en json. Je pense que ce sera le format le plus efficace (et surtout le plus simple à coder) tant qu'on n'a pas plus que quelques miliers de lieux privés connus.
 - des listes correspondant à des sous parties de l'écran (par tuiles, comme dans Google Map, c'est quelque chose que j'ai déjà fait dans une autre vie). Ca implique un peu plus de traitement mais ce sera sans doute nécessaire
 
 (notons qu'au niveau du navigateur une map (x,y)->case sera utilisée pour accélérer l'affichage, ce n'est pas un problème)
 
 Pour ce qui est du stockage côté serveur, en rappelant qu'on a au plus un lieu privé par case, les problèmes sont de :
 - lister rapidement tous les lieux privés vus par un braldûn
 - effacer rapidement tous les lieux les lieux privés pour un braldûn dans un (petit) rectangle donné (sa vue)
 - écrire rapidement tous les lieux privés pour un braldûn dans un rectangle donné (sa vue)
 - lister rapidement tous les lieux privés vus par un braldûn et dans un rectangle donné (en option)
 Notons que dans tous ces cas on utilise cette clef d'accès au maximum : (idBraldun.x.y). Cette clef tient en 64 bits.
 
 On peut utiliser pour le stockage des lieux privés vus soit mysql soit une base un peu plus adaptée comme Redis. Dans un premier temps, tant que le nombre de joueurs reste faible, mysql semble acceptable et les traitements peuvent être faits efficacement en php. Si nécessaire on peut cacher sur disque le fichier json des lieux privés visibles d'un joueur.
 
Les objets
----------
 Aucun stockage spécifique (au problème de vue/carte) ne semble nécessaire : il suffit d'envoyer une vue, en json de façon semblable à ce qui est fait dans la maquette braldop, construite dynamiquement à partir de ce qui est en vue.

Affichage côté client
=====================

Les techniques de base sont déjà plus qu'ébauchées dans [la carte actuelle](http://canop.org/braldop/map.html) donc je vais surtout parler des différences.

Affichage des fonds
-------------------

La version actuelle (au 20/09/2011) itère sur l'ensemble des fonds reçus sous forme de liste json et affiche ceux dont le rectangle interesecte le canvas à l'écran. Ceci ne sera pas efficace quand la liste grandira. Ce que je prévois de faire est ceci :
- pour les résolutions 0.5, 1 et 2, simplement afficher l'image brute reçue par dessus l'image de carte déjà utilisée
- pour les résolutions supérieures, simplement itérer sur les cases visibles à l'écran et aller chercher l'image correspondant au fond encodé par la "couleur".
En principe ceci sera fluide même sur grand écran.

Affichage des lieux
-------------------

Je pense mettre en place le même système dans le javascript que pour la vue : précompiler la liste reçue pour associer rapidement à chaque position une case (uniquement s'il y a quelque chose dans la case). La seule différence a priori sera l'utilisation d'une map au lieu d'un tableau (mais ça ne change rien à part des économies de RAM).

Intégration de commandes
========================

La carte affichée doit présenter des actions. Certaines, purement documentaires, peuvent être définies simplement (ouverture de la fiche d'un braldun ou d'un monstre, affichage d'une page du wiki). D'autres nécessitent la prise en compte de règles complexes (charge, déplacement, etc.).

L'API de la carte offre cette méthode l'intégration :

	Map.setActions(idBraldun, list, callback)

La liste est un tableau d'objets présentant les champs suivants
action : {
	Type : un type numérique permettant de faire certains traitements (comme afficher l'icône de pas sur la carte par exemple) et de retrouver le label et l'icône
	X : x du point d'application
	Y : y du point d'application
	IdBraldun : cible éventuelle
	IdMonstre : cible éventuelle
	PA : coût en PA
}

Les noms de ces champs sont tous capitalisés, comme les champs de tous les paramètres externes passés à Braldop.

Le callback a entre autres responsabilités celle du dialogue de confirmation

Lors de la sélection dans un menu, la fonction définie par setActionFonction est appelée avec pour paramètres l'id du braldun acteur et l'objet action.

Points en suspens
=================
 
* Et si l'on causait d'autre chose que de la surface, ça donnerait quoi ?
  - il y a les mines, en z=-10 ou -11 ou -12, accessibles via une entrée de mine dans les villes ou via un puit (un puit = un lieu de type puit)
  - les donjons, avec z < 0 et z > -10 (une entrée de donjon = un lieu, cette entrée permet de descendre en z < 0)

* Fait-on une image (moins grande) pour chaque niveau ? Ca me semble le mieux, non ?