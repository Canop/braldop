
function changePageParamètres() {
	var $container = $('div.box_parametres');
	var $innerDiv = $container.find('div.inner');
	$innerDiv.detach();
	html = "<br><br>paras braldop";
	html += "<br><br><input type=checkbox id=activation_envoi_braldop><label for=activation_envoi_braldop>Activer le système cartographique</for>";
	html += "<br><br>Si vous activez le système cartographique, vos données de vue et votre mot de passe restreint seront envoyés au serveur Braldop.";
	ajouteOnglets($container, {
		'Paramètres Braldahim': $innerDiv.find('div.contenu'),
		'Paramètres Braldop':html
	});
	
}
