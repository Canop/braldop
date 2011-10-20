/*
 ce fichier contient des fonctions destinées à tester l'API de Braldop pour une utilisation intégrée
*/


function test_marcheAutour(actions, acteur, xc, yc, d) {
	for (var x=xc-d; x<=xc+d; x++) {
		for (var y=yc-d; y<=yc+d; y++) {
			if (x!=xc||y!=yc) actions.push({Acteur:acteur, Type:'Marcher', X:x, Y:y, PA:2});
		}
	}
}

function test_injecteActions(mapData, map) {
	var actions = [];
	var v = mapData.Vues[0];
	test_marcheAutour(actions, v.Voyeur, (v.XMin+v.XMax)/2, (v.YMin+v.YMax)/2, 2);
	map.setCallback('Marcher', function(a) {console.log('Action:'); console.log(a);});
	mapData.Actions = actions;
}

function teste(mapData, map) {
	test_injecteActions(mapData, map);
}
