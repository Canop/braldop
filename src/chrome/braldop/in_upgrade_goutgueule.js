/*
 * Ce fichier met à jour la version de Gout-Gueule intégrée à l'interface de Braldahim
 *  (actuellement en retard par rapport à la version Braldop)
 * 
 */ 


color2envs = {
	0: null,
	4877977: "profonde",
	5128548: "caverne-crevasse",
	6197187: "lac",
	6589635: "peuprofonde",
	7117918: "chenes",
	7916106: "gazon",
	8029769: "hetres-gr",
	9551989: "plaine-gr",
	9951580: "peupliers",
	10255739: "mine",
	10719647: "caverne",
	11185755: "hetres",
	11310219: "montagne-gr",
	11830875: "erables",
	12049793: "plaine",
	12097633: "tunnel",
	12116936: "marais",
	12381606: "gazon-gr",
	13026235: "route",
	13810090: "montagne",
	13880266: "pave"
};

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
			//if (pixels[i+3]==0) return null;
			var color = (pixels[i]<<16) + (pixels[i+1]<<8) + (pixels[i+2]);
			//~ console.log(pixels[i], pixels[i+1], pixels[i+2], '  ->  ', color, color2envs[color]);
			return color2envs[color];
		};
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
			if (this.mapData.Vues) {
				if (this.zoom>30) {
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
