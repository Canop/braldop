Présentation :
==============

Braldop est un système cartographique pour les joueurs de [Braldahim](http://www.braldahim.com)

Il s'agit également d'un prototype pour une version future de la vue dans ce jeu.

Le projet est géré par cano.petrole@gmail.com. N'hésitez pas à appeler au secours si nécessaire.



Structure :
===========

* scripts : les scripts qui permettent de maintenir la carte à jour
* src/server : les codes sources


Installation :
==============

	Cas 1 :
	-------

	Vous disposez d'une machine linux visible sur internet et sur laquelle vous pouvez installer des logiciels.
	
	* vous installez le [go](http://golang.org)
	* vous récupérez le présent machin via git
	* vous configurez scripts/config.sh
	* vous lancez scripts/update.sh
	* si ça marche, vous le mettez dans un cron, avec exécution toutes les 8 ou 12h
	* sinon, ben vous me contactez en irc/gtalk/fofo...
	
	Cas 2 :
	-------
	
	Vous disposez d'une part d'une machine linux sur laquelle vous pouvez installer des choses et d'autre part d'un hébergement de fichiers sur internet (on trouve ça gratuitement).
	
	Dans ce cas vous procédez comme dans le cas 1 sauf qu'au lieu de servir en http le contenu de votre répertoire de déploiement, vous envoyez (par exemple via rsync ou ftp) ce répertoire sur la zone publique de votre hébergeur.
	
