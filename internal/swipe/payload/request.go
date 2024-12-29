package payload

type (
	CreateSwipeReq struct {
		SwipedUserID string `json:"swiped_user_id" binding:"required" gorm:"column:swiped_user_id"`
		SwipeType    string `json:"swipe_type" binding:"required" gorm:"column:swipe_type"`
	}
)
