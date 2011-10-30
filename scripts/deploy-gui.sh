#!/bin/bash

# Ce script déploie les fichiers de l'interface.
# Il ne compile pas la carte mais ne supprime pas celle qui existe déjà (éventullement.
# Le but est de tester une nouvelle version de l'interface (côté client uniquement) sans
# devoir attendre la compilation de la carte (et sans recompiler le serveur)


CHEMIN=`dirname $0`
cd $CHEMIN

source config.sh


for (( i = 0 ; i < ${#NOM_GROUPE[@]} ; i++ ))
do
echo "======================= DEPLOYEMENT GROUPE ${NOM_GROUPE[$i]} ======================="
mkdir -p $CHEMIN_DEPLOIEMENT_WEB/groupes/${NOM_GROUPE[$i]}
rsync -az $CHEMIN_BRALDOP/src/gui/* $CHEMIN_DEPLOIEMENT_WEB/groupes/${NOM_GROUPE[$i]}
done
