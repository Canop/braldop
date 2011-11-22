
// envoie au serveur un message authentifié par le mdp restreint
function sendToBraldopServer(message) {
	//~ if (!compteChrallActif()) return false
	console.log('Message sortant de '+pageName+' vers '+SERVEUR_BRALDOP+' : ');
	console.log(message);
	$.ajax(
		{
			url: SERVEUR_BRALDOP + '?in='+JSON.stringify(message),
			crossDomain: true,
			dataType: "jsonp"
		}
	);
	return true;
}
