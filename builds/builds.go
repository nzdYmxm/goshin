package builds

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	genshinapi "github.com/darylhjd/genshin-api"
)

func TestBuild() {
	characters, err := genshinapi.GetCharacters()
	if err == nil {
		rand.Seed(time.Now().Unix())
		fmt.Println("Your character for test build is", characters[rand.Int()%len(characters)])
	} else {
		// some err
		log.Fatal(err)
	}
}
