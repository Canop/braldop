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
	stmt, err := con.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.FreeResult()
	err = stmt.BindParams(idBraldun, idBraldun)
	if err != nil {
		return nil, err
	}
	err = stmt.Execute()
	if err != nil {
		return nil, err
	}
	cb := new(CompteBraldop)
	stmt.BindResult(&cb.IdBraldun, &cb.Mdpr, &cb.X, &cb.Y, &cb.Z)
	for {
		eof, _err := stmt.Fetch()
		if _err != nil || eof {
			return amis, _err
		}
		amis = append(amis, cb.Clone())
	}
	return amis, nil // je ne crois pas qu'on puisse arriver là mais cette ligne permet la compilation...
}

// récupère toutes les infos de partage, acceptés ou non, impliquant un braldun
func (con ConnexionMysql) Partages(idBraldun uint) ([]*Partage, error) {
	sql := "select a_id, b_id, a_ok, b_ok from partage where a_id=? or b_id=?"
	stmt, err := con.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.FreeResult()
	err = stmt.BindParams(idBraldun, idBraldun)
	if err != nil {
		return nil, err
	}
	err = stmt.Execute()
	if err != nil {
		return nil, err
	}
	r := new(Partage)
	var aok, bok int
	stmt.BindResult(&r.IdA, &r.IdB, &aok, &bok)
	partages := make([]*Partage, 0, 10)
	for {
		eof, _err := stmt.Fetch()
		if _err != nil || eof {
			return partages, _err
		}
		p := &Partage{r.IdA, r.IdB, aok==1, bok==1}
		partages = append(partages, p)
	}
	return partages, nil
}

func (con ConnexionMysql) ModifiePartage(proposant uint, cible uint, action string) (error) {
	var sql string
	doubleParas := false
	switch action {
	case "accepter" :
		sql = "update partage set b_ok=1 where b_id=? and a_id=?"
	case "annuler", "refuser", "rompre" :
		sql = "delete from partage where (a_id=? and b_id=?) or (b_id=? and a_id=?)"
		doubleParas = true
	case "proposer" :
		sql = "insert into partage (a_id, b_id, a_ok, b_ok) values (?, ?, 1, 0)"
	default:
		log.Println("Action partage inconnue : ", action)
		return nil
	}
	stmt, err := con.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.FreeResult()
	if (doubleParas) {
		err = stmt.BindParams(proposant, cible, proposant, cible)
	} else {
		err = stmt.BindParams(proposant, cible)
	}
	if err != nil {
		return err
	}
	err = stmt.Execute()
	return err
}
