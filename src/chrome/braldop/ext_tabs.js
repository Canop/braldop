

braldop.ajouteOnglets = function($target, tabs) {
	var html = '<table class=tab_holder cellspacing=0><tr>';
	var i=0;
	for (var key in tabs) {
		html += '<td class=inter>&nbsp;</td><th class='+(i==0?'active':'inactive')+' num='+i+'>'+key+'</th>';
		i++;
	}
	html += '<td class=inter>&nbsp;</td></tr>';
	html += '<tr><td class=tab_page_holder colspan='+(2*i+1)+'>';
	var i=0;
	for (var key in tabs) {
		html += '<div class=tab_page id=tab_page_'+i+'></div>'; // pour l'instant on ne gère qu'un jeu d'onglet par page
		i++;
	}
	html += '</td></tr></table>';
	var $tabs=$(html);
	$tabs.appendTo($target);
	var i=0;
	for (var key in tabs) {
		var page=tabs[key];
		$('#tab_page_'+i).append(page);
		i++;
	}
	$('#tab_page_0').show();
	$tabs.delegate('th.inactive', 'click', function(){
		$tabs.find('th.active').removeClass('active').addClass('inactive');
		$(this).removeClass('inactive').addClass('active');
		$('div.tab_page').hide();
		$('div.tab_page#tab_page_'+$(this).attr('num')).show();
	});
}
