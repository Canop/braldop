package bra
/*
Frontal pour la BD

j'exploite ce connecteur mysql : https://github.com/Philio/GoMySQL

Pour le mettre à jour, je me mets dans son répertoire et je fais
   git fetch origin
   git merge origin/master
puis
	make
	make install

Note : on dirait que ce connecteur a des problèmes avec les int64. Il vaut mieux, donc, éviter d'en utiliser.

*/

import "mysql"

type BaseMysql struct {
	user     string
	password string
	database string
}

type ConnexionMysql struct {
	*mysql.Client
}


func NewBaseMysql(user string, password string, database string) *BaseMysql {
	store := new(BaseMysql)
	store.user = user
	store.password = password
	store.database = database
	return store
}

// renvoie une instance de mysql.Client connectée. Il est impératif de faire suivre
// l'appel à cette méthode d'un defer con.Close()
func (bd *BaseMysql) Open() (ConnexionMysql, error) {
	db, err := mysql.DialUnix(mysql.DEFAULT_SOCKET, bd.user, bd.password, bd.database)
	return ConnexionMysql{db}, err
}
