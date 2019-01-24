#!/usr/bin/env swift

import Foundation
import AppKit

// NSSpeachRecognizer
class Dispatcher: NSObject, NSSpeechRecognizerDelegate {
  var stop: Bool
  override init (){ stop = false }
  func speechRecognizer(didRecognizeCommand command: String) {
    if (command == "がいどさんつぎ") {
      print("s")
    } else if (command == "がいどさんまって") {
      print("3")
    }
  }
}

var commands: [String] = []
commands.append("がいどさんつぎ")
commands.append("がいどさんまって")

let dispatcher = Dispatcher()
let recognizer = NSSpeechRecognizer()!
recognizer.delegate = dispatcher
recognizer.commands = commands
recognizer.startListening()

let loop = RunLoop.current
let mode = loop.currentMode ?? RunLoopMode.defaultRunLoopMode
while ( loop.run(mode: mode,before: Date(timeIntervalSinceNow: 0.5) ) && !dispatcher.stop ) {
}
