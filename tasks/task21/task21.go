// Реализовать паттерн «адаптер» на любом примере.

package main

import "fmt"

// Пусть есть интерфейс MediaPlayer с методом Play:

type MediaPlayer interface {
	Play(audioType string, fileName string) error
}

// Нам хочется создать новую версию плеера под названием
// AdvancedMediaPlayer, которая будет воспроизводить
// не только аудио, но и видео:
type AdvancedMediaPlayer interface {
	PlayAudio(filename string) error
	PlayVideo(filename string) error
}

// Допустим, есть устройство VLCPlayer, которое поддерживает
// новый плеер, но к нему нельзя подключить старый:
type VLCPlayer struct{}

func (v *VLCPlayer) PlayVideo(fileName string) error {
	fmt.Println("Воспроизведение видеофайла под названием: ", fileName)
	return nil
}

func (v *VLCPlayer) PlayAudio(fileName string) error {
	fmt.Println("Воспроизведение аудиофайла под названием: ", fileName)
	return nil
}

// Чтобы подключить старый плеер к этому устройству,
// потребуется адаптер:
type MediaAdapter struct {
	advancedMusicPlayer AdvancedMediaPlayer
}

func (m *MediaAdapter) Play(audioType string, filename string) error {
	if audioType == "vlc" {
		return m.advancedMusicPlayer.PlayVideo(filename)
	} else if audioType == "mp4" {
		return m.advancedMusicPlayer.PlayAudio(filename)
	}
	return nil
}

type AudioPlayer struct {
	mediaAdapter MediaPlayer
}

func (a *AudioPlayer) Play(audioType string, fileName string) error {
	// Inbuilt support for mp3 music files
	if audioType == "mp3" {
		fmt.Printf("Playing mp3 file. Name: %s\n", fileName)
		return nil
	} else if audioType == "vlc" || audioType == "mp4" {
		a.mediaAdapter = &MediaAdapter{&VLCPlayer{}}
		return a.mediaAdapter.Play(audioType, fileName)
	}

	return fmt.Errorf("invalid media. %s format not supported", audioType)
}

func main() {

}
