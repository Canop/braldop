

// vérifie que le compte permet l'envoi de données au serveur Braldop
// Renvoie le mdp s'il est valide et si l'utilisateur a autorisé l'envoi
// de données.
function getMdprPourServeurBraldop() {
	if (localStorage['braldop/carte/activation']=='non') return null;
	var mdpr = localStorage['braldop/mdpr'];
	if (!mdpr || mdpr.length!=64) return null;
	return mdpr;
}


// envoie au serveur un message authentifié par le mdp restreint
function sendToBraldopServer(message) {
	message.Mdpr = getMdprPourServeurBraldop();
	if (!message.Mdpr) {
		console.log('Envoi au serveur braldop non authorisé');
		return;
	}
	
	if (!localStorage['braldop/braldun/id']) {
		console.log('ID braldun introuvable');
		return;
	}
	message.IdBraldun = parseInt(localStorage['braldop/braldun/id'], 10);
	message.Version = extVersion;
	//~ console.log('Message sortant depuis le contexte de la page vers '+SERVEUR_BRALDOP+' : ', message);
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
	//~ console.log("Message entrant :", message);
	if (message.Text && message.Text.length>0) {
		var $messageDiv = $('#braldop_message_content');
		if ($messageDiv.length==0) {
			var html = '<div id=braldop_message><div id=braldop_message_content>';
			html += '<hr><a id=>OK</a>'
			html += '</div><div id=braldop_message_controls><table width=100%><tr><td><span id=braldop_message_opener>!</span></td><td align=right id=braldop_message_deleter>Supprimer ce message</td></tr></table></div></div>';
			$(html).appendTo($('body'));
			$messageDiv =  $('#braldop_message_content');
			$('#braldop_message_opener').click(function(){$('#braldop_message_content, #braldop_message_deleter').toggle()});
			$('#braldop_message_deleter').click(function(){$('#braldop_message').remove()});
		}
		$messageDiv.html(message.Text);
	}
	if (message.PngCouche && message.PngCouche.length>5 && map && map.mapData) {
		map.mapData.Couches[0].fond.src = message.PngCouche;
		map.mapData.Couches[0].fond.onload = function() {
			map.mapData.Couches[0].Cases = null;
			map.mapData.Couches[0].getFond = null;
			//~ console.log("fond chargé");
			map.displayFog = true;
			map.redraw();
		}
	}
}



