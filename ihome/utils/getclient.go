package utils

func GetStatusCode(key string) map[string]interface{} {
	temp := make(map[string]interface{})
	data := make(map[string]string)
	data["errno"] = key
	data["errmsg"] = RecodeText(key)
	temp["data"] = data
	return temp
}
