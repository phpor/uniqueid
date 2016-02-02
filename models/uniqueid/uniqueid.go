package uniqueid
import (
	"strconv"
)

type Uniqueid struct {

}

var topics = []string{"test1", "test2"}
var uniqueidChans = map[string](chan int32){}

func Get(topic string, num ...int) (uids []string,err error) {
	i := 0
	cnt := 1
	if len(num) > 0 {
		cnt = num[0]
	}
	for i < cnt {
		//todo: 这里需要处理chan已关闭的情况
		uids = append(uids, strconv.Itoa(int(<-uniqueidChans[topic])))
		i++
	}
	return
}

func init() {
	for _, topic := range topics {
		uniqueidChans[topic] = make(chan int32, 10)
		println(topic)
		go func(topic string) {
			i := int32(0)
			for {
				println(topic, "len:" , len(uniqueidChans[topic]))
				uniqueidChans[topic] <- i
				i++
			}
		}(topic)
	}
}