package bra

// persistence des partages sur mysql

import (
	"log"
)

type Partage struct {
	IdA uint // id du braldun A
	IdB uint
	AOk bool // le braldun A a accepté (ou proposé) le partage
	BOk bool
}

// renvoie la liste des amis (les bralduns avec qui un partage est établi)
// Seuls les comptes ayant mdpr_ok à 1 sont pris en compte.
func (con ConnexionMysql) Amis(idBraldun uint) ([]*CompteBraldop, error) {
	amis := make([]*CompteBraldop, 0, 10)
	sql := "select id, mdpr, x, y, z from compte, partage where ((a_id=? and id=b_id) or (b_id=? and id=a_id)) and a_ok=1 and b_ok=1 and mdpr_ok=1"
	rows, err := con.Query(sql, idBraldun, idBraldun)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		cb := new(CompteBraldop)
		err = rows.Scan(&cb.IdBraldun, &cb.Mdpr, &cb.X, &cb.Y, &cb.Z)
		if err != nil {
			return nil, err
		}
		amis = append(amis, cb)
	}
	return amis, nil
}

// récupère toutes les infos de partage, acceptés ou non, impliquant un braldun
func (con ConnexionMysql) Partages(idBraldun uint) ([]*Partage, error) {
	sql := "select a_id, b_id, a_ok, b_ok from partage where a_id=? or b_id=?"
	rows, err := con.Query(sql, idBraldun, idBraldun)
	if err != nil {
		return nil, err
	}
	partages := make([]*Partage, 0, 10)
	var ida, idb uint
	var aok, bok int
	for rows.Next() {
		err = rows.Scan(&ida, &idb, &aok, &bok)
		if err != nil {
			return nil, err
		}
		p := &Partage{ida, idb, aok == 1, bok == 1}
		partages = append(partages, p)
	}
	return partages, nil
}

func (con ConnexionMysql) ModifiePartage(proposant uint, cible uint, action string) (err error) {
	var sql string
	doubleParas := false
	switch action {
	case "accepter":
		sql = "update partage set b_ok=1 where b_id=? and a_id=?"
	case "annuler", "refuser", "rompre":
		sql = "delete from partage where (a_id=? and b_id=?) or (b_id=? and a_id=?)"
		doubleParas = true
	case "proposer":
		sql = "insert into partage (a_id, b_id, a_ok, b_ok) values (?, ?, 1, 0)"
	default:
		log.Println("Action partage inconnue : ", action)
		return nil
	}
	if doubleParas {
		_, err = con.Exec(sql, proposant, cible, proposant, cible)
	} else {
		_, err = con.Exec(sql, proposant, cible)
	}
	return err
}
