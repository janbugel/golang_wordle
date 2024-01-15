package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var wordList = []string{
	"able", "acid", "also", "aunt", "back", "bath", "bean", "bell", "bend", "bind",
	"bird", "bite", "blue", "boat", "bold", "bone", "book", "boot", "bump", "burn",
	"cake", "call", "care", "cart", "case", "cast", "cell", "chat", "chip", "city",
	"clip", "coal", "coat", "code", "coin", "cold", "come", "cool", "crop", "cure",
	"dart", "date", "dawn", "dear", "deck", "deep", "deer", "desk", "dice", "diet",
	"dine", "dirt", "dish", "disk", "doll", "door", "dose", "draw", "drop", "drum",
	"duck", "dust", "earn", "east", "edge", "else", "even", "ever", "exit", "face",
	"fade", "fail", "fair", "fake", "fall", "farm", "fast", "fear", "feel", "file",
	"fill", "find", "fine", "fire", "fish", "flap", "flip", "flow", "foam", "fold",
	"food", "fool", "foot", "fork", "form", "free", "fuel", "gain", "game", "gate",
	"gaze", "gear", "gift", "glow", "gold", "good", "gray", "grip", "grow", "hair",
	"hand", "hard", "harm", "have", "heal", "hear", "heat", "help", "here", "hide",
	"high", "hill", "hold", "hope", "huge", "hunt", "hurt", "idea", "inch", "into",
	"iron", "jazz", "join", "jump", "just", "keep", "kick", "kind", "kiss", "knee",
	"knit", "lady", "lamp", "land", "late", "lead", "leaf", "lean", "left", "life",
	"lift", "like", "line", "live", "load", "lock", "long", "look", "loop", "loud",
	"love", "luck", "mail", "make", "mark", "mask", "melt", "mend", "mild", "mind",
	"mist", "moon", "more", "moss", "move", "near", "nest", "news", "node",
	"none", "note", "once", "open", "over", "pace", "page", "pair", "pale", "park",
	"part", "pass", "path", "pick", "pink", "play", "plot", "plus", "post", "pour",
	"pull", "pure", "push", "quit", "race", "rain", "read", "real", "ride", "ring",
	"rise", "road", "roof", "room", "root", "rose", "rule", "safe", "said", "salt",
	"sand", "save", "seat", "seed", "seek", "seem", "self", "send", "ship", "shoe",
	"shop", "shot", "show", "shut", "silk", "sing", "sink", "size", "skip", "slip",
	"slow", "snap", "snow", "soap", "soft", "soil", "sole", "song", "sort", "soup",
	"spin", "spot", "star", "stay", "step", "stop", "sure", "swim", "take", "talk",
	"tall", "tape", "tear", "tell", "tide", "tiny", "told", "tool", "trap", "tree",
	"trim", "true", "turn", "twig", "type", "unit", "urge", "view", "vivid",
	"walk", "warm", "wave", "wear", "weep", "whip", "wild", "wind", "wish", "wood",
	"word", "work", "wrap", "year", "yoga", "zero", "zone",
}

func generateRandomWord() string {
	rand.Seed(time.Now().UnixNano())
	return wordList[rand.Intn(len(wordList))]
}

func provideFeedback(secretWord, guessedWord string) string {
	var feedback strings.Builder

	for i := range secretWord {
		if secretWord[i] == guessedWord[i] {
			feedback.WriteByte(secretWord[i])
		} else {
			feedback.WriteByte('#')
		}
	}

	return feedback.String()
}

func playGame() {
	secretWord := generateRandomWord()
	attempts := 0

	for {
		fmt.Print("Enter your guess (4-letter word): ")
		var guessedWord string
		fmt.Scanln(&guessedWord)

		if isValidGuess(guessedWord, secretWord) {
			attempts++
			feedback := provideFeedback(secretWord, guessedWord)
			fmt.Printf("Feedback: %s\n", feedback)

			if feedback == secretWord {
				fmt.Printf("Congratulations! You guessed the correct word '%s' in %d attempts.\n", secretWord, attempts)
				break
			}
		} else {
			fmt.Println("Invalid input. Please enter a valid 4-letter word from the list.")
		}
	}
}

func isValidGuess(guessedWord, secretWord string) bool {
	return len(guessedWord) == 4 && isAllLetters(guessedWord) && containsWord(guessedWord, wordList)
}

func isAllLetters(s string) bool {
	for _, char := range s {
		if !isLetter(char) {
			return false
		}
	}
	return true
}

func isLetter(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}

func containsWord(word string, list []string) bool {
	for _, w := range list {
		if w == word {
			return true
		}
	}
	return false
}

func askToPlayAgain() bool {
	var playAgain string
	fmt.Print("Do you want to play again? (yes/no): ")
	fmt.Scanln(&playAgain)
	return playAgain == "yes" || playAgain == "y"
}

func main() {
	fmt.Println("Welcome to the Word Guessing Game!")

	var playAgain bool

	for playAgain = true; playAgain; playAgain = askToPlayAgain() {
		playGame()
	}

	fmt.Println("Thank you for playing! Goodbye.")
}
