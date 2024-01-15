import kotlin.random.Random

val wordList = listOf(
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
    "word", "work", "wrap", "year", "yoga", "zero", "zone"
)

fun generateRandomWord(): String {
    return wordList.random()
}

fun provideFeedback(secretWord: String, guessedWord: String): String {
    return secretWord.zip(guessedWord).map { (secret, guessed) ->
        if (secret == guessed) secret.toString() else "#"
    }.joinToString("")
}

fun main() {
    println("Welcome to the Word Guessing Game!")

    var playAgain = true

    while (playAgain) {
        playGame()
        playAgain = askToPlayAgain()
    }

    println("Thank you for playing! Goodbye.")
}

fun playGame() {
    val secretWord = generateRandomWord()
    var attempts = 0

    while (true) {
        print("Enter your guess (4-letter word): ")
        val guessedWord = readLine()?.toLowerCase()

        if (isValidGuess(guessedWord, secretWord)) {
            attempts++
            val feedback = provideFeedback(secretWord, guessedWord!!)
            println("Feedback: $feedback")

            if (feedback == secretWord) {
                println("Congratulations! You guessed the correct word '$secretWord' in $attempts attempts.")
                break
            }
        } else {
            println("Invalid input. Please enter a valid 4-letter word from the list.")
        }
    }
}

fun isValidGuess(guessedWord: String?, secretWord: String): Boolean {
    return guessedWord != null && guessedWord.length == 4 && guessedWord.all { it.isLetter() } && guessedWord in wordList
}

fun askToPlayAgain(): Boolean {
    print("Do you want to play again? (yes/no): ")
    val playAgainInput = readLine()?.toLowerCase()
    return playAgainInput == "yes" || playAgainInput == "y"
}
