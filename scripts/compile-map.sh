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

rsync -avz --stats --exclude="map.html" --exclude="carte.json" $CHEMIN_BRALDOP/src/web/* $CHEMIN_DEPLOIEMENT_WEB


for (( i = 0 ; i < ${#NOM_GROUPE[@]} ; i++ ))
do
echo "======================= COMPILATION GROUPE ${NOM_GROUPE[$i]} ======================="
mkdir -p $CHEMIN_DEPLOIEMENT_WEB/groupes/${NOM_GROUPE[$i]}
rsync -avz --stats --exclude="index.html" --exclude="*.crx" $CHEMIN_BRALDOP/src/web/* $CHEMIN_DEPLOIEMENT_WEB/groupes/${NOM_GROUPE[$i]}
./mapper $CHEMIN_REPERTOIRE_DONNEES ${BRALDUNS_GROUPE[$i]} $CHEMIN_DEPLOIEMENT_WEB/groupes/${NOM_GROUPE[$i]}
done
