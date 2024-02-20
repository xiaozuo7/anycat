package util

import (
	"strings"
	"testing"
	"time"
)

func TestXorDecryStrAll(t *testing.T) {
	data := `{"clientName":"本地客户端组","clientDsc":"","clientNodesId":0,"clientNodes":"1","nodeName":"本地客户端","groupIds":"2","agentTypes":"","ip":"fe80::1","port":6102,"cpu":"7.302632","memory":"73.395252","systemVersion":"arm64","netHeartbeat":"2023-06-01 03:29:15","heartbeatRate":30,"heartbeatTimeoutNum":3,"osType":"darwin","uuid":"0ef7446cbf5d34c54d0e34a6c1b5d6b1","maxConcurrentJobs":0,"clientNodeId":1,"clientId":2,"clientMode":0}`
	t.Logf("origin data len: %d\n", len(data))
	t1 := time.Now()
	encrypt := XorEncryptStr(data)
	t2 := time.Now()
	t.Logf("encry data: %s, spend time: %s\n", encrypt, t2.Sub(t1))
	decrypt, err := XorDecryptStr(encrypt)
	if err != nil {
		t.Error(err)
	}
	t3 := time.Now()
	t.Logf("decry data: %s, spend time: %s\n", decrypt, t3.Sub(t2))
	t.Logf("totoal spend time: %s", t3.Sub(t1))
}

func TestXorEncryptStr(t *testing.T) {
	data := "root"
	t.Logf("origin data len: %d\n", len(data))
	encrypt := XorEncryptStr(data)
	t.Logf("encry data: %s", encrypt)
}

func TestXorDecryptStr(t *testing.T) {
	data := "qm9dvS=="
	decrypt, err := XorDecryptStr(data)
	if err != nil {
		t.Error(err)
	}
	t.Logf("decry data: %s", decrypt)
}

func BenchmarkXorDecryStr(b *testing.B) {
	data := strings.Repeat("testStr", 10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		encry := XorEncryptStr(data)
		_, err := XorDecryptStr(encry)
		if err != nil {
			b.Error(err)
		}
	}
}
