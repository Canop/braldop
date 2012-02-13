#!/bin/bash

# Ce script se contente de redéployer le site web


CHEMIN=`dirname $0`
cd $CHEMIN

source config.sh

rsync -az --delete $CHEMIN_BRALDOP/src/site/* $CHEMIN_DEPLOIEMENT_WEB


