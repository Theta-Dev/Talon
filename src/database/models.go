package database

import (
	"database/sql"
	"time"
)

var AllModels = []interface{}{
	&Website{},
	&Version{},
	&VersionFile{},
	&File{},
	&User{},
	&ApiUser{},
	&Permission{},
}

var TableNames = []string{
	"version_files",
	"versions",
	"websites",
	"files",
	"api_users",
	"users",
	"permissions",
}

const (
	VISIBILITY_HIDDEN     = 1
	VISIBILITY_SEARCHABLE = 2
	VISIBILITY_VISIBLE    = 3
)

type Website struct {
	ID          uint          `gorm:"primary_key;unique_index;not null;auto_increment"`
	Name        string        `gorm:"type:varchar(200);not null"`
	Path        string        `gorm:"type:varchar(200);not null"`
	PathLower   string        `gorm:"type:varchar(200);not null;unique_index"`
	Logo        *File         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	LogoID      sql.NullInt64 `gorm:"type:uint"`
	Color       string        `gorm:"type:varchar(20);not null"`
	Visibility  int           ``
	User        *User         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	UserID      uint          ``
	CreatedAt   time.Time     ``
	MaxVersions int           ``
	SourceUrl   string        `gorm:"type:varchar(200);not null"`
	SourceType  string        `gorm:"type:varchar(20);not null"`
	Versions    []Version     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type Version struct {
	ID        uint          `gorm:"primary_key;unique_index;not null;auto_increment"`
	Name      string        `gorm:"type:varchar(200);not null"`
	Tags      string        `gorm:"type:varchar(1000);not null"`
	Website   *Website      ``
	WebsiteID uint          ``
	User      *User         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	UserID    uint          ``
	CreatedAt time.Time     ``
	Files     []VersionFile `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type VersionFile struct {
	ID        uint     `gorm:"primary_key;unique_index;not null;auto_increment"`
	Path      string   `gorm:"type:varchar(200);not null"`
	Version   *Version `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	VersionID uint     ``
	File      *File    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	FileID    uint     ``
}

type File struct {
	ID   uint   `gorm:"primary_key;unique_index;not null;auto_increment"`
	Hash string `gorm:"type:varchar(64);unique_index;not null"`
}

type User struct {
	ID           uint        `gorm:"primary_key;unique_index;not null;auto_increment"`
	Name         string      `gorm:"type:varchar(50);not null;unique_index"`
	PasswordHash string      `gorm:"type:varchar(200);not null"`
	CreatedAt    time.Time   ``
	Permission   *Permission `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	PermissionID uint        ``
}

type ApiUser struct {
	ID           uint        `gorm:"primary_key;unique_index;not null;auto_increment"`
	KeyHash      string      `gorm:"type:varchar(200);not null"`
	CreatedAt    time.Time   ``
	Creator      *User       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatorID    uint        ``
	Permission   *Permission `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	PermissionID uint        ``
}

type Permission struct {
	ID            uint   `gorm:"primary_key;unique_index;not null;auto_increment"`
	AllowedPaths  string `gorm:"type:varchar(500);not null"`
	IsAdmin       bool   ``
	CanCreate     bool   ``
	MaxSize       int    ``
	MaxVersions   int    ``
	MaxVisibility int    ``
}
