//Le script suivant me permet de mesurer l'utilisation de Chrall via Google Analytics.
function setgoogleAnalytics() {
var _gaq = _gaq || [];
_gaq.push(['_setAccount', 'UA-15064357-4']);
_gaq.push(['_trackPageview']);

(function() {
var ga = document.createElement('script'); ga.type = 'text/javascript'; ga.async = true;
    ga.src = ('https:' == document.location.protocol ? 'https://ssl' : 'http://www') + '.google-analytics.com/ga.js';
    var s = document.getElementsByTagName('script')[0]; s.parentNode.insertBefore(ga, s);
  })();
}

function setAlarms() {
var alarmTimeouts = {};

// on répond aux requêtes de mise en place d'alarmes
chrome.extension.onRequest.addListener(
	function(request, sender, sendResponse) {
		for (var key in request) {
			if (alarmTimeouts[key]) {
				clearTimeout(alarmTimeouts[key]);
			}
			var delay = request[key];
			alarmTimeouts[key] = setTimeout(function() {	
				document.getElementById('alarm').play();
			}, delay);
			
		}
	}
);
}


// Add event listeners once the DOM has fully loaded by listening for the
// `DOMContentLoaded` event on the document, and adding your listeners to
// specific elements when it triggers.
document.addEventListener('DOMContentLoaded', function () {
	setgoogleAnalytics();
	setAlarms();
});
