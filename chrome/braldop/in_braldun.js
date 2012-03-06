// analyse du braldun


// cette fonction ne doit être appelée que lorsque les données ont été chargées dans l'interface
// (elles le sont de façon asynchrone).
braldop.litDonnéesBraldun = function() {
	//> récupération pv, pvmax et faim
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
	//> récupération durée du tour
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
	//> récupération PA et DLA
	var $cockpit = $('#tourCockpit');
	var str = $cockpit.text();
	var i = str.indexOf('Point');
	if (i>0) {
		braldop.braldun.PA = parseInt(str.substring(0, i-1).trim(), 10);
	}
	i = str.indexOf('DLA');
	if (i>0) {
		str = str.substring(i+4).trim();
		var heureJour = str.split('le');
		var heureMinuteSeconde = heureJour[0].trim().split(':');
		var jourMoisAnnée = heureJour[1].split('/');
		console.log('DLA :', heureMinuteSeconde, jourMoisAnnée);
		var nhms = heureMinuteSeconde.length;
		if (nhms>=3 && jourMoisAnnée.length>=3) {
			var dateDLA = new Date(
				parseInt(jourMoisAnnée[2], 10)+2000,
				parseInt(jourMoisAnnée[1], 10)-1,
				parseInt(jourMoisAnnée[0], 10),
				parseInt(heureMinuteSeconde[nhms-3], 10),
				parseInt(heureMinuteSeconde[nhms-2], 10),
				parseInt(heureMinuteSeconde[nhms-1], 10)
			);
			console.log('dateDLA:', dateDLA);
			braldop.braldun.DLA = dateDLA.getTime()/1000
		}
	}
}
