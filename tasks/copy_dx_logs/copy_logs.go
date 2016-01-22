package copy_dx_logs
import (
	"github.schq.secious.com/Logrhythm/easy/conf"
	"github.schq.secious.com/Logrhythm/easy/tasks"
)


type Copier struct {
}

func (c *Copier) Exec(cf *conf.Conf) chan tasks.Feedback {
	return make(chan tasks.Feedback)
}

func (c *Copier) Rise(cf *conf.Conf) chan tasks.Feedback {
	return make(chan tasks.Feedback)
}

func (c *Copier) Usage() string {
	return "Copies DX logs to target."
}

func RiseUp(cf *conf.Conf) chan tasks.Feedback {
	return make(chan tasks.Feedback)
}