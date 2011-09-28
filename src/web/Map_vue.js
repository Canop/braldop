// parcoure les vues affichées pour trouver les bralduns visibles sur la case x,y.
Map.prototype.getBralduns = function(x, y) {
	if (this.mapData.Vues) {
		for (var i=this.mapData.Vues.length; i-->0;) {
			var vue = this.mapData.Vues[i];
			if (vue.active) {
				var cell = getCellVue(vue, x, y);
				if (cell) {
					return cell.bralduns;
				}
			}
		}
	}
	return [];
}

// renvoie la première cellule de vue trouvée (en ne cherchant que parmi les vues affichées)
Map.prototype.getCellVueVisible = function(x, y) {
	if (this.mapData.Vues) {
		for (var i=this.mapData.Vues.length; i-->0;) {
			var vue = this.mapData.Vues[i];
			if (vue.active) {
				var cell = getCellVue(vue, x, y);
				if (cell) {
					return cell;
				}
			}
		}
	}
	return null;
}

// renvoie la cellule de la vue ou null (hors vue ou vide)
function getCellVue(vue, x, y) {
	var W = vue.XMax-vue.XMin+1;
	var index = ((x-vue.XMin)%W)+(W*(y-vue.YMin));
	//~ console.log('('+x+','+y+') -> '+index);
	return vue.matrix ? vue.matrix[index] : null;
}

// renvoie une cellule (en la créant si nécessaire, ne pas utiliser cette méthode en simple lecture)
function getCellVueCreate(vue, x, y) {
	var W = vue.XMax-vue.XMin+1;
	var index = ((x-vue.XMin)%W)+(W*(y-vue.YMin));
	var cell = vue.matrix[index];
	if (!cell) {
		cell = {};
		cell.bralduns = [];
		cell.objets = [];
		cell.monstres = [];
		cell.nbBraldunsFémininsNonKO=0; 
		cell.nbBraldunsMasculinsNonKO=0;
		vue.matrix[index] = cell;
	}
	return cell;
}

// dessine de 1 à 5 icônes sur ou autour d'un point de l'écran
Map.prototype.drawIcons = function(c, sx, sy, icons, hover) {
	var x=[]; var y=[];
	switch (icons.length) {
	case 1:
		x[0]=sx; y[0]=sy;
		break;
	case 2:
		var d = 0.17*this.zoom;
		x[0]=sx; y[0]=sy-d;
		x[1]=sx; y[1]=sy+d;
		break;
	case 3:
		var d = 0.17*this.zoom;
		x[0]=sx;   y[0]=sy-d;
		x[1]=sx-d; y[1]=sy+d;
		x[2]=sx-d; y[2]=sy+d;
		break;
	case 4:
	default: // si le cas se présente, il faut implémenter le cas 5
		var d = 0.17*this.zoom;
		x[0]=sx-d; y[0]=sy-d;
		x[1]=sx+d; y[1]=sy-d;
		x[2]=sx-d; y[2]=sy+d;
		x[3]=sx+d; y[3]=sy+d;
	}
	var imgh;
	if (this.zoom!=64) imgh=this.zoom*0.35;
	for (var i=x.length; i-->0;) {
		if (hover) {
			drawCenteredImage(c, this.getOutlineImg(icons[i]), x[i], y[i], null, imgh?imgh+4:null);
		}
		drawCenteredImage(c, icons[i], x[i], y[i], null, imgh?imgh+4:null);
	}
}

// dessine la vue d'un Braldun (la partie intersectant xMin, xMax, etc.)
Map.prototype.drawVue = function(vue, xMin, xMax, yMin, yMax) {
	var c = this.context;
	
	if (this.zoom>30) {

		//> on compile les objets de la vue sous forme matricielle et on compte les 
		//   bralduns de chaque sexe
		if (!vue.matrix) {
			vue.matrix = [];
			for (ib in vue.Bralduns) {
				var b = vue.Bralduns[ib];
				var cell = getCellVueCreate(vue, b.X, b.Y)
				cell.bralduns.push(b);
			}
			for (io in vue.Objets) {
				var o = vue.Objets[io];
				var cell = getCellVueCreate(vue, o.X, o.Y);
				cell.objets.push(o);
			}
			for (io in vue.Monstres) {
				var o = vue.Monstres[io];
				var cell = getCellVueCreate(vue, o.X, o.Y);
				cell.monstres.push(o);
			}
			//> pour chaque cellule on construit les tableaux d'images par zones
			for (var x=vue.XMin; x<=vue.XMax; x++) {
				for (var y=vue.YMin; y<=vue.YMax; y++) {
					var cell = getCellVue(vue, x, y);
					if (cell) {
						cell.zones = [[], [], [], []]; // 4 zones : haut-gauche, centre, bas-gauche et bas-droit (haut-droit n'est pas géré dans la vue et correspond au lieu)
						var nbBraldunsFémininsNonKO=0; 
						var nbBraldunsMasculinsNonKO=0;
						var nbBraldunsKO=0;
						for (var i=0; i<cell.bralduns.length; i++) {
							var b = cell.bralduns[i];
							if (b.KO) {
								nbBraldunsKO++;
							} else {
								if (b.Sexe=='f') nbBraldunsFémininsNonKO++;
								else nbBraldunsMasculinsNonKO++;
							}
						}
						//-- zone 0 : bralduns
						var imgb = null;
						if (nbBraldunsFémininsNonKO>0 && nbBraldunsMasculinsNonKO>0) {
							imgb = this.img_bralduns_masculin_feminin;					
						} else if (nbBraldunsFémininsNonKO>1) {
							imgb = this.img_bralduns_feminin;
						} else if (nbBraldunsMasculinsNonKO>1) {
							imgb = this.img_bralduns_masculin;
						} else if (nbBraldunsFémininsNonKO>0) {
							imgb = this.img_braldun_feminin;
						} else if (nbBraldunsMasculinsNonKO>0) {
							imgb = this.img_braldun_masculin;
						}
						if (imgb) cell.zones[0].push(imgb);
						//-- zone 0 : monstres
						if (cell.monstres.length>0) {
							var imgbase =  this.imgMonstres[cell.monstres[0].IdType];							
							var img = imgbase ? (cell.monstres.length==1 ? imgbase.a : imgbase.b) : this.imgMonstreInconnu;
							// on vérifie l'homogénéïté
							for (var i=1; i<cell.monstres.length; i++) {
								if (cell.monstres[0].IdType!=cell.monstres[i].IdType) {
									img = this.imgMultiMonstres;
									break;
								}
							}
							cell.zones[0].push(img);
						}
						//-- zone 1 : buisson
						//   TODO
						//-- zone 2 : braldun KO
						if (nbBraldunsKO) cell.zones[2].push(this.img_braldun_ko);
						//-- zone 2 : cadavre
						//   TODO
						//-- zones 2 et 3 : objets, triés suivant leur type et orientés dans l'une des deux zones
						if (cell.objets.length>0) {
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
								var img = this.imgObjets[o.Type];
								if (img) {
									dest.push(img);
								} else {
									console.log("pas d'image pour l'objet " + o.Type);
								}
							}
						}
					}
				}
			}
		} // fin compilation vue
		
		//> on dessine les trucs en vue
		xMin = Math.max(xMin, vue.XMin);
		xMax = Math.min(xMax, vue.XMax);
		yMin = Math.max(yMin, vue.YMin);
		yMax = Math.min(yMax, vue.YMax);
		for (var x=xMin; x<=xMax; x++) {
			for (var y=yMin; y<=yMax; y++) {
				var cell = getCellVue(vue, x, y);
				if (cell) {
					var hover = (this.pointerX==x && this.pointerY==y);
					var d = this.zoom*0.25;
					var cx = this.zoom*(this.originX+x);
					var cy = this.zoom*(this.originY-y);
					if (cell.zones[0].length>0) this.drawIcons(c, cx+d, cy+d, cell.zones[0], hover);
					if (cell.zones[1].length>0) this.drawIcons(c, cx+2*d, cy+2*d, cell.zones[1], hover);
					if (cell.zones[2].length>0) this.drawIcons(c, cx+d, cy+3*d, cell.zones[2], hover);
					if (cell.zones[3].length>0) this.drawIcons(c, cx+3*d, cy+3*d, cell.zones[3], hover);
					if (hover) {
						if (cell.bralduns.length>0) {
							this.bubbleText.push('Bralduns :');
							for (var ib=0; ib<cell.bralduns.length; ib++) {
								var b = cell.bralduns[ib];
								this.bubbleText.push('  '+b.Prénom+' '+b.Nom+' (niveau '+b.Niveau+')');
							}
						}
						if (cell.monstres.length>0) {
							this.bubbleText.push('Monstres :');
							for (var ib=0; ib<cell.monstres.length; ib++) {
								var o = cell.monstres[ib];
								this.bubbleText.push('  '+o.Nom+' '+o.Taille+' (niveau '+o.Niveau+')');
							}
						}
						if (cell.objets.length>0) {
							this.bubbleText.push('Au sol :');
							for (var ib=0; ib<cell.objets.length; ib++) {
								var o = cell.objets[ib];
								if (o.Quantité) this.bubbleText.push('  '+o.Quantité+' '+o.Type+(o.Quantité>1?'s':''));
								else this.bubbleText.push('  '+o.Type);
							}
						}
					}
				}
			}
		}
		
	}
}
