/*
ce fichier contient des fonctions destinées à tester l'API de Braldop pour une utilisation intégrée
*/


function test_marcheAutour(actions, xc, yc, d) {
	for (var x=xc-d; x<=xc+d; x++) {
		for (var y=yc-d; y<=yc+d; y++) {
			if (x!=xc||y!=yc) actions.push({type:0, x:x, y:y});
		}
	}
}

function test_injecteActions() {
	
	var actions = [];
	test_marcheAutour(actions, -110, 336, 2);
	map.setActions(22, actions, function(id, a) {console.log(a);});
}

function teste() {
	test_injecteActions();
}
