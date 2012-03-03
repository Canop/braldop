// Ce fichier est exécuté à la fois dans l'extension (ext) et dans le contexte de la page (in)
// Il déclare braldop et doit donc être exécuté en premier

var braldop = {};

// version de l'extension
braldop.extVersion = "3.2.1";

braldop.serveur = 'http://canop.org:8001/';
//~ braldop.serveur = 'http://localhost:8001/';


braldop.depths = null;
braldop.braldun = {};
