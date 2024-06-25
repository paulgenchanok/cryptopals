package set2

import (
	"fmt"
	"strings"
)

type fakeCookie struct {
	values map[string]string
}

func parseCookie(parseme string) fakeCookie {
	// Do all of my prep work here

	parsemeSplit := strings.Split(parseme, "&")

	outMap := make(map[string]string, len(parsemeSplit))

	for s := range parsemeSplit {

		equalSplit := strings.Split(parsemeSplit[s], "=")

		outMap[equalSplit[0]] = equalSplit[1]
	}

	ret := fakeCookie{values: outMap}

	return ret

}

func printCookie(cookie fakeCookie) {

	fmt.Printf("{\n")

	for k, v := range cookie.values {
		fmt.Printf("  %s: '%s'\n", k, v)
	}

	fmt.Printf("}\n")
}

func printCookieOneliner(cookie fakeCookie) {

	outs := []string{}

	for k, v := range cookie.values {
		out := k
		out += "="
		out += v
		outs = append(outs, out)
	}

	outString := strings.Join(outs, "&")

	fmt.Printf("%s\n", outString)

}

func profileFor(email string) fakeCookie {
	// Also have it print. Error check email here

	cookieString := "email=" + email + "&uid=10&role=user"

	return parseCookie(cookieString)

}

// func profileFor(email string) string {
// 	// TODO: Error checking for the email string.

// }

func Thirteen() {
	parseme := "foo=bar&baz=qux&zap=zazzle"

	cookie := parseCookie(parseme)
	printCookie(cookie)

	foo := profileFor("foo@bar.com")
	printCookie(foo)
	printCookieOneliner(foo)
}
