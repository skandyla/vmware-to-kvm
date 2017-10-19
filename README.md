# vmware-to-kvm
Helper scripts to migrate vmware VM to kvm.  

## Info
This script just generates shell command for virt-install and print it to the screen (you should just execute it).  

Also there is a dedicated tool [virt-v2v](https://access.redhat.com/articles/1353223), unfortunatelly it's not working as expected.   
So this scripts helps to migrate VMs in [classic way](https://www.linux-kvm.org/page/How_To_Migrate_From_Vmware_To_KVM)  

## Prerequirements
GO installed

## Usage:

1. shutdown and copy VM to kvm server (use scp)
2. launch  
`./parser.sh path_to_VM.vmx` - it's optional, but allows to parse some info like Uuid, mac, etc. This script generates vmx.conf
3. launch   
`go run main.go` - just to generate defaults  
`go run main.go -config vmx.conf -netDev1 br11 -vcpus 4 -ram 4096` - example to parse vmx.conf and ajust some variables  


## Tested:
migration from vmware ESXi 5.0, 5.5 to kvm srv running CentOS 7  
