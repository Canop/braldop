--
-- Structure de la table `compte`
--

CREATE TABLE IF NOT EXISTS `compte` (
  `id` int(11) NOT NULL,
  `mdpr` varchar(64) COLLATE utf8_unicode_ci DEFAULT NULL,
  `mdpr_ok` tinyint(4) NOT NULL COMMENT '1=OK, 0=NOK',
  `x` int(11) NOT NULL,
  `y` int(11) NOT NULL,
  `z` int(11) NOT NULL,
  `maj` int(11) NOT NULL COMMENT 'secondes depuis 1970',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
