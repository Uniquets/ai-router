package runpodapi

type GenerateReq struct {
	Input   map[InputColumn]interface{} `json:"input"`
	WebHook string                      `json:"webhook"` // 用于接收结果的 Webhook URL。
}

type InputColumn string

var InputColumns = struct {
	Prompt                InputColumn `json:"prompt"`                  // 输入提示。
	NegativePrompt        InputColumn `json:"negative_prompt"`         // 反向提示词
	NumInferenceSteps     InputColumn `json:"num_inference_steps"`     // 去噪步数。范围 1 - 500
	RefinerInferenceSteps InputColumn `json:"refiner_inference_steps"` // 精简程序模型的步骤数。
	Width                 InputColumn `json:"width"`                   // 输出图像的宽度。
	Height                InputColumn `json:"height"`                  // 输出图像的高度。
	GuidanceScale         InputColumn `json:"guidance_scale"`          // 无分类器指导的缩放。范围1 - 20
	Strength              InputColumn `json:"strength"`                // 使用 Refiner 模型时的权重。范围 0 - 1
	Seed                  InputColumn `json:"seed"`                    // 随机种子。
	NumImages             InputColumn `json:"num_images"`              // 要生成的图像数量，最多 2 张。
	ImageUrl              InputColumn `json:"image_url"`               // 要使用精简程序模型优化的初始图像的 URL。
}{
	Prompt:                "prompt",
	NegativePrompt:        "negative_prompt",
	NumInferenceSteps:     "num_inference_steps",
	RefinerInferenceSteps: "refiner_inference_steps",
	Width:                 "width",
	Height:                "height",
	GuidanceScale:         "guidance_scale",
	Strength:              "strength",
	Seed:                  "seed",
	NumImages:             "num_images",
	ImageUrl:              "image_url",
}

type GenerateResp struct {
	DelayTime     int    `json:"delayTime"`
	ExecutionTime int    `json:"executionTime"`
	Id            string `json:"id"`
	Output        struct {
		ImageUrl string   `json:"image_url"` // 图片路径
		Images   []string `json:"images"`
		Seed     int      `json:"seed"`
	} `json:"output"`
	Status string `json:"status"` // IN_QUEUE=队列中,IN_PROGRESS=进程中,COMPLETED=已完成
}
