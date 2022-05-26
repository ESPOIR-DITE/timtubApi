package util

import (
	role_repo "timtubeApi/storage/user/role-repo"
	user_account_repo "timtubeApi/storage/user/user-account-repo"
	user_repo "timtubeApi/storage/user/user-repo"
	user_sub_repo "timtubeApi/storage/user/user-sub-repo"
	user_video_repo "timtubeApi/storage/user/user-video-repo"
	video_repo2 "timtubeApi/storage/video/category"
	video_category "timtubeApi/storage/video/video-category"
	video_category2 "timtubeApi/storage/video/video-comment"
	video_repo "timtubeApi/storage/video/video-repo"
)

func TableSetUp() []TableSetUpReport {
	var result []TableSetUpReport
	result = append(result, TableSetUpReport{"USER TABLE", user_repo.CreateUserTable()})
	result = append(result, TableSetUpReport{"USER- ACCOUNT TABLE", user_account_repo.CreateUserAccountTable()})
	result = append(result, TableSetUpReport{"USER- SUBSCRIPTION TABLE", user_sub_repo.CreateUserSubscriptionTable()})
	result = append(result, TableSetUpReport{"USER- VIDEO TABLE", user_video_repo.CreateUserVideoTable()})
	result = append(result, TableSetUpReport{"ROLE TABLE", role_repo.CreateRoleTable()})

	result = append(result, TableSetUpReport{"VIDEO TABLE", video_repo.CreateVideoTable()})
	result = append(result, TableSetUpReport{"VIDEO- CATEGORY TABLE", video_category.CreateVideoCategoryTable()})
	result = append(result, TableSetUpReport{"VIDEO- COMMENT TABLE", video_category2.CreateVideoCommentTable()})
	result = append(result, TableSetUpReport{"CATEGORY TABLE", video_repo2.CreateCategoryTable()})
	return result
}
