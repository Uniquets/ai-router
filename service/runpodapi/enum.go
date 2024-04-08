package runpodapi

type ModelName string

var (
	SDv1          ModelName = "stable-diffusion-v1"
	SDv2          ModelName = "stable-diffusion-v2"
	SDXL          ModelName = "sdxl"
	SDAnythingV3  ModelName = "sd-anything-v3"
	SDAnythingV4  ModelName = "sd-anything-v4"
	SDOpenJourney ModelName = "sd-openjourney"
	DreamBoothV1  ModelName = "dream-booth-v1"
	KandinskyV2   ModelName = "kandinsky-v2"
)
