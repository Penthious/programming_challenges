// singly linked list can only go forward, cannot go backwards
package main

import "fmt"

type song struct {
	name     string
	artist   string
	previous *song // this makes it a doubly linked list
	next     *song
}

type playlist struct {
	name       string
	tail       *song // this makes it a doubly linked list
	head       *song // a pointer to the head of our linked list
	nowPlaying *song // a pointer to the current playing song
}

func createPlaylist(name string) *playlist {
	return &playlist{
		name: name,
	}
}

func (p *playlist) addSong(name, artist string) {
	s := &song{
		name:   name,
		artist: artist,
	}

	if p.head == nil {
		p.head = s
	} else {
		currentNode := p.tail
		currentNode.next = s
		s.previous = currentNode
	}
	p.tail = s
}

func (p *playlist) showAllSongs() {
	currentNode := p.head
	if currentNode == nil {
		fmt.Println("Playlist is empty.")
	}
	fmt.Printf("%+v\n", *currentNode)
	for currentNode.next != nil {
		currentNode = currentNode.next
		fmt.Printf("%+v\n", *currentNode)
	}
}

func (p *playlist) startPlaying() *song {
	p.nowPlaying = p.head
	return p.nowPlaying
}

func (p *playlist) nextSong() *song {
	if p.nowPlaying.next == nil { // circular linked list
		p.nowPlaying = p.head
	} else {
		p.nowPlaying = p.nowPlaying.next
	}
	return p.nowPlaying
}

func (p *playlist) previousSong() *song {
	if p.nowPlaying.previous == nil {
		p.nowPlaying = p.tail
	} else {
		p.nowPlaying = p.nowPlaying.previous
	}

	return p.nowPlaying
}

func main() {
	playlist := "new Playlist"
	myPlaylist := createPlaylist(playlist)

	fmt.Println("Created Playlist: ", myPlaylist)
	fmt.Println()

	fmt.Println("Adding songs to the playlist")
	myPlaylist.addSong("Ophelia", "The Lumineers")
	myPlaylist.addSong("Shape of you", "Ed Sheeran")
	myPlaylist.addSong("Stubborn Love", "The Lumineers")
	myPlaylist.addSong("Feels", "Calvin Harris")

	fmt.Println("Showing all songs in playlist...")
	myPlaylist.showAllSongs()
	fmt.Println()

	myPlaylist.startPlaying()
	fmt.Printf("Now playing: %s by %s\n", myPlaylist.nowPlaying.name, myPlaylist.nowPlaying.artist)
	fmt.Println()

	fmt.Println("Changing to next song")
	myPlaylist.nextSong()
	fmt.Printf("Now playing: %s by %s\n", myPlaylist.nowPlaying.name, myPlaylist.nowPlaying.artist)
	fmt.Println()

	myPlaylist.previousSong()
	fmt.Println("Changing previous song...")
	fmt.Printf("Now playing: %s by %s\n", myPlaylist.nowPlaying.name, myPlaylist.nowPlaying.artist)
	fmt.Println()

}
