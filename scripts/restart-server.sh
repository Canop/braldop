# tue, recompile et relance le mapserver
# les sorties sont redirigées vers log/chrallserver.log

CHEMIN=`dirname $0`
cd $CHEMIN

source config.sh

go clean bra
go install bra
go clean braldopserver
go install braldopserver
killall -q braldopserver

cd ..
mkdir -p log
mv log/chrallserver.log log/chrallserver.old.log

echo "*** Ctrl C stoppe l'affichage de la trace mais pas le serveur ***" > log/chrallserver.log

nohup go/bin/braldopserver ${MAPSERVER_ARGS} -mysqluser ${USER_MYSQL} -mysqlmdp ${MDP_MYSQL} -datadir $CHEMIN_REPERTOIRE_DONNEES >> log/chrallserver.log 2>&1 < /dev/null &

tail -f log/chrallserver.log
