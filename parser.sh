#!/bin/bash
# parse vmware exsi vmx conf files to go-flag config structure
if [ -z $1 ]; then
	echo "Usage: $0 filename.vmx"
	exit 1
fi
inputFile=$1
conf=vmx.conf

if [ -f $conf ]; then
	rm $conf
fi

extract(){
	destVarName=$1
	srcVarName=$2

    var=$(grep "$srcVarName" $inputFile | cut -d= -f2 | sed -e 's/"//g' -e 's/^ //')
	echo "$destVarName $var" >> $conf
}


#function kvmname  vmwarename
extract vcpus numvcpus
extract ram memSize
extract uuid uuid.bios
extract mac1 'ethernet0.generatedAddress ='
extract mac2 'ethernet1.generatedAddress ='
extract vmdk scsi0:0.fileName
extract name displayName

