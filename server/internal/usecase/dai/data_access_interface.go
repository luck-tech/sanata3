package dai

type DataAccessInterface interface {
	GitHubService
	User
	Session
	Skill
	UsedSkill
	WantLearnSkill
	AimSkill
	Room
	RoomMember
}
