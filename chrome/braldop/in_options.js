
braldop.remplitTablePartages = function(partages) {
	var html = '';
	var idb = braldop.getIdBraldun();
	if (partages.length==0) {
		html += "<i>Aucun partage ni demande de partage pour l'instant</i>";
	} else {
		for (var ip=0; ip<partages.length; ip++) {
			var p = partages[ip];
			var cible = idb==p.IdA ? p.IdB : p.IdA;
			if (p.AOk && p.BOk) {
				html += '<br>Le braldun numéro ' + cible + ' et vous partagez vos vues';
				html += '<input class="button braldop_partage" action=rompre cible='+cible+' type=button value=Rompre>';
			} else if (p.IdA==idb) {
				html += "<br>Vous avez proposé un partage au braldun " + cible + " qui n'a pour l'instant ni accepté ni refusé";
				html += '<input class="button braldop_partage" action=annuler cible='+cible+' type=button value=Annuler>';
			} else {
				html += '<br>Le braldun numéro ' + cible + ' vous a proposé un partage';
				html += '<input class="button braldop_partage" action=refuser cible='+cible+' type=button value=Refuser>';
				html += '<input class="button braldop_partage" action=accepter cible='+cible+' type=button value=Accepter>';
			}
		}
	}
	html += '<br>Proposer un partage au Braldun numéro <input id=prop_partage_id_braldun> <input class="button braldop_partage" action=proposer type=button value=Proposer>';
	$('#partages_braldop').html(html);
	$('input.braldop_partage').click(function() {
		var $this = $(this);
		var strcible = $this.attr('cible');
		if (!strcible) strcible = $('#prop_partage_id_braldun').val();
		var cible = 0;
		try {
			cible = parseInt(strcible, 10);
		} catch (e) {}
		if (cible>0) {
			var action = $this.attr('action');
			braldop.sendToBraldopServer({Cmd:"partages", Action:action, Cible:cible});
		} else {
			alert('Il faut saisir un numéro de Braldun');
		}
	});
};
