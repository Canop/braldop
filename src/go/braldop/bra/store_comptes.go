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
	cb.Authentifié = mdprok==1
	return cb, nil
}

