/*
pour tester des trucs...

*/



Map.prototype.dessineCasePalissade = function (x, y, sides) {
	var c = this.context;
	var cx = this.zoom*(this.originX+x);
	var cy = this.zoom*(this.originY-y);
	var r = 0.68;
	var lt = imgTroncPalissade.width*r; // largeur tronc
	var rz = this.zoom/64;
	if (lt==0) {
		console.log('not yet loaded');
		return;
	}
	//~ console.log('lt='+lt);
	var ht = imgTroncPalissade.height*r;
	switch (sides) {
		
		case B_LEFT|B_RIGHT:
		var nbt = Math.ceil(64/lt);
		var by = cy + 0.5*this.zoom;
		var bx = cx;
		var lta = this.zoom/nbt;
		for (var i=0; i<nbt; i++) {
			bx += lta;
			c.drawImage(imgTroncPalissade, bx-lt*0.5, (by-ht), rz*lt, rz*ht);
		}
		break;
		
		case B_LEFT|B_BOTTOM:
		var nbt = Math.round(64/lt)*1.57;
		var lta = lt*0.5;
		var angle = Math.PI*1.5;
		for (var i=0; i<=nbt; i++) {
			angle += Math.PI*0.5/nbt;
			var bx = cx+(Math.cos(angle))*this.zoom*0.5;
			var by = cy+(1+Math.sin(angle)*0.5)*this.zoom;
			c.drawImage(imgTroncPalissade, bx-lta, by-ht, rz*lt, rz*ht);
		}
		break;

		case B_TOP|B_BOTTOM:
		var nbt = Math.ceil(64/lt);
		var bx = cx + 0.5*this.zoom;
		var by = cy;
		var lta = this.zoom/nbt;
		for (var i=0; i<nbt; i++) {
			by += lta;
			c.drawImage(imgTroncPalissade, bx-lt*0.5, by-ht, rz*lt, rz*ht);
		}
		break;
		
		case B_TOP|B_RIGHT:
		var nbt = Math.round(64/lt)*1.57;
		var lta = lt*0.5;
		var angle = Math.PI;
		for (var i=0; i<=nbt; i++) {
			angle -= Math.PI*0.5/nbt;
			var bx = cx+(1+Math.cos(angle)*0.5)*this.zoom;
			var by = cy+(Math.sin(angle)*0.5)*this.zoom;
			c.drawImage(imgTroncPalissade, bx-lta, by-ht, rz*lt, rz*ht);
		}
		break;

		case B_TOP|B_LEFT:
		var nbt = Math.round(64/lt)*1.57;
		var lta = lt*0.5;
		var angle = 0;
		for (var i=0; i<=nbt; i++) {
			angle += Math.PI*0.5/nbt;
			var bx = cx+(Math.cos(angle)*0.5)*this.zoom;
			var by = cy+(Math.sin(angle)*0.5)*this.zoom;
			c.drawImage(imgTroncPalissade, bx-lta, by-ht, rz*lt, rz*ht);
		}
		break;

		case B_RIGHT|B_BOTTOM:
		var nbt = Math.round(64/lt)*1.57;
		var lta = lt*0.5;
		var angle = Math.PI*1.5;
		for (var i=0; i<=nbt; i++) {
			var bx = cx+(1+Math.cos(angle)*0.5)*this.zoom;
			var by = cy+(1+Math.sin(angle)*0.5)*this.zoom;
			c.drawImage(imgTroncPalissade, bx-lta, by-ht, rz*lt, rz*ht);
			angle -= Math.PI*0.5/nbt;
		}
		break;
	}
}


Map.prototype.dessinePalissades = function() {
	var x0 = -17; var y0 = 3;
	this.dessineCasePalissade(x0-1, y0, B_LEFT|B_RIGHT);
	this.dessineCasePalissade(x0, y0, B_LEFT|B_RIGHT);
	this.dessineCasePalissade(x0+1, y0, B_LEFT|B_BOTTOM);

	this.dessineCasePalissade(x0-2, y0, B_RIGHT|B_BOTTOM);
	
	this.dessineCasePalissade(x0+1, y0-1, B_TOP|B_BOTTOM);

	this.dessineCasePalissade(x0-4, y0-1, B_RIGHT|B_BOTTOM);
	this.dessineCasePalissade(x0-3, y0-1, B_LEFT|B_RIGHT);
	this.dessineCasePalissade(x0-2, y0-1, B_LEFT|B_TOP);
	
	this.dessineCasePalissade(x0+1, y0-2, B_TOP|B_RIGHT);
	this.dessineCasePalissade(x0+2, y0-2, B_LEFT|B_BOTTOM);

	this.dessineCasePalissade(x0-4, y0-2, B_TOP|B_BOTTOM);

	this.dessineCasePalissade(x0+2, y0-3, B_TOP|B_BOTTOM);

	this.dessineCasePalissade(x0-2, y0-3, B_LEFT|B_BOTTOM);
	this.dessineCasePalissade(x0-3, y0-3, B_LEFT|B_RIGHT);
	this.dessineCasePalissade(x0-4, y0-3, B_TOP|B_RIGHT);

	this.dessineCasePalissade(x0+2, y0-4, B_TOP|B_LEFT);
	this.dessineCasePalissade(x0+1, y0-4, B_LEFT|B_RIGHT);
	this.dessineCasePalissade(x0, y0-4, B_LEFT|B_RIGHT);
	this.dessineCasePalissade(x0-1, y0-4, B_LEFT|B_RIGHT);
	this.dessineCasePalissade(x0-2, y0-4, B_TOP|B_RIGHT);
	
}
