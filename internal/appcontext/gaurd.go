package appcontext

import "com/example/gobot/internal/guard"


func (self *Context) GuardWith(f guard.GuardFunc) error {
	return f(self.RawUpdate)
}
