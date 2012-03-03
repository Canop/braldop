

braldop.$mapSettings = null;
braldop.$depthMenu = null;
braldop.$globalMenu = null;
braldop.globalMenuOpen = false;


braldop.ensureMapSettings = function() {
	if (!braldop.$mapSettings) {
		var html = '<div id=bra_map_settings>';
		html += '<select id=bra_depth>';
		html += '</select>';
		html += '<span id="bra_triangle_vues">&#x25bc;</span>';
		html += '<div id="bra_menu">BIP!</div>';
		html += '</div>';
		
		braldop.$mapSettings = $(html);
		$('body').append(braldop.$mapSettings);
		braldop.$depthMenu = $('#bra_depth');
		braldop.$depthMenu.change(function(){
			var z = parseInt(braldop.$depthMenu.val(), 10);
			console.log("nouveau z:", z);
			braldop.sendToBraldopServer({Cmd:'carte', ZRequis:z});
		});
		braldop.$globalMenu = $('#bra_menu');
		$triangle=$('#bra_triangle_vues');
		$triangle.click(function() {
			braldop.globalMenuOpen = !braldop.globalMenuOpen;
			if (braldop.globalMenuOpen) {
				braldop.$globalMenu.show();
				$triangle.html('&#x25b2;');
			} else {
				braldop.$globalMenu.hide();
				$triangle.html('&#x25bc;');
			}
		});
	}
}


braldop.updateMapSettings = function() {
	console.log("Profondeurs disponibles :", braldop.depths);
	braldop.ensureMapSettings();
	if (braldop.depths) {
		var html = '';
		for (var i=0; i<braldop.depths.length; i++) {
			html += '<option';
			if (map.z==braldop.depths[i]) html += ' selected';
			html += ' value='+braldop.depths[i]+'>'+braldop.depths[i]+'</option>';
		}
		braldop.$depthMenu.html(html);
		braldop.$depthMenu.show();
		$('a[maj_braldun]').live('click', function(){
			var cible = parseInt($(this).attr('maj_braldun'),10);
			braldop.sendToBraldopServer({Cmd:"carte", Action:"maj", Cible:cible});
			$(this).html('en cours...');
		});
	} else {
		braldop.$depthMenu.hide();
	}
	if (map.mapData.Vues && map.mapData.Vues.length>0) {
		var troisHeuresAvantUnix = ((new Date()).getTime()/1000)-3*60*60; // le timestamp unix en secondes correspondant à il y a 3h
		var html = '<table cellpadding=4 cellspacing=4>';
		for (var i=0; i<map.mapData.Vues.length; i++) {
			var v = map.mapData.Vues[i];
			html += '<tr>';
			html += '<td><a target=profil href="http://jeu.braldahim.com/voir/braldun/?braldun='+v.Voyeur+'&direct=evenements">'+v.PrénomVoyeur+'</a></td>';
			html += '<td>';
			html += '<a class="petit-bouton" href="javascript:if (map.zoom<32) {map.zoom=32;} map.goto('+(v.XMin+v.XMax)/2+','+(v.YMin+v.YMax)/2+','+v.Z+');" >Centrer</a>';
			html += '</td>';
			html += '<td>'+formatDate(1000*v.Time)+'</td>';
			html += '<td>';
			if (v.Time<troisHeuresAvantUnix) html += '<a class="petit-bouton" maj_braldun='+v.Voyeur+'>Mettre à jour</a>';
			html += '</td>';
			html += '</tr>';
		}
		html += '</table>';
		braldop.$globalMenu.html(html);
		$('#bra_triangle_vues').show();
	} else {
		$('#bra_triangle_vues').hide();		
	}
}

/*
 * préparation de l'interception des apparitions des blocs de l'interface pour récupérer dés
 * que possible les infos du braldun et les envoyer au serveur
 */ 
var originalShowResponse = showResponse;
showResponse = function(response) {
	console.log('showResponse called with ', arguments);
	originalShowResponse(response);
	var $alarmHolder = $('div.img_tour_activite span.braltexte');
	var ok = ($alarmHolder.length>0);
	if (ok) {
		braldop.litDonnéesBraldun();
		console.log('braldun:', braldop.braldun);
		console.log('on a les données du braldun, on peut débrancher le hook');
		showResponse = originalShowResponse; // pas forcément immédiat
		braldop.sendToBraldopServer({Etat:braldop.braldun});
	} else {
		console.log("on n'a toujours pas les données du braldun");
	}
}


