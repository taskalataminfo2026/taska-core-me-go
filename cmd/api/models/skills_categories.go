package models

type SkillCategory struct {
	ID         int64
	SkillID    int64
	CategoryID int64
	IsPrimary  bool
	IsActive   bool
}

type ParamsSkillsCategorySearch struct {
	ID         int64
	SkillID    int64
	CategoryID int64
	IsPrimary  bool
	IsActive   bool
	Limit      int
	Offset     int
}

type ParamsSkillsCategorySave struct {
	SkillID    int64
	CategoryID int64
	IsPrimary  bool
	IsActive   bool
}
