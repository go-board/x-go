package xunsafe

import (
	"log"
	"testing"
)

var gStr = `
Duis ullamco eu non ad commodo voluptate anim. Consequat adipisicing ipsum nisi officia. Magna aliqua pariatur commodo mollit voluptate nisi proident duis qui laboris non excepteur proident ea. Voluptate sint minim commodo ea voluptate incididunt exercitation ea do anim cillum nisi. Adipisicing aute officia et deserunt aute.
Amet sit est tempor et. Mollit laborum labore aliqua nulla fugiat sit veniam ex commodo dolore Lorem. Qui velit Lorem excepteur aliquip nisi cillum consectetur dolore culpa mollit mollit.
Dolore et ad laborum adipisicing laborum elit enim. Nulla irure enim occaecat est mollit dolor occaecat commodo voluptate ipsum ad anim ad. Nisi nisi culpa proident minim et exercitation.
Aute do dolor ipsum ut est cillum elit cupidatat esse. Incididunt tempor eiusmod non reprehenderit do sint est adipisicing in proident dolor ut qui minim. Eu officia ullamco do ullamco excepteur labore ex do id labore. Esse do officia ea nulla reprehenderit ullamco ut consectetur dolore aliquip. Non sit nostrud ipsum qui minim reprehenderit irure do Lorem labore sunt irure. Ut sunt cillum sint Lorem aliquip officia velit et labore eiusmod eiusmod laborum sit.
Minim consectetur non irure esse nisi veniam ex ea esse nisi deserunt magna laborum. Laborum culpa incididunt tempor incididunt aliqua ad nisi culpa. Officia enim eiusmod dolore ipsum nisi ad velit.
Cupidatat anim aute non velit eu irure ullamco ut voluptate ut fugiat reprehenderit. Do cillum minim exercitation laboris incididunt eu. Deserunt officia magna culpa adipisicing irure aliqua anim irure magna minim cillum et minim. Sint incididunt aliquip proident laboris irure aliqua elit consequat adipisicing laborum in. Elit aliquip cillum pariatur incididunt ut Lorem duis quis Lorem sunt excepteur elit officia aliqua. Officia ea velit cupidatat ad laborum laborum incididunt et est anim labore reprehenderit.
Ad id cupidatat elit amet sit nisi commodo. In mollit est fugiat eu do sit mollit aute culpa. Pariatur cillum quis ipsum magna. Dolor laboris consequat minim fugiat tempor enim proident laboris labore esse in labore tempor sunt. Adipisicing cillum consectetur adipisicing reprehenderit ut. Laborum ex ex consectetur fugiat ipsum.
In tempor dolore nulla ipsum velit velit nostrud magna. Incididunt nulla excepteur dolore occaecat pariatur amet id irure culpa ut pariatur dolor laborum. Nulla deserunt excepteur ipsum sit occaecat occaecat tempor sunt laborum consequat Lorem nostrud. Ea excepteur ut duis amet laboris magna velit pariatur nostrud ullamco minim occaecat.
Enim cupidatat commodo eu consequat sunt qui anim et ipsum fugiat voluptate ex sit qui. Ipsum pariatur id qui exercitation amet eu est ipsum qui aliquip sit labore ullamco enim. Enim adipisicing elit laboris exercitation ullamco adipisicing ad deserunt occaecat ullamco veniam dolor consectetur. Do voluptate elit elit minim mollit sit sint proident minim et nulla ea. Cupidatat ad laboris ut do et et nostrud irure irure cillum aute aute adipisicing. Do aliqua excepteur non sunt do in ad cillum pariatur incididunt voluptate nostrud. Lorem excepteur velit eiusmod amet incididunt occaecat eiusmod nulla tempor laborum est mollit.
Nostrud cillum cupidatat ad deserunt qui ipsum amet reprehenderit aliqua minim. Voluptate excepteur eiusmod adipisicing commodo dolor magna enim nostrud magna ex. Lorem mollit aliquip veniam ex cupidatat do ea occaecat. Cupidatat deserunt anim nisi occaecat Lorem id mollit.
Commodo laboris amet aute cupidatat incididunt sit magna reprehenderit culpa eu do non laboris duis. Velit est proident anim labore esse ut quis laboris commodo occaecat est sit. Aliqua magna dolor dolore velit ullamco laborum aute et id ipsum sunt. Amet mollit non ea magna minim est et incididunt duis ea magna sunt.
Ullamco id laboris excepteur amet cillum est enim duis irure officia irure sunt. Sint nostrud laboris occaecat sunt sint. Aliquip enim veniam incididunt duis. Exercitation quis occaecat consectetur exercitation anim culpa culpa. Laborum in mollit Lorem veniam do qui commodo occaecat nostrud elit ipsum. Occaecat pariatur adipisicing culpa duis incididunt ut irure aute culpa. Dolore aliqua proident eu qui fugiat laborum quis est Lorem laborum anim duis.
Adipisicing quis fugiat sunt consequat sunt aliqua id aliquip ut. Cupidatat magna ullamco aute minim eu adipisicing. Velit officia ipsum do labore exercitation aute esse officia labore nulla. Sunt in eu excepteur consectetur duis elit eu sunt velit nisi.
Sit veniam labore ipsum ipsum do labore pariatur occaecat. Magna deserunt deserunt nulla laborum eiusmod non pariatur duis elit cupidatat consequat commodo mollit ad. Incididunt tempor ullamco eu sint. Ullamco dolor ullamco et culpa fugiat officia culpa incididunt consectetur non anim non aliqua. Velit esse sint quis et qui pariatur deserunt. Et do ex fugiat voluptate ea fugiat incididunt ullamco laborum velit excepteur incididunt labore anim. Eiusmod laboris non enim laboris culpa laborum irure.
Qui est dolore nisi occaecat officia ea occaecat. Ea adipisicing cupidatat commodo cillum qui laboris minim cillum non consequat. Proident quis aliquip aliqua eu ex voluptate labore esse et excepteur elit. Labore non ullamco ex veniam.
Officia minim nisi ut id sunt in ea Lorem. Laborum exercitation in occaecat incididunt sint voluptate excepteur enim velit cupidatat minim commodo. Sit amet aliqua tempor excepteur quis veniam magna elit enim.
Laborum magna elit Lorem ut non sit nulla nulla sunt fugiat aute voluptate. Ullamco magna esse velit aliqua nulla et proident aute nisi eu. Aliquip ea aute proident ut ullamco elit fugiat officia aute commodo nostrud nisi laboris aliquip. Et est sint ut esse tempor do occaecat mollit Lorem minim deserunt aliqua.
Velit sint culpa pariatur nisi duis nulla eiusmod elit. Elit voluptate nostrud irure cupidatat reprehenderit. Ut proident nisi consequat veniam magna qui.
Aliqua id pariatur culpa sunt enim exercitation dolore anim cillum ullamco ad. Enim excepteur proident proident consequat laboris ea duis mollit dolor. Excepteur consequat ea esse fugiat irure consequat dolor mollit. Consectetur nostrud quis ex ullamco proident id nostrud commodo consequat Lorem in aute.
Aliqua eiusmod id exercitation cillum velit labore. Magna excepteur deserunt sunt proident nostrud. Quis amet tempor dolore eu incididunt consectetur officia.
`

var gBytes []byte

func init() {
	gBytes = []byte(gStr)
	log.Printf("testcase length is %d", len(gStr))
}

func TestUnsafeStringToBytes(t *testing.T) {
	bytes := StringToBytes("Hello,world")
	t.Logf("%s\n", bytes)
}

func TestUnsafeBytesToString(t *testing.T) {
	bytes := StringToBytes("Hello,world")
	str := BytesToString(bytes)
	t.Logf("%+v, %s\n", bytes, str)
}

func BenchmarkStringToBytes(b *testing.B) {
	// BenchmarkStringToBytes/safe-12	1000000	1232 ns/op	8192 B/op	1 allocs/op
	b.Run("safe", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = []byte(gStr)
		}
	})
	//BenchmarkStringToBytes/unsafe-12	2000000000	0.35 ns/op	0 B/op	0 allocs/op
	b.Run("unsafe", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			StringToBytes(gStr)
		}
	})
}

func BenchmarkBytesToString(b *testing.B) {
	//BenchmarkBytesToString/safe-12	1000000	1115 ns/op	8192 B/op	1 allocs/op
	b.Run("safe", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = string(gBytes)
		}
	})
	//BenchmarkBytesToString/unsafe-12	2000000000	0.35 ns/op	0 B/op	0 allocs/op
	b.Run("unsafe", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			BytesToString(gBytes)
		}
	})
}
