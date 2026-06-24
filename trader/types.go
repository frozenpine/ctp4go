package trader

import (
	"fmt"
)

//go:generate stringer -type traderState -linecomment
type traderState uint8

const (
	Created      traderState = iota // 已创建
	Initialized                     // 已初始化
	Connected                       // 已连接
	Disconnected                    // 已断连
	AuthSuccess                     // 授权成功
	AuthFailed                      // 授权失败
	LoginSuccess                    // 登录成功
	LoginFailed                     // 登录失败
	LoggedOut                       // 已登出
	Finalized                       // 已终止
)

func (curr *traderState) Migrate(next traderState) error {
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
	case AuthSuccess, AuthFailed:
		if *curr != Connected {
			goto ERROR
		}
	case LoginSuccess, LoginFailed:
		if *curr != AuthSuccess {
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
