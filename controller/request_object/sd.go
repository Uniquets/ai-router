package request_object

type GenerateReq struct {
	Input struct {
		Prompt string `json:"prompt" validate:"required"` // 文生图描述语句
		//ImageNum int    `json:"image_num" validate:"omitempty,min=1,max=2" `  // 输出数量
		Width    int    `json:"width" validate:"omitempty,min=512,max=1024"`  // 宽 px
		Height   int    `json:"height" validate:"omitempty,min=512,max=1024"` // 高 px
		ImageUrl string `json:"image_url"`                                    // 用于图生图的图片url
		Seed     int    `json:"seed"`                                         // 种子ID
	}
	Async bool `json:"async"` // 是否异步
}

type CheckStatusReq struct {
	JobId string `json:"job_id" validate:"required"` // 任务ID
}
