#!/bin/bash

# Ce script recompile l'application qui produit la carte puis l'exécute
# On déploie la carte dans le répertoire des sources web pour faciliter
#  le développement sur place tant du go que du json en évitant les transferts
#  manuels de fichiers.
# Le but de ce script est donc de tester le développement go sans incrémenter
#  votre compteur d'appels des scripts publics de Braldahim.

CHEMIN=`dirname $0`
cd $CHEMIN

source config.sh

cd $CHEMIN_BRALDOP/src/server/mapper
gomake
./mapper $CHEMIN_REPERTOIRE_DONNEES

rsync -avz --stats $CHEMIN_REPERTOIRE_DONNEES/carte.json $CHEMIN_BRALDOP/src/web

