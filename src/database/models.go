package database

import (
	"time"

	"gorm.io/gorm"
)

var allModels = []interface{}{
	&Website{},
	&Version{},
	&File{},
	&User{},
	&ApiUser{},
	&Permission{},
	&TalonInfo{},
}

var tableNames = []string{
	"websites",
	"versions",
	"files",
	"users",
	"api_users",
	"permissions",
	"talon_infos",
}

type Website struct {
	ID          uint      `gorm:"primary_key;unique_index;not null;auto_increment"`
	Name        string    `gorm:"type:varchar(100);not null"`
	Path        string    `gorm:"type:varchar(200);not null"`
	PathLower   string    `gorm:"type:varchar(200);not null;unique_index"`
	Logo        *File     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	LogoID      *uint     ``
	Color       *string   `gorm:"type:varchar(20)"`
	Visibility  uint      ``
	User        User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	UserID      uint      ``
	CreatedAt   time.Time ``
	MaxVersions uint      ``
	SourceUrl   *string   `gorm:"type:varchar(200)"`
	SourceType  *string   `gorm:"type:varchar(20)"`
	Versions    []Version `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type Version struct {
	ID        uint      `gorm:"primary_key;unique_index;not null;auto_increment"`
	Name      string    `gorm:"type:varchar(100);not null"`
	Website   Website   ``
	WebsiteID uint      ``
	User      User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	UserID    uint      ``
	CreatedAt time.Time ``
	Files     []File    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type File struct {
	ID        uint    `gorm:"primary_key;unique_index;not null;auto_increment"`
	Path      string  `gorm:"type:varchar(200)"`
	Hash      string  `gorm:"type:varchar(64)"`
	Version   Version ``
	VersionID uint    ``
}

type User struct {
	ID           uint           `gorm:"primary_key;unique_index;not null;auto_increment"`
	Name         string         `gorm:"type:varchar(50)"`
	PasswordHash string         `gorm:"type:varchar(200)"`
	CreatedAt    time.Time      ``
	Deleted      gorm.DeletedAt ``
	Permission   Permission     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	PermissionID uint           ``
}

type ApiUser struct {
	ID           uint       `gorm:"primary_key;unique_index;not null;auto_increment"`
	KeyHash      string     `gorm:"type:varchar(200)"`
	CreatedAt    time.Time  ``
	Creator      User       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatorID    uint       ``
	Permission   Permission `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	PermissionID uint       ``
}

type Permission struct {
	ID            uint   `gorm:"primary_key;unique_index;not null;auto_increment"`
	AllowedPaths  string `gorm:"type:varchar(500)"`
	CanCreate     bool   ``
	MaxSize       uint   ``
	MaxVersions   uint   ``
	MaxVisibility uint   ``
}

type TalonInfo struct {
	Key   string `gorm:"type:varchar(20);primary_key;unique_index;not null"`
	Value string `gorm:"type:varchar(200)"`
}
