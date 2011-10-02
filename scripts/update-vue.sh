#!/bin/bash

# Ce script récupère les données des scripts publics puis appele la compilation et le déployement
#  de la carte et du site web.
# Il est recommandé d'appeler en cron le présent script.
#  
# Ce script est une variante du script update.sh. On ne récupère ici que la vue

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
wget -N "http://public.braldahim.com/communautes.csv"


echo "= récupération des données dynamiques ============================="
FIN_REPERTOIRE_JOUR=`date +%Y/%m/%d`
for (( i = 0 ; i < ${#ID_BRALDUN[@]} ; i++ ))
do
CHEMIN_JOUR="$CHEMIN_REPERTOIRE_DONNEES/private/${ID_BRALDUN[$i]}/$FIN_REPERTOIRE_JOUR"
mkdir -p $CHEMIN_JOUR
cd $CHEMIN_JOUR
DEBUT_NOM_FICHIERS_JOUR=`date +%Hh%M`
wget "http://sp.braldahim.com/scripts/vue/?idBraldun=${ID_BRALDUN[$i]}&mdpRestreint=${MDP_RESTREINT_BRALDUN[$i]}&version=5" -O "$DEBUT_NOM_FICHIERS_JOUR-vue.csv"
done

echo "= compilation de la carte et déploiement =========================="
cd $CHEMIN_BRALDOP/scripts/
./compile-map.sh

echo "==================================================================="
