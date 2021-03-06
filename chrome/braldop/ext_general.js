

braldop.getUrlParameter = function(name, defaultValue) {
  name = name.replace(/[\[]/,"\\\[").replace(/[\]]/,"\\\]");
  var regexS = "[\\?&]"+name+"=([^&#]*)";
  var regex = new RegExp( regexS );
  var results = regex.exec( document.location.href );
  if( results == null )
    return defaultValue;
  else
    return results[1];
}

/**
 * injecte un fichier javascript présent dans l'extension, de telle sorte
 *  qu'il soit exécuté dans le contexte de la page et non celui de l'extension
 */
braldop.inject = function(fileName) {
	$.getScript(chrome.extension.getURL(fileName));
}

/**
 * découpe en mots (un nombre peut être un mot).
 * 
 * Note : comme je ne suis pas fort en expressions régulières, si un "-" est isolé, il sort comme un mot...
 * Attention : si vous corrigez le comportement de la ligne ci-dessus il faudra modifier Chrall_extractBasicInfos et pas mal d'autres méthodes
 */
braldop.tokenize = function(text) {
	return text.trim().split(new RegExp("[ /|\t\n\r\f,.:=()]+", "g"));	
}




