package goroutines

/*

Concurrency nedir?

Birden fazla iş ile aynı anda ilgilenebilme becerisidir diyebiliriz. Bir programı çalıştırdığınızda öncelikle belleğe yüklenir ve bir process olarak nitelendirilir.
Processler kendi içlerinde bir ya da birden fazla thread barındırabilirler. Thread, uygulamamızın çalışan bir parçası anlamına gelmektedir.
Örneğin “maaş hesaplama fonksiyonumuz bir thread üzerinde çalışır”, “veri tabanına sorgu gönderdiğimiz fonksiyon bir thread üzerinde çalışır” gibi.
Her bir thread T anında sadece bir iş yapabilme becerisine sahiptir. Modern bilgisayarlarda CPU içerisinde birden fazla çekirdek bulunur.
Process içerisinde bulunan bu Thread’ler bir CPU çekirdeği üzerinde çalışırlar. Her bir core aynı anda sadece bir Thread çalıştırabilir.
Bir çekirdeğe atanmış birden fazla thread olabilir ve aslında bu noktada yukarıda bahsettiğimiz “aynı anda birden fazla iş ile ilgilenebilme” tanımına da varmış oluyoruz.
Elbette çekirdekten daha fazla thread oluşturulabilir. CPU her bir thread arasında geçiş yaparak Thread’lerin çalıştırılmasını sağlamaktadır. Bu işleme context-switch adı verilir.
Bir Thread I/O (input/output - giriş/çıkış) işlemi yaptığında CPU thread’in o anki tüm bilgilerini kaydeder ve o Thread işlemini bitirene kadar bir köşeye alır.
Onun yerine ise çalıştırılmaya hazır olan diğer Thread’i çalıştırmaya başlar.
Köşeye alınan thread işlemini bitirdiğinde ise kuyruğun en sonuna geçerek çalıştırılmaya hazır bir şekilde sırasının gelmesini bekler.
Asenkron programlama ve multi-threading kavramları ise concurrency kavramının formlarındandır.


Goroutine nedir?

Goroutineler Go Runtime tarafından yönetilen hafif bir sistemdir. Bir işlemi eşzamanlı olarak yapmak istiyorsak, Goroutine'den faydalanabiliriz.
Bu sayede aynı çalışma zamanı içerisinde birden fazla iş parçacığı oluşturabiliriz. Goroutine, diğer işlevlerle eşzamanlı olarak çalışabilen bir işlevdir.
Genellikle bir fonksiyon çağrıldığında kontrol çağrılan fonksiyona aktarılır ve yürütme kontrolü tamamlandığında çağıran fonksiyona geri döner.
Çağıran işlev daha sonra yürütülmesine devam eder. Çağıran işlev, çağrılan işlevin geri kalan ifadelere devam etmeden önce yürütmeyi tamamlamasını bekler.
Ancak goroutine durumunda, çağıran fonksiyon, çağrılan fonksiyonun yürütülmesinin tamamlanmasını beklemeyecektir.
Sonraki ifadelerle çalışmaya devam edecektir. Bir programda birden fazla goroutine sahip olabilirsiniz.
Ayrıca ana program, ifadelerini yürütmeyi tamamladıktan sonra çıkacak ve çağrılan goroutinlerin tamamlanmasını beklemeyecektir.

*/

import (
	"fmt"
)

func CiftSayilar() {

	for i := 0; i <= 10; i += 2 {
		fmt.Println("Çift sayı: ", i)
	}
}

func TekSayilar() {

	for i := 1; i <= 10; i += 2 {
		fmt.Println("Tek sayı: ", i)
	}
}