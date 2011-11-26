

// bâtit une copie indépendante et transmissible en json à un serveur go :
// effectue une deep-copy de l'objet source mais en ignorant tous les
//  champs dont le nom ne commence pas par une majuscule. Ne copie pas
//  les prototypes ni les champs nulls (ou 0).
function goclone(source) {
	if ($.isArray(source)) {
		var clone = [];
		for (var i=0; i<source.length; i++) {
			if (source[i]) clone[i] = goclone(source[i]);
		}
		return clone;
	} else if (typeof(source)=="object") {
		var clone = {};
		for (var prop in source) {
			if (source[prop]) {
				var firstChar = prop.charAt(0);
				if (firstChar!=firstChar.toUpperCase()) continue;
				clone[prop] = goclone(source[prop]);
			}
		}
		return clone;
	} else {
		return source;
	}
}


var timer = null;
function waitForMap(callback) {
	if (map!='undefined' && map && map.mapData) {
		console.log('map déjà là');
		callback();
	} else {
		console.log('attente nécessaire');
		timer = window.setInterval(
			function(){
				if (map!='undefined' && map && map.mapData) {
					console.log('la map est là et remplie');
					window.clearInterval(timer);
					callback();
				} else {
					console.log('...');
				}
			}, 500
		);
	}
}

waitForMap(function(){
	console.log('OK : ', map.mapData);
	
	//> récupération et stockage de l'ID du braldun
	localStorage['braldop/braldun/id']=map.mapData.Vues[0].Voyeur;
	
	//> on assemble les données qu'on veut envoyer au serveur
	var data = {
		Couches: goclone(map.mapData.Couches),
		Vues: goclone(map.mapData.Vues),
		Position: goclone(map.mapData.Position)
	};
	
	sendToBraldopServer({Vue:data});
});

	
