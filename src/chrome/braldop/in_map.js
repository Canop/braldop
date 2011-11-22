


// effectue une deep-copy de l'objet source mais en ignorant tous les
//  champs dont le nom ne commence pas par une majuscule. Ne copie pas
//  les prototypes.
function deepUpperClone(source) {
	if ($.isArray(source)) {
		var clone = [];
		for (var i=0; i<source.length; i++) clone[i] = deepUpperClone(source[i]);
		return clone;
	} else if (typeof(source)=="object") {
		var clone = {};
		for (var prop in source) {
			var firstChar = prop.charAt(0);
			if (firstChar!=firstChar.toUpperCase()) continue;
			clone[prop] = deepUpperClone(source[prop]);
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
	console.log('OK : ', map);
	//~ console.log('map.mapData : ', map.mapData, 'clone : ', deepUpperClone(map.mapData));
	var data = {
		Couches: deepUpperClone(map.mapData.Couches),
		Vues: deepUpperClone(map.mapData.Vues),
		Position: deepUpperClone(map.mapData.Position)
	};
	
	sendToBraldopServer({Vue:data});
});

	
