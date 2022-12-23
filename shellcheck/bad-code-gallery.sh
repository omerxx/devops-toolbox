#!/bin/bash

beginner_mistakes() {
	var=42                          
	echo $var
	foo=42                           
	echo $foo
	for var in *; do echo;done
	local n
	var "$n"="Hello"                     
	echo "${var"$n"}"                     
	var=(1 2 3)                     
	array=( [index]=value )         
	echo "${array[@]}"
	echo "${var[14]}"                     
	echo "Argument 10 is ${10}"         
	# if $(myfunction); then ..; fi     
	# else if othercondition; then ..   
	# f; f() { echo "hello world; }     
	false                         
	# if ( -f file )   
}

qouting() {
	echo "$1"
	find . -name "./*ogg"
	rm "$HOME"
	# v='--verbose="true"'; cmd $v      
	# for f in "*.ogg"; do 
	# 	echo "X"
	# done
	#
	touch "$@"                          
	#
	echo "Don't forget to restart!"   
	echo 'Don'\''t try this at home'
	#
	echo "Path is $PATH"              
	trap 'echo Took ${SECONDS}s' 0    
	unset "var[i]"                      
}

conditionals() {
	[[ $n != 0 ]] && echo "ok"                      
	# [[ -e *.mpg ]]                    
	# [[ $foo==0 ]]                     
	# [[ -n "$foo " ]]                  
	# [[ $foo =~ "fo+" ]]               
	# [ foo =~ re ]                     
	[ "$1" = "shellcheck" ]           
	# [ $n && $m ]                      
	# [ grep -q foo file ]              
	# [[ "$$file" == *.jpg ]]           
	(( 1 < 2 ))                     
	# [ x ] & [ y ] | [ z ]             
}
