package haproxy

import (
	"encoding/csv"
	"fmt"

	"github.com/gocarina/gocsv"
)

// Response from HAProxy "show stat" command.
type Stat struct {
	PxName        string `csv:"# pxname"`
	SvName        string `csv:"svname"`
	Qcur          uint64 `csv:"qcur"`
	Qmax          uint64 `csv:"qmax"`
	Scur          uint64 `csv:"scur"`
	Smax          uint64 `csv:"smax"`
	Slim          uint64 `csv:"slim"`
	Stot          uint64 `csv:"stot"`
	Bin           uint64 `csv:"bin"`
	Bout          uint64 `csv:"bout"`
	Dreq          uint64 `csv:"dreq"`
	Dresp         uint64 `csv:"dresp"`
	Ereq          uint64 `csv:"ereq"`
	Econ          uint64 `csv:"econ"`
	Eresp         uint64 `csv:"eresp"`
	Wretr         uint64 `csv:"wretr"`
	Wredis        uint64 `csv:"wredis"`
	Status        string `csv:"status"`
	Weight        uint64 `csv:"weight"`
	Act           uint64 `csv:"act"`
	Bck           uint64 `csv:"bck"`
	ChkFail       uint64 `csv:"chkfail"`
	ChkDown       uint64 `csv:"chkdown"`
	Lastchg       uint64 `csv:"lastchg"`
	Downtime      uint64 `csv:"downtime"`
	Qlimit        uint64 `csv:"qlimit"`
	Pid           uint64 `csv:"pid"`
	Iid           uint64 `csv:"iid"`
	Sid           uint64 `csv:"sid"`
	Throttle      uint64 `csv:"throttle"`
	Lbtot         uint64 `csv:"lbtot"`
	Tracked       uint64 `csv:"tracked"`
	Type          uint64 `csv:"type"`
	Rate          uint64 `csv:"rate"`
	RateLim       uint64 `csv:"rate_lim"`
	RateMax       uint64 `csv:"rate_max"`
	CheckStatus   string `csv:"check_status"`
	CheckCode     uint64 `csv:"check_code"`
	CheckDuration uint64 `csv:"check_duration"`
	Hrsp1xx       uint64 `csv:"hrsp_1xx"`
	Hrsp2xx       uint64 `csv:"hrsp_2xx"`
	Hrsp3xx       uint64 `csv:"hrsp_3xx"`
	Hrsp4xx       uint64 `csv:"hrsp_4xx"`
	Hrsp5xx       uint64 `csv:"hrsp_5xx"`
	HrspOther     uint64 `csv:"hrsp_other"`
	Hanafail      uint64 `csv:"hanafail"`
	ReqRate       uint64 `csv:"req_rate"`
	ReqRateMax    uint64 `csv:"req_rate_max"`
	ReqTot        uint64 `csv:"req_tot"`
	CliAbrt       uint64 `csv:"cli_abrt"`
	SrvAbrt       uint64 `csv:"srv_abrt"`
	CompIn        uint64 `csv:"comp_in"`
	CompOut       uint64 `csv:"comp_out"`
	CompByp       uint64 `csv:"comp_byp"`
	CompRsp       uint64 `csv:"comp_rsp"`
	LastSess      int64  `csv:"lastsess"`
	LastChk       string `csv:"last_chk"`
	LastAgt       uint64 `csv:"last_agt"`
	Qtime         uint64 `csv:"qtime"`
	Ctime         uint64 `csv:"ctime"`
	Rtime         uint64 `csv:"rtime"`
	Ttime         uint64 `csv:"ttime"`
	AgentStatus   uint64 `csv:"agent_status"`
	AgentCode     uint64 `csv:"agent_code"`
	AgentDuration uint64 `csv:"agent_duration"`
	CheckDesc     uint64 `csv:"check_desc"`
	AgentDesc     uint64 `csv:"agent_desc"`
	CheckRise     uint64 `csv:"check_rise"`
	CheckFall     uint64 `csv:"check_fall"`
	CheckHealth   uint64 `csv:"check_health"`
	AgentRise     uint64 `csv:"agent_rise"`
	AgentFall     uint64 `csv:"agent_fall"`
	AgentHealth   uint64 `csv:"agent_health"`
	Addr          string `csv:"addr"`
	Cookie        uint64 `csv:"cookie"`
	Mode          string `csv:"mode"`
	Algo          string `csv:"algo"`
	ConnRate      uint64 `csv:"conn_rate"`
	ConnRateMax   uint64 `csv:"conn_rate_max"`
	ConnTot       uint64 `csv:"conn_tot"`
	Intercepted   uint64 `csv:"intercepted"`
	Dcon          uint64 `csv:"dcon"`
	Dses          uint64 `csv:"dses"`
	Wrew          uint64 `csv:"wrew"`
	Connect       uint64 `csv:"connect"`
	Reuse         uint64 `csv:"reuse"`
	CacheLookups  uint64 `csv:"cache_lookups"`
	CacheHits     uint64 `csv:"cache_hits"`
	SrvIcur       uint64 `csv:"srv_icur"`
	SrcIlim       uint64 `csv:"src_ilim"`
	QtimeMax      uint64 `csv:"qtime_max"`
	CtimeMax      uint64 `csv:"ctime_max"`
	RtimeMax      uint64 `csv:"rtime_max"`
	TtimeMax      uint64 `csv:"ttime_max"`
	Eint          uint64 `csv:"eint"`
	IdleConnCur   uint64 `csv:"idle_conn_cur"`
	SafeConnCur   uint64 `csv:"safe_conn_cur"`
	UsedConnCur   uint64 `csv:"used_conn_cur"`
	NeedConnEst   uint64 `csv:"need_conn_est"`
}

// Equivalent to HAProxy "show stat" command.
func (h *HAProxyClient) Stats() (stats []*Stat, err error) {
	res, err := h.RunCommand("show stat")
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(res)
	reader.TrailingComma = true
	err = gocsv.UnmarshalCSV(reader, &stats)
	if err != nil {
		return nil, fmt.Errorf("error reading csv: %s", err)
	}

	//	for _, s := range allStats {
	//		switch s.SvName {
	//		case "FRONTEND":
	//			services.Frontends = append(services.Frontends, s)
	//		case "BACKEND":
	//			services.Backends = append(services.Backends, s)
	//		default:
	//			services.Listeners = append(services.Listeners, s)
	//		}
	//	}

	return stats, nil
}
