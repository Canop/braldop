
function Map(canvasId, posmarkid) {
	this.canvas = document.getElementById(canvasId);
	this.context = this.canvas.getContext("2d");
	this.posmarkdiv = document.getElementById(posmarkid);
	this.initTiles();
	this.screenRect = new Rect();
	this.rect = new Rect(); // le rectangle englobant le contenu que l'on veut montrer
	this.originX=0; // coin haut gauche de la grotte au centre de l'écran
	this.originY=0;
	this.scales = [0.5, 1, 2, 4, 8, 16, 32, 48, 64]; // éliminé : 92
	this.scaleIndex = 8; // -> zoom "naturel" de 64
	this.zoom = this.scales[this.scaleIndex];
	this.W = 900; // simplement un nombre supérieur à la demi-largeur ou demi-hauteur de la carte mais pas trop
	this.mouseIsDown=false;
	this.pointerX = 0; // coordonnées du pointeur dans l'univers Braldahim
	this.pointerY = 0;
	this.pointerScreenX = 0; // coordonnées du pointeur dans le référentiel de l'écran
	this.pointerScreenY = 0;
	this.hoverObject = null; // cette notion sera remplacée à terme par une cellule (mais restera null si la cellule ne contient rien d'intéressant)
	this.photoSatellite = new Image();
	this.displayPhotoSatellite = true;
	this.displayRégions = false;
	this.recomputeCanvasPosition();
	var _this = this;
	this.photoSatelliteOK = false;
	this.photoSatellite.src = "http://static.braldahim.com/images/sources/harilinn/braldahim_carte4.png";
	this.photoSatellite.onload = function(){
		_this.photoSatelliteRect = new Rect();
		var ps = _this.photoSatellite;
		var ratioSatellite = 1.5;
		var psw = ps.width*ratioSatellite;
		var psh = ps.height*ratioSatellite;
		_this.photoSatelliteRect.x = -0.5*psw -1; // dernier nombre : ajustement manuel
		_this.photoSatelliteRect.y = 0.5*psh - 27; // dernier nombre : ajustement manuel
		_this.photoSatelliteRect.w = psw;
		_this.photoSatelliteRect.h = psh;
		_this.photoSatelliteScreenRect = new Rect();		
		_this.photoSatelliteOK = true;
		_this.redraw();
	};

	// Gestion de la molette, au dessus ou non, de la carte
	this.mouseOnMap = false;
	$('#'+canvasId).mouseover(function() {
		_this.mouseOnMap = true;
		document.body.style.overflow = "hidden";
	}).mouseout(function() {
		_this.mouseOnMap = false;
		document.body.style.overflow = "";
	});

	this.canvas.addEventListener("mousedown", function(e) {_this.mouseDown(e)}, false);
	this.canvas.addEventListener("mouseup", function(e) {_this.mouseUp(e)}, false);
	this.canvas.addEventListener("mouseleave", function(e) {_this.mouseLeave(e)}, false);
	this.canvas.addEventListener("mousemove", function(e) {_this.mouseMove(e)}, false);
	if (window.addEventListener) window.addEventListener("DOMMouseScroll", function(e) {_this.mouseWheel(e)}, false); // firefox
	window.onmousewheel = function(e) {_this.mouseWheel(e)}; // chrome
	$(window).resize(function(){
		_this.recomputeCanvasPosition();
		_this.redraw();
	});
}

// renvoie une cellule (en la créant si nécessaire, ne pas utiliser cette méthode en simple lecture)
Map.prototype.getCellCreate = function(x, y) {
	var index = ((x+this.W)%(2*this.W))+2*this.W*(y+this.W);
	//console.log("("+x+","+y+") -> "+index);
	var cell = this.matrix[index];
	if (!cell) {
		cell = {};
		this.matrix[index] = cell;
	}
	return cell;
}
// renvoie une cellule (en la créant si nécessaire, ne pas utiliser cette méthode en simple lecture)
Map.prototype.getCell = function(x, y) {
	var index = ((x+this.W)%(2*this.W))+2*this.W*(y+this.W);
	return this.matrix[index];
}

Map.prototype.recomputeCanvasPosition = function() {
	var pos = $(this.canvas).position();
	this.canvas_position_x = pos.left;
	this.canvas_position_y = pos.top;
	this.screenRect = new Rect();
	this.screenRect.x = 0;
	this.screenRect.y = 0;
	this.screenRect.w = this.canvas.clientWidth;
	this.screenRect.h = this.canvas.clientHeight;
	this.canvas.width = this.screenRect.w;
	this.canvas.height = this.screenRect.h;
	this.originX = (this.screenRect.w/2)/this.zoom;
	this.originY = (this.screenRect.h/2)/this.zoom;
}

// renvoie l'objet aux coordonnées (univers Braldahim) x et y.
// On optimisera ça plus tard via une matrice (on utilise déjà une matrice pour la partie vue)
Map.prototype.objectOn = function(x, y) {
	if (this.mapData.Vues) {
		for (var i=this.mapData.Vues.length; i-->0;) {
			var vue = this.mapData.Vues[i];
			if (vue.active && vue.matrix) {
				var W = vue.XMax-vue.XMin;
				var index = ((x-vue.XMin)%W)+(W*(y-vue.YMin));
				var cell = vue.matrix[index];
				if (cell) return cell; // la cellule n'est définie que si elle contient quelque chose
			}
		}
	}
	if (this.mapData.LieuxVilles && this.zoom>25) {
		for (var i=this.mapData.LieuxVilles.length; i-->0;) {
			var lv = this.mapData.LieuxVilles[i];
			if (lv.X==x && lv.Y==y) return lv;
		}
	}
	if (this.mapData.Champs && this.zoom>25) {
		for (var i=this.mapData.Champs.length; i-->0;) {
			var o = this.mapData.Champs[i];
			if (o.X==x && o.Y==y) return o;
		}
	}
	if (this.mapData.Echoppes && this.zoom>25) {
		for (var i=this.mapData.Echoppes.length; i-->0;) {
			var o = this.mapData.Echoppes[i];
			if (o.X==x && o.Y==y) return o;
		}
	}
	return null;
}
// l'objet passé, reçu en json, devient le fournisseur des données de carte et de vue.
// Les champs dans le nom commence par une minuscule sont définis localement et
//  ceux dont le nom commence par une majuscule proviennent du serveur (cette norme
//  est valable sur toute la hiérarchie des objets de mapData).
// Les données sont copiées dans une structure qui donne un accès par les coordonnées des cases.
Map.prototype.setData = function(mapData) {
	this.mapData = mapData;
	console.log("carte reçue");
	this.matrix = {};//new Array(); // todo benchmarker pour comparer les effets en ram et cpu
	if (this.mapData.Cases) {
		for (var i=this.mapData.Cases.length; i-->0;) {
			var o = this.mapData.Cases[i];
			var c = this.getCell(o.X, o.Y);
			if (c) {
				console.log("doublon!");
			}
			this.getCellCreate(o.X, o.Y).fond = o.Fond;
		}
	}
	if (this.mapData.Champs) {
		for (var i=this.mapData.Champs.length; i-->0;) {
			var o = this.mapData.Champs[i];
			o.Nom = "Champ";
			o.Details = "Propriétaire : "+o.IdBraldun;
			this.getCellCreate(o.X, o.Y).champ=o;
		}
	}
	if (this.mapData.Echoppes) {
		for (var i=this.mapData.Echoppes.length; i-->0;) {
			var o = this.mapData.Echoppes[i];
			o.Details = o.Métier+" : "+o.IdBraldun;
			this.getCellCreate(o.X, o.Y).échoppe=o;
		}
	}
	if (this.mapData.LieuxVilles) {
		for (var i=this.mapData.LieuxVilles.length; i-->0;) {
			var o = this.mapData.LieuxVilles[i];
			this.getCellCreate(o.X, o.Y).lieu=o;
		}
	}
	console.log("carte compilée");
}
// redessine la page. Peut-être appelée de n'importe quel contexte, y compris depuis une méthode de dessin (pour par exemple faire une animation)
Map.prototype.redraw = function() {
	if (this.drawInProgress) {
		this.redrawStacked = true;
		return;
	}
	this.redrawStacked = false;
	try {
		this.drawInProgress = true;
		this.context.fillStyle="#343";
		this.context.fillRect(0, 0, this.screenRect.w, this.screenRect.h);
		this.bubbleText = [];
		if (this.mapData) {
			if (this.displayPhotoSatellite && this.photoSatelliteOK) {
				this.naturalRectToScreenRect(this.photoSatelliteRect, this.photoSatelliteScreenRect);
				this.photoSatelliteScreenRect.drawImage(this.context, this.photoSatellite);
			}
			var xMin = Math.floor(-this.originX);
			var xMax = Math.ceil(this.screenRect.w/this.zoom-this.originX);
			var yMin = -Math.floor(this.screenRect.h/this.zoom-this.originY);
			var yMax = Math.ceil(this.originY);
			
			//~ console.log("xMin="+xMin);
			//~ console.log("xMax="+xMax);
			//~ console.log("yMin="+yMin);
			//~ console.log("yMax="+yMax);
			
			if (this.zoom>1) {
				var screenRect = new Rect();
				screenRect.w = this.zoom;
				screenRect.h = this.zoom;
				for (var x=xMin; x<=xMax; x++) {
					for (var y=yMin; y<=yMax; y++) {
						var cell = this.getCell(x, y);
						if (cell) {
							screenRect.x = this.zoom*(this.originX+x);
							screenRect.y = this.zoom*(this.originY-y);
							var hover = this.pointerX==x && this.pointerY==y;
							if (cell.fond) this.drawFond(screenRect, cell.fond);
							if (cell.champ) this.drawLieu(screenRect, cell.champ, this.champImg, hover);
							else if (cell.échoppe) this.drawLieu(screenRect, cell.échoppe, this.echoppeImg[cell.échoppe.Métier], hover);
							else if (cell.lieu) this.drawLieu(screenRect, cell.lieu, this.placeImg[cell.lieu.IdTypeLieu], hover);
						}
					}
				}
			}
			if (this.mapData.Vues) {
				for (var i=this.mapData.Vues.length; i-->0;) {
					var vue = this.mapData.Vues[i];
					if (vue.active) this.drawVue(vue);
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
			if (this.bubbleText.length>0) {
				this.bubbleText.splice(0, 0, this.pointerX+','+this.pointerY);
				this.drawBubble();
			}
		}
	} finally {
		this.drawInProgress = false;
	}
	if (this.redrawStacked) {
		setTimeout(this.redraw, 40); 
	}
}
Map.prototype.mouseWheel = function(e) {
	if (this.mouseIsDown) return;
	if (!this.mouseOnMap) return;
	//console.log('scaleIndex before = '+this.scaleIndex);
	var delta = 0;
	if (!e) e=window.e;
	if (e.wheelDelta) {
		delta = e.wheelDelta / 120;
	} else if (e.detail) {
		delta = -e.detail / 3;
	}
	var oldZoom = this.zoom;
	if (delta>0) {
		if (this.scaleIndex<this.scales.length-1) {
			this.zoom = this.scales[++this.scaleIndex];
		}
	} else if (this.scaleIndex>0){
		this.zoom = this.scales[--this.scaleIndex];
	}
	var zr = (1/this.zoom-1/oldZoom);
	this.zoomChangedSinceLastRedraw = true;
	var mouseX = e.offsetX; // Chrome
	var mouseY = e.offsetY; // Chrome
	if (!mouseX) {
		mouseX = e.layerX; // FF
		mouseY = e.layerY; // FF
	}
	this.originX += (mouseX-this.canvas_position_x)*zr; 
	this.originY += (mouseY-this.canvas_position_y)*zr;
	this.posmarkdiv.innerHTML='Zoom='+this.zoom+' &nbsp; X='+this.pointerX+' &nbsp; Y='+this.pointerY;
	this.hoverObject = null;
	this.redraw();
}
Map.prototype.mouseDown = function(e) {
	var mouseX = e.offsetX; // Chrome
	var mouseY = e.offsetY; // Chrome
	if (!mouseX) {
		mouseX = e.layerX; // FF
		mouseY = e.layerY; // FF
	}
	this.mouseIsDown = true;
	this.dragStartPageX = mouseX;
	this.dragStartPageY = mouseY;
	this.dragStartOriginX = this.originX;
	this.dragStartOriginY = this.originY;
	this.zoomChangedSinceLastRedraw = true;
	this.hoverObject = null;
	this.redraw();
}
Map.prototype.mouseUp = function(e) {
	this.mouseIsDown = false;
	this.hoverObject = null;
	this.redraw();
}

Map.prototype.mouseLeave = function(e) {
	this.mouseIsDown = false;
	this.hoverObject = null;
	this.redraw();
}

Map.prototype.mouseMove = function(e) {
	if (!this.mapData) return;
	var mouseX = e.offsetX; // Chrome
	var mouseY = e.offsetY; // Chrome
	if (!mouseX) {
		mouseX = e.layerX; // FF
		mouseY = e.layerY; // FF
	}
	this.pointerScreenX = mouseX;
	this.pointerScreenY = mouseY;
	this.pointerX = Math.floor(mouseX/this.zoom-this.originX);
	this.pointerY = -Math.floor(mouseY/this.zoom-this.originY);
	this.posmarkdiv.innerHTML='Zoom='+this.zoom+' &nbsp; X='+this.pointerX+' &nbsp; Y='+this.pointerY;
	if (this.mouseIsDown) {
		var dx = (mouseX-this.dragStartPageX)/this.zoom;
		var dy = (mouseY-this.dragStartPageY)/this.zoom;
		this.originX = this.dragStartOriginX + dx;
		this.originY = this.dragStartOriginY + dy;
		this.redraw();		
	} else {
		var newHoverObject = this.objectOn(this.pointerX, this.pointerY);
		if (newHoverObject!=this.hoverObject) {
			this.hoverObject = newHoverObject;
			this.redraw();
		}
	}
}

Map.prototype.naturalToScreen = function(naturalPoint, screenPoint) {
	screenPoint.x = this.zoom*(this.originX+naturalPoint.x+0.5);
	screenPoint.y = this.zoom*(this.originY-naturalPoint.y+0.5);
};

Map.prototype.screenToNatural = function(screenPoint, naturalPoint) {
	naturalPoint.x = screenPoint.x/this.zoom - this.originX;	
	naturalPoint.y = screenPoint.y/this.zoom - this.originY;
};

Map.prototype.screenRectToNaturalRect = function(screenRect, naturalRect) {
	naturalRect.x = screenRect.x/this.zoom - this.originX;	
	naturalRect.y = screenRect.y/this.zoom - this.originY;
	naturalRect.w = screenRect.w/this.zoom;
	naturalRect.h = screenRect.h/this.zoom;
};

Map.prototype.naturalRectToScreenRect = function(naturalRect, screenRect) {
	screenRect.x = this.zoom*(this.originX+naturalRect.x);
	screenRect.y = this.zoom*(this.originY-naturalRect.y);
	screenRect.w = this.zoom*naturalRect.w;
	screenRect.h = this.zoom*naturalRect.h;
};
