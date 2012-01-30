

// vérifie que le compte permet l'envoi de données au serveur Braldop
// Renvoie le mdp s'il est valide et si l'utilisateur a autorisé l'envoi
// de données.
braldop.getMdprPourServeurBraldop = function() {
	if (localStorage['braldop/carte/activation']=='non') return null;
	var mdpr = localStorage['braldop/mdpr'];
	if (!mdpr || mdpr.length!=64) return null;
	return mdpr;
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
	$.ajax(
		{
			url: braldop.serveur + '?in='+JSON.stringify(message),
			crossDomain: true,
			dataType: "jsonp"
		}
	);
	return true;
}


// réception (intégrée à la page) du message de réponse du serveur braldop
receiveFromMapServer = function(message) {
	console.log("Message entrant :", message);
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
	if (message.DV && message.DV.Vues) {
		// ce sont les vues supplémentaires (celles des copains), on les ajoute
		for (var i=0; i<message.DV.Vues.length; i++) {
			message.DV.Vues[i].active = true;
			var found = false;
			for (var iv=0; iv<map.mapData.Vues.length; iv++) {
				if (message.DV.Vues[i].Voyeur==map.mapData.Vues[iv].Voyeur) {
					map.mapData.Vues[iv] = message.DV.Vues[i];
					found = true;
					break;
				}
			}
			if (!found) {
				map.mapData.Vues.push(message.DV.Vues[i]);
			}
		}
		map.compileLesVues();
	}
	if (!message.Z) message.Z = 0; // juste le temps de la transition, avant mise à jour du serveur
	if (message.PngCouche && message.PngCouche.length>5 && map && map.mapData) {
		var couche = null;
		console.log("A");
		for (var ic=0; ic<map.mapData.Couches.length; ic++) {
			var c = map.mapData.Couches[ic];
			if (c.Z==message.Z) {
				console.log('couche = c');
				couche = c;
				break;
			}
		}
		console.log("B couche :", c);
		if (couche==null) {
			console.log('new couche');
			couche = {};
			couche.Z = message.Z;
			couche.matrix = {};
			map.mapData.Couches.push(couche);
		}
		console.log("C");
		couche.fond = new Image();
		couche.fond.src = message.PngCouche;
		couche.fond.onload = function() {
			console.log("couche.fond.onload");
			couche.Cases = null;
			couche.getFond = null;
			map.displayFog = true;
			map.changeProfondeur(couche.Z);
			braldop.depths = message.ZConnus;
			braldop.updateMapSettings();
			map.redraw();
			console.log("Couche active : ", map.couche);
		};
	}
}



