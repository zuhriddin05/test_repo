// 1) 12 chi uy vazifada ValidateUser funksiya yaratilgan edi.
// 2) Bugungi test_repo repository yaratilgan ichida o'zgarishlar kiritilishi kerak
// 3) ValidateUser dasturri test_repo ni ichiga joylang
// 4) user package yaratib uning ichida Validate digan funksiyani e'lon qiling
// 6) main package da user.Validate qilingan holda chaqirinsin
// 7) dastur ishlashini hosil qiling (go run yoki go build o'rqali)
// 8) barchasi ishlagan taqdirda, hamma o'zgarishlarni git add, git commit va git push commandar o'rqali bajaring
// 9) natijada local codelaringiz github ga joylangan bolishi kerak
// 10) Repository linkini jonating

package main

import (
	"fmt"

	"github.com/zuhriddin05/test_repo/user"
)

func main() {
	for {
		var son int
		us := user.User{}
		fmt.Printf("Ismingizni kiriting: ")
		fmt.Scanln(&us.Name)
		fmt.Printf("Yoshingizni kiriting: ")
		fmt.Scanln(&us.Age)
		fmt.Printf("Emailingizni kiriting: ")
		fmt.Scanln(&us.Email)
		errors := us.ValidateUser()
		if len(errors) > 0 {
			for _, err := range errors {
				fmt.Println("--> ", err)
			}
		} else {
			fmt.Println("-------------------------")
			fmt.Println("FOYDALANUVCHI MA'LUMOTLARI:")
			fmt.Println("-------------------------")
			fmt.Println("Ism:", us.Name)
			fmt.Println("Yosh:", us.Age)
			fmt.Println("Email:", us.Email)
			fmt.Println("-------------------------")

		}

		fmt.Printf(`-------------------------
[1]-START
[2]-EXIT
-------------------------
kiriting: `)
		fmt.Scanln(&son)
		if son == 1 {
			continue
		} else if son == 2 {
			break
		} else {
			fmt.Printf(`-------------------------
Xatolikga yo'l qo'dingiz !
-------------------------`)
		}

	}
}
