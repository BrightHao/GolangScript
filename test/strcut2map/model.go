package strcut2map

type WriteData struct {
	DstIpType    string `json:"dst_ip_type"`
	SrcIpType    string `json:"src_ip_type"`
	SrcIp        string `json:"src_ip"`
	DstIp        string `json:"dst_ip"`
	SrcPort      string `json:"src_port"`
	DstPort      string `json:"dst_port"`
	HttpsDomain  string `json:"https_domain"`
	HttpDomain   string `json:"http_domain"`
	InBytes      int    `json:"in_bytes"`
	OutputValue  string `json:"output_value"`
	InputValue   string `json:"input_value"`
	AgentIp      string `json:"agent_ip"`
	SamplingRate string `json:"sampling_rate"`
	LogType      string `json:"log_type"`
	FlowSeqNum   string `json:"flow_seq_num"`
	TcpFlags     string `json:"tcp_flags"`
	HttpRespCode string `json:"http_resp_code"`
	Ttl          string `json:"ttl"`
	SslName      string `json:"ssl_name"`
	RequestUrl   string `json:"request_url"`
	AgentType    string `json:"agent_type"`
	HttpMethod   string `json:"http_method"`
	DscpName     string `json:"dscp_name"`
	TimeSet      string `json:"time_set"` // 从这里开始，都是处理生成的字段
	Flow         string `json:"flow"`
	SrcIpDstIp   string `json:"src_ip_dst_ip"`
	NetWorkLine  string `json:"net_work_line"`
	SourceNode1  string `json:"source_node_1"`
	SourceNode2  string `json:"source_node_2"`
	SourceNode3  string `json:"source_node_3"`
	SourceNode4  string `json:"source_node_4"`
	SourceNode5  string `json:"source_node_5"`
	DestNode1    string `json:"dest_node_1"`
	DestNode2    string `json:"dest_node_2"`
	DestNode3    string `json:"dest_node_3"`
	DestNode4    string `json:"dest_node_4"`
	DestNode5    string `json:"dest_node_5"`
	DestAPPId    string `json:"dest_app_id"`
	SourceAPPId  string `json:"source_app_id"`
	SourceSite   string `json:"source_site"`
	DestSite     string `json:"dest_site"`
	SpecialLine  string `json:"special_line"`
}
