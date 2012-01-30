package bra

// persistence des comptes braldop sur mysql

// renvoie un compte braldop pris en bd
func (con ConnexionMysql) AuthentifieCompte(idBraldun uint, mdpr string) (*CompteBraldop, error) {
	sql := "select mdpr_ok, x, y, z from compte where id=? and mdpr=?"
	stmt, err := con.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.FreeResult()
	err = stmt.BindParams(idBraldun, mdpr)
	if err != nil {
		return nil, err
	}
	err = stmt.Execute()
	if err != nil {
		return nil, err
	}
	cb := new(CompteBraldop)
	var mdprok int
	stmt.BindResult(&mdprok, &cb.X, &cb.Y, &cb.Z)
	eof, err := stmt.Fetch()
	if err != nil || eof {
		return nil, err
	}
	cb.IdBraldun = idBraldun
	cb.Mdpr = mdpr
	cb.Authentifié = mdprok == 1
	return cb, nil
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
