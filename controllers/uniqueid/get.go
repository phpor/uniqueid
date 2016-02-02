package uniqueid
import "github.com/astaxie/beego"

type UniqueidController struct {
	beego.Controller
}

func (c *UniqueidController) Get() {

	var uids []string
	uids[0] = "123"
	c.Data["json"] = map[string][]string{"uids":uids}
	c.ServeJSON()
}
