package models

import (
	"fmt"

	"github.com/Jackkakaya/NeteaseCloudMusicGoApi/pkg/request"
)

func (m *MusicObain) CommentAlbum(query map[string]interface{}) map[string]interface{} {
	data := map[string]interface{}{
		"rid": query["id"],
	}
	if val, ok := query["limit"]; ok {
		data["limit"] = val
	} else {
		data["limit"] = 20
	}
	if val, ok := query["offset"]; ok {
		data["offset"] = val
	} else {
		data["offset"] = 0
	}
	if val, ok := query["before"]; ok {
		data["beforeTime"] = val
	} else {
		data["beforeTime"] = 0
	}

	if val, ok := query["cookie"]; ok {
		valMapper := val.(map[string]interface{})
		valMapper["os"] = "pc"
		query["cookie"] = valMapper
	} else {
		query["cookie"] = map[string]interface{}{
			"os": "pc",
		}
	}

	options := map[string]interface{}{
		"crypto": "weapi",
		"cookie": query["cookie"],
		"proxy":  query["proxy"],
	}

	return request.CreateRequest(
		"POST", "https://music.163.com/weapi/v1/resource/comments/R_AL_3_"+fmt.Sprintf("%v", query["id"]),
		data,
		options)
}
