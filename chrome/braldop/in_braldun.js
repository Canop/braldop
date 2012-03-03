// analyse du braldun


// cette fonction ne doit être appelée que lorsque les données ont été chargées dans l'interface
// (elles le sont de façon asynchrone)
braldop.litDonnéesBraldun = function() {
	var $boutonBraldun = $('button.butPersonnage');
	var $tableEtatBraldun = $boutonBraldun.find('span.ui-button-text table');
	if ($tableEtatBraldun.length==0) {
		console.log('table état braldun introuvable');
	} else {
		console.log('table état braldun trouvée');
		$tableEtatBraldun.find('tr').each(function(){
			$td = $(this).find('td');			
			if ($td.length>=2) {
				var str = $($td[1]).text().trim();
				var ip = str.indexOf('%'); 
				if (ip>0) {
					// il s'agit sans doute de la faim
					braldop.braldun.Faim = parseInt(str.substring(0, ip-1),10);
				} else {
					// on va supposer que c'est le niveau de vie (PV/PVmax)
					var mots = str.split('/');
					braldop.braldun.PV = parseInt(mots[0], 10);
					braldop.braldun.PVMax = parseInt(mots[1], 10);
				}
			}
		});
	}
	var $alarmHolder = $('div.img_tour_activite span.braltexte');
	if ($alarmHolder.length==0) return;
	var lines = $alarmHolder.html().split('<br>');
	var alarms = {};
	for (var i=0; i<lines.length; i++) {
		var lineParts = lines[i].split(' : ');
		if (lineParts.length==2) {
			var key = lineParts[0].trim();
			if (key=="Durée du tour") {
				var hourTokens = lineParts[1].trim().split(':');
				braldop.braldun.DuréeTour = parseInt(hourTokens[0])*3600+parseInt(hourTokens[1])*60+parseInt(hourTokens[2]);
			}
		}
	}
}
