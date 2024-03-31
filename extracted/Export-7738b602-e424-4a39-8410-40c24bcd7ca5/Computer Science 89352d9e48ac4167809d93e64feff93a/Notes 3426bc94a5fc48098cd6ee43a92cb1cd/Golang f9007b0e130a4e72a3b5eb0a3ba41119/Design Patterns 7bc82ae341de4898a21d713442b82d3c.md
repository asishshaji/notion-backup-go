# Design Patterns

[Desing Patters in Golang](https://medium.com/@josueparra2892/desing-patters-in-golang-24a142d2cc91)

## Adapter Pattern

It’s a **structural design pattern** that enables collaboration between objects with incompatible interfaces**.**

You have to charge your phone, but your charger doesn't fit the outlets. ***You will need an adapter that bridges the gap between your charger and the outlet.***

In software development, we have components that don't fit together directly. For these situations, we can use an Adapter to resolve these compatibility issues.

[Adapter Design Pattern in Go (GoLang) - Welcome To Golang By Example](https://golangbyexample.com/adapter-design-pattern-go/)

Imagine a car and a boat, both are different objects If we implement the Adapter pattern we can achieve that the boat goes to the road and the car goes to the water.

[Adapter Pattern in Go](https://medium.com/@josueparra2892/adapter-pattern-in-go-d77e08abd526)

![Untitled](Design%20Patterns%207bc82ae341de4898a21d713442b82d3c/Untitled.png)

![Untitled](Design%20Patterns%207bc82ae341de4898a21d713442b82d3c/Untitled%201.png)

*Now, we need to adapt the car to go ahead in the water to do that we need the adapter struct that defines a method “navigate To Destination” that going to adapt the car to do the needed actions.*

### ChatGPT

The Adapter Design Pattern is a structural design pattern that allows two incompatible interfaces to work together. It acts as a bridge between two interfaces that are not compatible but need to collaborate.

Here's a breakdown of the main components and concepts of the Adapter Design Pattern:

1. Target Interface: This is the interface that the client code expects to work with. It defines the methods and properties that the client code relies on.
2. Adaptee: The Adaptee is the class or interface that is incompatible with the client code. It's the interface that needs to be adapted to work with the client.
3. Adapter: The Adapter is the class that bridges the gap between the Target Interface and the Adaptee. It implements the Target Interface and delegates calls to the Adaptee, effectively translating the calls from the client code into a format that the Adaptee can understand.
4. Client: The Client is the code that uses the Target Interface. It is unaware of the Adapter's presence and interacts directly with the Target Interface.

Here's a simplified example in Python to illustrate the Adapter Design Pattern:

```python
# Target Interface
class MediaPlayer:
    def play(self, audio_type, file_name):
        pass

# Adaptee
class AdvancedMediaPlayer:
    def play_vlc(self, file_name):
        pass

    def play_mp4(self, file_name):
        pass

# Adapter
class MediaAdapter(MediaPlayer):
    def __init__(self, audio_type):
        self.advanced_player = AdvancedMediaPlayer()
        self.audio_type = audio_type

    def play(self, audio_type, file_name):
        if audio_type == "vlc":
            self.advanced_player.play_vlc(file_name)
        elif audio_type == "mp4":
            self.advanced_player.play_mp4(file_name)

# Client
class AudioPlayer(MediaPlayer):
    def play(self, audio_type, file_name):
        if audio_type == "mp3":
            print(f"Playing MP3 file: {file_name}")
        elif audio_type == "vlc" or audio_type == "mp4":
            media_adapter = MediaAdapter(audio_type)
            media_adapter.play(audio_type, file_name)
        else:
            print(f"Unsupported audio type: {audio_type}")

# Usage
audio_player = AudioPlayer()
audio_player.play("mp3", "song.mp3")
audio_player.play("vlc", "movie.vlc")
audio_player.play("mp4", "video.mp4")

```

In this example, `MediaPlayer` is the Target Interface, `AdvancedMediaPlayer` is the Adaptee, and `MediaAdapter` is the Adapter. The `AudioPlayer` is the Client code that interacts with the Target Interface. The `MediaAdapter` adapts the `AdvancedMediaPlayer` to work with the `MediaPlayer` interface.

The Adapter Design Pattern is particularly useful when integrating legacy code or third-party libraries into a new system or when you want to decouple components with incompatible interfaces.

![Untitled](Design%20Patterns%207bc82ae341de4898a21d713442b82d3c/Untitled%202.png)

![Untitled](Design%20Patterns%207bc82ae341de4898a21d713442b82d3c/Untitled%203.png)