package main

import (
	"fmt"
	"syscall"

	"github.com/moby/ipvs"
)

type ipvsWrapper struct {
	handler *ipvs.Handle
}

//NewIpvsWrapper an ipvs wrapper to interact with kernel module ip_vs
func NewIpvsWrapper() (*ipvsWrapper, error) {
	if handle, err := ipvs.New(""); err != nil {
		return nil, err
	} else {
		return &ipvsWrapper{
			handler: handle,
		}, nil
	}
}

//GetServices fetch all vs service instances
func (s *ipvsWrapper) GetServices() ([]*ipvs.Service, error) {
	return s.handler.GetServices()
}

//GetDestinations fetch all destination of one vs
func (s *ipvsWrapper) GetDestinations(svc *ipvs.Service) ([]*ipvs.Destination, error) {
	return s.handler.GetDestinations(svc)
}

//Close close ip_vs handler
func (s *ipvsWrapper) Close() {
	s.handler.Close()
}

//Protocol proto string represent
func (s *ipvsWrapper) Protocol(proto uint16) string {
	switch proto {
	case syscall.IPPROTO_TCP:
		return "TCP"
	case syscall.IPPROTO_UDP:
		return "UDP"
	default:
		return fmt.Sprintf("IP(%d)", proto)
	}
}
