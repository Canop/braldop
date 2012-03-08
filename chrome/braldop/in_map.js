


// bâtit une copie indépendante et transmissible en json à un serveur go :
// effectue une deep-copy de l'objet source mais en ignorant tous les
//  champs dont le nom ne commence pas par une majuscule. Ne copie pas
//  les prototypes ni les champs nulls (ou 0).
braldop.goclone = function(source) {
	if ($.isArray(source)) {
		var clone = [];
		for (var i=0; i<source.length; i++) {
			if (source[i]) clone[i] = braldop.goclone(source[i]);
		}
		return clone;
	} else if (typeof(source)=="object") {
		var clone = {};
		for (var prop in source) {
			if (source[prop]) {
				var firstChar = prop.charAt(0);
				if (firstChar!=firstChar.toUpperCase()) continue;
				clone[prop] = braldop.goclone(source[prop]);
			}
		}
		return clone;
	} else {
		return source;
	}
}


braldop.handleNewMapData = function() {
	localStorage['braldop/braldun/id']=map.mapData.Vues[0].Voyeur;
	var message = {Cmd:'carte'};
	message.Vue = {
		Couches: braldop.goclone(map.mapData.Couches),
		Vues: [braldop.goclone(map.mapData.Vues[0])],
		Position: braldop.goclone(map.mapData.Position)
	};
	braldop.sendToBraldopServer(message);
}



