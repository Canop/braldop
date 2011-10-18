Map.prototype.openDialog = function(startingRectInCanvas, title, content) {
	this.dialopIsOpen = true;
	var $canvas = $(this.canvas);
	var $win = $(window);
	var winWidth = $(window).width();
	var winHeight = $(window).height();
	var width = $canvas.width();
	if (width>400) width=400;
	var wx = $canvas.offset().left+this.pointerScreenX;
	var wy = $canvas.offset().top+this.pointerScreenY;
	var maxHeight;
	if (wx<winWidth/2) {
		this.$dialog.css('left', (wx+40)+'px');
		this.$dialog.css('right', (winWidth-wx-width)+'px');
	} else {
		this.$dialog.css('right', (winWidth-wx+40)+'px');
		this.$dialog.css('left', (wx-width)+'px');
	}
	if (wy<winHeight/2) {
		maxHeight = winHeight-wy-90;
		this.$dialog.css('top', (wy-20)+'px');
		this.$dialog.css('bottom', '');
	} else {
		maxHeight = wy-90;
		this.$dialog.css('top', '');
		this.$dialog.css('bottom', (winHeight-wy+20)+'px');
	}
	var html = [];
	var h=0;
	html[h++] = '<span class=dialog_title>';
	html[h++] = title;
	html[h++] = '</span><hr>';
	html[h++] = '<div id=dialog_content></div>';
	html[h++] = '<hr><small>Cliquez pour fermer ce menu</small>';
	this.$dialog.html(html.join(''));
	this.$dialog.show();
	$content = $('#dialog_content');
	$content.css('max-height', maxHeight);
	$content.css('overflow', 'auto');
	$content.html(content);
}

Map.prototype.openCellDialog = function(x, y) {
	var cell = this.getCell(this.couche, x, y);
	var screenRect = new Rect();
	screenRect.w = this.zoom;
	screenRect.h = this.zoom;
	screenRect.x = this.zoom*(this.originX+x);
	screenRect.y = this.zoom*(this.originY-y);
	var html = [];
	var h=0;
	var empty = false;
	if (cell.champ) {
		html[h++] = '<table><tr><td><span class="champ"/></td><td>';
		html[h++] = 'Champ de <a target=winprofil href="http://jeu.braldahim.com/voir/braldun/?braldun='+cell.champ.IdBraldun+'&direct=profil">'+cell.champ.NomCompletBraldun+'</a></td></tr></table>';
		html[h++] = '</td></tr></table>';
	} else if (cell.échoppe) {
		html[h++] = '<table><tr><td><span class="'+cell.échoppe.Métier+'"/></td><td>';
		html[h++] = cell.échoppe.Nom+'<br>';
		html[h++] = cell.échoppe.Métier+' : <a target=winprofil href="http://jeu.braldahim.com/voir/braldun/?braldun='+cell.échoppe.IdBraldun+'&direct=profil">'+cell.échoppe.NomCompletBraldun+'</a></td></tr></table>';
		html[h++] = '</td></tr></table>';
	} else if (cell.lieu) {
		html[h++] = '<table><tr><td><span class="'+this.typesBatiments[cell.lieu.IdTypeLieu]+'"/></td><td> '+cell.lieu.Nom+'</td></tr></table>';
	} else {
		empty = true;
	}
	var cellVue = this.getCellVue(x, y);
	if (cellVue) {
		if (cellVue.action) {
			empty = false;
			var t = this.typesActions[cellVue.action.Type];
			html[h++] = '<table><tr><td>';
			if (t.iconeCase) html[h++] = '<img hspace=5 vspace=2 src="'+t.iconeCase.src+'">'; // le hspace et le vspace là sont paresseux, on changera si plusieurs actions ont des icônes
			html[h++] = '</td><td><a href="javascript:mapDoAction('+cellVue.action.key+');">'+t.nom+'</a></td><td>('+cellVue.action.PA+' PA)</td></tr></table>';
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
				html[h++] = '<span class="'+this.spritesVueTypes.css('monstre_'+o.IdType+'a')+'"/>';
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
	this.openDialog(screenRect, x+","+y, html.join(''));
}
