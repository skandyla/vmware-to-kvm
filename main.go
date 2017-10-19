package main

import (
	"fmt"
	"strings"

	"github.com/namsral/flag"
)

func main() {
	var (
		config    string
		name      string
		ram       string
		vcpus     string
		osType    string
		osVariant string
		diskPath  string
		diskBus   string
		netDev1   string
		netDev2   string
		mac1      string
		mac2      string
		uuid      string
		vmdk      string
	)
	flag.StringVar(&config, "config", "", "Config file with variables - optional. Can parse both variables from config and CLI")
	flag.StringVar(&name, "name", "example-ve", "Name of the VE")
	flag.StringVar(&ram, "ram", "2048", "Memory")
	flag.StringVar(&vcpus, "vcpus", "2", "vCpu")
	flag.StringVar(&osType, "osType", "linux", "Os Type: Linux|windows")
	flag.StringVar(&osVariant, "osVariant", "rhel7", "Os Variant: rhel7,debian8 etc")
	flag.StringVar(&diskPath, "diskPath", "/srv/kvm", "Disk path to kvm images")
	flag.StringVar(&diskBus, "diskBus", "ide", "Disk bus: scsi,ide,virtio")
	flag.StringVar(&netDev1, "netDev1", "br11", "Netwrok interface 1")
	flag.StringVar(&netDev2, "netDev2", "", "Netwrok interface 2 - optional ")
	flag.StringVar(&mac1, "mac1", "", "Mac1 address - optional")
	flag.StringVar(&mac2, "mac2", "", "Mac2 address - optional")
	flag.StringVar(&uuid, "uuid", "", "Uuid address -optional")
	flag.StringVar(&vmdk, "vmdk", "", "Vmware original disk(vmdk) - optional")
	flag.Parse()

	// kvm command template
	tmpl := "virt-install --connect qemu:///system --name %s --ram %s --vcpus %s --os-type=%s  --os-variant=%s --accelerate --hvm --disk  %s/%s.qcow2,device=disk,bus=%s,format=qcow2 --graphics vnc,keymap=en-us --noautoconsole --import %s %s"

	// network template configuration
	net := "--network bridge=" + netDev1
	if mac1 != "" {
		net += ",mac='" + mac1 + "'"
	}
	if netDev2 != "" {
		net += " --network bridge=" + netDev2
		if mac2 != "" {
			net += ",mac='" + mac2 + "'"
		}
	}
	// uuid template
	// we extract it from vmx configuration
	uuidTmpl := ""
	if uuid != "" {
		words := strings.Fields(uuid)
		uuid := strings.Join(words[:4], "") + "-" + strings.Join(words[4:6], "") + "-" + strings.Join(words[6:9], "") + "-" + strings.Join(words[9:], "")
		uuidTmpl = "--uuid '" + uuid + "'"
	}

	//vmdk convertation
	if vmdk != "" {
		convert := fmt.Sprintf("time qemu-img convert -O qcow2 %s %s/%s.qcow2", vmdk, diskPath, name)
		fmt.Printf("#Convert image:\n%s\n\n", convert)
	}

	kvmCommand := fmt.Sprintf(tmpl, name, ram, vcpus, osType, osVariant, diskPath, name, diskBus, net, uuidTmpl)
	fmt.Println("#Import image:")
	fmt.Println(kvmCommand)
}
