#!/bin/bash

branch_atual() {
    echo $(git branch | grep '*' | awk -F ' +' '! /\(no branch\)/ {print $2}')
}

BRANCH_ATUAL=$(branch_atual)
DEV='desenvolvimento'
MASTER='master'

if [ $BRANCH_ATUAL == $DEV ]; then
    clear
    echo 'Atualizando master com as alterações de desenvolvimento'
    git rebase desenvolvimento master

    BRANCH_ATUAL=$(branch_atual)

    if [ $BRANCH_ATUAL == $MASTER ]; then
        echo
        git branch
        echo

        echo 'Enviando alterações para origin'
        git push origin master

        echo
        echo 'Alterado branch de master para desenvolvimento'
        git checkout desenvolvimento

    else
        echo -e '\nAdicione(add) e grave(commit) as alterações de' $DEV '\n'
    fi

else
    echo -e '\nAltere para o branch' $DEV '\n'
fi

