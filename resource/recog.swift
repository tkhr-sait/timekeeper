#!/usr/bin/env swift

import Foundation
import AppKit

// NSSpeachRecognizer
class Dispatcher: NSObject, NSSpeechRecognizerDelegate {
  func speechRecognizer(_ sender: NSSpeechRecognizer,didRecognizeCommand command: String) {
    if (command == "がいどさんつぎ") {
      print("s")
    } else if (command == "がいどさんまて") {
      print("3")
    }
		fflush(stdout)
  }
}

var commands: [String] = []
commands.append("がいどさんつぎ")
commands.append("がいどさんまて")

let dispatcher = Dispatcher()
let recognizer = NSSpeechRecognizer()!
recognizer.delegate = dispatcher
recognizer.commands = commands
recognizer.startListening()

let loop = RunLoop.current
let mode = loop.currentMode ?? RunLoop.Mode.default
while ( loop.run(mode: mode,before: Date(timeIntervalSinceNow: 0.5) ) ) {
}
