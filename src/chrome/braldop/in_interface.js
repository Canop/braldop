

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
			braldop.sendToBraldopServer({ZRequis:z});
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
	} else {
		braldop.$depthMenu.hide();
	}
	if (map.mapData.Vues && map.mapData.Vues.length>0) {
		var html = '<table cellpadding=4 cellspacing=4>';
		for (var i=0; i<map.mapData.Vues.length; i++) {
			var v = map.mapData.Vues[i];
			html += '<tr><td><a href="javascript:if (map.zoom<32) {map.zoom=32;} map.goto('+(v.XMin+v.XMax)/2+','+(v.YMin+v.YMax)/2+','+v.Z+');">'+v.PrénomVoyeur+'</a></td><td>'+formatDate(1000*v.Time)+'</td></tr>';
		}
		html += '</table>';
		braldop.$globalMenu.html(html);
		$('#bra_triangle_vues').show();
	} else {
		$('#bra_triangle_vues').hide();		
	}
}
