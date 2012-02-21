// remplacement des fonctions de l'interface que l'on trouve normalement dans Map_dialog.js

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
