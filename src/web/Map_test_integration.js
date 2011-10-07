/*
ce fichier contient des fonctions destinées à tester l'API de Braldop pour une utilisation intégrée
*/


function test_marcheAutour(actions, xc, yc, d) {
	for (var x=xc-d; x<=xc+d; x++) {
		for (var y=yc-d; y<=yc+d; y++) {
			if (x!=xc||y!=yc) actions.push({Type:0, X:x, Y:y, PA:2});
		}
	}
}

function test_injecteActions() {
	
	var actions = [];
	test_marcheAutour(actions, -110, 336, 2);
	map.setActions(22, actions, function(id, a) {console.log('Action:');console.log(a);});
}

function teste() {
	test_injecteActions();
}