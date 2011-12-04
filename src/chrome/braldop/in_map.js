


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
		//~ console.log('map déjà là');
		callback();
	} else {
		//~ console.log('attente nécessaire pour la carte');
		timer = window.setInterval(
			function(){
				if (map!='undefined' && map && map.mapData) {
					window.clearInterval(timer);
					callback();
				} else {
					//console.log('...');
				}
			}, 500
		);
	}
}

function handleNewMapData() {
	var data = {
		Couches: goclone(map.mapData.Couches),
		Vues: goclone(map.mapData.Vues),
		Position: goclone(map.mapData.Position)
	};
	sendToBraldopServer({Vue:data});
}

waitForMap(function(){
	//~ console.log('Vue fournie par Braldahim : ', map.mapData);
	
	//> récupération et stockage de l'ID du braldun
	localStorage['braldop/braldun/id']=map.mapData.Vues[0].Voyeur;
	
	//> traitement des données
	handleNewMapData();
	
	//> on met un hook pour les prochaines modifs
	map.onSetData = handleNewMapData;
});



