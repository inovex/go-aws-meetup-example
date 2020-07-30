#!/bin/sh

print_help () {
    printf "This is a terraform wrapper script.\n"
    printf "Usage: ./terraformw [STAGE] [COMMAND]\n"
    printf "\n"
    printf "STAGE:  \t One of 'dev', 'int', 'prd'\n"
    printf "COMMAND:\t A terraform command. Feel free to use parameters as well.\n"
}

check_input_args() {
    if [ -z $stage ] || [ -z $command ];
    then
      print_help
      exit 1
    fi

    STAGE_REGEX='(dev|int|prd)$'

    IS_VALID_STAGE=$(expr $stage : $STAGE_REGEX)

    if [ -z ${IS_VALID_STAGE} ];
    then
        print_help
        echo ""
        echo "Stage was \"${stage}\" but it may only be one of dev,int,prd"
        exit 1
    fi
}

add_args_to_init() {
    if [ "${command}" = "init" ];
    then
        additional_args="-backend-config=\"bucket=go-meetup-tf-state-${stage}\" -backend-config=\"dynamodb_table=tf-state-lock-${stage}\""
    else
        additional_args="-var=\"stage=${stage}\""
    fi
}

stage=$1
command=$2
p1=$3
p2=$4
p3=$5

check_input_args

add_args_to_init

eval "terraform ${command} ${additional_args} ${p1} ${p2} ${p3}"
