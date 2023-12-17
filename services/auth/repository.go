package main

type UserProfile struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UserProfiles []*UserProfile

func CreateUserProfile(userProfile *UserProfile) {
	userProfiles = append(userProfiles, userProfile)
}

func GetUserProfiles() UserProfiles {
	return userProfiles
}

var userProfiles = UserProfiles{
	{
		FirstName: "Roman",
		LastName:  "Tkachenko",
	},
	{
		FirstName: "Kate",
		LastName:  "Krut",
	},
}
