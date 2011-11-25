/*
 * ce script lance les opérations.
 * Comme tous les scripts "ext_", il est exécuté dans le contexte de l'extension et non de la page.
 * Les scripts dont le nom commence par "in_" sont injectés pour être exécuté dans le contexte de la page.
 */

inject('inext_const.js');
inject('in_com.js');

switch (document.location.pathname) {
case '/auth/login':
case '/auth/login/':
	traiteLogin();
	break;
case '/interface':
case '/interface/':
	inject('in_map.js');
	setTimeout(setAlarms, 5000); // on n'exécute pas tout de suite car les éléments mettent du temps à se charger
	break;
case '/Parametres':
case '/Parametres/':
	changePageParamètres();
	break;
default:
	console.log('page ignorée');
}

