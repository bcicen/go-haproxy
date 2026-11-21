package haproxy

import (
	"encoding/csv"
	"fmt"

	"github.com/gocarina/gocsv"
)

// Response from HAProxy "show stat" command.
type Stat struct {
	// stat_idx_px in include/haproxy/stats-t.h
	PxName               string `csv:"# pxname"`
	SvName               string `csv:"svname"`
	Qcur                 uint64 `csv:"qcur"`
	Qmax                 uint64 `csv:"qmax"`
	Scur                 uint64 `csv:"scur"`
	Smax                 uint64 `csv:"smax"`
	Slim                 uint64 `csv:"slim"`
	Stot                 uint64 `csv:"stot"`
	Bin                  uint64 `csv:"bin"`
	Bout                 uint64 `csv:"bout"`
	Dreq                 uint64 `csv:"dreq"`
	Dresp                uint64 `csv:"dresp"`
	Ereq                 uint64 `csv:"ereq"`
	Econ                 uint64 `csv:"econ"`
	Eresp                uint64 `csv:"eresp"`
	Wretr                uint64 `csv:"wretr"`
	Wredis               uint64 `csv:"wredis"`
	Status               string `csv:"status"`
	Weight               uint64 `csv:"weight"`
	Act                  uint64 `csv:"act"`
	Bck                  uint64 `csv:"bck"`
	ChkFail              uint64 `csv:"chkfail"`
	ChkDown              uint64 `csv:"chkdown"`
	Lastchg              uint64 `csv:"lastchg"`
	Downtime             uint64 `csv:"downtime"`
	Qlimit               uint64 `csv:"qlimit"`
	Pid                  uint64 `csv:"pid"`
	Iid                  uint64 `csv:"iid"`
	Sid                  uint64 `csv:"sid"`
	Throttle             uint64 `csv:"throttle"`
	Lbtot                uint64 `csv:"lbtot"`
	Tracked              uint64 `csv:"tracked"`
	Type                 uint64 `csv:"type"`
	Rate                 uint64 `csv:"rate"`
	RateLim              uint64 `csv:"rate_lim"`
	RateMax              uint64 `csv:"rate_max"`
	CheckStatus          string `csv:"check_status"`
	CheckCode            uint64 `csv:"check_code"`
	CheckDuration        uint64 `csv:"check_duration"`
	Hrsp1xx              uint64 `csv:"hrsp_1xx"`
	Hrsp2xx              uint64 `csv:"hrsp_2xx"`
	Hrsp3xx              uint64 `csv:"hrsp_3xx"`
	Hrsp4xx              uint64 `csv:"hrsp_4xx"`
	Hrsp5xx              uint64 `csv:"hrsp_5xx"`
	HrspOther            uint64 `csv:"hrsp_other"`
	Hanafail             uint64 `csv:"hanafail"`
	ReqRate              uint64 `csv:"req_rate"`
	ReqRateMax           uint64 `csv:"req_rate_max"`
	ReqTot               uint64 `csv:"req_tot"`
	CliAbrt              uint64 `csv:"cli_abrt"`
	SrvAbrt              uint64 `csv:"srv_abrt"`
	CompIn               uint64 `csv:"comp_in"`
	CompOut              uint64 `csv:"comp_out"`
	CompByp              uint64 `csv:"comp_byp"`
	CompRsp              uint64 `csv:"comp_rsp"`
	LastSess             int64  `csv:"lastsess"`
	LastChk              string `csv:"last_chk"`
	LastAgt              uint64 `csv:"last_agt"`
	Qtime                uint64 `csv:"qtime"`
	Ctime                uint64 `csv:"ctime"`
	Rtime                uint64 `csv:"rtime"`
	Ttime                uint64 `csv:"ttime"`
	AgentStatus          uint64 `csv:"agent_status"`
	AgentCode            uint64 `csv:"agent_code"`
	AgentDuration        uint64 `csv:"agent_duration"`
	CheckDesc            string `csv:"check_desc"`
	AgentDesc            string `csv:"agent_desc"`
	CheckRise            uint64 `csv:"check_rise"`
	CheckFall            uint64 `csv:"check_fall"`
	CheckHealth          uint64 `csv:"check_health"`
	AgentRise            uint64 `csv:"agent_rise"`
	AgentFall            uint64 `csv:"agent_fall"`
	AgentHealth          uint64 `csv:"agent_health"`
	Addr                 string `csv:"addr"`
	Cookie               string `csv:"cookie"`
	Mode                 string `csv:"mode"`
	Algo                 string `csv:"algo"`
	ConnRate             uint64 `csv:"conn_rate"`
	ConnRateMax          uint64 `csv:"conn_rate_max"`
	ConnTot              uint64 `csv:"conn_tot"`
	Intercepted          uint64 `csv:"intercepted"`
	Dcon                 uint64 `csv:"dcon"`
	Dses                 uint64 `csv:"dses"`
	Wrew                 uint64 `csv:"wrew"`
	Connect              uint64 `csv:"connect"`
	Reuse                uint64 `csv:"reuse"`
	CacheLookups         uint64 `csv:"cache_lookups"`
	CacheHits            uint64 `csv:"cache_hits"`
	SrvIcur              uint64 `csv:"srv_icur"`
	SrcIlim              uint64 `csv:"src_ilim"`
	QtimeMax             uint64 `csv:"qtime_max"`
	CtimeMax             uint64 `csv:"ctime_max"`
	RtimeMax             uint64 `csv:"rtime_max"`
	TtimeMax             uint64 `csv:"ttime_max"`
	Eint                 uint64 `csv:"eint"`
	IdleConnCur          uint64 `csv:"idle_conn_cur"`
	SafeConnCur          uint64 `csv:"safe_conn_cur"`
	UsedConnCur          uint64 `csv:"used_conn_cur"`
	NeedConnEst          uint64 `csv:"need_conn_est"`
	ConnEst              uint64 `csv:"conn_est"`
	Uweight              uint64 `csv:"uweight"`
	AggServerStatus      uint64 `csv:"agg_server_status"`
	AggServerCheckStatus uint64 `csv:"agg_server_check_status"`
	AggCheckStatus       uint64 `csv:"agg_check_status"`
	Srid                 uint64 `csv:"srid"`
	SessOther            uint64 `csv:"sess_other"`
	H1sess               uint64 `csv:"h1sess"`
	H2sess               uint64 `csv:"h2sess"`
	H3sess               uint64 `csv:"h3sess"`
	ReqOther             uint64 `csv:"req_other"`
	H1req                uint64 `csv:"h1req"`
	H2req                uint64 `csv:"h2req"`
	H3req                uint64 `csv:"h3req"`
	Proto                string `csv:"proto"`

	// Special delimiter on proxy stats to mark end of static fields
	// see stats_dump_csv_header in src/stats.c
	// -

	// Below all stats are defined by stats_register_module function in
	// HAProxy.

	// h2_counters in src/mux_h2.c
	H2HeadersRcvd                uint64 `csv:"h2_headers_rcvd"`
	H2DataRcvd                   uint64 `csv:"h2_data_rcvd"`
	H2SettingsRcvd               uint64 `csv:"h2_settings_rcvd"`
	H2RstStreamRcvd              uint64 `csv:"h2_rst_stream_rcvd"`
	H2GoawayRcvd                 uint64 `csv:"h2_goaway_rcvd"`
	H2DetectedConnProtocolErrors uint64 `csv:"h2_detected_conn_protocol_errors"`
	H2DetectedStrmProtocolErrors uint64 `csv:"h2_detected_strm_protocol_errors"`
	H2RstStreamResp              uint64 `csv:"h2_rst_stream_resp"`
	H2GoawayResp                 uint64 `csv:"h2_goaway_resp"`
	H2OpenConnections            uint64 `csv:"h2_open_connections"`
	H2BackendOpenStreams         uint64 `csv:"h2_backend_open_streams"`
	H2TotalConnections           uint64 `csv:"h2_total_connections"`
	H2BackendTotalStreams        uint64 `csv:"h2_backend_total_streams"`

	// h1_counters in src/mux_h1.c
	H1OpenConnections  uint64 `csv:"h1_open_connections"`
	H1OpenStreams      uint64 `csv:"h1_open_streams"`
	H1TotalConnections uint64 `csv:"h1_total_connections"`
	H1TotalStreams     uint64 `csv:"h1_total_streams"`
	H1BytesIn          uint64 `csv:"h1_bytes_in"`
	H1BytesOut         uint64 `csv:"h1_bytes_out"`
	H1SplicedBytesIn   uint64 `csv:"h1_spliced_bytes_in"`
	H1SplicedBytesOut  uint64 `csv:"h1_spliced_bytes_out"`

	// ssl_stats in src/ssl_sock.c
	SslSess             uint64 `csv:"ssl_sess"`
	SslReusedSess       uint64 `csv:"ssl_reused_sess"`
	SslFailedHandshake  uint64 `csv:"ssl_failed_handshake"`
	SslOcspStaple       uint64 `csv:"ssl_ocsp_staple"`
	SslFailedOcspStaple uint64 `csv:"ssl_failed_ocsp_staple"`

	// quic_stats in src/quic_stats.c include/haproxy/quic_stats-t.h
	QuicRxbufFull                        uint64 `csv:"quic_rxbuf_full"`
	QuicDroppedPkt                       uint64 `csv:"quic_dropped_pkt"`
	QuicDroppedPktBufoverrun             uint64 `csv:"quic_dropped_pkt_bufoverrun"`
	QuicDroppedParsingPkt                uint64 `csv:"quic_dropped_parsing_pkt"`
	QuicSocketFull                       uint64 `csv:"quic_socket_full"`
	QuicSendtoErr                        uint64 `csv:"quic_sendto_err"`
	QuicSendtoErrUnknwn                  uint64 `csv:"quic_sendto_err_unknwn"`
	QuicSentPkt                          uint64 `csv:"quic_sent_pkt"`
	QuicLostPkt                          uint64 `csv:"quic_lost_pkt"`
	QuicTooShortDgram                    uint64 `csv:"quic_too_short_dgram"`
	QuicRetrySent                        uint64 `csv:"quic_retry_sent"`
	QuicRetryValidated                   uint64 `csv:"quic_retry_validated"`
	QuicRetry_error                      uint64 `csv:"quic_retry_error"`
	QuicHalf_open_conn                   uint64 `csv:"quic_half_open_conn"`
	QuicHdshk_fail                       uint64 `csv:"quic_hdshk_fail"`
	QuicStless_rst_sent                  uint64 `csv:"quic_stless_rst_sent"`
	QuicConn_migration_done              uint64 `csv:"quic_conn_migration_done"`
	QuicTranspErrNo_error                uint64 `csv:"quic_transp_err_no_error"`
	QuicTranspErrInternalError           uint64 `csv:"quic_transp_err_internal_error"`
	QuicTranspErrConnectionRefused       uint64 `csv:"quic_transp_err_connection_refused"`
	QuicTranspErrFlowControlError        uint64 `csv:"quic_transp_err_flow_control_error"`
	QuicTranspErrStreamlimitError        uint64 `csv:"quic_transp_err_stream_limit_error"`
	QuicTranspErrStreamStateError        uint64 `csv:"quic_transp_err_stream_state_error"`
	QuicTranspErrFinalSizeError          uint64 `csv:"quic_transp_err_final_size_error"`
	QuicTranspErrFrame_encoding_error    uint64 `csv:"quic_transp_err_frame_encoding_error"`
	QuicTranspErrTransportParameterError uint64 `csv:"quic_transp_err_transport_parameter_error"`
	QuicTranspErrConnectionIdLimit       uint64 `csv:"quic_transp_err_connection_id_limit"`
	QuicTranspErrProtocoViolationError   uint64 `csv:"quic_transp_err_protocol_violation_error"`
	QuicTranspErrInvalidToken            uint64 `csv:"quic_transp_err_invalid_token"`
	QuicTranspErrApplication_error       uint64 `csv:"quic_transp_err_application_error"`
	QuicTranspErrCryptoBufferExceeded    uint64 `csv:"quic_transp_err_crypto_buffer_exceeded"`
	QuicTranspErrKeyUpdateError          uint64 `csv:"quic_transp_err_key_update_error"`
	QuicTranspErrAeadLimitReached        uint64 `csv:"quic_transp_err_aead_limit_reached"`
	QuicTranspErrNoViablePath            uint64 `csv:"quic_transp_err_no_viable_path"`
	QuicTranspErrCryptoError             uint64 `csv:"quic_transp_err_crypto_error"`
	QuicTranspErrUnknownError            uint64 `csv:"quic_transp_err_unknown_error"`
	QuicDataBlocked                      uint64 `csv:"quic_data_blocked"`
	QuicStreamDataBlocked                uint64 `csv:"quic_stream_data_blocked"`
	QuicStreamsBlockedBidi               uint64 `csv:"quic_streams_blocked_bidi"`
	QuicStreamsBlockedUni                uint64 `csv:"quic_streams_blocked_uni"`

	// h3_stats in src/h3_stats.c
	H3Data                  uint64 `csv:"h3_data"`
	H3Headers               uint64 `csv:"h3_headers"`
	H3CancelPush            uint64 `csv:"h3_cancel_push"`
	H3PushPromise           uint64 `csv:"h3_push_promise"`
	H3MaxPushId             uint64 `csv:"h3_max_push_id"`
	H3Goaway                uint64 `csv:"h3_goaway"`
	H3Settings              uint64 `csv:"h3_settings"`
	H3NoError               uint64 `csv:"h3_no_error"`
	H3GeneralProtocolError  uint64 `csv:"h3_general_protocol_error"`
	H3InternalError         uint64 `csv:"h3_internal_error"`
	H3StreamCreationError   uint64 `csv:"h3_stream_creation_error"`
	H3ClosedCriticalStream  uint64 `csv:"h3_closed_critical_stream"`
	H3FrameUnexpected       uint64 `csv:"h3_frame_unexpected"`
	H3FrameError            uint64 `csv:"h3_frame_error"`
	H3ExcessiveLoad         uint64 `csv:"h3_excessive_load"`
	H3IdError               uint64 `csv:"h3_id_error"`
	H3SettingsError         uint64 `csv:"h3_settings_error"`
	H3MissingSettings       uint64 `csv:"h3_missing_settings"`
	H3RequestRejected       uint64 `csv:"h3_request_rejected"`
	H3RequestCancelled      uint64 `csv:"h3_request_cancelled"`
	H3RequestIncomplete     uint64 `csv:"h3_request_incomplete"`
	H3MessageError          uint64 `csv:"h3_message_error"`
	H3ConnectError          uint64 `csv:"h3_connect_error"`
	H3VersionFallback       uint64 `csv:"h3_version_fallback"`
	PackDecompressionFailed uint64 `csv:"pack_decompression_failed"`
	QpackEncoderStreamError uint64 `csv:"qpack_encoder_stream_error"`
	QpackDecoderStreamError uint64 `csv:"qpack_decoder_stream_error"`
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
