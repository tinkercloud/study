package main

// 2020-10-09 实现ip地址位置查询
import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lionsoul2014/ip2region/binding/golang/ip2region"
)

func main() {

	r := gin.Default()

	r.POST("/secip", getIP)

	r.Run(":8082")
}

type SearchData struct {
	IP string `form:"ip"`
}

func getIP(c *gin.Context) {

	var sd SearchData

	c.ShouldBindQuery(&sd)

	region, err := ip2region.New("c:\\download\\202010\\ip2region.db")

	if err != nil {
		log.Fatal(err)
	}

	ipdata, err := region.BinarySearch(sd.IP)

	if err != nil {
		log.Fatal(err)
	}

	c.String(200, ipdata.String())
}
