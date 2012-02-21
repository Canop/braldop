// remplacement des fonctions de l'interface que l'on trouve normalement dans Map_env.js

color2envs = {
	9951580: "peupliers", 9885787: "peupliers",
	12381606: "gazon-gr", 12315813: "gazon-gr",
	7183454: "chenes-gr", 7117661: "chenes-gr",
	7117918: "chenes", 7052125: "chenes",
	12097633: "tunnel", 12031840: "tunnel",
	10255739: "mine", 10189946: "mine",
	5128548: "caverne-crevasse", 5062755: "caverne-crevasse",
	11958351: "erables-gr", 11892558: "erables-gr",
	13880266: "pave", 13814473: "pave",
	12116936: "marais", 12051143: "marais",
	11589980: "peupliers-gr", 11524187: "peupliers-gr",
	13026235: "route", 12960442: "route",
	13810090: "montagne", 13744297: "montagne",
	6197187: "lac", 6131394: "lac",
	4877977: "profonde", 4812184: "profonde",
	11310219: "montagne-gr", 11244426: "montagne-gr",
	12049793: "plaine", 11984000: "plaine",
	11830875: "erables", 11765082: "erables",
	10719647: "caverne", 10653854: "caverne",
	6589635: "peuprofonde", 6523842: "peuprofonde",
	7916106: "gazon", 7850313: "gazon",
	8029769: "hetres-gr", 7963976: "hetres-gr",
	11185755: "hetres", 11119962: "hetres",
	9551989: "plaine-gr", 9486196: "plaine-gr",
	8698773: "marais-gr", 8632980: "marais-gr",
	0: null
};


// assure, si cela est possible, que la couche contient les pixels du fond (imageData)
//  lesquels pourront être utilisés pour déterminer l'environnement de chaque case.
// Renvoie true ssi les pixels du fond sont disponibles.
Map.prototype.initializePixelsFond = function(couche) {
	if (couche.getFond) return true;
	if (couche.fond.width) {
		var tempCanvas = document.createElement('canvas');
		tempCanvas.width = 1600;
		tempCanvas.height = 1000;
		var context = tempCanvas.getContext('2d');
		context.drawImage(couche.fond, 0, 0);
		var pixels = context.getImageData(0, 0, 1600, 1000).data;
		couche.getFond = function(x, y) {
			var i = 4*(x+800+1600*(500-y)); // 4 pour les 4 composantes rgba
			var color = (pixels[i]<<16) + (pixels[i+1]<<8) + (pixels[i+2]);
			return color2envs[color];
		};
		couche.aPalissade = function(x, y) { // les points à palissade sont ceux dont l'alpha vaut 254/255
			var i = 4*(x+800+1600*(500-y)) + 3;
			return pixels[i]==254;
		}
		return true;
	}
	return false;
}
