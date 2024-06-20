package sysmodel

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
)
import sq "github.com/Masterminds/squirrel"

var _ SysUserModel = (*customSysUserModel)(nil)

type (
	// SysUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysUserModel.
	SysUserModel interface {
		sysUserModel
		withSession(session sqlx.Session) SysUserModel

		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error

		UpdateBuilder() sq.UpdateBuilder
		UpdateByQuery(ctx context.Context, updateBuilder sq.UpdateBuilder) error

		RowBuilder() sq.SelectBuilder
		FindOneByQuery(ctx context.Context, rowBuilder sq.SelectBuilder) (*SysUser, error)
		FindRowsByQuery(ctx context.Context, rowBuilder sq.SelectBuilder, orderBy string) ([]*SysUser, error)

		CountBuilder(field string) sq.SelectBuilder
		FindCount(ctx context.Context, countBuilder sq.SelectBuilder) (int64, error)

		FindAll(ctx context.Context, rowBuilder sq.SelectBuilder, orderBy string) ([]*SysUserList, error)

		TableName() string
	}

	customSysUserModel struct {
		*defaultSysUserModel
	}

	SysUserList struct {
		Id         int64     `db:"id"`          // 编号
		Name       string    `db:"name"`        // 账号
		NickName   string    `db:"nick_name"`   // 名称
		Avatar     string    `db:"avatar"`      // 头像
		Password   string    `db:"password"`    // 密码
		Salt       string    `db:"salt"`        // 加密盐
		Email      string    `db:"email"`       // 邮箱
		Mobile     string    `db:"mobile"`      // 手机号
		Status     int64     `db:"status"`      // 状态  -1：禁用   1：正常
		CreateBy   string    `db:"create_by"`   // 创建人
		CreateTime time.Time `db:"create_time"` // 创建时间
		UpdateBy   string    `db:"update_by"`   // 更新人
		UpdateTime time.Time `db:"update_time"` // 更新时间
		DelFlag    int64     `db:"del_flag"`    // 是否删除  1：已删除  0：正常
		RoleId     int64     `db:"role_id"`
		RoleName   string    `db:"role_name"`
	}
)

func (m *customSysUserModel) UpdateByQuery(ctx context.Context, updateBuilder sq.UpdateBuilder) error {
	query, values, err := updateBuilder.Where("del_flag = ?", 0).ToSql()
	if err != nil {
		return err
	}
	_, err = m.conn.ExecCtx(ctx, query, values...)
	return err
}

func (m *customSysUserModel) UpdateBuilder() sq.UpdateBuilder {
	return sq.Update(m.table)
}

func (m *customSysUserModel) Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error {
	return m.conn.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})
}

func (m *customSysUserModel) TableName() string {
	return m.table
}

func (m *customSysUserModel) FindAll(ctx context.Context, rowBuilder sq.SelectBuilder, orderBy string) ([]*SysUserList, error) {
	if orderBy == "" {
		rowBuilder = rowBuilder.OrderBy("id AEC")
	} else {
		rowBuilder = rowBuilder.OrderBy(orderBy)
	}

	query, values, err := rowBuilder.Where("del_flag = ?", 0).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*SysUserList
	err = m.conn.QueryRowsCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return nil, errors.New("查询记录为空")
	default:
		return nil, err
	}
}

func (m *customSysUserModel) FindCount(ctx context.Context, countBuilder sq.SelectBuilder) (int64, error) {
	query, values, err := countBuilder.Where("del_flag = ?", 0).ToSql()
	if err != nil {
		return 0, err
	}

	var resp int64
	err = m.conn.QueryRowCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *customSysUserModel) CountBuilder(field string) sq.SelectBuilder {
	return sq.Select("COUNT(" + field + ")").From(m.table)
}

func (m *customSysUserModel) FindRowsByQuery(ctx context.Context, rowBuilder sq.SelectBuilder, orderBy string) ([]*SysUser, error) {
	if orderBy == "" {
		rowBuilder = rowBuilder.OrderBy("id DESC")
	} else {
		rowBuilder = rowBuilder.OrderBy(orderBy)
	}

	query, values, err := rowBuilder.Where("del_flag = ?", 0).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*SysUser
	err = m.conn.QueryRowCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return nil, errors.New("查询记录为空")
	default:
		return nil, err
	}
}

func (m *customSysUserModel) FindOneByQuery(ctx context.Context, rowBuilder sq.SelectBuilder) (*SysUser, error) {
	query, values, err := rowBuilder.Where("del_flag = ?", 0).Limit(1).ToSql()
	if err != nil {
		return nil, err
	}

	var resp SysUser
	err = m.conn.QueryRowCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return &resp, nil
	default:
		return nil, err
	}
}

func (m *customSysUserModel) RowBuilder() sq.SelectBuilder {
	return sq.Select(sysUserRows).From(m.table)
}

// NewSysUserModel returns a model for the database table.
func NewSysUserModel(conn sqlx.SqlConn) SysUserModel {
	return &customSysUserModel{
		defaultSysUserModel: newSysUserModel(conn),
	}
}

func (m *customSysUserModel) withSession(session sqlx.Session) SysUserModel {
	return NewSysUserModel(sqlx.NewSqlConnFromSession(session))
}
