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

rsync -az --delete $CHEMIN_BRALDOP/src/site/* $CHEMIN_DEPLOIEMENT_WEB


for (( i = 0 ; i < ${#NOM_GROUPE[@]} ; i++ ))
do
echo "======================= COMPILATION GROUPE ${NOM_GROUPE[$i]} ======================="
mkdir -p $CHEMIN_DEPLOIEMENT_WEB/groupes/${NOM_GROUPE[$i]}
rsync -az --delete $CHEMIN_BRALDOP/src/gui/* $CHEMIN_DEPLOIEMENT_WEB/groupes/${NOM_GROUPE[$i]}
./mapper ${MAPPER_ARGS_GROUPE[$i]} -in $CHEMIN_REPERTOIRE_DONNEES -bralduns ${BRALDUNS_GROUPE[$i]} -out $CHEMIN_DEPLOIEMENT_WEB/groupes/${NOM_GROUPE[$i]}
done
