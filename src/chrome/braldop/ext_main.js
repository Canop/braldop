/*
 * ce script lance les opérations.
 * Comme tous les scripts "ext_", il est exécuté dans le contexte de l'extension et non de la page.
 * Les scripts dont le nom commence par "in_" ou inext_ sont injectés pour être exécuté dans le contexte de la page.
 */


braldop.inject('inext_const.js');
braldop.inject('inext_com.js');

switch (document.location.pathname) {
case '/auth/login':
case '/auth/login/':
	braldop.traiteLogin();
	break;
case '/interface':
case '/interface/':
	braldop.initInterface();
	if (braldop.getMdprPourServeurBraldop()) { // on ne traite la carte que si l'utilisateur l'a autorisé
		braldop.inject('in_upgrade_goutgueule.js');
		braldop.inject('in_interface.js');
		braldop.inject('in_map.js');
	} else {
		console.log('pas de mdpr ou pas de authorisation');
	}
	setTimeout(braldop.setAlarms, 5000); // on n'exécute pas tout de suite car les éléments mettent du temps à se charger
	break;
case '/Parametres':
case '/Parametres/':
	braldop.inject('in_options.js');
	braldop.changePageParamètres();
	break;
default:
	//~ console.log('page ignorée');
}

