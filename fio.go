package dumbfioparser

type Fio struct {
	DiskUtil []struct {
		InQueue     float64 `json:"in_queue"`
		Name        string  `json:"name"`
		ReadIos     float64 `json:"read_ios"`
		ReadMerges  float64 `json:"read_merges"`
		ReadTicks   float64 `json:"read_ticks"`
		Util        float64 `json:"util"`
		WriteIos    float64 `json:"write_ios"`
		WriteMerges float64 `json:"write_merges"`
		WriteTicks  float64 `json:"write_ticks"`
	} `json:"disk_util"`
	Fio_version    string `json:"fio version"`
	Global_options struct {
		Bs               string `json:"bs"`
		Buffered         string `json:"buffered"`
		Direct           string `json:"direct"`
		Filesize         string `json:"filesize"`
		Iodepth          string `json:"iodepth"`
		Ioengine         string `json:"ioengine"`
		Nrfiles          string `json:"nrfiles"`
		Numjobs          string `json:"numjobs"`
		PercentageRandom string `json:"percentage_random"`
		Randrepeat       string `json:"randrepeat"`
		Runtime          string `json:"runtime"`
		Rw               string `json:"rw"`
		Rwmixwrite       string `json:"rwmixwrite"`
	} `json:"global options"`
	Jobs []struct {
		Jobname           string  `json:"jobname"`
		Ctx               float64 `json:"ctx"`
		Elapsed           float64 `json:"elapsed"`
		Error             float64 `json:"error"`
		Eta               float64 `json:"eta"`
		Groupid           float64 `json:"groupid"`
		LatencyDepth      float64 `json:"latency_depth"`
		LatencyPercentile float64 `json:"latency_percentile"`
		LatencyTarget     float64 `json:"latency_target"`
		LatencyWindow     float64 `json:"latency_window"`
		Majf              float64 `json:"majf"`
		Minf              float64 `json:"minf"`
		Read              struct {
			Bw          float64 `json:"bw"`
			BwAgg       float64 `json:"bw_agg"`
			BwBytes     float64 `json:"bw_bytes"`
			BwDev       float64 `json:"bw_dev"`
			BwMax       float64 `json:"bw_max"`
			BwMean      float64 `json:"bw_mean"`
			BwMin       float64 `json:"bw_min"`
			BwSamples   float64 `json:"bw_samples"`
			DropIos     float64 `json:"drop_ios"`
			IoBytes     float64 `json:"io_bytes"`
			IoKbytes    float64 `json:"io_kbytes"`
			Iops        float64 `json:"iops"`
			IopsMax     float64 `json:"iops_max"`
			IopsMean    float64 `json:"iops_mean"`
			IopsMin     float64 `json:"iops_min"`
			IopsSamples float64 `json:"iops_samples"`
			IopsStddev  float64 `json:"iops_stddev"`
			LatNs       struct {
				N      float64 `json:"N"`
				Max    float64 `json:"max"`
				Mean   float64 `json:"mean"`
				Min    float64 `json:"min"`
				Stddev float64 `json:"stddev"`
			} `json:"lat_ns"`
			Runtime  float64 `json:"runtime"`
			ShortIos float64 `json:"short_ios"`
			SlatNs   struct {
				N      float64 `json:"N"`
				Max    float64 `json:"max"`
				Mean   float64 `json:"mean"`
				Min    float64 `json:"min"`
				Stddev float64 `json:"stddev"`
			} `json:"slat_ns"`
			TotalIos float64 `json:"total_ios"`
		} `json:"read"`
		Sync struct {
			LatNs struct {
				N      float64 `json:"N"`
				Max    float64 `json:"max"`
				Mean   float64 `json:"mean"`
				Min    float64 `json:"min"`
				Stddev float64 `json:"stddev"`
			} `json:"lat_ns"`
			TotalIos float64 `json:"total_ios"`
		} `json:"sync"`
		SysCPU float64 `json:"sys_cpu"`
		Trim   struct {
			Bw        float64 `json:"bw"`
			BwAgg     float64 `json:"bw_agg"`
			BwBytes   float64 `json:"bw_bytes"`
			BwDev     float64 `json:"bw_dev"`
			BwMax     float64 `json:"bw_max"`
			BwMean    float64 `json:"bw_mean"`
			BwMin     float64 `json:"bw_min"`
			BwSamples float64 `json:"bw_samples"`
			ClatNs    struct {
				N      float64 `json:"N"`
				Max    float64 `json:"max"`
				Mean   float64 `json:"mean"`
				Min    float64 `json:"min"`
				Stddev float64 `json:"stddev"`
			} `json:"clat_ns"`
			DropIos     float64 `json:"drop_ios"`
			IoBytes     float64 `json:"io_bytes"`
			IoKbytes    float64 `json:"io_kbytes"`
			Iops        float64 `json:"iops"`
			IopsMax     float64 `json:"iops_max"`
			IopsMean    float64 `json:"iops_mean"`
			IopsMin     float64 `json:"iops_min"`
			IopsSamples float64 `json:"iops_samples"`
			IopsStddev  float64 `json:"iops_stddev"`
			LatNs       struct {
				N      float64 `json:"N"`
				Max    float64 `json:"max"`
				Mean   float64 `json:"mean"`
				Min    float64 `json:"min"`
				Stddev float64 `json:"stddev"`
			} `json:"lat_ns"`
			Runtime  float64 `json:"runtime"`
			ShortIos float64 `json:"short_ios"`
			SlatNs   struct {
				N      float64 `json:"N"`
				Max    float64 `json:"max"`
				Mean   float64 `json:"mean"`
				Min    float64 `json:"min"`
				Stddev float64 `json:"stddev"`
			} `json:"slat_ns"`
			TotalIos float64 `json:"total_ios"`
		} `json:"trim"`
		UsrCPU float64 `json:"usr_cpu"`
		Write  struct {
			Bw        float64 `json:"bw"`
			BwAgg     float64 `json:"bw_agg"`
			BwBytes   float64 `json:"bw_bytes"`
			BwDev     float64 `json:"bw_dev"`
			BwMax     float64 `json:"bw_max"`
			BwMean    float64 `json:"bw_mean"`
			BwMin     float64 `json:"bw_min"`
			BwSamples float64 `json:"bw_samples"`
			ClatNs    struct {
				N          float64 `json:"N"`
				Max        float64 `json:"max"`
				Mean       float64 `json:"mean"`
				Min        float64 `json:"min"`
				Percentile struct {
					One_000000    float64 `json:"1.000000"`
					One0_000000   float64 `json:"10.000000"`
					Two0_000000   float64 `json:"20.000000"`
					Three0_000000 float64 `json:"30.000000"`
					Four0_000000  float64 `json:"40.000000"`
					Five_000000   float64 `json:"5.000000"`
					Five0_000000  float64 `json:"50.000000"`
					Six0_000000   float64 `json:"60.000000"`
					Seven0_000000 float64 `json:"70.000000"`
					Eight0_000000 float64 `json:"80.000000"`
					Nine0_000000  float64 `json:"90.000000"`
					Nine5_000000  float64 `json:"95.000000"`
					Nine9_000000  float64 `json:"99.000000"`
					Nine9_500000  float64 `json:"99.500000"`
					Nine9_900000  float64 `json:"99.900000"`
					Nine9_950000  float64 `json:"99.950000"`
					Nine9_990000  float64 `json:"99.990000"`
				} `json:"percentile"`
				Stddev float64 `json:"stddev"`
			} `json:"clat_ns"`
			DropIos     float64 `json:"drop_ios"`
			IoBytes     float64 `json:"io_bytes"`
			IoKbytes    float64 `json:"io_kbytes"`
			Iops        float64 `json:"iops"`
			IopsMax     float64 `json:"iops_max"`
			IopsMean    float64 `json:"iops_mean"`
			IopsMin     float64 `json:"iops_min"`
			IopsSamples float64 `json:"iops_samples"`
			IopsStddev  float64 `json:"iops_stddev"`
			LatNs       struct {
				N      float64 `json:"N"`
				Max    float64 `json:"max"`
				Mean   float64 `json:"mean"`
				Min    float64 `json:"min"`
				Stddev float64 `json:"stddev"`
			} `json:"lat_ns"`
			Runtime  float64 `json:"runtime"`
			ShortIos float64 `json:"short_ios"`
			SlatNs   struct {
				N      float64 `json:"N"`
				Max    float64 `json:"max"`
				Mean   float64 `json:"mean"`
				Min    float64 `json:"min"`
				Stddev float64 `json:"stddev"`
			} `json:"slat_ns"`
			TotalIos float64 `json:"total_ios"`
		} `json:"write"`
	} `json:"jobs"`
	Time        string  `json:"time"`
	Timestamp   float64 `json:"timestamp"`
	TimestampMs float64 `json:"timestamp_ms"`
}
