// remplacement des fonctions de l'interface que l'on trouve normalement dans Map_vue.js


// construit l'objet matriceVues qui contient les infos de toutes les vues visibles
Map.prototype.compileLesVues = function() {
	if (!(this.spritesVueTypes.ready&&this.spritesEnv.ready)) {
		//~ console.log('not ready for compilation');
		return;
	}
	this.matricesVuesParZ = {};
	this.matriceVues = {};
	if (!this.mapData) return;
	var déjàVus = {};
	var nouveau = function(id) {
		if (déjàVus[id]) return false;
		déjàVus[id] = 1;
		return true;
	}
	var vuesPlusRécentes = [];
	var vue;
	celeste = function(x, y) {  // une fonction(x,y) qui renvoie true si la case n'est pas dans le rectangle d'une vue plus récente
		for (var jv=vuesPlusRécentes.length; jv-->0;) {
			var vr = vuesPlusRécentes[jv];
			if (!vr.active || vue.Z!=vr.Z) continue;
			if (x>=vr.XMin && x<=vr.XMax && y>=vr.YMin && y<=vr.YMax) return false;
		}
		return true;
	};
	for (var iv=this.mapData.Vues.length; iv-->0;) { // les plus récentes en premier			
		vue = this.mapData.Vues[iv];
		if (!vue.active) continue;
		this.matriceVues = this.matricesVuesParZ[vue.Z];
		if (!this.matriceVues) {
			this.matriceVues = {};
			this.matricesVuesParZ[vue.Z] = this.matriceVues;
		}
		for (ib in vue.Bralduns) {
			var b = vue.Bralduns[ib];
			if (celeste(b.X, b.Y) && nouveau(b.Id)) this.getCellVueCreate(b.X, b.Y).bralduns.push(b);
		}
		for (io in vue.Objets) {
			var o = vue.Objets[io];
			if (celeste(o.X, o.Y)) this.getCellVueCreate(o.X, o.Y).objets.push(o);
		}
		for (io in vue.Monstres) {
			var o = vue.Monstres[io];
			if (celeste(o.X, o.Y) && nouveau(o.Id)) this.getCellVueCreate(o.X, o.Y).monstres.push(o);
		}
		for (io in vue.Cadavres) {
			var o = vue.Cadavres[io];
			if (celeste(o.X, o.Y)) this.getCellVueCreate(o.X, o.Y).cadavres.push(o);
		}
		//> on ajoute les actions aux cellules
		if (vue.actions) {
			for (var i=0; i<vue.actions.length; i++) {
				var a = vue.actions[i];
				var cell = this.getCellVueCreate(a.X, a.Y);
				cell.actions.push(a); // pour la popup, plusieurs actions possibles
				if (this.typesActions[a.Type].isIconeMap) { // affichage de l'icône ou non sur la case
					cell.zones[1].push(this.typesActions[a.Type].icone);
				}
			}
		}
		//> pour chaque cellule on construit les tableaux d'images par zones
		for (var x=vue.XMin; x<=vue.XMax; x++) {
			for (var y=vue.YMin; y<=vue.YMax; y++) {
				var cell = this.getCellVue(x, y);
				if (cell && celeste(x,y)) {
					var nbBraldunsFémininsNonKO=0; 
					var nbBraldunsMasculinsNonKO=0;
					var nbBraldunsKO=0;
					//-- zone 0 : bralduns
					if (cell.bralduns.length) {
						var hasBraldunsCampA = false;
						var hasBraldunsCampB = false;
						for (var i=0; i<cell.bralduns.length; i++) {
							var b = cell.bralduns[i];
							if (b.KO) {
								nbBraldunsKO++;
							} else {
								if (b.Sexe=='f') nbBraldunsFémininsNonKO++;
								else nbBraldunsMasculinsNonKO++;
								if (b.Camp=='a') hasBraldunsCampA=true;
								else if (b.Camp=='b') hasBraldunsCampB=true;
							}
						}
						if (nbBraldunsFémininsNonKO+nbBraldunsMasculinsNonKO>0) {
							var key = 'braldun';
							if (nbBraldunsFémininsNonKO+nbBraldunsMasculinsNonKO>1) key+='s';
							if (nbBraldunsMasculinsNonKO>0) key += '_masculin';
							if (nbBraldunsFémininsNonKO>0) key += '_feminin';
							if (hasBraldunsCampA && hasBraldunsCampB) key += '-combat';
							else if (hasBraldunsCampA) key += '-a';
							else if (hasBraldunsCampB) key += '-b';
							var img = this.spritesVueTypes.get(key);
							if (img) cell.zones[0].push(img);
							//~ else console.log("pas d'image de braldun pour la clé '" +key+"'");
						}
					}
					//-- zone 0 : monstres
					if (cell.monstres.length) {
						var nbByType = {};
						var nbTypes=0;
						var t;
						for (var i=cell.monstres.length; i-->0;) {
							t = cell.monstres[i].IdType;
							if (nbByType[t]) {
								nbByType[t]++;
							} else {
								nbByType[t] = 1;
								nbTypes++;
							}
						}
						if (nbTypes==1 && cell.monstres.length==2) {
							var img = this.spritesVueTypes.get('monstre_'+t+'a', 'monstre');
							cell.zones[0].push(img);
							cell.zones[0].push(img);
						} else if (nbTypes==1 && cell.monstres.length==3) {
							cell.zones[0].push(this.spritesVueTypes.get('monstre_'+t+'b', 'monstres'));
							cell.zones[0].push(this.spritesVueTypes.get('monstre_'+t+'a', 'monstre'));
						} else if (nbTypes<3) {
							for (t in nbByType) {
								cell.zones[0].push(nbByType[t]==1 ? this.spritesVueTypes.get('monstre_'+t+'a', 'monstre') : this.spritesVueTypes.get('monstre_'+t+'b', 'monstres'));
							}
						} else {
							cell.zones[0].push(this.spritesVueTypes.get('monstres'));
						}
					}
					//-- zone 2 : braldun KO
					if (nbBraldunsKO>0) {
						cell.zones[2].push(this.spritesVueTypes.get('braldunko'));
					}
					//-- zone 2 : cadavre
					if (cell.cadavres.length) {
						cell.zones[2].push(this.spritesVueTypes.get('cadavre'));
					}
					//-- zones 1, 2 et 3 : objets, triés suivant leur type et orientés dans l'une des deux zones
					if (cell.objets.length) {
						for (var i=0; i<cell.objets.length; i++) {
							var o = cell.objets[i];
							var typeDéjàPrésent = false;
							for (var j=0; j<i; j++) {
								if (o.Type==cell.objets[j].Type) {
									typeDéjàPrésent = true;
									break;
								}
							}
							if (typeDéjàPrésent) continue;
							var dest = cell.zones[3];
							if (o.Type=='castar'||o.Type=='rune') dest = cell.zones[2];
							else if (o.Type=="ballon"||o.Type=="buisson") dest = cell.zones[1];
							var img = this.spritesVueTypes.get(this.getObjectImgKey(o));
							if (img) {
								dest.push(img);
							} else {
								console.log("pas d'image pour cet objet :", o);
							}
						}
					}
				}
			}
		}
		vuesPlusRécentes.push(vue);
	}	
}
