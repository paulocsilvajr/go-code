#!/bin/bash

echo 'Atualizando master com as alterações de desenvolvimento'
git rebase desenvolvimento master

echo
git branch
echo

echo 'Enviando alterações para origin'
git push origin master

echo 'Alterado branch de master para desenvolvimento'
git checkout desenvolvimento
