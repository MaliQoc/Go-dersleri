package restful

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Todo struct {
	UserId    int    `json: "userId"`
	Id        int    `json: "id"`
	Title     string `json: "title"`
	Completed bool   `json: "completed"`
}

func Demo1() {

	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)

	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	var todo Todo
	json.Unmarshal(bodyBytes, &todo)
	fmt.Println(todo)
}

/*

Bu kod, basit bir RESTful API istemcisi örneğidir.
Demo1 fonksiyonu, belirli bir URL'den (bu durumda "https://jsonplaceholder.typicode.com/todos/1") bir HTTP GET isteği yapar, yanıtı alır ve bu yanıtı işler.

Kodun işleyişi şu şekildedir:

HTTP GET İsteği:

http.Get fonksiyonu, belirli bir URL'ye GET isteği yapar ve sunucunun döndürdüğü HTTP yanıtını response değişkeninde saklar.
Eğer bir hata oluşursa, hata mesajı konsola yazdırılır.

Yanıtın Okunması:

response.Body.Close() satırı, HTTP yanıtının gövdesini kapatır. Bu, bellekteki kaynakların serbest bırakılması için gereklidir.
ioutil.ReadAll(response.Body) fonksiyonu, yanıtın gövdesini okuyarak bunu bir byte dizisi ([]byte) olarak döndürür.
Bu byte dizisi önce string'e çevrilir (bodyString) ve bu string konsola yazdırılır.

JSON Verisinin Serileştirilmesi:

Todo adında bir struct tanımlanmıştır ve bu struct, API'den alınan JSON verisini tutmak için kullanılır.
json.Unmarshal fonksiyonu, JSON verisini struct'a dönüştürür. bodyBytes byte dizisi olarak JSON verisini tutar ve bu veri, todo adlı Todo struct'ına dönüştürülür.
Son olarak, struct'ın içeriği (todo) konsola yazdırılır. */

func Demo2() {

	todo := Todo{1, 2, "Alışveriş yapılacak", false}
	jsonTodo, err := json.Marshal(todo)
	response, err := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json;charset=utf-8", bytes.NewBuffer(jsonTodo))

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)

	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	var todoReponse Todo
	json.Unmarshal(bodyBytes, &todo)
	fmt.Println(todoReponse)
}

/* Bizim burada görevimiz aslen backend kısmına yazıp fronttan gelen bilgiyi orayla eşleştirmektir.
Yani Whatsapp'ta yazılan bir mesaj nasıl gönderiliyorsa post işlemi odur. Ekrana gönderilen bilgi backend kısmında kodda işlenir ve karşı tarafa gönderilir. Mantık budur. */

/*

Bu kod, bir RESTful API'ye JSON formatında bir veri POST eden basit bir Go örneğidir.
Demo2 fonksiyonu, bir Todo nesnesi oluşturur, bu nesneyi JSON formatına çevirir ve bir HTTP POST isteği yaparak sunucuya gönderir.

İşleyişi şu şekildedir:

Todo Nesnesinin Oluşturulması:

Todo struct'ı kullanılarak yeni bir todo nesnesi oluşturulur.
Bu nesne, kullanıcı ID'si (UserId: 1), görev ID'si (Id: 2), başlık (Title: "Alışveriş yapılacak") ve tamamlanma durumu (Completed: false) gibi özellikleri içerir.

JSON'a Dönüştürme:

json.Marshal fonksiyonu, todo nesnesini JSON formatına dönüştürür ve bunu jsonTodo adlı bir byte dizisinde saklar.
Bu işlem sırasında bir hata oluşursa, err değişkeni aracılığıyla kontrol edilir.

HTTP POST İsteği:

http.Post fonksiyonu, JSON formatındaki veriyi https://jsonplaceholder.typicode.com/todos adresine bir POST isteği olarak gönderir.
Bu işlem sırasında içeriğin türü olarak "application/json;charset=utf-8" belirtilir ve jsonTodo byte dizisi bu isteğin gövdesi olarak gönderilir.
Eğer bir hata oluşursa, hata mesajı konsola yazdırılır.

Yanıtın Okunması:

response.Body.Close() satırı, HTTP yanıtının gövdesini kapatır.
ioutil.ReadAll(response.Body) fonksiyonu, sunucudan gelen yanıtın gövdesini okur ve bu yanıtı bir byte dizisi ([]byte) olarak döndürür.
Bu byte dizisi bir string'e dönüştürülür (bodyString) ve bu string konsola yazdırılır.

JSON Yanıtının Serileştirilmesi:

json.Unmarshal fonksiyonu, sunucudan dönen JSON verisini todoResponse adlı bir Todo struct'ına dönüştürür.
Son olarak, todoResponse nesnesi konsola yazdırılır. */