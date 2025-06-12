# notice
注册通知服务
## 用法
服务端
~~~go
import (
	"fmt"
	"github.com/goodluckxu-go/notice"
)

func main() {
	addr := "127.0.0.1:8080"
	fmt.Println("Listening on " + addr)
	err := notice.Listen(addr)
	if err != nil {
		panic(err)
	}
}
~~~
客户端
~~~go
import (
	"github.com/goodluckxu-go/notice"
	"net/http"
)

func main() {
	// 注册服务
	service, err := notice.Dail("127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	defer service.Close()
	go func() {
		// 循环接受通知
		err = service.RecvMessage(func(id string, message []byte) {
			// 服务接受通知该服务下注册的客户端消息
		})
	}()
	err = service.AddClient("id", map[string]any{"uid": 1, "name": "zs"})
	if err != nil {
		panic(err)
	}

	// 可直接更加id给客户端发消息，RecvMessage可接受到消息
	// 也可更加condition条件给筛选的客户端发消息，RecvMessage可接受到消息
	err = service.SendMessage([]byte("消息"), nil, nil)
	// 删除客户端
	err = service.DelClient("id")
}
~~~
## 可用场景
1. websocket集群通知
2. 任务集群通知
3. 其他通知应用