package bra

// accès à la BD
// Pour l'instant je ne fais pas de déconnexion, on reste connecté en permanence

import (
	"database/sql"
	_ "github.com/ziutek/mymysql/godrv"
)

type ConnexionMysql struct {
	*sql.DB
}

type BaseMysql struct {
	user     string
	password string
	database string
	con ConnexionMysql
}


func NewBaseMysql(user string, password string, database string) *BaseMysql {
	store := new(BaseMysql)
	store.user = user
	store.password = password
	store.database = database
	return store
}

// renvoie une instance de DB connectée.
func (store *BaseMysql) DB() (ConnexionMysql, error) {
	if store.con.DB == nil {
		db, err := sql.Open("mymysql", store.database+"/"+store.user+"/"+store.password)
		if err == nil {
			store.con.DB = db
		}
	}
	return store.con, nil
}
