#!/bin/bash

# Ce script récupère les données des scripts publics puis appele la compilation et le déployement
#  de la carte et du site web.
# Il est recommandé d'appeler en cron le présent script.


echo "==================================================================="
echo "= MISE A JOUR CARTE BRALDOP                                       ="
echo "==================================================================="

CHEMIN=`dirname $0`
cd $CHEMIN

echo "= récupération de la config ======================================="
source config.sh
echo "= récupération des données statiques =============================="
mkdir -p $CHEMIN_REPERTOIRE_DONNEES/public
cd $CHEMIN_REPERTOIRE_DONNEES/public
wget -N "http://public.braldahim.com/lieux_villes.csv"
wget -N "http://public.braldahim.com/regions.csv"
wget -N "http://public.braldahim.com/villes.csv"
wget -N "http://public.braldahim.com/zones.csv"
wget -N "http://public.braldahim.com/bralduns.csv"


echo "= récupération des données dynamiques ============================="
FIN_REPERTOIRE_JOUR=`date +%Y/%m/%d`
CHEMIN_JOUR="$CHEMIN_REPERTOIRE_DONNEES/private/$ID_BRALDUN/$FIN_REPERTOIRE_JOUR"
mkdir -p $CHEMIN_JOUR
cd $CHEMIN_JOUR
DEBUT_NOM_FICHIERS_JOUR=`date +%Hh%M`
wget "http://sp.braldahim.com/scripts/vue/?idBraldun=$ID_BRALDUN&mdpRestreint=$MDP_RESTREINT_BRALDUN&version=3" -O "$DEBUT_NOM_FICHIERS_JOUR-vue.csv"
wget "http://sp.braldahim.com/scripts/profil/?idBraldun=$ID_BRALDUN&mdpRestreint=$MDP_RESTREINT_BRALDUN&version=2" -O "$DEBUT_NOM_FICHIERS_JOUR-profil.csv"
wget "http://sp.braldahim.com/scripts/evenements/?idBraldun=$ID_BRALDUN&mdpRestreint=$MDP_RESTREINT_BRALDUN&version=2" -O "$DEBUT_NOM_FICHIERS_JOUR-evenements.csv"

echo "= compilation de la carte ========================================="
cd $CHEMIN_BRALDOP/scripts/
./compile-map.sh

echo "= déployement du site web ========================================="
rsync -avz --stats $CHEMIN_BRALDOP/src/web/* $CHEMIN_DEPLOIEMENT_WEB

echo "==================================================================="
