package main

import (
	"flag"
	"fmt"
)


func main() {
	useGpu := flag.Bool("use_gpu", true, "GPUを使うか")
	cpuNumThreads := flag.Int("cpu_num_threads", 0, "CPU処理のスレッド数(0: 自動)")
	text := flag.String("text", "", "読ませる文字列")
	speakerId := flag.Int("speaker_id", 0, "スピーカー(0: 四国めたん, 1: ずんだもん)")
	openjtalkDict := flag.String("openjtalk_dict", "open_jtalk_dic_utf_8-1.11", "")
	flag.Parse()
	fmt.Println(*useGpu, *cpuNumThreads, *text, *speakerId, *openjtalkDict)
	h, err := Initialize(*useGpu, *cpuNumThreads)
	fmt.Println("Initialize Call: Handle:", h, "Error:", err)

	m, err := Metas()
	fmt.Println("Metas Call: Handle:", m, "Error:", err)

	o, err := VoicevoxLoadOpenjtalkDict(*openjtalkDict)
	fmt.Println("VoicevoxLoadOpenjtalkDict Call: Handle:", o, "Error:", err)
}
