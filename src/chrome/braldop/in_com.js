
// envoie au serveur un message authentifié par le mdp restreint
function sendToBraldopServer(message) {
	//~ if (!compteChrallActif()) return false
	console.log('Message sortant depuis le contexte de la page vers '+SERVEUR_BRALDOP+' : ');
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


// réception (intégrée à la page) du message de réponse du serveur braldop
function receiveFromMapServer(message) {
	console.log("Message entrant :", message);
}



