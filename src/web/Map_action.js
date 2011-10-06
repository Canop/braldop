/*
 * Champs de l'objet action :
 *  type : numérique, voir map.typesActions
 *  x
 *  y
 *  z : optionnel car implicitement le z du braldun
 *  idBraldun : id, le cas échéant, de la cible s'il s'agit d'un braldun
 *  idMonstre :  id, le cas échéant, de la cible s'il s'agit d'un monstre
 * 
 * 
 * Champs du type d'action :
 *  nom
 *  iconeCase : optionnel, à afficher sur la case où l'action est disponible
 */
 
 



Map.prototype.initTypesActions = function() {
	var icon = function(s) {
		var img = new Image();
		img.src="http://static.braldahim.com/images/"+s+".png";
		return img;
	}
	this.typesActions = [];
	// 0 : marcher
	this.typesActions[0] = {nom:'Marcher', iconeCase:icon('vue/pas2')}; 
}

// champs :
//  - idBraldun : id du braldun pouvant réaliser les actions
//  - actions  : une liste d'actions
//  - callback : la méthode à appeler en cas de demande de réalisation de l'action. 
//               cette méthode sera appelée avec pour paramètres idBraldun et l'action.
//               elle peut être nulle (cas par exemple d'une interface tactique ne faisant
//               que lister les actions possibles
// L'implémentation ne permet pas pour l'instant de supprimer des actions
Map.prototype.setActions = function(idBraldun, actions, callback) {
	if (!this.actionsParBraldun) this.actionsParBraldun = {};
	this.actionsParBraldun[idBraldun] = {
		actions: actions,
		callback: callback
	}
}
