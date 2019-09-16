package lib

import (
	"net"
	"os/exec"
	"strings"
)

func RunShell(command string) (string, error) {

	cmd := exec.Command("/bin/bash", "-c", command)
	bytes, err := cmd.Output()
	if err != nil {
		return "", nil
	}
	resp := string(bytes)

	return resp, nil
}

func ExeSysCommand(cmdStr string) string {
	cmd := exec.Command("sh", "-c", cmdStr)
	opBytes, err := cmd.Output()
	if err != nil {
		return ""
	}
	return string(opBytes)
}

func GetPublicIpAddr() string {

	// ifconfig.me
	// icanhazip.com
	// ident.me
	// ipecho.net/plain
	// whatismyip.akamai.com
	// tnx.nl/ip
	// myip.dnsomatic.com
	// ip.appspot.com
	// curl -s checkip.dyndns.org | sed 's/.*IP Address: [0−9\.]∗.*/\1/g'
	// curl ipinfo.io/ip

	//$ curl ifconfig.me
	//$ curl icanhazip.com
	//$ curl ident.me
	//$ curl ipecho.net/plain
	//$ curl whatismyip.akamai.com
	//$ curl tnx.nl/ip
	//$ curl myip.dnsomatic.com
	//$ curl ip.appspot.com
	//$ curl -s checkip.dyndns.org | sed 's/.*IP Address: \([0-9\.]*\).*/\1/g'
	//$ curl http://ip.3322.net


	tmp := ExeSysCommand("curl icanhazip.com")
	if len(tmp) == 0 {
		return ""
	}

	ipAddr := strings.Trim(tmp, "\n") // 去除尾部换行符
	return ipAddr
}

//广播地址列表
func GetInterfaceIps() []string {

	addRs, err := net.InterfaceAddrs()

	interfaceIps := []string{}

	if err != nil {
		return interfaceIps
	}

	for _, address := range addRs {

		// 检查ip地址判断是否回环地址
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				interfaceIps = append(interfaceIps, ipNet.IP.String())
			}
		}
	}

	return interfaceIps
}

func GetLocalIp() string {
	tmp := ExeSysCommand("ifconfig eth0 | grep 'inet addr' | cut -d : -f 2 | cut -d ' ' -f 1")
	if len(tmp) == 0 {
		return ""
	}

	localIp := strings.Trim(tmp, "\n") // 去除尾部换行符
	return localIp
}