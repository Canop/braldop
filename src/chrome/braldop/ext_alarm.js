
function durationMsToStr(delay) {
	delay /= 1000; // on passe en secondes 
	var delayHours = Math.floor(delay/3600);
	var delayMinutes = Math.floor((delay-3600*delayHours)/60);
	var delaySeconds = Math.floor(delay-60*delayMinutes-3600*delayHours);
	var str = "";
	if (delayHours>1) str = delayHours+" heures, ";
	else if (delayHours==1) str = "1 heure, ";
	if (delayMinutes>1) str += delayMinutes+" minutes, ";
	else if (delayMinutes==1) str += "1 minute, ";
	if (delaySeconds==1) str += "1 seconde";
	else str+=delaySeconds+" secondes";
	return str;
}

// règle un réveil
function setAlarm(key, date) {
	var now = (new Date()).getTime();
	var delay = date.getTime()-now-5*60*1000;
	if (delay>0) {
		//~ console.log('Alarme programmée "'+key+'" : '+date + ' (dans '+durationMsToStr(delay)+')');
		var request = {};
		request[key] = delay;
		chrome.extension.sendRequest(request);
	}
}

// positionne des alarmes 5 minutes avant la fin de chaque période du cycle, ainsi que pour le cycle suivant
function setAlarms() {
	var $alarmHolder = $('div.img_tour_activite span.braltexte');
	if ($alarmHolder.length==0) return;
	var lines = $alarmHolder.html().split('<br>');
	var turnDurationSeconds = 0; // en millisecondes 
	var alarms = {};
	for (var i=0; i<lines.length; i++) {
		var lineParts = lines[i].split(' : ');
		if (lineParts.length==2) {
			var key = lineParts[0].trim();
			if (key=="Durée du tour") {
				var hourTokens = lineParts[1].trim().split(':');
				turnDurationSeconds = parseInt(hourTokens[0])*3600+parseInt(hourTokens[1])*60+parseInt(hourTokens[2]);
			} else {
				var dateParts = lineParts[1].split(' le ');
				if (dateParts.length==2) { // date
					var hourTokens = dateParts[0].split(':');
					var dayTokens = dateParts[1].split('/');
					var date = new Date(
						2000+parseInt(dayTokens[2]),
						parseInt(dayTokens[1]-1), // les mois sont indexés à partir de zéro...
						parseInt(dayTokens[0]),
						parseInt(hourTokens[0]),
						parseInt(hourTokens[1]),
						parseInt(hourTokens[2])
					);
					alarms[key]=date;
				}
			}
		}
	}
	//console.log('Durée du tour en secondes : '+turnDurationSeconds);
	for (var key in alarms) {
		setAlarm(key, alarms[key]);
		if (key!='Début tour') {
			setAlarm(key+" tour suivant", new Date(alarms[key].getTime()+turnDurationSeconds*1000));
		}
	}
}
