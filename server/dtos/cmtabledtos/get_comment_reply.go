package cmtabledtos

import "outstagram/server/dtos/dtomodels"

type GetCommentRepliesRequest struct {
	Limit  uint `form:"limit"`
	Offset uint `form:"offset"`
}

type GetCommentRepliesResponse struct {
	Replies    []dtomodels.Reply `json:"replies"`
	ReplyCount int               `json:"replyCount"`
}
