package analytics

import (
	// "errors"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
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

	// fmt.Printf("Req:\n%+v\n\n", req)
	// fmt.Printf("Header:\n%+v\n\n", header)
	// fmt.Printf("PageReq:\n%+v\n\n", PR)

	return &PR, nil
}

func SaveVisitorPageRequest(db *gorm.DB, vid int64, now time.Time, req *http.Request) error {
	pr, err := ParsePageRequest(req)
	if err != nil {
		return err
	}
	v := &VisitorPageRequest{
		VisitorId:  vid,
		Time:       now,
		Method:     pr.Method,
		RequestURI: pr.RequestURI,
		Host:       pr.Host,
		XRealIp:    pr.XRealIp,
		Referer:    pr.Referer,
	}

	return db.Save(v).Error
}

func SaveUserPageRequest(db *gorm.DB, uId int64, now time.Time, req *http.Request) error {
	pr, err := ParsePageRequest(req)
	if err != nil {
		return err
	}
	v := &UserPageRequest{
		UserId:     uId,
		Time:       now,
		Method:     pr.Method,
		RequestURI: pr.RequestURI,
		Host:       pr.Host,
		XRealIp:    pr.XRealIp,
		Referer:    pr.Referer,
	}

	return db.Save(v).Error
}

func GetAllVisitorPageRequests(db *gorm.DB) ([]VisitorPageRequest, error) {
	var prs []VisitorPageRequest
	err := db.Find(&prs).Error
	if err != nil {
		return nil, err
	}
	return prs, nil
}

func GetVisitorPageRequestsByVisitorId(db *gorm.DB) ([]VisitorPageRequest, error) {
	var prs []VisitorPageRequest
	err := db.Where(&VisitorPageRequest{}).Find(&prs).Error
	if err != nil {
		return nil, err
	}
	return prs, nil
}

func GetAllUserPageRequests(db *gorm.DB) ([]UserPageRequest, error) {
	var prs []UserPageRequest
	err := db.Find(&prs).Error
	if err != nil {
		return nil, err
	}
	return prs, nil
}

func GetUserPageRequestsByUserId(db *gorm.DB) ([]UserPageRequest, error) {
	var prs []UserPageRequest
	err := db.Where(&UserPageRequest{}).Find(&prs).Error
	if err != nil {
		return nil, err
	}
	return prs, nil
}
