
braldop.rafraichitParagraphePartages = function() {
	var html = '';
	if (localStorage['braldop/carte/activation']=='oui') {
		html += 'en attente...';
		braldop.sendToBraldopServer({Cmd:'partages'}); // la méthode receiveFromMapServer from map server (appelée dans le contexte de la page) s'occupera de l'affichage
	} else {
		html += "Les partages de vue nécessitent l'activation de la carte.";		
	}
	$('#partages_braldop').html(html);
};


braldop.changePageParamètres = function() {
	//> récupération et stockage du mot de passe restreint
	var mdpr = $('pre').text().trim();
	if (mdpr.length==64) {
		localStorage['braldop/mdpr'] = mdpr;
	} else {
		console.log("mot de passe restreint non reconnu");
	}
	
	//> décomposition de l'écran en 2 onglets : l'existant dans "Paramètres Braldahim" et les miens dans "Paramètres Braldop"
	//  + ajout de commentaires et d'une case à cocher pour l'activation de la carte
	var $container = $('div.box_parametres');
	var $innerDiv = $container.find('div.inner');
	$innerDiv.detach();
	html = '<br>';
	html += '<span class="titre textalic textsizeb">Activation</span>'
	html += '<p>';
	html += "<br>Activer le système cartographique permet de disposer dans votre vue des terrains vus précédemment.";
	html += "<br><br><input type=checkbox id=activation_envoi_braldop "+(localStorage['braldop/carte/activation']=='oui'?'checked':'')+"><label for=activation_envoi_braldop>Activer le système cartographique</for>";
	html += "<br><br>Si vous activez le système cartographique, vos données de vue et votre mot de passe restreint seront envoyés au serveur Braldop.";
	html += '</p><br><br>';
	html += '<span class="titre textalic textsizeb">Partages de vue</span>'
	html += '<p id=partages_braldop>';
	html += '</p><br><br>';

	braldop.ajouteOnglets($container, {
		'Paramètres Braldahim': $innerDiv.find('div.contenu'),
		'Paramètres Braldop':html
	});
	$('#activation_envoi_braldop').change(function(){
		localStorage['braldop/carte/activation'] = this.checked ? 'oui' : 'non';
		braldop.rafraichitParagraphePartages();
	});
	
	braldop.rafraichitParagraphePartages();
};
