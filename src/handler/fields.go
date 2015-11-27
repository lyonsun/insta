package handler

type Counts struct {
	Media      int "json:`media`"
	Follows    int "json:`follows`"
	FollowedBy int "json:`followed_by`"
}

type User struct {
	ID             string "json:`id`"
	UserName       string "json:`username`"
	FullName       string "json:`full_name`"
	ProfilePicture string "json:`profile_picture`"
	Bio            string "json:`bio`"
	Website        string "json:`website`"
	Counts         Counts
}

type UserInfo struct {
	Data User "json:`data`"
}
