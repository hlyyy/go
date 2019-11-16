package main
import(
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"io/ioutil"
)
func main() {
	requesturl :="https://accounts.douban.com/j/mobile/login/basic"
	data :=url.Values{}
	data.Set("name","13476099548")
	data.Set("password","huanglingyun0130")
	data.Set("remember","false")
	data.Set("ticket","")
	data.Set("ck","")
	payload :=strings.NewReader(data.Encode())
	req,err :=http.NewRequest("POST",requesturl,payload)
	if err!=nil {
		panic(err)	
	}
	req.Header.Add("Accept","application/json")
	req.Header.Add("Referer","https://accounts.douban.com/passport/login_popup?login_source=anony")
	req.Header.Add("Cookie","login_start_time=1573635953584; __utma=30149280.303356969.1573633004.1573633004.1573635671.2; __utmb=30149280.11.10.1573635671; __utmc=30149280; __utmv=30149280.20668; __utmz=30149280.1573635671.2.2.utmcsr=accounts.douban.com|utmccn=(referral)|utmcmd=referral|utmcct=/passport/login_popup; apiKey=; push_doumail_num=0; push_noty_num=0; __utmt=1; __gads=Test; last_login_way=account; ap_v=0,6.0; user_data={%22area_code%22:%22+86%22%2C%22number%22:%2213476099548%22%2C%22code%22:%226291%22}; vtoken=phone_register%20236cca2a6311406c9c310cddfc8a6c5d; bid=Kie0MQHvsGk;")
	req.Header.Add("User-Agent","Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.1.2 Safari/605.1.15")
	req.Header.Add("Content-Type","application/x-www-form-urlencoded")
	res,err :=http.DefaultClient.Do(req)
	if err!=nil {
		panic(err)
	}
	defer res.Body.Close()
	body,err :=ioutil.ReadAll(res.Body)
	if err !=nil {
		panic(err)
	}
	fmt.Println(string(body))
}
