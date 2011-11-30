echo "Attention : le site sera écrasée à la prochaine compilation, il faut déployer toute l'application"
rsync -avz --stats --exclude="deploy.sh" * dys@canop.org:/var/www/canop/braldop

