package main

type UserProfile struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func CreateUserProfile(userProfile *UserProfile) {
	userProfiles = append(userProfiles, userProfile)
}

func GetUserProfiles() []*UserProfile {
	return userProfiles
}

var userProfiles = []*UserProfile{
	{
		FirstName: "Roman",
		LastName:  "Tkachenko",
	},
	{
		FirstName: "Kate",
		LastName:  "Krut",
	},
}
