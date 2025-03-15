type UserToken struct {
    ID        uint      `gorm:"primaryKey;autoIncrement"`
    UserID    uint      `gorm:"not null"`
    Token     string    `gorm:"unique;not null"`
    IsValid   bool      `gorm:"not null;default:true"`
    CreatedAt time.Time
}
