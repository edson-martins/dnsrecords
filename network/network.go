package network

/* - ------------------------------------------------------------------------------------------------------------------
   - @author: Edson Martins
   -
   - Module: network.go
   -
   - Description: This module implement a structure used to store the IP and MX data and a method used to retrieve the
   -              data based in the domain input
   - -----------------------------------------------------------------------------------------------------------------*/

import (
	"net"

	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("network")

// - ------------------------------------------------------------------------------------------------------------------
// - Public struct use to store the IP addresses and the MX
// - ------------------------------------------------------------------------------------------------------------------
type Dns struct {
	Ip []net.IP  `json:"ip"`
	Mx []*net.MX `json:"host"`
}

// - ------------------------------------------------------------------------------------------------------------------
// - Public method used to retrieve the IP address and MX data
// - @parameter: [in] domain
// - @return   : [out] Dns struct, error
// - ------------------------------------------------------------------------------------------------------------------
func GetDns(domain string) (dns *Dns, err error) {

	log.Debugf("Started process to retrieve the IP and MX data based on the url %s", domain)

	dns = new(Dns)

	// - --------------------------------------------------------------------------------------------------------------
	// - Lookup IP address
	// - --------------------------------------------------------------------------------------------------------------
	dns.Ip, err = net.LookupIP(domain)
	if err != nil {
		log.Errorf("The IP lookup failed with error %s ", err.Error())
		return nil, err
	}

	// - --------------------------------------------------------------------------------------------------------------
	// - Lookup MX data
	// - --------------------------------------------------------------------------------------------------------------
	dns.Mx, err = net.LookupMX(domain)
	if err != nil {
		log.Errorf("The MX lookup failed with error %s ", err.Error())
	}

	return dns, nil
}
