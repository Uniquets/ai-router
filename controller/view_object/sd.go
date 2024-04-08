package view_object

type GenerateResp struct {
	DelayTime     int    `json:"delayTime"`     // 已执行时间
	ExecutionTime int    `json:"executionTime"` // 总执行时间,Status=COMPLETED时有值
	Id            string `json:"id"`            // 进程ID
	Output        struct {
		ImageUrl string   `json:"image_url"` // 图片路径
		Images   []string `json:"images"`
		Seed     int      `json:"seed"`
	} `json:"output"` // 输出,Status=COMPLETED时有值
	Status string `json:"status"` // IN_QUEUE=队列中,IN_PROGRESS=进程中,COMPLETED=已完成
}
