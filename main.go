package main

import (
	"fmt"

	"github.com/TheCacophonyProject/audiobait/audiofilelibrary"
	"github.com/TheCacophonyProject/audiobait/playlist"
)

func main() {

	audio := audioFileLibrary.AudioFileLibrary{}
	fmt.Println(audio.soundsDirectory)
	sched := playlist.Schedule{}
	fmt.Println(sched.Description)

}
