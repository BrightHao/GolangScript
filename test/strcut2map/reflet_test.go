package strcut2map

import (
	"fmt"
	"testing"
)

func TestReflectTrans(t *testing.T) {
	obj := WriteData{
		DstIpType:    "",
		SrcIpType:    "",
		SrcIp:        "",
		DstIp:        "",
		SrcPort:      "",
		DstPort:      "",
		HttpsDomain:  "",
		HttpDomain:   "",
		InBytes:      114,
		OutputValue:  "",
		InputValue:   "",
		AgentIp:      "",
		SamplingRate: "",
		LogType:      "",
		FlowSeqNum:   "",
		TcpFlags:     "",
		HttpRespCode: "",
		Ttl:          "",
		SslName:      "",
		RequestUrl:   "",
		AgentType:    "",
		HttpMethod:   "",
		DscpName:     "",
		TimeSet:      "2022-02-17T18:57:01",
		Flow:         "10.36.255.10:5247_10.43.7.150:5256",
		SrcIpDstIp:   "10.36.255.10_10.43.7.150",
		NetWorkLine:  "",
		SourceNode1:  "办公网",
		SourceNode2:  "",
		SourceNode3:  "",
		SourceNode4:  "",
		SourceNode5:  "",
		DestNode1:    "",
		DestNode2:    "",
		DestNode3:    "",
		DestNode4:    "",
		DestNode5:    "",
		DestAPPId:    "",
		SourceAPPId:  "",
		SourceSite:   "",
		DestSite:     "DFDZ",
		SpecialLine:  "",
	}
	reflectTrans(obj)
}

func BenchmarkReflectTrans(b *testing.B) {
	obj := WriteData{
		DstIpType:    "",
		SrcIpType:    "",
		SrcIp:        "",
		DstIp:        "",
		SrcPort:      "",
		DstPort:      "",
		HttpsDomain:  "",
		HttpDomain:   "",
		InBytes:      114,
		OutputValue:  "",
		InputValue:   "",
		AgentIp:      "",
		SamplingRate: "",
		LogType:      "",
		FlowSeqNum:   "",
		TcpFlags:     "",
		HttpRespCode: "",
		Ttl:          "",
		SslName:      "",
		RequestUrl:   "",
		AgentType:    "",
		HttpMethod:   "",
		DscpName:     "",
		TimeSet:      "2022-02-17T18:57:01",
		Flow:         "10.36.255.10:5247_10.43.7.150:5256",
		SrcIpDstIp:   "10.36.255.10_10.43.7.150",
		NetWorkLine:  "",
		SourceNode1:  "办公网",
		SourceNode2:  "",
		SourceNode3:  "",
		SourceNode4:  "",
		SourceNode5:  "",
		DestNode1:    "",
		DestNode2:    "",
		DestNode3:    "",
		DestNode4:    "",
		DestNode5:    "",
		DestAPPId:    "",
		SourceAPPId:  "",
		SourceSite:   "",
		DestSite:     "DFDZ",
		SpecialLine:  "",
	}
	n := b.N
	fmt.Println("次数：", n)
	for i := 0; i < n; i++ {
		reflectTrans(obj)
	}
}
