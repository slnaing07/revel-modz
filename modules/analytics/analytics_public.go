package analytics

import (
	"fmt"
	"net/http"
)

type PageRequest struct {
	Method     string
	RequestURI string
	Host       string
	XRealIp    string
	Referer    string
}

func ParsePageRequest(req *http.Request) (*PageRequest, error) {
	var PR PageRequest
	header := req.Header

	var ip string
	ips := header["X-Real-Ip"]
	for _, i := range ips {
		if i[0:3] != "127" {
			ip = i
			break
		}
	}
	ref := header["Referer"]

	PR.Method = req.Method
	PR.RequestURI = req.RequestURI
	PR.Host = req.Host
	PR.XRealIp = ip
	if len(ref) > 0 {
		PR.Referer = ref[0]
	}

	fmt.Printf("Req:\n%+v\n\n", req)
	fmt.Printf("Header:\n%+v\n\n", header)
	fmt.Printf("PageReq:\n%+v\n\n", PR)

	return &PR, nil
}
