

// vérifie que le compte permet l'envoi de données au serveur Braldop
// Renvoie le mdp s'il est valide et si l'utilisateur a autorisé l'envoi
// de données.
braldop.getMdprPourServeurBraldop = function() {
	if (localStorage['braldop/carte/activation']=='non') return null;
	var mdpr = localStorage['braldop/mdpr'];
	if (!mdpr || mdpr.length!=64) return null;
	return mdpr;
}

// renvoie 0 s'il n'est pas disponible
braldop.getIdBraldun = function() {
	if (!localStorage['braldop/braldun/id']) {
		console.log('ID braldun introuvable');
		return 0;
	}
	return parseInt(localStorage['braldop/braldun/id'], 10);
}

// envoie au serveur un message authentifié par le mdp restreint
braldop.sendToBraldopServer = function(message) {
	message.Mdpr = braldop.getMdprPourServeurBraldop();
	if (!message.Mdpr) {
		console.log('Envoi au serveur braldop non authorisé');
		return;
	}
	if (!localStorage['braldop/braldun/id']) {
		console.log('ID braldun introuvable');
		return;
	}
	message.IdBraldun = parseInt(localStorage['braldop/braldun/id'], 10);
	message.Version = braldop.extVersion;
	console.log('Message sortant depuis le contexte de la page vers '+braldop.serveur+' : ', message);
	braldop.messageTimeout = setTimeout(function(){
		var html = 'Problèmes de connexion au serveur Braldop.';
		html += '<br>Ce peut être lié à des problèmes réseau, à un verrouillage de votre accès au port 8001, à un problème sur le serveur lui-même.';
		html += "<br>Si le problème persiste vous devriez désactiver l'extension Braldop afin de jouer normalement.";
		html += "<br>Avant d'en arriver à de telles extrémités, essayez de recharger la page et d'en causer sur le forum.";
		braldop.alertUser(html);
	}, 10000);
	$.ajax(
		{
			url: braldop.serveur + '?in='+JSON.stringify(message),
			crossDomain: true,
			dataType: "jsonp"
		}
	);
	return true;
}

// affiche un message pour l'utilisateur
braldop.alertUser = function(text) {
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
	$messageDiv.html(text);
}

// réception (intégrée à la page) du message de réponse du serveur braldop
receiveFromMapServer = function(message) {
	if (braldop.messageTimeout) clearTimeout(braldop.messageTimeout);
	console.log("Message entrant :", message);
	if (message.Text && message.Text.length>0) {
		braldop.alertUser(message.Text);
	}
	if (message.DV && message.DV.Vues) {
		// les vues recues du serveur remplacent les vues présentes, mais on garde les actions locales
		for (var i=0; i<message.DV.Vues.length; i++) {
			message.DV.Vues[i].active = true;
			var found = false;
			for (var iv=0; iv<map.mapData.Vues.length; iv++) {
				if (map.mapData.Vues[iv].actions) {
					message.DV.Vues[i].actions = map.mapData.Vues[iv].actions; // reprise des actions reçues de Braldahim
					break;
				}
			}
		}
		map.mapData.Vues = message.DV.Vues;
		map.compileLesVues();
	}
	if (message.PngCouche && message.PngCouche.length>5 && map && map.mapData) {
		if (!message.Z) message.Z = 0; // juste le temps de la transition, avant mise à jour du serveur
		var couche = null;
		for (var ic=0; ic<map.mapData.Couches.length; ic++) {
			var c = map.mapData.Couches[ic];
			if (c.Z==message.Z) {
				couche = c;
				break;
			}
		}
		if (couche==null) {
			couche = {};
			couche.Z = message.Z;
			couche.matrix = {};
			map.mapData.Couches.push(couche);
		}
		couche.fond = new Image();
		couche.fond.src = message.PngCouche;
		couche.fond.onload = function() {
			couche.Cases = null;
			couche.getFond = null;
			map.displayFog = true;
			map.changeProfondeur(couche.Z);
			braldop.depths = message.ZConnus;
			braldop.updateMapSettings();
			map.redraw();
		};
	}
	if (message.Partages) {
		braldop.remplitTablePartages(message.Partages);
	}
}



