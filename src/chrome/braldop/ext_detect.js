
var paragraph = $("#browserInfos");
if (paragraph.length>0) {
	var html = "Braldop est déjà installé dans ce navigateur et fonctionne.";
	if (braldop.extVersion != $("#version").text()) {
		html += "<br>Votre version est plus ancienne (" + extVersion + "). Vous devriez cliquer sur le bouton ci-dessus afin de mettre à jour l'extension.";
		setTimeout(function() {
			document.location.reload();
		}, 5000);
	} else {
		html += "<br>De plus, vous disposez de la dernière version. Il n'est donc a priori pas nécessaire de cliquer sur le bouton ci-dessus.";
	}
	paragraph.html(html);
}
