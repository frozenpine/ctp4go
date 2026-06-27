package mduser

import (
	"fmt"
)

//go:generate stringer -type traderState -linecomment
type mduserState uint8

const (
	Created      mduserState = iota // 已创建
	Initialized                     // 已初始化
	Connected                       // 已连接
	Disconnected                    // 已断连
	LoginSuccess                    // 登录成功
	LoginFailed                     // 登录失败
	LoggedOut                       // 已登出
	Finalized                       // 已终止
)

func (curr *mduserState) Migrate(next mduserState) error {
	switch next {
	case Created:
	case Initialized:
		switch *curr {
		case Created, Finalized:
		default:
			goto ERROR
		}
	case Connected:
		switch *curr {
		case Initialized, Disconnected:
		default:
			goto ERROR
		}
	case Disconnected:
		switch *curr {
		case Created, Finalized:
			goto ERROR
		}
	case LoginSuccess, LoginFailed:
		if *curr != Connected {
			goto ERROR
		}
	case LoggedOut:
		if *curr != LoginSuccess {
			goto ERROR
		}
	case Finalized:
	}

	*curr = next
	return nil

ERROR:
	return fmt.Errorf(
		"%w: can not migrate to %s from %s",
		ErrInvalidState, next.String(), curr.String(),
	)
}
