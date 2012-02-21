package bra

// persistence des comptes braldop sur mysql

import (
	"database/sql"
	"log"
)

func erreurCompte(idBraldun uint, idError string, err error) (*CompteBraldop, string) {
	log.Printf(" Erreur %s sur braldun %d : %s\n", idError, idBraldun, err.Error())
	return nil, "Erreur " + idError // il s'agit de l'erreur affichable dans l'IHM	
}

// renvoie un compte trouvé en BD. Ne fait aucune vérification et renvoie donc false dans le champ Authentifié
func (con ConnexionMysql) GetCompteExistant(idBraldun uint) (*CompteBraldop, string) {
	row := con.QueryRow("select mdpr, x, y, z from compte where id=?", idBraldun)
	cb := new(CompteBraldop)
	cb.IdBraldun = idBraldun
	cb.Authentifié = false
	err := row.Scan(&cb.Mdpr, &cb.X, &cb.Y, &cb.Z)
	if err == nil {
		return cb, ""
	}
	if err == sql.ErrNoRows {
		return nil, ""
	}
	return erreurCompte(idBraldun, "Auth 01", err)
}

// Renvoie un compte braldop pris en bd, et un message d'erreur pour l'utilisateur (ou "")
// Si nécessaire crée le compte et/ou vérifie le mot de passe auprès du script public.
// Renvoie un compte ssi l'authentification est OK (donc cb.Authentifié vaut toujours true)
// Logique :
//  Si le couple id,mdpr est présent et identique en bd, on le renvoie
//  Si le couple id,mdpr est présent mais mdpr différent, on tente l'authentification par le service web
//     Si le nouveau id,mdpr est correct, on met à jour la bd et on renvoie le nouveau cb
//     Si le nouveau id,mdpr est incorrect, on ne met pas à jour la bd et on renvoie nil avec un message d'erreur
//  Si le braldun n'a pas de compte en bd, on tente l'authentification par le service web
//     Si le nouveau id,mdpr est correct, on met à jour la bd et on renvoie le nouveau cb
//     Si le nouveau id,mdpr est incorrect, on ne met pas à jour la bd et on renvoie nil avec un message d'erreur
func (con ConnexionMysql) AuthentifieCompte(idBraldun uint, mdpr string, créeSiNécessaire bool) (*CompteBraldop, string) {
	cb, errmess := con.GetCompteExistant(idBraldun)
	if errmess != "" {
		return nil, errmess
	}
	if cb == nil { //- nouveau compte
		log.Println(" compte absent de la BD pour braldun ", idBraldun)
		auth, err := AuthentifieCompteParScriptPublic(idBraldun, mdpr)
		if err != nil {
			return erreurCompte(idBraldun, "Auth 05", err)
		}
		if !auth {
			log.Println(" Non validation mot de passe restreint braldun ", idBraldun)
			return nil, "Non validation mot de passe restreint (peut-être un dépassement du nombre de tentatives)"
		}
		//- compte OK : à insérer
		_, err = con.Exec("insert into compte (id, mdpr, mdpr_ok) values (?, ?, 1)", idBraldun, mdpr)
		if err != nil {
			return erreurCompte(idBraldun, "Auth 07", err)
		}
		cb := new(CompteBraldop)
		cb.IdBraldun = idBraldun
		cb.Authentifié = true
		cb.Mdpr = mdpr
		log.Println(" Nouveau compte : braldun ", idBraldun)
		return cb, ""
	} else { //- compte déjà présent
		if mdpr == cb.Mdpr { //- mot de passe identique à celui en bd : ok
			cb.Authentifié = true
			return cb, ""
		} else { //- mot de passe différent, on va regarder si le nouveau est bon
			auth, err := AuthentifieCompteParScriptPublic(idBraldun, mdpr)
			if err != nil {
				return erreurCompte(idBraldun, "Auth 09", err)
			}
			if !auth {
				log.Println(" Non validation nouveau mot de passe restreint braldun ", idBraldun)
				return nil, "Non validation nouveau mot de passe restreint (peut-être un dépassement du nombre de tentatives)"
			} else { //- nouveau mot de passe restreint, à mettre à jour (TODO -> transmettre info pour merge images)
				cb.Authentifié = true
				sql := "update compte set mdpr=?, mdpr_ok=1 where id=?"
				_, err := con.Exec(sql, mdpr, idBraldun)
				if err != nil {
					return erreurCompte(idBraldun, "Auth 11", err)
				}
				cb.Mdpr = mdpr
				log.Println(" Mot de passe modifié pour braldun ", idBraldun)
				return cb, ""
			}
		}
	}
	return nil, "" // ligne inateignable
}
