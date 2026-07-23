package impl

import (
	"blog/apps/user"
	"blog/common"
	"blog/conf"
	"blog/ioc"
	"context"

	"gorm.io/gorm"
)

func init() {
	ioc.Contorller.Registry(user.AppName, &UserServiceImpl{})
}

type UserServiceImpl struct {
	db *gorm.DB
}

func (i *UserServiceImpl) Init() error {
	i.db = conf.C().Mysql.GetDB()
	return nil
}

func (i *UserServiceImpl) CreateUser(ctx context.Context, in *user.CreateUserReuqest) (*user.User, error) {
	if err := common.Validate(in); err != nil {
		return nil, err
	}
	if err := in.HashPassword(); err != nil {
		return nil, err
	}
	us := user.NewUser(in)

	if err := i.db.WithContext(ctx).Save(us).Error; err != nil {
		return nil, err
	}
	return us, nil
}

func (i *UserServiceImpl) QueryUser(ctx context.Context, in *user.QueryUserRequest) (*user.UserSet, error) {
	set := user.NewUserSet()
	query := i.db.WithContext(ctx).Model(&user.User{})
	if in.Username != "" {
		query = query.Where("username=?", in.Username)
	}
	if err := query.Count(&set.Total).Error; err != nil {
		return nil, err
	}
	if err := query.Offset(in.Offset()).Limit(in.PageSize).Find(&set.Items).Error; err != nil {
		return nil, err
	}
	return set, nil
}
