# tue, recompile et relance le mapserver
# la sortie standard est redirigee vers gogo.out
# la sortie d'erreur est redirigee vers gogo.err
# les anciens fichiers sont renommes en -old

CHEMIN=`dirname $0`
cd $CHEMIN

source config.sh


cd $CHEMIN_BRALDOP/src/go/braldop/bra
rm -f _go_.6
rm -f _go_.8
gomake install

cd $CHEMIN_BRALDOP/src/go/braldop/mapserver
rm -f _go_.6
rm -f _go_.8
gomake

killall -q mapserver
mv mapserver.out mapserver.out-old
mv mapserver.err mapserver.err-old

echo "*** Ctrl C stoppe l'affichage de la trace mais pas le serveur ***" > mapserver.out

nohup ./mapserver -cartes $CHEMIN_REPERTOIRE_DONNEES/cartes >> mapserver.out 2> mapserver.err < /dev/null &

tail -f mapserver.out
