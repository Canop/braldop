/*
 * Ce fichier met à jour la version de Gout-Gueule intégrée à l'interface de Braldahim
 *  (actuellement en retard par rapport à la version Braldop)
 * 
 */ 
 

Map.prototype.setData = function(mapData) {
	this.mapData = mapData;
	this.matricesVuesParZ = {};
	this.matricesVuesParZ[0]={};
	this.z = 0; // on va basculer forcément sur la couche zéro
	this.couche = null;
	var _this = this;
	for (var ic=0; ic<this.mapData.Couches.length; ic++) {
		var couche = this.mapData.Couches[ic];
		this.couche = couche;
		couche.matrix = {};
		couche.fond = new Image();
		if (couche.Cases) {
			for (var i=couche.Cases.length; i-->0;) {
				var o = couche.Cases[i];
				this.getCellCreate(couche, o.X, o.Y).fond = o.Fond;
			}
		}
		if (couche.Champs) {
			for (var i=couche.Champs.length; i-->0;) {
				var o = couche.Champs[i];
				o.Nom = "Champ";
				this.getCellCreate(couche, o.X, o.Y).champ=o;
			}
		}
		if (couche.Echoppes) {
			for (var i=couche.Echoppes.length; i-->0;) {
				var o = couche.Echoppes[i];
				this.getCellCreate(couche, o.X, o.Y).échoppe=o;
			}
		}
		if (couche.Lieux) {
			for (var i=couche.Lieux.length; i-->0;) {
				var o = couche.Lieux[i];
				this.getCellCreate(couche, o.X, o.Y).lieu=o;
			}
		}
		if (couche.Palissades) {
			for (var i=couche.Palissades.length; i-->0;) {
				var o = couche.Palissades[i];
				o.sides = -1; // sera calculé lors du redraw
				this.getCellCreate(couche, o.X, o.Y).palissade=o; 
			}
		}
	}
	if (!this.mapData.Vues) this.mapData.Vues=[];
	this.mapData.Vues.sort(function(a, b) {
		return a.Time-b.Time;
	});
	//  les lieux de ville (pour l'instant ?) n'ont pas de profondeur explicite mais ne concernent que la surface. On les met dans la couche zéro
	if (this.mapData.LieuxVilles) {
		for (var i=this.mapData.LieuxVilles.length; i-->0;) {
			var o = this.mapData.LieuxVilles[i];
			this.getCellCreate(this.couche, o.X, o.Y).lieu=o;
		}
	}
	if (mapData.Actions) {
		for (var ia=mapData.Actions.length; ia-->0;) {
			var a = mapData.Actions[ia];
			a.key = this.actions.length; // on donne à l'action une clef pour la retrouver plus facilement
			this.actions.push(a);
			// on ajoute les actions à la vue (trouvée par l'acteur)
			var vue;
			if (this.mapData.Vues) {
				for (var i=this.mapData.Vues.length; i-->0;) {
					if (this.mapData.Vues[i].Voyeur==a.Acteur) {
						vue = this.mapData.Vues[i];
						break;
					}
				}
			}
			if (!vue) {
				console.log('Vue non trouvée pour action');
				continue;
			}
			if (!vue.actions) vue.actions = []; // les actions seront attachées à leur case d'effet éventuelle dans compileLesVues
			vue.actions.push(a);
		}
	}
	this.compileLesVues();
	this.matriceVues = this.matricesVuesParZ[0];
	if (this.onSetData) this.onSetData();
}


color2envs = {
	11185755: "hetres", 11119962: "hetres",
	9551989: "plaine-gr", 9486196: "plaine-gr",
	8698773: "marais-gr", 8632980: "marais-gr",
	9951580: "peupliers", 9885787: "peupliers",
	12381606: "gazon-gr", 12315813: "gazon-gr",
	7117918: "chenes", 7052125: "chenes",
	12097633: "tunnel", 12031840: "tunnel",
	10255739: "mine", 10189946: "mine",
	5128548: "caverne-crevasse", 5062755: "caverne-crevasse",
	11958351: "erables-gr", 11892558: "erables-gr",
	13880266: "pave", 13814473: "pave",
	12116936: "marais", 12051143: "marais",
	11589980: "peupliers-gr", 11524187: "peupliers-gr",
	13026235: "route", 12960442: "route",
	13810090: "montagne", 13744297: "montagne",
	6197187: "lac", 6131394: "lac",
	4877977: "profonde", 4812184: "profonde",
	11310219: "montagne-gr", 11244426: "montagne-gr",
	12049793: "plaine", 11984000: "plaine",
	11830875: "erables", 11765082: "erables",
	10719647: "caverne", 10653854: "caverne",
	6589635: "peuprofonde", 6523842: "peuprofonde",
	7916106: "gazon", 7850313: "gazon",
	8029769: "hetres-gr", 7963976: "hetres-gr",
	0: null
};


// assure, si cela est possible, que la couche contient les pixels du fond (imageData)
//  lesquels pourront être utilisés pour déterminer l'environnement de chaque case.
// Renvoie true ssi les pixels du fond sont disponibles.
Map.prototype.initializePixelsFond = function(couche) {
	if (couche.getFond) return true;
	if (couche.fond.width) {
		var tempCanvas = document.createElement('canvas');
		tempCanvas.width = 1600;
		tempCanvas.height = 1000;
		var context = tempCanvas.getContext('2d');
		context.drawImage(couche.fond, 0, 0);
		var pixels = context.getImageData(0, 0, 1600, 1000).data;
		couche.getFond = function(x, y) {
			var i = 4*(x+800+1600*(500-y)); // 4 pour les 4 composantes rgba
			var color = (pixels[i]<<16) + (pixels[i+1]<<8) + (pixels[i+2]);
			return color2envs[color];
		};
		couche.aPalissade = function(x, y) { // les points à palissade sont ceux dont l'alpha vaut 254/255
			var i = 4*(x+800+1600*(500-y)) + 3;
			return pixels[i]==254;
		}
		return true;
	}
	return false;
}

Map.prototype.redraw = function() {
	if (this.drawInProgress) {
		this.redrawStacked = true;
		return;
	}
	this.redrawStacked = false;
	if (!(this.spritesVueTypes.ready&&this.spritesEnv.ready)) {
		return;
	}
	try {
		this.drawInProgress = true;
		if (this.onload) {
			this.onload();
			this.onload = null;
		}
		if (this.mapData) {
			this.context.fillStyle="#343";
			this.context.fillRect(0, 0, this.screenRect.w, this.screenRect.h);
			if (this.displayPhotoSatellite && this.photoSatelliteOK) {
				this.naturalRectToScreenRect(this.photoSatelliteRect, this.photoSatelliteScreenRect);
				this.photoSatelliteScreenRect.drawImage(this.context, this.photoSatellite);
			}
			
			// un carambar au premier qui pourra me réduire le paragraphe qui suit sans diminuer les perfs
			this.xMin = Math.floor(-this.originX);
			this.xMax = Math.ceil(this.screenRect.w/this.zoom-this.originX);
			this.yMin = -Math.floor(this.screenRect.h/this.zoom-this.originY);
			this.yMax = Math.ceil(this.originY);
			if (this.xMin<-800) {
				this.xMin=-800;
				if (this.xMax<-800) this.xMax=-800;
			}
			if (this.xMax>800) {
				this.xMax=800;
				if (this.xMin>800) this.xMin=800;
			}
			if (this.yMin<-500) {
				this.yMin=-500;
				if (this.yMax<-500) this.xMax=-500;
			}
			if (this.yMax>500) {
				this.yMax=500;
				if (this.yMin>500) this.yMin=500;
			}
			if (this.zoom>2) {
				var envInPngAvailable = false;
				if (!this.couche.Cases) envInPngAvailable = this.initializePixelsFond(this.couche);
				var screenRect = new Rect(); // rectangle d'une cellule en coordonnées canvas
				screenRect.w = this.zoom;
				screenRect.h = this.zoom;
				for (var x=this.xMin; x<=this.xMax; x++) {
					for (var y=this.yMax; y>=this.yMin; y--) { // on balaie en commencant par le haut de l'écran (plus "loin" en perspective)
						screenRect.x = this.zoom*(this.originX+x);
						screenRect.y = this.zoom*(this.originY-y);
						if (envInPngAvailable) {
							var fond = this.couche.getFond(x, y);
							if (fond) this.drawFond(screenRect, fond);
						//~ } else if (cell.fond) {
							//~ this.drawFond(screenRect, cell.fond);
						}
						var cell = this.getCell(this.couche, x, y);
						if (cell) {
							var hover = this.zoom>20 && this.pointerX==x && this.pointerY==y;
							if (cell.champ) this.drawLieu(screenRect, cell.champ, this.spritesVueTypes.get('champ'), hover);
							else if (cell.échoppe) this.drawLieu(screenRect, cell.échoppe, this.spritesVueTypes.get(cell.échoppe.Métier), hover);
							else if (cell.lieu) this.drawLieu(screenRect, cell.lieu, this.spritesVueTypes.get('lieu_' + cell.lieu.IdTypeLieu), hover);
						}
					}
				}
				//~ console.log(colorMap);
			} else if (this.couche.fond.width) { // si l'image de fond obtenue du serveur est disponible, on l'utilise pour les basses résolutions
				var sw = this.xMax-this.xMin;
				var sh = this.yMax-this.yMin;
				this.context.drawImage(
					this.couche.fond,
					this.xMin+800, 500-this.yMax, sw, sh,
					this.zoom*(this.originX+this.xMin), this.zoom*(this.originY-this.yMax), this.zoom*sw, this.zoom*sh
				);
			}
			if (this.zoom>15 && this.displayGrid) {
				this.drawGrid();
			}
			if (this.zoom>2) { // on dessine les palissades après avoir dessiné la grille pour qu'elle ne les recouvre pas
				var screenRect = new Rect();
				screenRect.w = this.zoom;
				screenRect.h = this.zoom;
				if (this.couche.aPalissade) {
					for (var x=this.xMin; x<=this.xMax; x++) {
						for (var y=this.yMax; y>=this.yMin; y--) { // on balaie en commencant par le haut de l'écran (plus "loin" en perspective)
							if (this.couche.aPalissade(x, y)) {
								var cell = this.getCellCreate(this.couche, x, y);
								if ((!cell.palissade) || cell.palissade.sides==-1) {
									if (!cell.palissade) { // présence palissade reçue en png
										cell.palissade = {
											X:x, Y:y, Z:this.z, sides:0, png:true // png==true ==> pas de données
										};
									} else { // objet palissade reçu en json
										cell.palissade.sides = 0;
									}
									p = cell.palissade;
									var nb=0;
									if (this.couche.aPalissade(p.X+1, p.Y)) {p.sides |= B_RIGHT; nb++;}
									if (this.couche.aPalissade(p.X-1, p.Y)) {p.sides |= B_LEFT; nb++;}
									if (this.couche.aPalissade(p.X, p.Y+1)) {p.sides |= B_TOP; nb++;}
									if (this.couche.aPalissade(p.X, p.Y-1)) {p.sides |= B_BOTTOM; nb++;}
									if (nb==1) {
										if ((p.sides&B_LEFT)&&(!this.couche.getFond(p.X+1, p.Y))) p.sides|=B_RIGHT;
										else if ((p.sides&B_TOP)&&(!this.couche.getFond(p.X, p.Y-1))) p.sides|=B_BOTTOM;
										else if ((p.sides&B_RIGHT)&&(!this.couche.getFond(p.X-1, p.Y))) p.sides|=B_LEFT;
										else if ((p.sides&B_BOTTOM)&&(!this.couche.getFond(p.X, p.Y+1))) p.sides|=B_TOP;
									} else if (nb==0) {
										if ((!this.couche.getFond(p.X-1, p.Y))&&(!this.couche.getFond(p.X, p.Y+1))) p.sides|=B_LEFT|B_TOP;
										else if ((!this.couche.getFond(p.X-1, p.Y))&&(!this.couche.getFond(p.X, p.Y-1))) p.sides|=B_LEFT|B_BOTTOM;
										else if ((!this.couche.getFond(p.X+1, p.Y))&&(!this.couche.getFond(p.X, p.Y-1))) p.sides|=B_RIGHT|B_BOTTOM;
										else if ((!this.couche.getFond(p.X+1, p.Y))&&(!this.couche.getFond(p.X, p.Y+1))) p.sides|=B_RIGHT|B_TOP;										
									}
								}								
								screenRect.x = this.zoom*(this.originX+x);
								screenRect.y = this.zoom*(this.originY-y);
								this.drawPalissade(screenRect, cell.palissade);
							}
						}
					}
				} else {
					for (var x=this.xMin; x<=this.xMax; x++) {
						for (var y=this.yMax; y>=this.yMin; y--) { // on balaie en commencant par le haut de l'écran (plus "loin" en perspective)
							var cell = this.getCell(this.couche, x, y);
							if (cell && cell.palissade) {
								screenRect.x = this.zoom*(this.originX+x);
								screenRect.y = this.zoom*(this.originY-y);
								this.drawPalissade(screenRect, cell.palissade);
							}
						}
					}
				}
			}
			if (this.mapData.Vues) {
				if (this.zoom>10) {
					this.dessineLesVues();
				}
				if (this.displayFog) {
					this.drawFog();
				}
			}
			if (this.mapData.Villes && this.zoom<=60) {
				for (var i=this.mapData.Villes.length; i-->0;) {
					this.drawTown(this.mapData.Villes[i]);
				}
			}
			if (this.displayRégions && this.mapData.Régions) {
				for (var i=this.mapData.Régions.length; i-->0;) {
					this.drawRégion(this.mapData.Régions[i]);
				}
			}
		}
	} finally {
		this.drawInProgress = false;
	}
	if (this.redrawStacked) {
		var _this = this;
		setTimeout(function(){_this.redraw();}, 40); 
	}
}

Map.prototype.openCellDialog = function(x, y, fixed) {
	var cell = this.getCell(this.couche, x, y);
	var screenRect = new Rect();
	screenRect.w = this.zoom;
	screenRect.h = this.zoom;
	screenRect.x = this.zoom*(this.originX+x);
	screenRect.y = this.zoom*(this.originY-y);
	var html = [];
	var h=0;
	var empty = true;
	if (cell) {
		if (cell.palissade) {
			empty = false;
			html[h++] = "<b>Palissade";
			if (!cell.palissade.png) {
				if (!cell.palissade.Destructible) html[h++] = " indestructible";
				html[h++] = "</b>";
				if (cell.palissade.Destructible && cell.palissade.TimeFin) {
					html[h++] = ' (date de fin : ' + formatDate(cell.palissade.TimeFin*1000, true) + ')';
				}
			}
			html[h++] = '<br>';
			empty=false;
		} else if (cell.champ) {
			html[h++] = '<table><tr><td><span class="champ"/></td><td>';
			html[h++] = 'Champ de <a target=winprofil href="http://jeu.braldahim.com/voir/braldun/?braldun='+cell.champ.IdBraldun+'&direct=profil">'+cell.champ.NomCompletBraldun+'</a></td></tr></table>';
			html[h++] = '</td></tr></table>';
			empty=false;
		} else if (cell.échoppe) {
			html[h++] = '<table><tr><td><span class="'+cell.échoppe.Métier+'"/></td><td>';
			html[h++] = cell.échoppe.Nom+'<br>';
			html[h++] = cell.échoppe.Métier+' : <a target=winprofil href="http://jeu.braldahim.com/voir/braldun/?braldun='+cell.échoppe.IdBraldun+'&direct=profil">'+cell.échoppe.NomCompletBraldun+'</a></td></tr></table>';
			html[h++] = '</td></tr></table>';
			empty=false;
		} else if (cell.lieu) {
			html[h++] = '<table><tr><td><span class="lieu_'+cell.lieu.IdTypeLieu+'"/></td><td> '+cell.lieu.Nom+'</td></tr></table>';
			empty=false;
		}
	}
	var cellVue = this.getCellVue(x, y);
	if (cellVue) {
		if (cellVue.actions.length) {
			empty = false;
			html[h++] = '<table>';
			for (var ia=cellVue.actions.length; ia-->0;) {
				var a = cellVue.actions[ia];
				var t = this.typesActions[a.Type];
				html[h++] = '<tr><td>';
				if (t.icone) html[h++] = '<img hspace=5 vspace=2 src="'+t.icone.src+'">'; // le hspace et le vspace là sont paresseux, on changera si plusieurs actions ont des icônes
				html[h++] = '</td><td><a href="javascript:mapDoAction('+a.key+');">'+t.nom+'</a></td><td>('+a.PA+' PA)</td></tr>';
			}
			html[h++] = '</table>';
		}
		if (cellVue.bralduns.length) {
			empty = false;
			html[h++] = "<b>Braldûns :</b>";
			html[h++] = '<table>';
			for (var ib=0; ib<cellVue.bralduns.length; ib++) {
				var b = cellVue.bralduns[ib];
				var key = b.KO ? 'braldunKo' : ( b.Sexe=='f' ? 'braldun_feminin' : 'braldun_masculin' );
				if (b.Camp.length) key += '-'+b.Camp;
				html[h++] = '<tr><td>';
				html[h++] = '<span class="'+key+'"></span>';
				html[h++] = '</td><td><a target=winprofil href="http://jeu.braldahim.com/voir/braldun/?braldun='+b.Id+'&direct=profil">'+b.Prénom+' '+b.Nom+'</a></td><td>niv. '+b.Niveau;
				html[h++] = '</td><td>';
				if (b.IdCommunauté>0) html[h++] =  this.mapData.Communautés[b.IdCommunauté].Nom;
				html[h++] = '</td><td>';
				if (b.PointsGredin) html[h++] = '<span class=pointsGredin>'+b.PointsGredin+'</span>';
				if (b.PointsRedresseur) html[h++] = '<span class=pointsRedresseur>'+b.PointsRedresseur+'</span>';
				html[h++] = '</td></tr>';
			}
			html[h++] = '</table>';
		}
		if (cellVue.monstres.length) {
			empty = false;
			html[h++] = "<b>Monstres :</b>";
			html[h++] = '<table>';
			for (var ib=0; ib<cellVue.monstres.length; ib++) {
				var o = cellVue.monstres[ib];
				html[h++] = '<tr><td>';
				html[h++] = '<span class="'+this.spritesVueTypes.css('monstre_'+o.IdType+'a', 'monstre')+'"/>';
				html[h++] = '</td><td><a target=winprofil href="http://jeu.braldahim.com/voir/monstre/?monstre='+o.Id+'">'+o.Nom+' '+o.Taille+'</a>';
				html[h++] = '</td></tr>';
			}
			html[h++] = '</table>';
		}
		if (cellVue.cadavres.length) {
			empty = false;
			html[h++] = "<b>Cadavres :</b>";
			html[h++] = '<table>';
			for (var ib=0; ib<cellVue.cadavres.length; ib++) {
				var o = cellVue.cadavres[ib];
				html[h++] = '<tr><td>';
				html[h++] = '<span class="cadavre"></span>';
				html[h++] = '</td><td><a target=winprofil href="http://jeu.braldahim.com/voir/monstre/?monstre='+o.Id+'">'+o.Nom+' '+o.Taille+'</a>';
				if (o.Gibier) html[h++] = ' (gibier)';
				html[h++] = '</td></tr>';
			}
			html[h++] = '</table>';
		}
		if (cellVue.objets.length) {
			empty = false;
			html[h++] = "<b>Au sol :</b>";
			html[h++] = '<table>';
			for (var ib=0; ib<cellVue.objets.length; ib++) {
				var o = cellVue.objets[ib];
				html[h++] = '<tr><td>';
				html[h++] = '<span class="'+this.getObjectImgKey(o)+'"></span>';
				html[h++] = '</td><td>';
				html[h++] = '  '+o.Label;
				html[h++] = '</td></tr>';
			}
			html[h++] = '</table>';
		}
	}
	if (empty) html[h++] = "<i>Il n'y a rien ici</i>";
	this.openDialog(x+","+y, html.join(''), fixed);
}

Map.prototype.updatePosDiv = function() {
	var html = 'X='+this.pointerX+' &nbsp; Y='+this.pointerY+' &nbsp; Z='+this.z;
	var cell = this.getCell(this.couche, this.pointerX, this.pointerY);
	var fond;
	if (this.couche.getFond) {
		fond = this.couche.getFond(this.pointerX, this.pointerY);
	} else if (cell) {
		fond = cell.fond
	} else {
		console.log("Caractéristiques cases introuvables", this.pointerX, this.pointerY);
		return;
	}
	var env = this.environnements[fond];
	if (env) html += ' ' + env.nom + ', ' + env.description;
	else console.log('env inconnu : ' + fond); // notons qu'on a des undefined quand il n'y a pas de terrain sous des palissades par exemple
	this.$posmarkdiv.html(html);
}
