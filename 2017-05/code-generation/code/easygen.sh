#!/bin/bash
# usage: $0  <enum-name[.yaml]> [<template-root-directory>] // HL

# ARGUMENTS
ENUM_NAME=${1:?}	# which enum to generate (<ENUM_NAME>.yaml file)
ENUM_ROOT="${2:-$HOME/go/easygen/enums/}"

cp -f ${GOFILE} ${GOFILE}.bak
(
	set -e
	sed -n "1,${GOLINE}p" ${GOFILE}.bak
	easygen -tf ${ENUM_ROOT%/}/enum_go ${ENUM_ROOT%/}/${ENUM_NAME} // HL
	echo "// {{end-easygen}}" # marker
	sed -n "$((GOLINE+1)),\$p" ${GOFILE}.bak
) > ${GOFILE}
