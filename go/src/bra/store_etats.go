package bra

// persistence de l'état des bralduns sur mysql

import (
	"database/sql"
	"log"
	"time"
)

func (con ConnexionMysql) EtatBraldun(idBraldun uint) (*EtatBraldun, error) {
	row := con.QueryRow("select pv, pvmax, pa, tour, dla, faim from compte where id=?", idBraldun)
	état := new(EtatBraldun)
	err := row.Scan(&état.PV, &état.PVMax, &état.PA, &état.DuréeTour, &état.DLA, &état.Faim)
	état.IdBraldun = idBraldun
	if état.PVMax == 0 {
		return nil, nil
	}
	if err == nil {
		return état, nil
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return nil, err
}

func (con ConnexionMysql) StockeEtatBraldun(e *EtatBraldun) error {
	now := time.Now().Unix()
	sql := "update compte set pv=?, pvmax=?, pa=?, tour=?, dla=?, faim=?, maj=? where id=?"
	res, err := con.Exec(sql, e.PV, e.PVMax, e.PA, e.DuréeTour, e.DLA, e.Faim, now, e.IdBraldun)
	if res != nil {
		nbi, _ := res.RowsAffected()
		log.Println(" StockeEtatBraldun :", nbi, " compte modifié")
	}
	return err
}
