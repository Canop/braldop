# tue, recompile et relance le mapserver
# les sorties sont redirigées vers log/chrallserver.log

CHEMIN=`dirname $0`
cd $CHEMIN

source config.sh

go clean bra
go install bra
go clean braldopserver
go install braldopserver
killall -s SIGINT -q braldopserver

cd ..
mkdir -p log
mv log/chrallserver.log log/chrallserver.old.log
mv profiling/server.prof profiling/server.old.prof

echo "*** Ctrl C stoppe l'affichage de la trace mais pas le serveur ***" > log/chrallserver.log

echo go/bin/braldopserver ${MAPSERVER_ARGS} -mysqluser ${USER_MYSQL} -mysqlmdp ${MDP_MYSQL} -datadir $CHEMIN_REPERTOIRE_DONNEES
nohup go/bin/braldopserver ${MAPSERVER_ARGS} -mysqluser ${USER_MYSQL} -mysqlmdp ${MDP_MYSQL} -datadir $CHEMIN_REPERTOIRE_DONNEES >> log/server.log 2>&1 < /dev/null &

echo  tail -n 100 -f log/server.log
