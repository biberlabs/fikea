package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

var vowels = []rune("AEIOUÄÉÖÜÅ")
var consonants = []rune("KNBDFGLPCRSTYMVZHKNJMLTGS")
var sylrange = []int{1, 2, 3, 4}
var images = []string{
	"http://images4.emlakjet.com/emlak-haber/2011/08/16/foto_galeri/l/ikea-tv-unitesi-fiyatlari-1313492161.jpg",
	"http://www.ikea.com/PIAimages/19488_PE104775_S5.JPG",
	"http://www.canadianliving.com/blogs/home/files/2012/10/kaustby1.jpg",
	"http://g02.a.alicdn.com/kf/HTB18HPhHVXXXXb5XXXXq6xXFXXX4/Kennidiming-chair-presidential-chair-designer-fashion-wood-dining-chair-armrest-chair-Ikea-computer-desk.jpg",
	"http://www.dinaters.info/wp-content/uploads//ikea-swivel-chair-i9nro40t.jpg",
	"http://www.ikea.com/ca/en/images/products/ekne-mirror__0119704_PE276136_S4.JPG",
	"http://www.polyvore.com/cgi/img-thing?.out=jpg&size=l&tid=885889",
	"http://vanelibg.com/wp-content/uploads/2015/08/Standing-Mirror-Ikea-as-standing-mirror-jewelry-case-as-additional-suggestion-for-make-a-perfect-Interior-design-16815-19.jpg",
	"http://picture-cdn.wheretoget.it/lv19jc-i.jpg",
	"http://www.ikea.com/PIAimages/0099006_PE240451_S5.JPG",
	"https://confettistyle.files.wordpress.com/2011/05/ikea-expedit-white.jpg",
	"http://www.westernlivingmagazine.com/wp-content/uploads/2015/08/Terrariums_Ikea_SelfWatering.jpg",
	"https://s-media-cache-ak0.pinimg.com/736x/9b/d0/89/9bd089d9679c058692455b6d45844094.jpg",
	"http://www.ikea.com/au/en/images/products/kvarnvik-box-set-of-grey__0189737_PE343744_S4.JPG",
	"https://lh6.googleusercontent.com/-p_ncSZTyYcM/TXa3_dZ28tI/AAAAAAAABKo/mG5BFmQujDo/s1600/greno-pad-assorted-colors-checkered__0116372_PE270938_S4.JPG",
	"http://i.huffpost.com/gen/1885156/images/o-IKEA-FIRST-APARTMENT-facebook.jpg",
	"http://dekorasyontarz.com/wp-content/uploads/2015/05/Ikea-2015-A%C3%A7%C4%B1k-Dolap-ve-Raf-Modelleri-G%C3%B6rselleri.jpg",
	"https://gd2.alicdn.com/imgextra/i2/153010613/T2MrTTXjNXXXXXXXXX_!!153010613.jpg",
}

// Roll a dice
func dice(min int, max int) int {
	return min + rand.Intn(max-min)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// pick a random rune from given list
func pickRune(list []rune) rune {
	i := dice(0, len(list))
	return list[i]
}

// return boolean true if haystack contains needle
func contains(haystack []rune, needle rune) bool {
	for _, a := range haystack {
		if a == needle {
			return true
		}
	}
	return false
}

func syllable(last rune, total int) []rune {

	var s int

	switch {
		case total == 4:
			s = dice(1, 2)
		case total == 3:
			s = dice(2, 3)
		default:
			s = dice(1, 3)
	}

	var res []rune // slice of runes
	var start rune

	// Decide starting with vowel or consonant
	switch {
	case s%2 == 0 || contains(vowels, start) == true:
		res = append(res, pickRune(consonants))
	default:
		res = append(res, pickRune(vowels))
	}

	for i := 1; i <= s; i++ {
		switch {
		case contains(vowels, res[i-1]) == true:
			res = append(res, pickRune(consonants))
		default:
			res = append(res, pickRune(vowels))
		}
	}

	if total < 2 && contains([]rune("KNMLTGS"), res[s]) == true {
		res = append(res, res[s])
	}

	return res
}

func generateNames(w http.ResponseWriter, r *http.Request) {

	syllen := dice(1, 3)

	var name string
	var last rune

	for i := 1; i <= syllen; i++ {
		h := syllable(last, syllen)
		last = h[len(h)-1]
		name += string(h)
	}

	im := images[dice(0, len(images))]

	io.WriteString(w, "<html><body><h1>"+string(name)+"</h1><img src='"+im+"' style='max-width:500px; max-height:500px;'/></body></html>")

}

func main() {
	fmt.Println("Listening..", "Open your browser and go to http://localhost:8000")
	http.HandleFunc("/", generateNames)
	http.ListenAndServe(":8000", nil)
}