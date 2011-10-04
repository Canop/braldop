/*
 * ce script lance les opérations.
 * Comme tous les scripts "ext_", il est exécuté dans le contexte de l'extension et non de la page
 */


var splitedPathname = document.location.pathname.split('/');
var pageName = splitedPathname[splitedPathname.length-1];
console.log("pageName=\""+pageName+"\"");


if (pageName=='login') {
	traiteLogin();
} else {
	setTimeout(setAlarms, 5000); // on n'exécute pas tout de suite car les éléments mettent du temps à se charger
}

