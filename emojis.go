package epick

// emoji type used for individual emojis
type emoji struct {
	EMOJI          string
	SKINSUPPORT    bool
	NAME           string
	UNICODEVERSION string
	EMOJIVERSION   string
}

// emojiGroup type used for specific emoji groups
type emojiGroup struct {
	NAME   string
	EMOJIS []emoji
}

// rawEmojiData type used as a list of emoji groups
type rawEmojiData struct {
	LIST []emojiGroup
}

var emojiJson string = `
{
"list": [
	{
	"name": "Smileys and Emotion",
	"emojis": [
    {
      "emoji": "😀",
      "skin tone support": false,
      "name": "grinning face",
      "unicode version": "6.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "😃",
      "skin tone support": false,
      "name": "grinning face with big eyes",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "😄",
      "skin tone support": false,
      "name": "grinning face with smiling eyes",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "😁",
      "skin tone support": false,
      "name": "beaming face with smiling eyes",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "😆",
      "skin tone support": false,
      "name": "grinning squinting face",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "😅",
      "skin tone support": false,
      "name": "grinning face with sweat",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🤣",
      "skin tone support": false,
      "name": "rolling on the floor laughing",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "😂",
      "skin tone support": false,
      "name": "face with tears of joy",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🙂",
      "skin tone support": false,
      "name": "slightly smiling face",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🙃",
      "skin tone support": false,
      "name": "upside down face",
      "unicode version": "8.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "😉",
      "skin tone support": false,
      "name": "winking face",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "😊",
      "skin tone support": false,
      "name": "smiling face with smiling eyes",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "😇",
      "skin tone support": false,
      "name": "smiling face with halo",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🥰",
      "skin tone support": false,
      "name": "smiling face with hearts",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "😍",
      "skin tone support": false,
      "name": "smiling face with heart eyes",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🤩",
      "skin tone support": false,
      "name": "star struck",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "😘",
      "skin tone support": false,
      "name": "face blowing a kiss",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "😗",
      "skin tone support": false,
      "name": "kissing face",
      "unicode version": "6.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "☺️",
      "skin tone support": false,
      "name": "smiling face",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "😚",
      "skin tone support": false,
      "name": "kissing face with closed eyes",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "😙",
      "skin tone support": false,
      "name": "kissing face with smiling eyes",
      "unicode version": "6.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "😋",
      "skin tone support": false,
      "name": "face savoring food",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "😛",
      "skin tone support": false,
      "name": "face with tongue",
      "unicode version": "6.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "😜",
      "skin tone support": false,
      "name": "winking face with tongue",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🤪",
      "skin tone support": false,
      "name": "zany face",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "😝",
      "skin tone support": false,
      "name": "squinting face with tongue",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🤑",
      "skin tone support": false,
      "name": "money mouth face",
      "unicode version": "8.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🤗",
      "skin tone support": false,
      "name": "hugging face",
      "unicode version": "8.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🤭",
      "skin tone support": false,
      "name": "face with hand over mouth",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🤫",
      "skin tone support": false,
      "name": "shushing face",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🤔",
      "skin tone support": false,
      "name": "thinking face",
      "unicode version": "8.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🤐",
      "skin tone support": false,
      "name": "zipper mouth face",
      "unicode version": "8.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🤨",
      "skin tone support": false,
      "name": "face with raised eyebrow",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "😐",
      "skin tone support": false,
      "name": "neutral face",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "😑",
      "skin tone support": false,
      "name": "expressionless face",
      "unicode version": "6.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "😶",
      "skin tone support": false,
      "name": "face without mouth",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "😏",
      "skin tone support": false,
      "name": "smirking face",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "😒",
      "skin tone support": false,
      "name": "unamused face",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🙄",
      "skin tone support": false,
      "name": "face with rolling eyes",
      "unicode version": "8.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "😬",
      "skin tone support": false,
      "name": "grimacing face",
      "unicode version": "6.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "🤥",
      "skin tone support": false,
      "name": "lying face",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "😌",
      "skin tone support": false,
      "name": "relieved face",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "😔",
      "skin tone support": false,
      "name": "pensive face",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "😪",
      "skin tone support": false,
      "name": "sleepy face",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🤤",
      "skin tone support": false,
      "name": "drooling face",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "😴",
      "skin tone support": false,
      "name": "sleeping face",
      "unicode version": "6.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "😷",
      "skin tone support": false,
      "name": "face with medical mask",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🤒",
      "skin tone support": false,
      "name": "face with thermometer",
      "unicode version": "8.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🤕",
      "skin tone support": false,
      "name": "face with head bandage",
      "unicode version": "8.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🤢",
      "skin tone support": false,
      "name": "nauseated face",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🤮",
      "skin tone support": false,
      "name": "face vomiting",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🤧",
      "skin tone support": false,
      "name": "sneezing face",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🥵",
      "skin tone support": false,
      "name": "hot face",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "🥶",
      "skin tone support": false,
      "name": "cold face",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "🥴",
      "skin tone support": false,
      "name": "woozy face",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "😵",
      "skin tone support": false,
      "name": "dizzy face",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🤯",
      "skin tone support": false,
      "name": "exploding head",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🤠",
      "skin tone support": false,
      "name": "cowboy hat face",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🥳",
      "skin tone support": false,
      "name": "partying face",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "😎",
      "skin tone support": false,
      "name": "smiling face with sunglasses",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🤓",
      "skin tone support": false,
      "name": "nerd face",
      "unicode version": "8.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🧐",
      "skin tone support": false,
      "name": "face with monocle",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "😕",
      "skin tone support": false,
      "name": "confused face",
      "unicode version": "6.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "😟",
      "skin tone support": false,
      "name": "worried face",
      "unicode version": "6.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "🙁",
      "skin tone support": false,
      "name": "slightly frowning face",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "☹️",
      "skin tone support": false,
      "name": "frowning face",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "😮",
      "skin tone support": false,
      "name": "face with open mouth",
      "unicode version": "6.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "😯",
      "skin tone support": false,
      "name": "hushed face",
      "unicode version": "6.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "😲",
      "skin tone support": false,
      "name": "astonished face",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "😳",
      "skin tone support": false,
      "name": "flushed face",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🥺",
      "skin tone support": false,
      "name": "pleading face",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "😦",
      "skin tone support": false,
      "name": "frowning face with open mouth",
      "unicode version": "6.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "😧",
      "skin tone support": false,
      "name": "anguished face",
      "unicode version": "6.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "😨",
      "skin tone support": false,
      "name": "fearful face",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "😰",
      "skin tone support": false,
      "name": "anxious face with sweat",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "😥",
      "skin tone support": false,
      "name": "sad but relieved face",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "😢",
      "skin tone support": false,
      "name": "crying face",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "😭",
      "skin tone support": false,
      "name": "loudly crying face",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "😱",
      "skin tone support": false,
      "name": "face screaming in fear",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "😖",
      "skin tone support": false,
      "name": "confounded face",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "😣",
      "skin tone support": false,
      "name": "persevering face",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "😞",
      "skin tone support": false,
      "name": "disappointed face",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "😓",
      "skin tone support": false,
      "name": "downcast face with sweat",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "😩",
      "skin tone support": false,
      "name": "weary face",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "😫",
      "skin tone support": false,
      "name": "tired face",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🥱",
      "skin tone support": false,
      "name": "yawning face",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "😤",
      "skin tone support": false,
      "name": "face with steam from nose",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "😡",
      "skin tone support": false,
      "name": "pouting face",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "😠",
      "skin tone support": false,
      "name": "angry face",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🤬",
      "skin tone support": false,
      "name": "face with symbols on mouth",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "😈",
      "skin tone support": false,
      "name": "smiling face with horns",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👿",
      "skin tone support": false,
      "name": "angry face with horns",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "💀",
      "skin tone support": false,
      "name": "skull",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "☠️",
      "skin tone support": false,
      "name": "skull and crossbones",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "💩",
      "skin tone support": false,
      "name": "pile of poo",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🤡",
      "skin tone support": false,
      "name": "clown face",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "👹",
      "skin tone support": false,
      "name": "ogre",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👺",
      "skin tone support": false,
      "name": "goblin",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👻",
      "skin tone support": false,
      "name": "ghost",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👽",
      "skin tone support": false,
      "name": "alien",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👾",
      "skin tone support": false,
      "name": "alien monster",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🤖",
      "skin tone support": false,
      "name": "robot",
      "unicode version": "8.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "😺",
      "skin tone support": false,
      "name": "grinning cat",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "😸",
      "skin tone support": false,
      "name": "grinning cat with smiling eyes",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "😹",
      "skin tone support": false,
      "name": "cat with tears of joy",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "😻",
      "skin tone support": false,
      "name": "smiling cat with heart eyes",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "😼",
      "skin tone support": false,
      "name": "cat with wry smile",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "😽",
      "skin tone support": false,
      "name": "kissing cat",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🙀",
      "skin tone support": false,
      "name": "weary cat",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "😿",
      "skin tone support": false,
      "name": "crying cat",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "😾",
      "skin tone support": false,
      "name": "pouting cat",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🙈",
      "skin tone support": false,
      "name": "see no evil monkey",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🙉",
      "skin tone support": false,
      "name": "hear no evil monkey",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🙊",
      "skin tone support": false,
      "name": "speak no evil monkey",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "💋",
      "skin tone support": false,
      "name": "kiss mark",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "💌",
      "skin tone support": false,
      "name": "love letter",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "💘",
      "skin tone support": false,
      "name": "heart with arrow",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "💝",
      "skin tone support": false,
      "name": "heart with ribbon",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "💖",
      "skin tone support": false,
      "name": "sparkling heart",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "💗",
      "skin tone support": false,
      "name": "growing heart",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "💓",
      "skin tone support": false,
      "name": "beating heart",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "💞",
      "skin tone support": false,
      "name": "revolving hearts",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "💕",
      "skin tone support": false,
      "name": "two hearts",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "💟",
      "skin tone support": false,
      "name": "heart decoration",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "❣️",
      "skin tone support": false,
      "name": "heart exclamation",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "💔",
      "skin tone support": false,
      "name": "broken heart",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "❤️",
      "skin tone support": false,
      "name": "red heart",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "🧡",
      "skin tone support": false,
      "name": "orange heart",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "💛",
      "skin tone support": false,
      "name": "yellow heart",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "💚",
      "skin tone support": false,
      "name": "green heart",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "💙",
      "skin tone support": false,
      "name": "blue heart",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "💜",
      "skin tone support": false,
      "name": "purple heart",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🤎",
      "skin tone support": false,
      "name": "brown heart",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "🖤",
      "skin tone support": false,
      "name": "black heart",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🤍",
      "skin tone support": false,
      "name": "white heart",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "💯",
      "skin tone support": false,
      "name": "hundred points",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "💢",
      "skin tone support": false,
      "name": "anger symbol",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "💥",
      "skin tone support": false,
      "name": "collision",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "💫",
      "skin tone support": false,
      "name": "dizzy",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "💦",
      "skin tone support": false,
      "name": "sweat droplets",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "💨",
      "skin tone support": false,
      "name": "dashing away",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🕳️",
      "skin tone support": false,
      "name": "hole",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "💣",
      "skin tone support": false,
      "name": "bomb",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "💬",
      "skin tone support": false,
      "name": "speech balloon",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👁️‍🗨️",
      "skin tone support": false,
      "name": "eye in speech bubble",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🗨️",
      "skin tone support": false,
      "name": "left speech bubble",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🗯️",
      "skin tone support": false,
      "name": "right anger bubble",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "💭",
      "skin tone support": false,
      "name": "thought balloon",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "💤",
      "skin tone support": false,
      "name": "zzz",
      "unicode version": "6.0",
      "emoji version": "2.0"
    }
  ]
	},
	{
	"name": "People and Body",
	"emojis": [
    {
      "emoji": "👋",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "waving hand",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🤚",
      "skin tone support": true,
      "skin tone support unicode version": "9.0",
      "name": "raised back of hand",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🖐️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "hand with fingers splayed",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "✋",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "raised hand",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🖖",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "vulcan salute",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👌",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "ok hand",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🤏",
      "skin tone support": true,
      "skin tone support unicode version": "12.0",
      "name": "pinching hand",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "✌️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "victory hand",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "🤞",
      "skin tone support": true,
      "skin tone support unicode version": "9.0",
      "name": "crossed fingers",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🤟",
      "skin tone support": true,
      "skin tone support unicode version": "10.0",
      "name": "love you gesture",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🤘",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "sign of the horns",
      "unicode version": "8.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🤙",
      "skin tone support": true,
      "skin tone support unicode version": "9.0",
      "name": "call me hand",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "👈",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "backhand index pointing left",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👉",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "backhand index pointing right",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👆",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "backhand index pointing up",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🖕",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "middle finger",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👇",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "backhand index pointing down",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "☝️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "index pointing up",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "👍",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "thumbs up",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👎",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "thumbs down",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "✊",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "raised fist",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👊",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "oncoming fist",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🤛",
      "skin tone support": true,
      "skin tone support unicode version": "9.0",
      "name": "left facing fist",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🤜",
      "skin tone support": true,
      "skin tone support unicode version": "9.0",
      "name": "right facing fist",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "👏",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "clapping hands",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🙌",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "raising hands",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👐",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "open hands",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🤲",
      "skin tone support": true,
      "skin tone support unicode version": "10.0",
      "name": "palms up together",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🤝",
      "skin tone support": false,
      "name": "handshake",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🙏",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "folded hands",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "✍️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "writing hand",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "💅",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "nail polish",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🤳",
      "skin tone support": true,
      "skin tone support unicode version": "9.0",
      "name": "selfie",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "💪",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "flexed biceps",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🦾",
      "skin tone support": false,
      "name": "mechanical arm",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "🦿",
      "skin tone support": false,
      "name": "mechanical leg",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "🦵",
      "skin tone support": true,
      "skin tone support unicode version": "11.0",
      "name": "leg",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "🦶",
      "skin tone support": true,
      "skin tone support unicode version": "11.0",
      "name": "foot",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "👂",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "ear",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🦻",
      "skin tone support": true,
      "skin tone support unicode version": "12.0",
      "name": "ear with hearing aid",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "👃",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "nose",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🧠",
      "skin tone support": false,
      "name": "brain",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🦷",
      "skin tone support": false,
      "name": "tooth",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "🦴",
      "skin tone support": false,
      "name": "bone",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "👀",
      "skin tone support": false,
      "name": "eyes",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👁️",
      "skin tone support": false,
      "name": "eye",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👅",
      "skin tone support": false,
      "name": "tongue",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👄",
      "skin tone support": false,
      "name": "mouth",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👶",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "baby",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🧒",
      "skin tone support": true,
      "skin tone support unicode version": "10.0",
      "name": "child",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "👦",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "boy",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👧",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "girl",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🧑",
      "skin tone support": true,
      "skin tone support unicode version": "10.0",
      "name": "person",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "👱",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "person blond hair",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👨",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "man",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🧔",
      "skin tone support": true,
      "skin tone support unicode version": "10.0",
      "name": "man beard",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "👨‍🦰",
      "skin tone support": true,
      "skin tone support unicode version": "11.0",
      "name": "man red hair",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "👨‍🦱",
      "skin tone support": true,
      "skin tone support unicode version": "11.0",
      "name": "man curly hair",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "👨‍🦳",
      "skin tone support": true,
      "skin tone support unicode version": "11.0",
      "name": "man white hair",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "👨‍🦲",
      "skin tone support": true,
      "skin tone support unicode version": "11.0",
      "name": "man bald",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "👩",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "woman",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👩‍🦰",
      "skin tone support": true,
      "skin tone support unicode version": "11.0",
      "name": "woman red hair",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "🧑‍🦰",
      "skin tone support": true,
      "skin tone support unicode version": "11.0",
      "name": "person red hair",
      "unicode version": "11.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "👩‍🦱",
      "skin tone support": true,
      "skin tone support unicode version": "11.0",
      "name": "woman curly hair",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "🧑‍🦱",
      "skin tone support": true,
      "skin tone support unicode version": "11.0",
      "name": "person curly hair",
      "unicode version": "11.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "👩‍🦳",
      "skin tone support": true,
      "skin tone support unicode version": "11.0",
      "name": "woman white hair",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "🧑‍🦳",
      "skin tone support": true,
      "skin tone support unicode version": "11.0",
      "name": "person white hair",
      "unicode version": "11.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "👩‍🦲",
      "skin tone support": true,
      "skin tone support unicode version": "11.0",
      "name": "woman bald",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "🧑‍🦲",
      "skin tone support": true,
      "skin tone support unicode version": "11.0",
      "name": "person bald",
      "unicode version": "11.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "👱‍♀️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "woman blond hair",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "👱‍♂️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "man blond hair",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🧓",
      "skin tone support": true,
      "skin tone support unicode version": "10.0",
      "name": "older person",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "👴",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "old man",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👵",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "old woman",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🙍",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "person frowning",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🙍‍♂️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "man frowning",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🙍‍♀️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "woman frowning",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🙎",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "person pouting",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🙎‍♂️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "man pouting",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🙎‍♀️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "woman pouting",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🙅",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "person gesturing no",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🙅‍♂️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "man gesturing no",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🙅‍♀️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "woman gesturing no",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🙆",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "person gesturing ok",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🙆‍♂️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "man gesturing ok",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🙆‍♀️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "woman gesturing ok",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "💁",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "person tipping hand",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "💁‍♂️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "man tipping hand",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "💁‍♀️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "woman tipping hand",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🙋",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "person raising hand",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🙋‍♂️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "man raising hand",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🙋‍♀️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "woman raising hand",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🧏",
      "skin tone support": true,
      "skin tone support unicode version": "12.0",
      "name": "deaf person",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "🧏‍♂️",
      "skin tone support": true,
      "skin tone support unicode version": "12.0",
      "name": "deaf man",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "🧏‍♀️",
      "skin tone support": true,
      "skin tone support unicode version": "12.0",
      "name": "deaf woman",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "🙇",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "person bowing",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🙇‍♂️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "man bowing",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🙇‍♀️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "woman bowing",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🤦",
      "skin tone support": true,
      "skin tone support unicode version": "9.0",
      "name": "person facepalming",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🤦‍♂️",
      "skin tone support": true,
      "skin tone support unicode version": "9.0",
      "name": "man facepalming",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🤦‍♀️",
      "skin tone support": true,
      "skin tone support unicode version": "9.0",
      "name": "woman facepalming",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🤷",
      "skin tone support": true,
      "skin tone support unicode version": "9.0",
      "name": "person shrugging",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🤷‍♂️",
      "skin tone support": true,
      "skin tone support unicode version": "9.0",
      "name": "man shrugging",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🤷‍♀️",
      "skin tone support": true,
      "skin tone support unicode version": "9.0",
      "name": "woman shrugging",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🧑‍⚕️",
      "skin tone support": true,
      "skin tone support unicode version": "10.0",
      "name": "health worker",
      "unicode version": "10.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "👨‍⚕️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "man health worker",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "👩‍⚕️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "woman health worker",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🧑‍🎓",
      "skin tone support": true,
      "skin tone support unicode version": "10.0",
      "name": "student",
      "unicode version": "10.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "👨‍🎓",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "man student",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "👩‍🎓",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "woman student",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🧑‍🏫",
      "skin tone support": true,
      "skin tone support unicode version": "10.0",
      "name": "teacher",
      "unicode version": "10.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "👨‍🏫",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "man teacher",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "👩‍🏫",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "woman teacher",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🧑‍⚖️",
      "skin tone support": true,
      "skin tone support unicode version": "10.0",
      "name": "judge",
      "unicode version": "10.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "👨‍⚖️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "man judge",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "👩‍⚖️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "woman judge",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🧑‍🌾",
      "skin tone support": true,
      "skin tone support unicode version": "10.0",
      "name": "farmer",
      "unicode version": "10.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "👨‍🌾",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "man farmer",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "👩‍🌾",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "woman farmer",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🧑‍🍳",
      "skin tone support": true,
      "skin tone support unicode version": "10.0",
      "name": "cook",
      "unicode version": "10.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "👨‍🍳",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "man cook",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "👩‍🍳",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "woman cook",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🧑‍🔧",
      "skin tone support": true,
      "skin tone support unicode version": "10.0",
      "name": "mechanic",
      "unicode version": "10.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "👨‍🔧",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "man mechanic",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "👩‍🔧",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "woman mechanic",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🧑‍🏭",
      "skin tone support": true,
      "skin tone support unicode version": "10.0",
      "name": "factory worker",
      "unicode version": "10.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "👨‍🏭",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "man factory worker",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "👩‍🏭",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "woman factory worker",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🧑‍💼",
      "skin tone support": true,
      "skin tone support unicode version": "10.0",
      "name": "office worker",
      "unicode version": "10.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "👨‍💼",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "man office worker",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "👩‍💼",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "woman office worker",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🧑‍🔬",
      "skin tone support": true,
      "skin tone support unicode version": "10.0",
      "name": "scientist",
      "unicode version": "10.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "👨‍🔬",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "man scientist",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "👩‍🔬",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "woman scientist",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🧑‍💻",
      "skin tone support": true,
      "skin tone support unicode version": "10.0",
      "name": "technologist",
      "unicode version": "10.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "👨‍💻",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "man technologist",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "👩‍💻",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "woman technologist",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🧑‍🎤",
      "skin tone support": true,
      "skin tone support unicode version": "10.0",
      "name": "singer",
      "unicode version": "10.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "👨‍🎤",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "man singer",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "👩‍🎤",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "woman singer",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🧑‍🎨",
      "skin tone support": true,
      "skin tone support unicode version": "10.0",
      "name": "artist",
      "unicode version": "10.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "👨‍🎨",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "man artist",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "👩‍🎨",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "woman artist",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🧑‍✈️",
      "skin tone support": true,
      "skin tone support unicode version": "10.0",
      "name": "pilot",
      "unicode version": "10.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "👨‍✈️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "man pilot",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "👩‍✈️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "woman pilot",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🧑‍🚀",
      "skin tone support": true,
      "skin tone support unicode version": "10.0",
      "name": "astronaut",
      "unicode version": "10.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "👨‍🚀",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "man astronaut",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "👩‍🚀",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "woman astronaut",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🧑‍🚒",
      "skin tone support": true,
      "skin tone support unicode version": "10.0",
      "name": "firefighter",
      "unicode version": "10.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "👨‍🚒",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "man firefighter",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "👩‍🚒",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "woman firefighter",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "👮",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "police officer",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👮‍♂️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "man police officer",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "👮‍♀️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "woman police officer",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🕵️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "detective",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🕵️‍♂️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "man detective",
      "unicode version": "7.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🕵️‍♀️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "woman detective",
      "unicode version": "7.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "💂",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "guard",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "💂‍♂️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "man guard",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "💂‍♀️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "woman guard",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "👷",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "construction worker",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👷‍♂️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "man construction worker",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "👷‍♀️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "woman construction worker",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🤴",
      "skin tone support": true,
      "skin tone support unicode version": "9.0",
      "name": "prince",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "👸",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "princess",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👳",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "person wearing turban",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👳‍♂️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "man wearing turban",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "👳‍♀️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "woman wearing turban",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "👲",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "man with skullcap",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🧕",
      "skin tone support": true,
      "skin tone support unicode version": "10.0",
      "name": "woman with headscarf",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🤵",
      "skin tone support": true,
      "skin tone support unicode version": "9.0",
      "name": "man in tuxedo",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "👰",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "bride with veil",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🤰",
      "skin tone support": true,
      "skin tone support unicode version": "9.0",
      "name": "pregnant woman",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🤱",
      "skin tone support": true,
      "skin tone support unicode version": "10.0",
      "name": "breast feeding",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "👼",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "baby angel",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🎅",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "santa claus",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🤶",
      "skin tone support": true,
      "skin tone support unicode version": "9.0",
      "name": "mrs claus",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🦸",
      "skin tone support": true,
      "skin tone support unicode version": "11.0",
      "name": "superhero",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "🦸‍♂️",
      "skin tone support": true,
      "skin tone support unicode version": "11.0",
      "name": "man superhero",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "🦸‍♀️",
      "skin tone support": true,
      "skin tone support unicode version": "11.0",
      "name": "woman superhero",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "🦹",
      "skin tone support": true,
      "skin tone support unicode version": "11.0",
      "name": "supervillain",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "🦹‍♂️",
      "skin tone support": true,
      "skin tone support unicode version": "11.0",
      "name": "man supervillain",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "🦹‍♀️",
      "skin tone support": true,
      "skin tone support unicode version": "11.0",
      "name": "woman supervillain",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "🧙",
      "skin tone support": true,
      "skin tone support unicode version": "10.0",
      "name": "mage",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🧙‍♂️",
      "skin tone support": true,
      "skin tone support unicode version": "10.0",
      "name": "man mage",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🧙‍♀️",
      "skin tone support": true,
      "skin tone support unicode version": "10.0",
      "name": "woman mage",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🧚",
      "skin tone support": true,
      "skin tone support unicode version": "10.0",
      "name": "fairy",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🧚‍♂️",
      "skin tone support": true,
      "skin tone support unicode version": "10.0",
      "name": "man fairy",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🧚‍♀️",
      "skin tone support": true,
      "skin tone support unicode version": "10.0",
      "name": "woman fairy",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🧛",
      "skin tone support": true,
      "skin tone support unicode version": "10.0",
      "name": "vampire",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🧛‍♂️",
      "skin tone support": true,
      "skin tone support unicode version": "10.0",
      "name": "man vampire",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🧛‍♀️",
      "skin tone support": true,
      "skin tone support unicode version": "10.0",
      "name": "woman vampire",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🧜",
      "skin tone support": true,
      "skin tone support unicode version": "10.0",
      "name": "merperson",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🧜‍♂️",
      "skin tone support": true,
      "skin tone support unicode version": "10.0",
      "name": "merman",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🧜‍♀️",
      "skin tone support": true,
      "skin tone support unicode version": "10.0",
      "name": "mermaid",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🧝",
      "skin tone support": true,
      "skin tone support unicode version": "10.0",
      "name": "elf",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🧝‍♂️",
      "skin tone support": true,
      "skin tone support unicode version": "10.0",
      "name": "man elf",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🧝‍♀️",
      "skin tone support": true,
      "skin tone support unicode version": "10.0",
      "name": "woman elf",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🧞",
      "skin tone support": false,
      "name": "genie",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🧞‍♂️",
      "skin tone support": false,
      "name": "man genie",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🧞‍♀️",
      "skin tone support": false,
      "name": "woman genie",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🧟",
      "skin tone support": false,
      "name": "zombie",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🧟‍♂️",
      "skin tone support": false,
      "name": "man zombie",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🧟‍♀️",
      "skin tone support": false,
      "name": "woman zombie",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "💆",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "person getting massage",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "💆‍♂️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "man getting massage",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "💆‍♀️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "woman getting massage",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "💇",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "person getting haircut",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "💇‍♂️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "man getting haircut",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "💇‍♀️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "woman getting haircut",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🚶",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "person walking",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🚶‍♂️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "man walking",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🚶‍♀️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "woman walking",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🧍",
      "skin tone support": true,
      "skin tone support unicode version": "12.0",
      "name": "person standing",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "🧍‍♂️",
      "skin tone support": true,
      "skin tone support unicode version": "12.0",
      "name": "man standing",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "🧍‍♀️",
      "skin tone support": true,
      "skin tone support unicode version": "12.0",
      "name": "woman standing",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "🧎",
      "skin tone support": true,
      "skin tone support unicode version": "12.0",
      "name": "person kneeling",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "🧎‍♂️",
      "skin tone support": true,
      "skin tone support unicode version": "12.0",
      "name": "man kneeling",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "🧎‍♀️",
      "skin tone support": true,
      "skin tone support unicode version": "12.0",
      "name": "woman kneeling",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "🧑‍🦯",
      "skin tone support": true,
      "skin tone support unicode version": "12.0",
      "name": "person with probing cane",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "👨‍🦯",
      "skin tone support": true,
      "skin tone support unicode version": "12.0",
      "name": "man with probing cane",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "👩‍🦯",
      "skin tone support": true,
      "skin tone support unicode version": "12.0",
      "name": "woman with probing cane",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "🧑‍🦼",
      "skin tone support": true,
      "skin tone support unicode version": "12.0",
      "name": "person in motorized wheelchair",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "👨‍🦼",
      "skin tone support": true,
      "skin tone support unicode version": "12.0",
      "name": "man in motorized wheelchair",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "👩‍🦼",
      "skin tone support": true,
      "skin tone support unicode version": "12.0",
      "name": "woman in motorized wheelchair",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "🧑‍🦽",
      "skin tone support": true,
      "skin tone support unicode version": "12.0",
      "name": "person in manual wheelchair",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "👨‍🦽",
      "skin tone support": true,
      "skin tone support unicode version": "12.0",
      "name": "man in manual wheelchair",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "👩‍🦽",
      "skin tone support": true,
      "skin tone support unicode version": "12.0",
      "name": "woman in manual wheelchair",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "🏃",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "person running",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🏃‍♂️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "man running",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🏃‍♀️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "woman running",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "💃",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "woman dancing",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🕺",
      "skin tone support": true,
      "skin tone support unicode version": "9.0",
      "name": "man dancing",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🕴️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "man in suit levitating",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👯",
      "skin tone support": false,
      "name": "people with bunny ears",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👯‍♂️",
      "skin tone support": false,
      "name": "men with bunny ears",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "👯‍♀️",
      "skin tone support": false,
      "name": "women with bunny ears",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🧖",
      "skin tone support": true,
      "skin tone support unicode version": "10.0",
      "name": "person in steamy room",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🧖‍♂️",
      "skin tone support": true,
      "skin tone support unicode version": "10.0",
      "name": "man in steamy room",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🧖‍♀️",
      "skin tone support": true,
      "skin tone support unicode version": "10.0",
      "name": "woman in steamy room",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🧗",
      "skin tone support": true,
      "skin tone support unicode version": "10.0",
      "name": "person climbing",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🧗‍♂️",
      "skin tone support": true,
      "skin tone support unicode version": "10.0",
      "name": "man climbing",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🧗‍♀️",
      "skin tone support": true,
      "skin tone support unicode version": "10.0",
      "name": "woman climbing",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🤺",
      "skin tone support": false,
      "name": "person fencing",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🏇",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "horse racing",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "⛷️",
      "skin tone support": false,
      "name": "skier",
      "unicode version": "5.2",
      "emoji version": "2.0"
    },
    {
      "emoji": "🏂",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "snowboarder",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🏌️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "person golfing",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🏌️‍♂️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "man golfing",
      "unicode version": "7.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🏌️‍♀️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "woman golfing",
      "unicode version": "7.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🏄",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "person surfing",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🏄‍♂️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "man surfing",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🏄‍♀️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "woman surfing",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🚣",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "person rowing boat",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🚣‍♂️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "man rowing boat",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🚣‍♀️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "woman rowing boat",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🏊",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "person swimming",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🏊‍♂️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "man swimming",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🏊‍♀️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "woman swimming",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "⛹️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "person bouncing ball",
      "unicode version": "5.2",
      "emoji version": "2.0"
    },
    {
      "emoji": "⛹️‍♂️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "man bouncing ball",
      "unicode version": "5.2",
      "emoji version": "4.0"
    },
    {
      "emoji": "⛹️‍♀️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "woman bouncing ball",
      "unicode version": "5.2",
      "emoji version": "4.0"
    },
    {
      "emoji": "🏋️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "person lifting weights",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🏋️‍♂️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "man lifting weights",
      "unicode version": "7.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🏋️‍♀️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "woman lifting weights",
      "unicode version": "7.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🚴",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "person biking",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🚴‍♂️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "man biking",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🚴‍♀️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "woman biking",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🚵",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "person mountain biking",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🚵‍♂️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "man mountain biking",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🚵‍♀️",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "woman mountain biking",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🤸",
      "skin tone support": true,
      "skin tone support unicode version": "9.0",
      "name": "person cartwheeling",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🤸‍♂️",
      "skin tone support": true,
      "skin tone support unicode version": "9.0",
      "name": "man cartwheeling",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🤸‍♀️",
      "skin tone support": true,
      "skin tone support unicode version": "9.0",
      "name": "woman cartwheeling",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🤼",
      "skin tone support": false,
      "name": "people wrestling",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🤼‍♂️",
      "skin tone support": false,
      "name": "men wrestling",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🤼‍♀️",
      "skin tone support": false,
      "name": "women wrestling",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🤽",
      "skin tone support": true,
      "skin tone support unicode version": "9.0",
      "name": "person playing water polo",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🤽‍♂️",
      "skin tone support": true,
      "skin tone support unicode version": "9.0",
      "name": "man playing water polo",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🤽‍♀️",
      "skin tone support": true,
      "skin tone support unicode version": "9.0",
      "name": "woman playing water polo",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🤾",
      "skin tone support": true,
      "skin tone support unicode version": "9.0",
      "name": "person playing handball",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🤾‍♂️",
      "skin tone support": true,
      "skin tone support unicode version": "9.0",
      "name": "man playing handball",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🤾‍♀️",
      "skin tone support": true,
      "skin tone support unicode version": "9.0",
      "name": "woman playing handball",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🤹",
      "skin tone support": true,
      "skin tone support unicode version": "9.0",
      "name": "person juggling",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🤹‍♂️",
      "skin tone support": true,
      "skin tone support unicode version": "9.0",
      "name": "man juggling",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🤹‍♀️",
      "skin tone support": true,
      "skin tone support unicode version": "9.0",
      "name": "woman juggling",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🧘",
      "skin tone support": true,
      "skin tone support unicode version": "10.0",
      "name": "person in lotus position",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🧘‍♂️",
      "skin tone support": true,
      "skin tone support unicode version": "10.0",
      "name": "man in lotus position",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🧘‍♀️",
      "skin tone support": true,
      "skin tone support unicode version": "10.0",
      "name": "woman in lotus position",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🛀",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "person taking bath",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🛌",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "person in bed",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🧑‍🤝‍🧑",
      "skin tone support": true,
      "skin tone support unicode version": "10.0",
      "name": "people holding hands",
      "unicode version": "10.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "👭",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "women holding hands",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👫",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "woman and man holding hands",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👬",
      "skin tone support": true,
      "skin tone support unicode version": "8.0",
      "name": "men holding hands",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "💏",
      "skin tone support": false,
      "name": "kiss",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👩‍❤️‍💋‍👨",
      "skin tone support": false,
      "name": "kiss woman man",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👨‍❤️‍💋‍👨",
      "skin tone support": false,
      "name": "kiss man man",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👩‍❤️‍💋‍👩",
      "skin tone support": false,
      "name": "kiss woman woman",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "💑",
      "skin tone support": false,
      "name": "couple with heart",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👩‍❤️‍👨",
      "skin tone support": false,
      "name": "couple with heart woman man",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👨‍❤️‍👨",
      "skin tone support": false,
      "name": "couple with heart man man",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👩‍❤️‍👩",
      "skin tone support": false,
      "name": "couple with heart woman woman",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👪",
      "skin tone support": false,
      "name": "family",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👨‍👩‍👦",
      "skin tone support": false,
      "name": "family man woman boy",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👨‍👩‍👧",
      "skin tone support": false,
      "name": "family man woman girl",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👨‍👩‍👧‍👦",
      "skin tone support": false,
      "name": "family man woman girl boy",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👨‍👩‍👦‍👦",
      "skin tone support": false,
      "name": "family man woman boy boy",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👨‍👩‍👧‍👧",
      "skin tone support": false,
      "name": "family man woman girl girl",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👨‍👨‍👦",
      "skin tone support": false,
      "name": "family man man boy",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👨‍👨‍👧",
      "skin tone support": false,
      "name": "family man man girl",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👨‍👨‍👧‍👦",
      "skin tone support": false,
      "name": "family man man girl boy",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👨‍👨‍👦‍👦",
      "skin tone support": false,
      "name": "family man man boy boy",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👨‍👨‍👧‍👧",
      "skin tone support": false,
      "name": "family man man girl girl",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👩‍👩‍👦",
      "skin tone support": false,
      "name": "family woman woman boy",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👩‍👩‍👧",
      "skin tone support": false,
      "name": "family woman woman girl",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👩‍👩‍👧‍👦",
      "skin tone support": false,
      "name": "family woman woman girl boy",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👩‍👩‍👦‍👦",
      "skin tone support": false,
      "name": "family woman woman boy boy",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👩‍👩‍👧‍👧",
      "skin tone support": false,
      "name": "family woman woman girl girl",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👨‍👦",
      "skin tone support": false,
      "name": "family man boy",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "👨‍👦‍👦",
      "skin tone support": false,
      "name": "family man boy boy",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "👨‍👧",
      "skin tone support": false,
      "name": "family man girl",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "👨‍👧‍👦",
      "skin tone support": false,
      "name": "family man girl boy",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "👨‍👧‍👧",
      "skin tone support": false,
      "name": "family man girl girl",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "👩‍👦",
      "skin tone support": false,
      "name": "family woman boy",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "👩‍👦‍👦",
      "skin tone support": false,
      "name": "family woman boy boy",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "👩‍👧",
      "skin tone support": false,
      "name": "family woman girl",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "👩‍👧‍👦",
      "skin tone support": false,
      "name": "family woman girl boy",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "👩‍👧‍👧",
      "skin tone support": false,
      "name": "family woman girl girl",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🗣️",
      "skin tone support": false,
      "name": "speaking head",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👤",
      "skin tone support": false,
      "name": "bust in silhouette",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👥",
      "skin tone support": false,
      "name": "busts in silhouette",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👣",
      "skin tone support": false,
      "name": "footprints",
      "unicode version": "6.0",
      "emoji version": "2.0"
    }
  ]
	},
	{
	"name": "Animals and Nature",
	"emojis": [
    {
      "emoji": "🐵",
      "skin tone support": false,
      "name": "monkey face",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🐒",
      "skin tone support": false,
      "name": "monkey",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🦍",
      "skin tone support": false,
      "name": "gorilla",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🦧",
      "skin tone support": false,
      "name": "orangutan",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "🐶",
      "skin tone support": false,
      "name": "dog face",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🐕",
      "skin tone support": false,
      "name": "dog",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🦮",
      "skin tone support": false,
      "name": "guide dog",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "🐕‍🦺",
      "skin tone support": false,
      "name": "service dog",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "🐩",
      "skin tone support": false,
      "name": "poodle",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🐺",
      "skin tone support": false,
      "name": "wolf",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🦊",
      "skin tone support": false,
      "name": "fox",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🦝",
      "skin tone support": false,
      "name": "raccoon",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "🐱",
      "skin tone support": false,
      "name": "cat face",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🐈",
      "skin tone support": false,
      "name": "cat",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🦁",
      "skin tone support": false,
      "name": "lion",
      "unicode version": "8.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🐯",
      "skin tone support": false,
      "name": "tiger face",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🐅",
      "skin tone support": false,
      "name": "tiger",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🐆",
      "skin tone support": false,
      "name": "leopard",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🐴",
      "skin tone support": false,
      "name": "horse face",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🐎",
      "skin tone support": false,
      "name": "horse",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🦄",
      "skin tone support": false,
      "name": "unicorn",
      "unicode version": "8.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🦓",
      "skin tone support": false,
      "name": "zebra",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🦌",
      "skin tone support": false,
      "name": "deer",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🐮",
      "skin tone support": false,
      "name": "cow face",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🐂",
      "skin tone support": false,
      "name": "ox",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🐃",
      "skin tone support": false,
      "name": "water buffalo",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🐄",
      "skin tone support": false,
      "name": "cow",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🐷",
      "skin tone support": false,
      "name": "pig face",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🐖",
      "skin tone support": false,
      "name": "pig",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🐗",
      "skin tone support": false,
      "name": "boar",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🐽",
      "skin tone support": false,
      "name": "pig nose",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🐏",
      "skin tone support": false,
      "name": "ram",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🐑",
      "skin tone support": false,
      "name": "ewe",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🐐",
      "skin tone support": false,
      "name": "goat",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🐪",
      "skin tone support": false,
      "name": "camel",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🐫",
      "skin tone support": false,
      "name": "two hump camel",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🦙",
      "skin tone support": false,
      "name": "llama",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "🦒",
      "skin tone support": false,
      "name": "giraffe",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🐘",
      "skin tone support": false,
      "name": "elephant",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🦏",
      "skin tone support": false,
      "name": "rhinoceros",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🦛",
      "skin tone support": false,
      "name": "hippopotamus",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "🐭",
      "skin tone support": false,
      "name": "mouse face",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🐁",
      "skin tone support": false,
      "name": "mouse",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🐀",
      "skin tone support": false,
      "name": "rat",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🐹",
      "skin tone support": false,
      "name": "hamster",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🐰",
      "skin tone support": false,
      "name": "rabbit face",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🐇",
      "skin tone support": false,
      "name": "rabbit",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🐿️",
      "skin tone support": false,
      "name": "chipmunk",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🦔",
      "skin tone support": false,
      "name": "hedgehog",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🦇",
      "skin tone support": false,
      "name": "bat",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🐻",
      "skin tone support": false,
      "name": "bear",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🐨",
      "skin tone support": false,
      "name": "koala",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🐼",
      "skin tone support": false,
      "name": "panda",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🦥",
      "skin tone support": false,
      "name": "sloth",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "🦦",
      "skin tone support": false,
      "name": "otter",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "🦨",
      "skin tone support": false,
      "name": "skunk",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "🦘",
      "skin tone support": false,
      "name": "kangaroo",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "🦡",
      "skin tone support": false,
      "name": "badger",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "🐾",
      "skin tone support": false,
      "name": "paw prints",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🦃",
      "skin tone support": false,
      "name": "turkey",
      "unicode version": "8.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🐔",
      "skin tone support": false,
      "name": "chicken",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🐓",
      "skin tone support": false,
      "name": "rooster",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🐣",
      "skin tone support": false,
      "name": "hatching chick",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🐤",
      "skin tone support": false,
      "name": "baby chick",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🐥",
      "skin tone support": false,
      "name": "front facing baby chick",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🐦",
      "skin tone support": false,
      "name": "bird",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🐧",
      "skin tone support": false,
      "name": "penguin",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🕊️",
      "skin tone support": false,
      "name": "dove",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🦅",
      "skin tone support": false,
      "name": "eagle",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🦆",
      "skin tone support": false,
      "name": "duck",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🦢",
      "skin tone support": false,
      "name": "swan",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "🦉",
      "skin tone support": false,
      "name": "owl",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🦩",
      "skin tone support": false,
      "name": "flamingo",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "🦚",
      "skin tone support": false,
      "name": "peacock",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "🦜",
      "skin tone support": false,
      "name": "parrot",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "🐸",
      "skin tone support": false,
      "name": "frog",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🐊",
      "skin tone support": false,
      "name": "crocodile",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🐢",
      "skin tone support": false,
      "name": "turtle",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🦎",
      "skin tone support": false,
      "name": "lizard",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🐍",
      "skin tone support": false,
      "name": "snake",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🐲",
      "skin tone support": false,
      "name": "dragon face",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🐉",
      "skin tone support": false,
      "name": "dragon",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🦕",
      "skin tone support": false,
      "name": "sauropod",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🦖",
      "skin tone support": false,
      "name": "t rex",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🐳",
      "skin tone support": false,
      "name": "spouting whale",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🐋",
      "skin tone support": false,
      "name": "whale",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🐬",
      "skin tone support": false,
      "name": "dolphin",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🐟",
      "skin tone support": false,
      "name": "fish",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🐠",
      "skin tone support": false,
      "name": "tropical fish",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🐡",
      "skin tone support": false,
      "name": "blowfish",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🦈",
      "skin tone support": false,
      "name": "shark",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🐙",
      "skin tone support": false,
      "name": "octopus",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🐚",
      "skin tone support": false,
      "name": "spiral shell",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🐌",
      "skin tone support": false,
      "name": "snail",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🦋",
      "skin tone support": false,
      "name": "butterfly",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🐛",
      "skin tone support": false,
      "name": "bug",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🐜",
      "skin tone support": false,
      "name": "ant",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🐝",
      "skin tone support": false,
      "name": "honeybee",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🐞",
      "skin tone support": false,
      "name": "lady beetle",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🦗",
      "skin tone support": false,
      "name": "cricket",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🕷️",
      "skin tone support": false,
      "name": "spider",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🕸️",
      "skin tone support": false,
      "name": "spider web",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🦂",
      "skin tone support": false,
      "name": "scorpion",
      "unicode version": "8.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🦟",
      "skin tone support": false,
      "name": "mosquito",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "🦠",
      "skin tone support": false,
      "name": "microbe",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "💐",
      "skin tone support": false,
      "name": "bouquet",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🌸",
      "skin tone support": false,
      "name": "cherry blossom",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "💮",
      "skin tone support": false,
      "name": "white flower",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🏵️",
      "skin tone support": false,
      "name": "rosette",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🌹",
      "skin tone support": false,
      "name": "rose",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🥀",
      "skin tone support": false,
      "name": "wilted flower",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🌺",
      "skin tone support": false,
      "name": "hibiscus",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🌻",
      "skin tone support": false,
      "name": "sunflower",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🌼",
      "skin tone support": false,
      "name": "blossom",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🌷",
      "skin tone support": false,
      "name": "tulip",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🌱",
      "skin tone support": false,
      "name": "seedling",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🌲",
      "skin tone support": false,
      "name": "evergreen tree",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🌳",
      "skin tone support": false,
      "name": "deciduous tree",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🌴",
      "skin tone support": false,
      "name": "palm tree",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🌵",
      "skin tone support": false,
      "name": "cactus",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🌾",
      "skin tone support": false,
      "name": "sheaf of rice",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🌿",
      "skin tone support": false,
      "name": "herb",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "☘️",
      "skin tone support": false,
      "name": "shamrock",
      "unicode version": "4.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "🍀",
      "skin tone support": false,
      "name": "four leaf clover",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🍁",
      "skin tone support": false,
      "name": "maple leaf",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🍂",
      "skin tone support": false,
      "name": "fallen leaf",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🍃",
      "skin tone support": false,
      "name": "leaf fluttering in wind",
      "unicode version": "6.0",
      "emoji version": "2.0"
    }
  ]
	},
	{
	"name": "Food and Drink",
	"emojis": [
    {
      "emoji": "🍇",
      "skin tone support": false,
      "name": "grapes",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🍈",
      "skin tone support": false,
      "name": "melon",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🍉",
      "skin tone support": false,
      "name": "watermelon",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🍊",
      "skin tone support": false,
      "name": "tangerine",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🍋",
      "skin tone support": false,
      "name": "lemon",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🍌",
      "skin tone support": false,
      "name": "banana",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🍍",
      "skin tone support": false,
      "name": "pineapple",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🥭",
      "skin tone support": false,
      "name": "mango",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "🍎",
      "skin tone support": false,
      "name": "red apple",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🍏",
      "skin tone support": false,
      "name": "green apple",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🍐",
      "skin tone support": false,
      "name": "pear",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🍑",
      "skin tone support": false,
      "name": "peach",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🍒",
      "skin tone support": false,
      "name": "cherries",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🍓",
      "skin tone support": false,
      "name": "strawberry",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🥝",
      "skin tone support": false,
      "name": "kiwi fruit",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🍅",
      "skin tone support": false,
      "name": "tomato",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🥥",
      "skin tone support": false,
      "name": "coconut",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🥑",
      "skin tone support": false,
      "name": "avocado",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🍆",
      "skin tone support": false,
      "name": "eggplant",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🥔",
      "skin tone support": false,
      "name": "potato",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🥕",
      "skin tone support": false,
      "name": "carrot",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🌽",
      "skin tone support": false,
      "name": "ear of corn",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🌶️",
      "skin tone support": false,
      "name": "hot pepper",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🥒",
      "skin tone support": false,
      "name": "cucumber",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🥬",
      "skin tone support": false,
      "name": "leafy green",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "🥦",
      "skin tone support": false,
      "name": "broccoli",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🧄",
      "skin tone support": false,
      "name": "garlic",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "🧅",
      "skin tone support": false,
      "name": "onion",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "🍄",
      "skin tone support": false,
      "name": "mushroom",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🥜",
      "skin tone support": false,
      "name": "peanuts",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🌰",
      "skin tone support": false,
      "name": "chestnut",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🍞",
      "skin tone support": false,
      "name": "bread",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🥐",
      "skin tone support": false,
      "name": "croissant",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🥖",
      "skin tone support": false,
      "name": "baguette bread",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🥨",
      "skin tone support": false,
      "name": "pretzel",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🥯",
      "skin tone support": false,
      "name": "bagel",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "🥞",
      "skin tone support": false,
      "name": "pancakes",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🧇",
      "skin tone support": false,
      "name": "waffle",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "🧀",
      "skin tone support": false,
      "name": "cheese wedge",
      "unicode version": "8.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🍖",
      "skin tone support": false,
      "name": "meat on bone",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🍗",
      "skin tone support": false,
      "name": "poultry leg",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🥩",
      "skin tone support": false,
      "name": "cut of meat",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🥓",
      "skin tone support": false,
      "name": "bacon",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🍔",
      "skin tone support": false,
      "name": "hamburger",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🍟",
      "skin tone support": false,
      "name": "french fries",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🍕",
      "skin tone support": false,
      "name": "pizza",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🌭",
      "skin tone support": false,
      "name": "hot dog",
      "unicode version": "8.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🥪",
      "skin tone support": false,
      "name": "sandwich",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🌮",
      "skin tone support": false,
      "name": "taco",
      "unicode version": "8.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🌯",
      "skin tone support": false,
      "name": "burrito",
      "unicode version": "8.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🥙",
      "skin tone support": false,
      "name": "stuffed flatbread",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🧆",
      "skin tone support": false,
      "name": "falafel",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "🥚",
      "skin tone support": false,
      "name": "egg",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🍳",
      "skin tone support": false,
      "name": "cooking",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🥘",
      "skin tone support": false,
      "name": "shallow pan of food",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🍲",
      "skin tone support": false,
      "name": "pot of food",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🥣",
      "skin tone support": false,
      "name": "bowl with spoon",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🥗",
      "skin tone support": false,
      "name": "green salad",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🍿",
      "skin tone support": false,
      "name": "popcorn",
      "unicode version": "8.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🧈",
      "skin tone support": false,
      "name": "butter",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "🧂",
      "skin tone support": false,
      "name": "salt",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "🥫",
      "skin tone support": false,
      "name": "canned food",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🍱",
      "skin tone support": false,
      "name": "bento box",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🍘",
      "skin tone support": false,
      "name": "rice cracker",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🍙",
      "skin tone support": false,
      "name": "rice ball",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🍚",
      "skin tone support": false,
      "name": "cooked rice",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🍛",
      "skin tone support": false,
      "name": "curry rice",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🍜",
      "skin tone support": false,
      "name": "steaming bowl",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🍝",
      "skin tone support": false,
      "name": "spaghetti",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🍠",
      "skin tone support": false,
      "name": "roasted sweet potato",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🍢",
      "skin tone support": false,
      "name": "oden",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🍣",
      "skin tone support": false,
      "name": "sushi",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🍤",
      "skin tone support": false,
      "name": "fried shrimp",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🍥",
      "skin tone support": false,
      "name": "fish cake with swirl",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🥮",
      "skin tone support": false,
      "name": "moon cake",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "🍡",
      "skin tone support": false,
      "name": "dango",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🥟",
      "skin tone support": false,
      "name": "dumpling",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🥠",
      "skin tone support": false,
      "name": "fortune cookie",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🥡",
      "skin tone support": false,
      "name": "takeout box",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🦀",
      "skin tone support": false,
      "name": "crab",
      "unicode version": "8.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🦞",
      "skin tone support": false,
      "name": "lobster",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "🦐",
      "skin tone support": false,
      "name": "shrimp",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🦑",
      "skin tone support": false,
      "name": "squid",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🦪",
      "skin tone support": false,
      "name": "oyster",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "🍦",
      "skin tone support": false,
      "name": "soft ice cream",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🍧",
      "skin tone support": false,
      "name": "shaved ice",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🍨",
      "skin tone support": false,
      "name": "ice cream",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🍩",
      "skin tone support": false,
      "name": "doughnut",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🍪",
      "skin tone support": false,
      "name": "cookie",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🎂",
      "skin tone support": false,
      "name": "birthday cake",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🍰",
      "skin tone support": false,
      "name": "shortcake",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🧁",
      "skin tone support": false,
      "name": "cupcake",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "🥧",
      "skin tone support": false,
      "name": "pie",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🍫",
      "skin tone support": false,
      "name": "chocolate bar",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🍬",
      "skin tone support": false,
      "name": "candy",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🍭",
      "skin tone support": false,
      "name": "lollipop",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🍮",
      "skin tone support": false,
      "name": "custard",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🍯",
      "skin tone support": false,
      "name": "honey pot",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🍼",
      "skin tone support": false,
      "name": "baby bottle",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🥛",
      "skin tone support": false,
      "name": "glass of milk",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "☕",
      "skin tone support": false,
      "name": "hot beverage",
      "unicode version": "4.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🍵",
      "skin tone support": false,
      "name": "teacup without handle",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🍶",
      "skin tone support": false,
      "name": "sake",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🍾",
      "skin tone support": false,
      "name": "bottle with popping cork",
      "unicode version": "8.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🍷",
      "skin tone support": false,
      "name": "wine glass",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🍸",
      "skin tone support": false,
      "name": "cocktail glass",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🍹",
      "skin tone support": false,
      "name": "tropical drink",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🍺",
      "skin tone support": false,
      "name": "beer mug",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🍻",
      "skin tone support": false,
      "name": "clinking beer mugs",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🥂",
      "skin tone support": false,
      "name": "clinking glasses",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🥃",
      "skin tone support": false,
      "name": "tumbler glass",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🥤",
      "skin tone support": false,
      "name": "cup with straw",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🧃",
      "skin tone support": false,
      "name": "beverage box",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "🧉",
      "skin tone support": false,
      "name": "mate",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "🧊",
      "skin tone support": false,
      "name": "ice",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "🥢",
      "skin tone support": false,
      "name": "chopsticks",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🍽️",
      "skin tone support": false,
      "name": "fork and knife with plate",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🍴",
      "skin tone support": false,
      "name": "fork and knife",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🥄",
      "skin tone support": false,
      "name": "spoon",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🔪",
      "skin tone support": false,
      "name": "kitchen knife",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🏺",
      "skin tone support": false,
      "name": "amphora",
      "unicode version": "8.0",
      "emoji version": "2.0"
    }
  ]
	},
	{
	"name": "Travel and Places",
	"emojis": [
    {
      "emoji": "🌍",
      "skin tone support": false,
      "name": "globe showing europe africa",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🌎",
      "skin tone support": false,
      "name": "globe showing americas",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🌏",
      "skin tone support": false,
      "name": "globe showing asia australia",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🌐",
      "skin tone support": false,
      "name": "globe with meridians",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🗺️",
      "skin tone support": false,
      "name": "world map",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🗾",
      "skin tone support": false,
      "name": "map of japan",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🧭",
      "skin tone support": false,
      "name": "compass",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "🏔️",
      "skin tone support": false,
      "name": "snow capped mountain",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "⛰️",
      "skin tone support": false,
      "name": "mountain",
      "unicode version": "5.2",
      "emoji version": "2.0"
    },
    {
      "emoji": "🌋",
      "skin tone support": false,
      "name": "volcano",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🗻",
      "skin tone support": false,
      "name": "mount fuji",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🏕️",
      "skin tone support": false,
      "name": "camping",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🏖️",
      "skin tone support": false,
      "name": "beach with umbrella",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🏜️",
      "skin tone support": false,
      "name": "desert",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🏝️",
      "skin tone support": false,
      "name": "desert island",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🏞️",
      "skin tone support": false,
      "name": "national park",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🏟️",
      "skin tone support": false,
      "name": "stadium",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🏛️",
      "skin tone support": false,
      "name": "classical building",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🏗️",
      "skin tone support": false,
      "name": "building construction",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🧱",
      "skin tone support": false,
      "name": "brick",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "🏘️",
      "skin tone support": false,
      "name": "houses",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🏚️",
      "skin tone support": false,
      "name": "derelict house",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🏠",
      "skin tone support": false,
      "name": "house",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🏡",
      "skin tone support": false,
      "name": "house with garden",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🏢",
      "skin tone support": false,
      "name": "office building",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🏣",
      "skin tone support": false,
      "name": "japanese post office",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🏤",
      "skin tone support": false,
      "name": "post office",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🏥",
      "skin tone support": false,
      "name": "hospital",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🏦",
      "skin tone support": false,
      "name": "bank",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🏨",
      "skin tone support": false,
      "name": "hotel",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🏩",
      "skin tone support": false,
      "name": "love hotel",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🏪",
      "skin tone support": false,
      "name": "convenience store",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🏫",
      "skin tone support": false,
      "name": "school",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🏬",
      "skin tone support": false,
      "name": "department store",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🏭",
      "skin tone support": false,
      "name": "factory",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🏯",
      "skin tone support": false,
      "name": "japanese castle",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🏰",
      "skin tone support": false,
      "name": "castle",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "💒",
      "skin tone support": false,
      "name": "wedding",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🗼",
      "skin tone support": false,
      "name": "tokyo tower",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🗽",
      "skin tone support": false,
      "name": "statue of liberty",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "⛪",
      "skin tone support": false,
      "name": "church",
      "unicode version": "5.2",
      "emoji version": "2.0"
    },
    {
      "emoji": "🕌",
      "skin tone support": false,
      "name": "mosque",
      "unicode version": "8.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🛕",
      "skin tone support": false,
      "name": "hindu temple",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "🕍",
      "skin tone support": false,
      "name": "synagogue",
      "unicode version": "8.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "⛩️",
      "skin tone support": false,
      "name": "shinto shrine",
      "unicode version": "5.2",
      "emoji version": "2.0"
    },
    {
      "emoji": "🕋",
      "skin tone support": false,
      "name": "kaaba",
      "unicode version": "8.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "⛲",
      "skin tone support": false,
      "name": "fountain",
      "unicode version": "5.2",
      "emoji version": "2.0"
    },
    {
      "emoji": "⛺",
      "skin tone support": false,
      "name": "tent",
      "unicode version": "5.2",
      "emoji version": "2.0"
    },
    {
      "emoji": "🌁",
      "skin tone support": false,
      "name": "foggy",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🌃",
      "skin tone support": false,
      "name": "night with stars",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🏙️",
      "skin tone support": false,
      "name": "cityscape",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🌄",
      "skin tone support": false,
      "name": "sunrise over mountains",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🌅",
      "skin tone support": false,
      "name": "sunrise",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🌆",
      "skin tone support": false,
      "name": "cityscape at dusk",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🌇",
      "skin tone support": false,
      "name": "sunset",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🌉",
      "skin tone support": false,
      "name": "bridge at night",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "♨️",
      "skin tone support": false,
      "name": "hot springs",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "🎠",
      "skin tone support": false,
      "name": "carousel horse",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🎡",
      "skin tone support": false,
      "name": "ferris wheel",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🎢",
      "skin tone support": false,
      "name": "roller coaster",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "💈",
      "skin tone support": false,
      "name": "barber pole",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🎪",
      "skin tone support": false,
      "name": "circus tent",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🚂",
      "skin tone support": false,
      "name": "locomotive",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🚃",
      "skin tone support": false,
      "name": "railway car",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🚄",
      "skin tone support": false,
      "name": "high speed train",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🚅",
      "skin tone support": false,
      "name": "bullet train",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🚆",
      "skin tone support": false,
      "name": "train",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🚇",
      "skin tone support": false,
      "name": "metro",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🚈",
      "skin tone support": false,
      "name": "light rail",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🚉",
      "skin tone support": false,
      "name": "station",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🚊",
      "skin tone support": false,
      "name": "tram",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🚝",
      "skin tone support": false,
      "name": "monorail",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🚞",
      "skin tone support": false,
      "name": "mountain railway",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🚋",
      "skin tone support": false,
      "name": "tram car",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🚌",
      "skin tone support": false,
      "name": "bus",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🚍",
      "skin tone support": false,
      "name": "oncoming bus",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🚎",
      "skin tone support": false,
      "name": "trolleybus",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🚐",
      "skin tone support": false,
      "name": "minibus",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🚑",
      "skin tone support": false,
      "name": "ambulance",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🚒",
      "skin tone support": false,
      "name": "fire engine",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🚓",
      "skin tone support": false,
      "name": "police car",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🚔",
      "skin tone support": false,
      "name": "oncoming police car",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🚕",
      "skin tone support": false,
      "name": "taxi",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🚖",
      "skin tone support": false,
      "name": "oncoming taxi",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🚗",
      "skin tone support": false,
      "name": "automobile",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🚘",
      "skin tone support": false,
      "name": "oncoming automobile",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🚙",
      "skin tone support": false,
      "name": "sport utility vehicle",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🚚",
      "skin tone support": false,
      "name": "delivery truck",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🚛",
      "skin tone support": false,
      "name": "articulated lorry",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🚜",
      "skin tone support": false,
      "name": "tractor",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🏎️",
      "skin tone support": false,
      "name": "racing car",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🏍️",
      "skin tone support": false,
      "name": "motorcycle",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🛵",
      "skin tone support": false,
      "name": "motor scooter",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🦽",
      "skin tone support": false,
      "name": "manual wheelchair",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "🦼",
      "skin tone support": false,
      "name": "motorized wheelchair",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "🛺",
      "skin tone support": false,
      "name": "auto rickshaw",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "🚲",
      "skin tone support": false,
      "name": "bicycle",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🛴",
      "skin tone support": false,
      "name": "kick scooter",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🛹",
      "skin tone support": false,
      "name": "skateboard",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "🚏",
      "skin tone support": false,
      "name": "bus stop",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🛣️",
      "skin tone support": false,
      "name": "motorway",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🛤️",
      "skin tone support": false,
      "name": "railway track",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🛢️",
      "skin tone support": false,
      "name": "oil drum",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "⛽",
      "skin tone support": false,
      "name": "fuel pump",
      "unicode version": "5.2",
      "emoji version": "2.0"
    },
    {
      "emoji": "🚨",
      "skin tone support": false,
      "name": "police car light",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🚥",
      "skin tone support": false,
      "name": "horizontal traffic light",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🚦",
      "skin tone support": false,
      "name": "vertical traffic light",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🛑",
      "skin tone support": false,
      "name": "stop sign",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🚧",
      "skin tone support": false,
      "name": "construction",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "⚓",
      "skin tone support": false,
      "name": "anchor",
      "unicode version": "4.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "⛵",
      "skin tone support": false,
      "name": "sailboat",
      "unicode version": "5.2",
      "emoji version": "2.0"
    },
    {
      "emoji": "🛶",
      "skin tone support": false,
      "name": "canoe",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🚤",
      "skin tone support": false,
      "name": "speedboat",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🛳️",
      "skin tone support": false,
      "name": "passenger ship",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "⛴️",
      "skin tone support": false,
      "name": "ferry",
      "unicode version": "5.2",
      "emoji version": "2.0"
    },
    {
      "emoji": "🛥️",
      "skin tone support": false,
      "name": "motor boat",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🚢",
      "skin tone support": false,
      "name": "ship",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "✈️",
      "skin tone support": false,
      "name": "airplane",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "🛩️",
      "skin tone support": false,
      "name": "small airplane",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🛫",
      "skin tone support": false,
      "name": "airplane departure",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🛬",
      "skin tone support": false,
      "name": "airplane arrival",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🪂",
      "skin tone support": false,
      "name": "parachute",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "💺",
      "skin tone support": false,
      "name": "seat",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🚁",
      "skin tone support": false,
      "name": "helicopter",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🚟",
      "skin tone support": false,
      "name": "suspension railway",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🚠",
      "skin tone support": false,
      "name": "mountain cableway",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🚡",
      "skin tone support": false,
      "name": "aerial tramway",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🛰️",
      "skin tone support": false,
      "name": "satellite",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🚀",
      "skin tone support": false,
      "name": "rocket",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🛸",
      "skin tone support": false,
      "name": "flying saucer",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🛎️",
      "skin tone support": false,
      "name": "bellhop bell",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🧳",
      "skin tone support": false,
      "name": "luggage",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "⌛",
      "skin tone support": false,
      "name": "hourglass done",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "⏳",
      "skin tone support": false,
      "name": "hourglass not done",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "⌚",
      "skin tone support": false,
      "name": "watch",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "⏰",
      "skin tone support": false,
      "name": "alarm clock",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "⏱️",
      "skin tone support": false,
      "name": "stopwatch",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "⏲️",
      "skin tone support": false,
      "name": "timer clock",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🕰️",
      "skin tone support": false,
      "name": "mantelpiece clock",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🕛",
      "skin tone support": false,
      "name": "twelve o clock",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🕧",
      "skin tone support": false,
      "name": "twelve thirty",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🕐",
      "skin tone support": false,
      "name": "one o clock",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🕜",
      "skin tone support": false,
      "name": "one thirty",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🕑",
      "skin tone support": false,
      "name": "two o clock",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🕝",
      "skin tone support": false,
      "name": "two thirty",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🕒",
      "skin tone support": false,
      "name": "three o clock",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🕞",
      "skin tone support": false,
      "name": "three thirty",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🕓",
      "skin tone support": false,
      "name": "four o clock",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🕟",
      "skin tone support": false,
      "name": "four thirty",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🕔",
      "skin tone support": false,
      "name": "five o clock",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🕠",
      "skin tone support": false,
      "name": "five thirty",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🕕",
      "skin tone support": false,
      "name": "six o clock",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🕡",
      "skin tone support": false,
      "name": "six thirty",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🕖",
      "skin tone support": false,
      "name": "seven o clock",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🕢",
      "skin tone support": false,
      "name": "seven thirty",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🕗",
      "skin tone support": false,
      "name": "eight o clock",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🕣",
      "skin tone support": false,
      "name": "eight thirty",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🕘",
      "skin tone support": false,
      "name": "nine o clock",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🕤",
      "skin tone support": false,
      "name": "nine thirty",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🕙",
      "skin tone support": false,
      "name": "ten o clock",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🕥",
      "skin tone support": false,
      "name": "ten thirty",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🕚",
      "skin tone support": false,
      "name": "eleven o clock",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🕦",
      "skin tone support": false,
      "name": "eleven thirty",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🌑",
      "skin tone support": false,
      "name": "new moon",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🌒",
      "skin tone support": false,
      "name": "waxing crescent moon",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🌓",
      "skin tone support": false,
      "name": "first quarter moon",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🌔",
      "skin tone support": false,
      "name": "waxing gibbous moon",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🌕",
      "skin tone support": false,
      "name": "full moon",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🌖",
      "skin tone support": false,
      "name": "waning gibbous moon",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🌗",
      "skin tone support": false,
      "name": "last quarter moon",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🌘",
      "skin tone support": false,
      "name": "waning crescent moon",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🌙",
      "skin tone support": false,
      "name": "crescent moon",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🌚",
      "skin tone support": false,
      "name": "new moon face",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🌛",
      "skin tone support": false,
      "name": "first quarter moon face",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🌜",
      "skin tone support": false,
      "name": "last quarter moon face",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🌡️",
      "skin tone support": false,
      "name": "thermometer",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "☀️",
      "skin tone support": false,
      "name": "sun",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "🌝",
      "skin tone support": false,
      "name": "full moon face",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🌞",
      "skin tone support": false,
      "name": "sun with face",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🪐",
      "skin tone support": false,
      "name": "ringed planet",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "⭐",
      "skin tone support": false,
      "name": "star",
      "unicode version": "5.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "🌟",
      "skin tone support": false,
      "name": "glowing star",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🌠",
      "skin tone support": false,
      "name": "shooting star",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🌌",
      "skin tone support": false,
      "name": "milky way",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "☁️",
      "skin tone support": false,
      "name": "cloud",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "⛅",
      "skin tone support": false,
      "name": "sun behind cloud",
      "unicode version": "5.2",
      "emoji version": "2.0"
    },
    {
      "emoji": "⛈️",
      "skin tone support": false,
      "name": "cloud with lightning and rain",
      "unicode version": "5.2",
      "emoji version": "2.0"
    },
    {
      "emoji": "🌤️",
      "skin tone support": false,
      "name": "sun behind small cloud",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🌥️",
      "skin tone support": false,
      "name": "sun behind large cloud",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🌦️",
      "skin tone support": false,
      "name": "sun behind rain cloud",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🌧️",
      "skin tone support": false,
      "name": "cloud with rain",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🌨️",
      "skin tone support": false,
      "name": "cloud with snow",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🌩️",
      "skin tone support": false,
      "name": "cloud with lightning",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🌪️",
      "skin tone support": false,
      "name": "tornado",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🌫️",
      "skin tone support": false,
      "name": "fog",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🌬️",
      "skin tone support": false,
      "name": "wind face",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🌀",
      "skin tone support": false,
      "name": "cyclone",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🌈",
      "skin tone support": false,
      "name": "rainbow",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🌂",
      "skin tone support": false,
      "name": "closed umbrella",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "☂️",
      "skin tone support": false,
      "name": "umbrella",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "☔",
      "skin tone support": false,
      "name": "umbrella with rain drops",
      "unicode version": "4.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "⛱️",
      "skin tone support": false,
      "name": "umbrella on ground",
      "unicode version": "5.2",
      "emoji version": "2.0"
    },
    {
      "emoji": "⚡",
      "skin tone support": false,
      "name": "high voltage",
      "unicode version": "4.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "❄️",
      "skin tone support": false,
      "name": "snowflake",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "☃️",
      "skin tone support": false,
      "name": "snowman",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "⛄",
      "skin tone support": false,
      "name": "snowman without snow",
      "unicode version": "5.2",
      "emoji version": "2.0"
    },
    {
      "emoji": "☄️",
      "skin tone support": false,
      "name": "comet",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "🔥",
      "skin tone support": false,
      "name": "fire",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "💧",
      "skin tone support": false,
      "name": "droplet",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🌊",
      "skin tone support": false,
      "name": "water wave",
      "unicode version": "6.0",
      "emoji version": "2.0"
    }
  ]
	},
	{
	"name": "Activities",
	"emojis": [
    {
      "emoji": "🎃",
      "skin tone support": false,
      "name": "jack o lantern",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🎄",
      "skin tone support": false,
      "name": "christmas tree",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🎆",
      "skin tone support": false,
      "name": "fireworks",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🎇",
      "skin tone support": false,
      "name": "sparkler",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🧨",
      "skin tone support": false,
      "name": "firecracker",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "✨",
      "skin tone support": false,
      "name": "sparkles",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🎈",
      "skin tone support": false,
      "name": "balloon",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🎉",
      "skin tone support": false,
      "name": "party popper",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🎊",
      "skin tone support": false,
      "name": "confetti ball",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🎋",
      "skin tone support": false,
      "name": "tanabata tree",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🎍",
      "skin tone support": false,
      "name": "pine decoration",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🎎",
      "skin tone support": false,
      "name": "japanese dolls",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🎏",
      "skin tone support": false,
      "name": "carp streamer",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🎐",
      "skin tone support": false,
      "name": "wind chime",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🎑",
      "skin tone support": false,
      "name": "moon viewing ceremony",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🧧",
      "skin tone support": false,
      "name": "red envelope",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "🎀",
      "skin tone support": false,
      "name": "ribbon",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🎁",
      "skin tone support": false,
      "name": "wrapped gift",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🎗️",
      "skin tone support": false,
      "name": "reminder ribbon",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🎟️",
      "skin tone support": false,
      "name": "admission tickets",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🎫",
      "skin tone support": false,
      "name": "ticket",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🎖️",
      "skin tone support": false,
      "name": "military medal",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🏆",
      "skin tone support": false,
      "name": "trophy",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🏅",
      "skin tone support": false,
      "name": "sports medal",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🥇",
      "skin tone support": false,
      "name": "1st place medal",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🥈",
      "skin tone support": false,
      "name": "2nd place medal",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🥉",
      "skin tone support": false,
      "name": "3rd place medal",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "⚽",
      "skin tone support": false,
      "name": "soccer ball",
      "unicode version": "5.2",
      "emoji version": "2.0"
    },
    {
      "emoji": "⚾",
      "skin tone support": false,
      "name": "baseball",
      "unicode version": "5.2",
      "emoji version": "2.0"
    },
    {
      "emoji": "🥎",
      "skin tone support": false,
      "name": "softball",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "🏀",
      "skin tone support": false,
      "name": "basketball",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🏐",
      "skin tone support": false,
      "name": "volleyball",
      "unicode version": "8.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🏈",
      "skin tone support": false,
      "name": "american football",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🏉",
      "skin tone support": false,
      "name": "rugby football",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🎾",
      "skin tone support": false,
      "name": "tennis",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🥏",
      "skin tone support": false,
      "name": "flying disc",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "🎳",
      "skin tone support": false,
      "name": "bowling",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🏏",
      "skin tone support": false,
      "name": "cricket game",
      "unicode version": "8.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🏑",
      "skin tone support": false,
      "name": "field hockey",
      "unicode version": "8.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🏒",
      "skin tone support": false,
      "name": "ice hockey",
      "unicode version": "8.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🥍",
      "skin tone support": false,
      "name": "lacrosse",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "🏓",
      "skin tone support": false,
      "name": "ping pong",
      "unicode version": "8.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🏸",
      "skin tone support": false,
      "name": "badminton",
      "unicode version": "8.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🥊",
      "skin tone support": false,
      "name": "boxing glove",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🥋",
      "skin tone support": false,
      "name": "martial arts uniform",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🥅",
      "skin tone support": false,
      "name": "goal net",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "⛳",
      "skin tone support": false,
      "name": "flag in hole",
      "unicode version": "5.2",
      "emoji version": "2.0"
    },
    {
      "emoji": "⛸️",
      "skin tone support": false,
      "name": "ice skate",
      "unicode version": "5.2",
      "emoji version": "2.0"
    },
    {
      "emoji": "🎣",
      "skin tone support": false,
      "name": "fishing pole",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🤿",
      "skin tone support": false,
      "name": "diving mask",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "🎽",
      "skin tone support": false,
      "name": "running shirt",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🎿",
      "skin tone support": false,
      "name": "skis",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🛷",
      "skin tone support": false,
      "name": "sled",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🥌",
      "skin tone support": false,
      "name": "curling stone",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🎯",
      "skin tone support": false,
      "name": "direct hit",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🪀",
      "skin tone support": false,
      "name": "yo yo",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "🪁",
      "skin tone support": false,
      "name": "kite",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "🎱",
      "skin tone support": false,
      "name": "pool 8 ball",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🔮",
      "skin tone support": false,
      "name": "crystal ball",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🧿",
      "skin tone support": false,
      "name": "nazar amulet",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "🎮",
      "skin tone support": false,
      "name": "video game",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🕹️",
      "skin tone support": false,
      "name": "joystick",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🎰",
      "skin tone support": false,
      "name": "slot machine",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🎲",
      "skin tone support": false,
      "name": "game die",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🧩",
      "skin tone support": false,
      "name": "puzzle piece",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "🧸",
      "skin tone support": false,
      "name": "teddy bear",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "♠️",
      "skin tone support": false,
      "name": "spade suit",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "♥️",
      "skin tone support": false,
      "name": "heart suit",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "♦️",
      "skin tone support": false,
      "name": "diamond suit",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "♣️",
      "skin tone support": false,
      "name": "club suit",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "♟️",
      "skin tone support": false,
      "name": "chess pawn",
      "unicode version": "1.1",
      "emoji version": "11.0"
    },
    {
      "emoji": "🃏",
      "skin tone support": false,
      "name": "joker",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🀄",
      "skin tone support": false,
      "name": "mahjong red dragon",
      "unicode version": "5.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "🎴",
      "skin tone support": false,
      "name": "flower playing cards",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🎭",
      "skin tone support": false,
      "name": "performing arts",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🖼️",
      "skin tone support": false,
      "name": "framed picture",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🎨",
      "skin tone support": false,
      "name": "artist palette",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🧵",
      "skin tone support": false,
      "name": "thread",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "🧶",
      "skin tone support": false,
      "name": "yarn",
      "unicode version": "11.0",
      "emoji version": "11.0"
    }
  ]
	},
	{
	"name": "Objects",
	"emojis": [
    {
      "emoji": "👓",
      "skin tone support": false,
      "name": "glasses",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🕶️",
      "skin tone support": false,
      "name": "sunglasses",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🥽",
      "skin tone support": false,
      "name": "goggles",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "🥼",
      "skin tone support": false,
      "name": "lab coat",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "🦺",
      "skin tone support": false,
      "name": "safety vest",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "👔",
      "skin tone support": false,
      "name": "necktie",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👕",
      "skin tone support": false,
      "name": "t shirt",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👖",
      "skin tone support": false,
      "name": "jeans",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🧣",
      "skin tone support": false,
      "name": "scarf",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🧤",
      "skin tone support": false,
      "name": "gloves",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🧥",
      "skin tone support": false,
      "name": "coat",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🧦",
      "skin tone support": false,
      "name": "socks",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "👗",
      "skin tone support": false,
      "name": "dress",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👘",
      "skin tone support": false,
      "name": "kimono",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🥻",
      "skin tone support": false,
      "name": "sari",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "🩱",
      "skin tone support": false,
      "name": "one piece swimsuit",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "🩲",
      "skin tone support": false,
      "name": "briefs",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "🩳",
      "skin tone support": false,
      "name": "shorts",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "👙",
      "skin tone support": false,
      "name": "bikini",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👚",
      "skin tone support": false,
      "name": "woman s clothes",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👛",
      "skin tone support": false,
      "name": "purse",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👜",
      "skin tone support": false,
      "name": "handbag",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👝",
      "skin tone support": false,
      "name": "clutch bag",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🛍️",
      "skin tone support": false,
      "name": "shopping bags",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🎒",
      "skin tone support": false,
      "name": "backpack",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👞",
      "skin tone support": false,
      "name": "man s shoe",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👟",
      "skin tone support": false,
      "name": "running shoe",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🥾",
      "skin tone support": false,
      "name": "hiking boot",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "🥿",
      "skin tone support": false,
      "name": "flat shoe",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "👠",
      "skin tone support": false,
      "name": "high heeled shoe",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👡",
      "skin tone support": false,
      "name": "woman s sandal",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🩰",
      "skin tone support": false,
      "name": "ballet shoes",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "👢",
      "skin tone support": false,
      "name": "woman s boot",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👑",
      "skin tone support": false,
      "name": "crown",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "👒",
      "skin tone support": false,
      "name": "woman s hat",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🎩",
      "skin tone support": false,
      "name": "top hat",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🎓",
      "skin tone support": false,
      "name": "graduation cap",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🧢",
      "skin tone support": false,
      "name": "billed cap",
      "unicode version": "10.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "⛑️",
      "skin tone support": false,
      "name": "rescue worker s helmet",
      "unicode version": "5.2",
      "emoji version": "2.0"
    },
    {
      "emoji": "📿",
      "skin tone support": false,
      "name": "prayer beads",
      "unicode version": "8.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "💄",
      "skin tone support": false,
      "name": "lipstick",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "💍",
      "skin tone support": false,
      "name": "ring",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "💎",
      "skin tone support": false,
      "name": "gem stone",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🔇",
      "skin tone support": false,
      "name": "muted speaker",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🔈",
      "skin tone support": false,
      "name": "speaker low volume",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🔉",
      "skin tone support": false,
      "name": "speaker medium volume",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🔊",
      "skin tone support": false,
      "name": "speaker high volume",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "📢",
      "skin tone support": false,
      "name": "loudspeaker",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "📣",
      "skin tone support": false,
      "name": "megaphone",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "📯",
      "skin tone support": false,
      "name": "postal horn",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🔔",
      "skin tone support": false,
      "name": "bell",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🔕",
      "skin tone support": false,
      "name": "bell with slash",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🎼",
      "skin tone support": false,
      "name": "musical score",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🎵",
      "skin tone support": false,
      "name": "musical note",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🎶",
      "skin tone support": false,
      "name": "musical notes",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🎙️",
      "skin tone support": false,
      "name": "studio microphone",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🎚️",
      "skin tone support": false,
      "name": "level slider",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🎛️",
      "skin tone support": false,
      "name": "control knobs",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🎤",
      "skin tone support": false,
      "name": "microphone",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🎧",
      "skin tone support": false,
      "name": "headphone",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "📻",
      "skin tone support": false,
      "name": "radio",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🎷",
      "skin tone support": false,
      "name": "saxophone",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🎸",
      "skin tone support": false,
      "name": "guitar",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🎹",
      "skin tone support": false,
      "name": "musical keyboard",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🎺",
      "skin tone support": false,
      "name": "trumpet",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🎻",
      "skin tone support": false,
      "name": "violin",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🪕",
      "skin tone support": false,
      "name": "banjo",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "🥁",
      "skin tone support": false,
      "name": "drum",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "📱",
      "skin tone support": false,
      "name": "mobile phone",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "📲",
      "skin tone support": false,
      "name": "mobile phone with arrow",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "☎️",
      "skin tone support": false,
      "name": "telephone",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "📞",
      "skin tone support": false,
      "name": "telephone receiver",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "📟",
      "skin tone support": false,
      "name": "pager",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "📠",
      "skin tone support": false,
      "name": "fax machine",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🔋",
      "skin tone support": false,
      "name": "battery",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🔌",
      "skin tone support": false,
      "name": "electric plug",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "💻",
      "skin tone support": false,
      "name": "laptop",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🖥️",
      "skin tone support": false,
      "name": "desktop computer",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🖨️",
      "skin tone support": false,
      "name": "printer",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "⌨️",
      "skin tone support": false,
      "name": "keyboard",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "🖱️",
      "skin tone support": false,
      "name": "computer mouse",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🖲️",
      "skin tone support": false,
      "name": "trackball",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "💽",
      "skin tone support": false,
      "name": "computer disk",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "💾",
      "skin tone support": false,
      "name": "floppy disk",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "💿",
      "skin tone support": false,
      "name": "optical disk",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "📀",
      "skin tone support": false,
      "name": "dvd",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🧮",
      "skin tone support": false,
      "name": "abacus",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "🎥",
      "skin tone support": false,
      "name": "movie camera",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🎞️",
      "skin tone support": false,
      "name": "film frames",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "📽️",
      "skin tone support": false,
      "name": "film projector",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🎬",
      "skin tone support": false,
      "name": "clapper board",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "📺",
      "skin tone support": false,
      "name": "television",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "📷",
      "skin tone support": false,
      "name": "camera",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "📸",
      "skin tone support": false,
      "name": "camera with flash",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "📹",
      "skin tone support": false,
      "name": "video camera",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "📼",
      "skin tone support": false,
      "name": "videocassette",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🔍",
      "skin tone support": false,
      "name": "magnifying glass tilted left",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🔎",
      "skin tone support": false,
      "name": "magnifying glass tilted right",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🕯️",
      "skin tone support": false,
      "name": "candle",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "💡",
      "skin tone support": false,
      "name": "light bulb",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🔦",
      "skin tone support": false,
      "name": "flashlight",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🏮",
      "skin tone support": false,
      "name": "red paper lantern",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🪔",
      "skin tone support": false,
      "name": "diya lamp",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "📔",
      "skin tone support": false,
      "name": "notebook with decorative cover",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "📕",
      "skin tone support": false,
      "name": "closed book",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "📖",
      "skin tone support": false,
      "name": "open book",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "📗",
      "skin tone support": false,
      "name": "green book",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "📘",
      "skin tone support": false,
      "name": "blue book",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "📙",
      "skin tone support": false,
      "name": "orange book",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "📚",
      "skin tone support": false,
      "name": "books",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "📓",
      "skin tone support": false,
      "name": "notebook",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "📒",
      "skin tone support": false,
      "name": "ledger",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "📃",
      "skin tone support": false,
      "name": "page with curl",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "📜",
      "skin tone support": false,
      "name": "scroll",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "📄",
      "skin tone support": false,
      "name": "page facing up",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "📰",
      "skin tone support": false,
      "name": "newspaper",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🗞️",
      "skin tone support": false,
      "name": "rolled up newspaper",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "📑",
      "skin tone support": false,
      "name": "bookmark tabs",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🔖",
      "skin tone support": false,
      "name": "bookmark",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🏷️",
      "skin tone support": false,
      "name": "label",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "💰",
      "skin tone support": false,
      "name": "money bag",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "💴",
      "skin tone support": false,
      "name": "yen banknote",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "💵",
      "skin tone support": false,
      "name": "dollar banknote",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "💶",
      "skin tone support": false,
      "name": "euro banknote",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "💷",
      "skin tone support": false,
      "name": "pound banknote",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "💸",
      "skin tone support": false,
      "name": "money with wings",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "💳",
      "skin tone support": false,
      "name": "credit card",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🧾",
      "skin tone support": false,
      "name": "receipt",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "💹",
      "skin tone support": false,
      "name": "chart increasing with yen",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "💱",
      "skin tone support": false,
      "name": "currency exchange",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "💲",
      "skin tone support": false,
      "name": "heavy dollar sign",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "✉️",
      "skin tone support": false,
      "name": "envelope",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "📧",
      "skin tone support": false,
      "name": "e mail",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "📨",
      "skin tone support": false,
      "name": "incoming envelope",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "📩",
      "skin tone support": false,
      "name": "envelope with arrow",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "📤",
      "skin tone support": false,
      "name": "outbox tray",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "📥",
      "skin tone support": false,
      "name": "inbox tray",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "📦",
      "skin tone support": false,
      "name": "package",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "📫",
      "skin tone support": false,
      "name": "closed mailbox with raised flag",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "📪",
      "skin tone support": false,
      "name": "closed mailbox with lowered flag",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "📬",
      "skin tone support": false,
      "name": "open mailbox with raised flag",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "📭",
      "skin tone support": false,
      "name": "open mailbox with lowered flag",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "📮",
      "skin tone support": false,
      "name": "postbox",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🗳️",
      "skin tone support": false,
      "name": "ballot box with ballot",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "✏️",
      "skin tone support": false,
      "name": "pencil",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "✒️",
      "skin tone support": false,
      "name": "black nib",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "🖋️",
      "skin tone support": false,
      "name": "fountain pen",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🖊️",
      "skin tone support": false,
      "name": "pen",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🖌️",
      "skin tone support": false,
      "name": "paintbrush",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🖍️",
      "skin tone support": false,
      "name": "crayon",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "📝",
      "skin tone support": false,
      "name": "memo",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "💼",
      "skin tone support": false,
      "name": "briefcase",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "📁",
      "skin tone support": false,
      "name": "file folder",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "📂",
      "skin tone support": false,
      "name": "open file folder",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🗂️",
      "skin tone support": false,
      "name": "card index dividers",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "📅",
      "skin tone support": false,
      "name": "calendar",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "📆",
      "skin tone support": false,
      "name": "tear off calendar",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🗒️",
      "skin tone support": false,
      "name": "spiral notepad",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🗓️",
      "skin tone support": false,
      "name": "spiral calendar",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "📇",
      "skin tone support": false,
      "name": "card index",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "📈",
      "skin tone support": false,
      "name": "chart increasing",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "📉",
      "skin tone support": false,
      "name": "chart decreasing",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "📊",
      "skin tone support": false,
      "name": "bar chart",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "📋",
      "skin tone support": false,
      "name": "clipboard",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "📌",
      "skin tone support": false,
      "name": "pushpin",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "📍",
      "skin tone support": false,
      "name": "round pushpin",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "📎",
      "skin tone support": false,
      "name": "paperclip",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🖇️",
      "skin tone support": false,
      "name": "linked paperclips",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "📏",
      "skin tone support": false,
      "name": "straight ruler",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "📐",
      "skin tone support": false,
      "name": "triangular ruler",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "✂️",
      "skin tone support": false,
      "name": "scissors",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "🗃️",
      "skin tone support": false,
      "name": "card file box",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🗄️",
      "skin tone support": false,
      "name": "file cabinet",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🗑️",
      "skin tone support": false,
      "name": "wastebasket",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🔒",
      "skin tone support": false,
      "name": "locked",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🔓",
      "skin tone support": false,
      "name": "unlocked",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🔏",
      "skin tone support": false,
      "name": "locked with pen",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🔐",
      "skin tone support": false,
      "name": "locked with key",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🔑",
      "skin tone support": false,
      "name": "key",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🗝️",
      "skin tone support": false,
      "name": "old key",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🔨",
      "skin tone support": false,
      "name": "hammer",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🪓",
      "skin tone support": false,
      "name": "axe",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "⛏️",
      "skin tone support": false,
      "name": "pick",
      "unicode version": "5.2",
      "emoji version": "2.0"
    },
    {
      "emoji": "⚒️",
      "skin tone support": false,
      "name": "hammer and pick",
      "unicode version": "4.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "🛠️",
      "skin tone support": false,
      "name": "hammer and wrench",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🗡️",
      "skin tone support": false,
      "name": "dagger",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "⚔️",
      "skin tone support": false,
      "name": "crossed swords",
      "unicode version": "4.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "🔫",
      "skin tone support": false,
      "name": "pistol",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🏹",
      "skin tone support": false,
      "name": "bow and arrow",
      "unicode version": "8.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🛡️",
      "skin tone support": false,
      "name": "shield",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🔧",
      "skin tone support": false,
      "name": "wrench",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🔩",
      "skin tone support": false,
      "name": "nut and bolt",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "⚙️",
      "skin tone support": false,
      "name": "gear",
      "unicode version": "4.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "🗜️",
      "skin tone support": false,
      "name": "clamp",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "⚖️",
      "skin tone support": false,
      "name": "balance scale",
      "unicode version": "4.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "🦯",
      "skin tone support": false,
      "name": "probing cane",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "🔗",
      "skin tone support": false,
      "name": "link",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "⛓️",
      "skin tone support": false,
      "name": "chains",
      "unicode version": "5.2",
      "emoji version": "2.0"
    },
    {
      "emoji": "🧰",
      "skin tone support": false,
      "name": "toolbox",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "🧲",
      "skin tone support": false,
      "name": "magnet",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "⚗️",
      "skin tone support": false,
      "name": "alembic",
      "unicode version": "4.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "🧪",
      "skin tone support": false,
      "name": "test tube",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "🧫",
      "skin tone support": false,
      "name": "petri dish",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "🧬",
      "skin tone support": false,
      "name": "dna",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "🔬",
      "skin tone support": false,
      "name": "microscope",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🔭",
      "skin tone support": false,
      "name": "telescope",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "📡",
      "skin tone support": false,
      "name": "satellite antenna",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "💉",
      "skin tone support": false,
      "name": "syringe",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🩸",
      "skin tone support": false,
      "name": "drop of blood",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "💊",
      "skin tone support": false,
      "name": "pill",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🩹",
      "skin tone support": false,
      "name": "adhesive bandage",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "🩺",
      "skin tone support": false,
      "name": "stethoscope",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "🚪",
      "skin tone support": false,
      "name": "door",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🛏️",
      "skin tone support": false,
      "name": "bed",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🛋️",
      "skin tone support": false,
      "name": "couch and lamp",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🪑",
      "skin tone support": false,
      "name": "chair",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "🚽",
      "skin tone support": false,
      "name": "toilet",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🚿",
      "skin tone support": false,
      "name": "shower",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🛁",
      "skin tone support": false,
      "name": "bathtub",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🪒",
      "skin tone support": false,
      "name": "razor",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "🧴",
      "skin tone support": false,
      "name": "lotion bottle",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "🧷",
      "skin tone support": false,
      "name": "safety pin",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "🧹",
      "skin tone support": false,
      "name": "broom",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "🧺",
      "skin tone support": false,
      "name": "basket",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "🧻",
      "skin tone support": false,
      "name": "roll of paper",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "🧼",
      "skin tone support": false,
      "name": "soap",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "🧽",
      "skin tone support": false,
      "name": "sponge",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "🧯",
      "skin tone support": false,
      "name": "fire extinguisher",
      "unicode version": "11.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "🛒",
      "skin tone support": false,
      "name": "shopping cart",
      "unicode version": "9.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🚬",
      "skin tone support": false,
      "name": "cigarette",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "⚰️",
      "skin tone support": false,
      "name": "coffin",
      "unicode version": "4.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "⚱️",
      "skin tone support": false,
      "name": "funeral urn",
      "unicode version": "4.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "🗿",
      "skin tone support": false,
      "name": "moai",
      "unicode version": "6.0",
      "emoji version": "2.0"
    }
  ]
	},
	{
	"name": "Symbols",
	"emojis": [
    {
      "emoji": "🏧",
      "skin tone support": false,
      "name": "atm sign",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🚮",
      "skin tone support": false,
      "name": "litter in bin sign",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🚰",
      "skin tone support": false,
      "name": "potable water",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "♿",
      "skin tone support": false,
      "name": "wheelchair symbol",
      "unicode version": "4.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "🚹",
      "skin tone support": false,
      "name": "men s room",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🚺",
      "skin tone support": false,
      "name": "women s room",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🚻",
      "skin tone support": false,
      "name": "restroom",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🚼",
      "skin tone support": false,
      "name": "baby symbol",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🚾",
      "skin tone support": false,
      "name": "water closet",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🛂",
      "skin tone support": false,
      "name": "passport control",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🛃",
      "skin tone support": false,
      "name": "customs",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🛄",
      "skin tone support": false,
      "name": "baggage claim",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🛅",
      "skin tone support": false,
      "name": "left luggage",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "⚠️",
      "skin tone support": false,
      "name": "warning",
      "unicode version": "4.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🚸",
      "skin tone support": false,
      "name": "children crossing",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "⛔",
      "skin tone support": false,
      "name": "no entry",
      "unicode version": "5.2",
      "emoji version": "2.0"
    },
    {
      "emoji": "🚫",
      "skin tone support": false,
      "name": "prohibited",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🚳",
      "skin tone support": false,
      "name": "no bicycles",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🚭",
      "skin tone support": false,
      "name": "no smoking",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🚯",
      "skin tone support": false,
      "name": "no littering",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🚱",
      "skin tone support": false,
      "name": "non potable water",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🚷",
      "skin tone support": false,
      "name": "no pedestrians",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "📵",
      "skin tone support": false,
      "name": "no mobile phones",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🔞",
      "skin tone support": false,
      "name": "no one under eighteen",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "☢️",
      "skin tone support": false,
      "name": "radioactive",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "☣️",
      "skin tone support": false,
      "name": "biohazard",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "⬆️",
      "skin tone support": false,
      "name": "up arrow",
      "unicode version": "4.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "↗️",
      "skin tone support": false,
      "name": "up right arrow",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "➡️",
      "skin tone support": false,
      "name": "right arrow",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "↘️",
      "skin tone support": false,
      "name": "down right arrow",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "⬇️",
      "skin tone support": false,
      "name": "down arrow",
      "unicode version": "4.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "↙️",
      "skin tone support": false,
      "name": "down left arrow",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "⬅️",
      "skin tone support": false,
      "name": "left arrow",
      "unicode version": "4.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "↖️",
      "skin tone support": false,
      "name": "up left arrow",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "↕️",
      "skin tone support": false,
      "name": "up down arrow",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "↔️",
      "skin tone support": false,
      "name": "left right arrow",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "↩️",
      "skin tone support": false,
      "name": "right arrow curving left",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "↪️",
      "skin tone support": false,
      "name": "left arrow curving right",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "⤴️",
      "skin tone support": false,
      "name": "right arrow curving up",
      "unicode version": "3.2",
      "emoji version": "2.0"
    },
    {
      "emoji": "⤵️",
      "skin tone support": false,
      "name": "right arrow curving down",
      "unicode version": "3.2",
      "emoji version": "2.0"
    },
    {
      "emoji": "🔃",
      "skin tone support": false,
      "name": "clockwise vertical arrows",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🔄",
      "skin tone support": false,
      "name": "counterclockwise arrows button",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🔙",
      "skin tone support": false,
      "name": "back arrow",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🔚",
      "skin tone support": false,
      "name": "end arrow",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🔛",
      "skin tone support": false,
      "name": "on arrow",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🔜",
      "skin tone support": false,
      "name": "soon arrow",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🔝",
      "skin tone support": false,
      "name": "top arrow",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🛐",
      "skin tone support": false,
      "name": "place of worship",
      "unicode version": "8.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "⚛️",
      "skin tone support": false,
      "name": "atom symbol",
      "unicode version": "4.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "🕉️",
      "skin tone support": false,
      "name": "om",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "✡️",
      "skin tone support": false,
      "name": "star of david",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "☸️",
      "skin tone support": false,
      "name": "wheel of dharma",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "☯️",
      "skin tone support": false,
      "name": "yin yang",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "✝️",
      "skin tone support": false,
      "name": "latin cross",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "☦️",
      "skin tone support": false,
      "name": "orthodox cross",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "☪️",
      "skin tone support": false,
      "name": "star and crescent",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "☮️",
      "skin tone support": false,
      "name": "peace symbol",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "🕎",
      "skin tone support": false,
      "name": "menorah",
      "unicode version": "8.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🔯",
      "skin tone support": false,
      "name": "dotted six pointed star",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "♈",
      "skin tone support": false,
      "name": "aries",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "♉",
      "skin tone support": false,
      "name": "taurus",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "♊",
      "skin tone support": false,
      "name": "gemini",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "♋",
      "skin tone support": false,
      "name": "cancer",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "♌",
      "skin tone support": false,
      "name": "leo",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "♍",
      "skin tone support": false,
      "name": "virgo",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "♎",
      "skin tone support": false,
      "name": "libra",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "♏",
      "skin tone support": false,
      "name": "scorpio",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "♐",
      "skin tone support": false,
      "name": "sagittarius",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "♑",
      "skin tone support": false,
      "name": "capricorn",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "♒",
      "skin tone support": false,
      "name": "aquarius",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "♓",
      "skin tone support": false,
      "name": "pisces",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "⛎",
      "skin tone support": false,
      "name": "ophiuchus",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🔀",
      "skin tone support": false,
      "name": "shuffle tracks button",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🔁",
      "skin tone support": false,
      "name": "repeat button",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🔂",
      "skin tone support": false,
      "name": "repeat single button",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "▶️",
      "skin tone support": false,
      "name": "play button",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "⏩",
      "skin tone support": false,
      "name": "fast forward button",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "⏭️",
      "skin tone support": false,
      "name": "next track button",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "⏯️",
      "skin tone support": false,
      "name": "play or pause button",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "◀️",
      "skin tone support": false,
      "name": "reverse button",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "⏪",
      "skin tone support": false,
      "name": "fast reverse button",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "⏮️",
      "skin tone support": false,
      "name": "last track button",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🔼",
      "skin tone support": false,
      "name": "upwards button",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "⏫",
      "skin tone support": false,
      "name": "fast up button",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🔽",
      "skin tone support": false,
      "name": "downwards button",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "⏬",
      "skin tone support": false,
      "name": "fast down button",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "⏸️",
      "skin tone support": false,
      "name": "pause button",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "⏹️",
      "skin tone support": false,
      "name": "stop button",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "⏺️",
      "skin tone support": false,
      "name": "record button",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "⏏️",
      "skin tone support": false,
      "name": "eject button",
      "unicode version": "4.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🎦",
      "skin tone support": false,
      "name": "cinema",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🔅",
      "skin tone support": false,
      "name": "dim button",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🔆",
      "skin tone support": false,
      "name": "bright button",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "📶",
      "skin tone support": false,
      "name": "antenna bars",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "📳",
      "skin tone support": false,
      "name": "vibration mode",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "📴",
      "skin tone support": false,
      "name": "mobile phone off",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "♀️",
      "skin tone support": false,
      "name": "female sign",
      "unicode version": "1.1",
      "emoji version": "4.0"
    },
    {
      "emoji": "♂️",
      "skin tone support": false,
      "name": "male sign",
      "unicode version": "1.1",
      "emoji version": "4.0"
    },
    {
      "emoji": "⚕️",
      "skin tone support": false,
      "name": "medical symbol",
      "unicode version": "4.1",
      "emoji version": "4.0"
    },
    {
      "emoji": "♾️",
      "skin tone support": false,
      "name": "infinity",
      "unicode version": "4.1",
      "emoji version": "11.0"
    },
    {
      "emoji": "♻️",
      "skin tone support": false,
      "name": "recycling symbol",
      "unicode version": "3.2",
      "emoji version": "2.0"
    },
    {
      "emoji": "⚜️",
      "skin tone support": false,
      "name": "fleur de lis",
      "unicode version": "4.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "🔱",
      "skin tone support": false,
      "name": "trident emblem",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "📛",
      "skin tone support": false,
      "name": "name badge",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🔰",
      "skin tone support": false,
      "name": "japanese symbol for beginner",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "⭕",
      "skin tone support": false,
      "name": "hollow red circle",
      "unicode version": "5.2",
      "emoji version": "2.0"
    },
    {
      "emoji": "✅",
      "skin tone support": false,
      "name": "check mark button",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "☑️",
      "skin tone support": false,
      "name": "check box with check",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "✔️",
      "skin tone support": false,
      "name": "check mark",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "✖️",
      "skin tone support": false,
      "name": "multiplication sign",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "❌",
      "skin tone support": false,
      "name": "cross mark",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "❎",
      "skin tone support": false,
      "name": "cross mark button",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "➕",
      "skin tone support": false,
      "name": "plus sign",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "➖",
      "skin tone support": false,
      "name": "minus sign",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "➗",
      "skin tone support": false,
      "name": "division sign",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "➰",
      "skin tone support": false,
      "name": "curly loop",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "➿",
      "skin tone support": false,
      "name": "double curly loop",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "〽️",
      "skin tone support": false,
      "name": "part alternation mark",
      "unicode version": "3.2",
      "emoji version": "2.0"
    },
    {
      "emoji": "✳️",
      "skin tone support": false,
      "name": "eight spoked asterisk",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "✴️",
      "skin tone support": false,
      "name": "eight pointed star",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "❇️",
      "skin tone support": false,
      "name": "sparkle",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "‼️",
      "skin tone support": false,
      "name": "double exclamation mark",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "⁉️",
      "skin tone support": false,
      "name": "exclamation question mark",
      "unicode version": "3.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "❓",
      "skin tone support": false,
      "name": "question mark",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "❔",
      "skin tone support": false,
      "name": "white question mark",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "❕",
      "skin tone support": false,
      "name": "white exclamation mark",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "❗",
      "skin tone support": false,
      "name": "exclamation mark",
      "unicode version": "5.2",
      "emoji version": "2.0"
    },
    {
      "emoji": "〰️",
      "skin tone support": false,
      "name": "wavy dash",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "©️",
      "skin tone support": false,
      "name": "copyright",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "®️",
      "skin tone support": false,
      "name": "registered",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "™️",
      "skin tone support": false,
      "name": "trade mark",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "#️⃣",
      "skin tone support": false,
      "name": "keycap ",
      "unicode version": "3.2",
      "emoji version": "0.0"
    },
    {
      "emoji": "*️⃣",
      "skin tone support": false,
      "name": "keycap ",
      "unicode version": "3.2",
      "emoji version": "0.0"
    },
    {
      "emoji": "0️⃣",
      "skin tone support": false,
      "name": "keycap 0",
      "unicode version": "3.2",
      "emoji version": "0.0"
    },
    {
      "emoji": "1️⃣",
      "skin tone support": false,
      "name": "keycap 1",
      "unicode version": "3.2",
      "emoji version": "0.0"
    },
    {
      "emoji": "2️⃣",
      "skin tone support": false,
      "name": "keycap 2",
      "unicode version": "3.2",
      "emoji version": "0.0"
    },
    {
      "emoji": "3️⃣",
      "skin tone support": false,
      "name": "keycap 3",
      "unicode version": "3.2",
      "emoji version": "0.0"
    },
    {
      "emoji": "4️⃣",
      "skin tone support": false,
      "name": "keycap 4",
      "unicode version": "3.2",
      "emoji version": "0.0"
    },
    {
      "emoji": "5️⃣",
      "skin tone support": false,
      "name": "keycap 5",
      "unicode version": "3.2",
      "emoji version": "0.0"
    },
    {
      "emoji": "6️⃣",
      "skin tone support": false,
      "name": "keycap 6",
      "unicode version": "3.2",
      "emoji version": "0.0"
    },
    {
      "emoji": "7️⃣",
      "skin tone support": false,
      "name": "keycap 7",
      "unicode version": "3.2",
      "emoji version": "0.0"
    },
    {
      "emoji": "8️⃣",
      "skin tone support": false,
      "name": "keycap 8",
      "unicode version": "3.2",
      "emoji version": "0.0"
    },
    {
      "emoji": "9️⃣",
      "skin tone support": false,
      "name": "keycap 9",
      "unicode version": "3.2",
      "emoji version": "0.0"
    },
    {
      "emoji": "🔟",
      "skin tone support": false,
      "name": "keycap 10",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🔠",
      "skin tone support": false,
      "name": "input latin uppercase",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🔡",
      "skin tone support": false,
      "name": "input latin lowercase",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🔢",
      "skin tone support": false,
      "name": "input numbers",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🔣",
      "skin tone support": false,
      "name": "input symbols",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🔤",
      "skin tone support": false,
      "name": "input latin letters",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🅰️",
      "skin tone support": false,
      "name": "a button",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🆎",
      "skin tone support": false,
      "name": "ab button",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🅱️",
      "skin tone support": false,
      "name": "b button",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🆑",
      "skin tone support": false,
      "name": "cl button",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🆒",
      "skin tone support": false,
      "name": "cool button",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🆓",
      "skin tone support": false,
      "name": "free button",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "ℹ️",
      "skin tone support": false,
      "name": "information",
      "unicode version": "3.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🆔",
      "skin tone support": false,
      "name": "id button",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "Ⓜ️",
      "skin tone support": false,
      "name": "circled m",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "🆕",
      "skin tone support": false,
      "name": "new button",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🆖",
      "skin tone support": false,
      "name": "ng button",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🅾️",
      "skin tone support": false,
      "name": "o button",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🆗",
      "skin tone support": false,
      "name": "ok button",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🅿️",
      "skin tone support": false,
      "name": "p button",
      "unicode version": "5.2",
      "emoji version": "2.0"
    },
    {
      "emoji": "🆘",
      "skin tone support": false,
      "name": "sos button",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🆙",
      "skin tone support": false,
      "name": "up button",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🆚",
      "skin tone support": false,
      "name": "vs button",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🈁",
      "skin tone support": false,
      "name": "japanese here button",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🈂️",
      "skin tone support": false,
      "name": "japanese service charge button",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🈷️",
      "skin tone support": false,
      "name": "japanese monthly amount button",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🈶",
      "skin tone support": false,
      "name": "japanese not free of charge button",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🈯",
      "skin tone support": false,
      "name": "japanese reserved button",
      "unicode version": "5.2",
      "emoji version": "2.0"
    },
    {
      "emoji": "🉐",
      "skin tone support": false,
      "name": "japanese bargain button",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🈹",
      "skin tone support": false,
      "name": "japanese discount button",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🈚",
      "skin tone support": false,
      "name": "japanese free of charge button",
      "unicode version": "5.2",
      "emoji version": "2.0"
    },
    {
      "emoji": "🈲",
      "skin tone support": false,
      "name": "japanese prohibited button",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🉑",
      "skin tone support": false,
      "name": "japanese acceptable button",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🈸",
      "skin tone support": false,
      "name": "japanese application button",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🈴",
      "skin tone support": false,
      "name": "japanese passing grade button",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🈳",
      "skin tone support": false,
      "name": "japanese vacancy button",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "㊗️",
      "skin tone support": false,
      "name": "japanese congratulations button",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "㊙️",
      "skin tone support": false,
      "name": "japanese secret button",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "🈺",
      "skin tone support": false,
      "name": "japanese open for business button",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🈵",
      "skin tone support": false,
      "name": "japanese no vacancy button",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🔴",
      "skin tone support": false,
      "name": "red circle",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🟠",
      "skin tone support": false,
      "name": "orange circle",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "🟡",
      "skin tone support": false,
      "name": "yellow circle",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "🟢",
      "skin tone support": false,
      "name": "green circle",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "🔵",
      "skin tone support": false,
      "name": "blue circle",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🟣",
      "skin tone support": false,
      "name": "purple circle",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "🟤",
      "skin tone support": false,
      "name": "brown circle",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "⚫",
      "skin tone support": false,
      "name": "black circle",
      "unicode version": "4.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "⚪",
      "skin tone support": false,
      "name": "white circle",
      "unicode version": "4.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "🟥",
      "skin tone support": false,
      "name": "red square",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "🟧",
      "skin tone support": false,
      "name": "orange square",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "🟨",
      "skin tone support": false,
      "name": "yellow square",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "🟩",
      "skin tone support": false,
      "name": "green square",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "🟦",
      "skin tone support": false,
      "name": "blue square",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "🟪",
      "skin tone support": false,
      "name": "purple square",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "🟫",
      "skin tone support": false,
      "name": "brown square",
      "unicode version": "12.0",
      "emoji version": "12.1"
    },
    {
      "emoji": "⬛",
      "skin tone support": false,
      "name": "black large square",
      "unicode version": "5.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "⬜",
      "skin tone support": false,
      "name": "white large square",
      "unicode version": "5.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "◼️",
      "skin tone support": false,
      "name": "black medium square",
      "unicode version": "3.2",
      "emoji version": "2.0"
    },
    {
      "emoji": "◻️",
      "skin tone support": false,
      "name": "white medium square",
      "unicode version": "3.2",
      "emoji version": "2.0"
    },
    {
      "emoji": "◾",
      "skin tone support": false,
      "name": "black medium small square",
      "unicode version": "3.2",
      "emoji version": "2.0"
    },
    {
      "emoji": "◽",
      "skin tone support": false,
      "name": "white medium small square",
      "unicode version": "3.2",
      "emoji version": "2.0"
    },
    {
      "emoji": "▪️",
      "skin tone support": false,
      "name": "black small square",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "▫️",
      "skin tone support": false,
      "name": "white small square",
      "unicode version": "1.1",
      "emoji version": "2.0"
    },
    {
      "emoji": "🔶",
      "skin tone support": false,
      "name": "large orange diamond",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🔷",
      "skin tone support": false,
      "name": "large blue diamond",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🔸",
      "skin tone support": false,
      "name": "small orange diamond",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🔹",
      "skin tone support": false,
      "name": "small blue diamond",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🔺",
      "skin tone support": false,
      "name": "red triangle pointed up",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🔻",
      "skin tone support": false,
      "name": "red triangle pointed down",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "💠",
      "skin tone support": false,
      "name": "diamond with a dot",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🔘",
      "skin tone support": false,
      "name": "radio button",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🔳",
      "skin tone support": false,
      "name": "white square button",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🔲",
      "skin tone support": false,
      "name": "black square button",
      "unicode version": "6.0",
      "emoji version": "2.0"
    }
  ]
	},
	{
	"name": "Flags",
	"emojis": [
    {
      "emoji": "🏁",
      "skin tone support": false,
      "name": "chequered flag",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🚩",
      "skin tone support": false,
      "name": "triangular flag",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🎌",
      "skin tone support": false,
      "name": "crossed flags",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🏴",
      "skin tone support": false,
      "name": "black flag",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🏳️",
      "skin tone support": false,
      "name": "white flag",
      "unicode version": "7.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🏳️‍🌈",
      "skin tone support": false,
      "name": "rainbow flag",
      "unicode version": "7.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🏴‍☠️",
      "skin tone support": false,
      "name": "pirate flag",
      "unicode version": "7.0",
      "emoji version": "11.0"
    },
    {
      "emoji": "🇦🇨",
      "skin tone support": false,
      "name": "flag ascension island",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇦🇩",
      "skin tone support": false,
      "name": "flag andorra",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇦🇪",
      "skin tone support": false,
      "name": "flag united arab emirates",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇦🇫",
      "skin tone support": false,
      "name": "flag afghanistan",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇦🇬",
      "skin tone support": false,
      "name": "flag antigua barbuda",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇦🇮",
      "skin tone support": false,
      "name": "flag anguilla",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇦🇱",
      "skin tone support": false,
      "name": "flag albania",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇦🇲",
      "skin tone support": false,
      "name": "flag armenia",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇦🇴",
      "skin tone support": false,
      "name": "flag angola",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇦🇶",
      "skin tone support": false,
      "name": "flag antarctica",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇦🇷",
      "skin tone support": false,
      "name": "flag argentina",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇦🇸",
      "skin tone support": false,
      "name": "flag american samoa",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇦🇹",
      "skin tone support": false,
      "name": "flag austria",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇦🇺",
      "skin tone support": false,
      "name": "flag australia",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇦🇼",
      "skin tone support": false,
      "name": "flag aruba",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇦🇽",
      "skin tone support": false,
      "name": "flag land islands",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇦🇿",
      "skin tone support": false,
      "name": "flag azerbaijan",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇧🇦",
      "skin tone support": false,
      "name": "flag bosnia herzegovina",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇧🇧",
      "skin tone support": false,
      "name": "flag barbados",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇧🇩",
      "skin tone support": false,
      "name": "flag bangladesh",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇧🇪",
      "skin tone support": false,
      "name": "flag belgium",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇧🇫",
      "skin tone support": false,
      "name": "flag burkina faso",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇧🇬",
      "skin tone support": false,
      "name": "flag bulgaria",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇧🇭",
      "skin tone support": false,
      "name": "flag bahrain",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇧🇮",
      "skin tone support": false,
      "name": "flag burundi",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇧🇯",
      "skin tone support": false,
      "name": "flag benin",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇧🇱",
      "skin tone support": false,
      "name": "flag st barth lemy",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇧🇲",
      "skin tone support": false,
      "name": "flag bermuda",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇧🇳",
      "skin tone support": false,
      "name": "flag brunei",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇧🇴",
      "skin tone support": false,
      "name": "flag bolivia",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇧🇶",
      "skin tone support": false,
      "name": "flag caribbean netherlands",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇧🇷",
      "skin tone support": false,
      "name": "flag brazil",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇧🇸",
      "skin tone support": false,
      "name": "flag bahamas",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇧🇹",
      "skin tone support": false,
      "name": "flag bhutan",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇧🇻",
      "skin tone support": false,
      "name": "flag bouvet island",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇧🇼",
      "skin tone support": false,
      "name": "flag botswana",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇧🇾",
      "skin tone support": false,
      "name": "flag belarus",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇧🇿",
      "skin tone support": false,
      "name": "flag belize",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇨🇦",
      "skin tone support": false,
      "name": "flag canada",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇨🇨",
      "skin tone support": false,
      "name": "flag cocos islands",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇨🇩",
      "skin tone support": false,
      "name": "flag congo kinshasa",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇨🇫",
      "skin tone support": false,
      "name": "flag central african republic",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇨🇬",
      "skin tone support": false,
      "name": "flag congo brazzaville",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇨🇭",
      "skin tone support": false,
      "name": "flag switzerland",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇨🇮",
      "skin tone support": false,
      "name": "flag c te d ivoire",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇨🇰",
      "skin tone support": false,
      "name": "flag cook islands",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇨🇱",
      "skin tone support": false,
      "name": "flag chile",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇨🇲",
      "skin tone support": false,
      "name": "flag cameroon",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇨🇳",
      "skin tone support": false,
      "name": "flag china",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇨🇴",
      "skin tone support": false,
      "name": "flag colombia",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇨🇵",
      "skin tone support": false,
      "name": "flag clipperton island",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇨🇷",
      "skin tone support": false,
      "name": "flag costa rica",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇨🇺",
      "skin tone support": false,
      "name": "flag cuba",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇨🇻",
      "skin tone support": false,
      "name": "flag cape verde",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇨🇼",
      "skin tone support": false,
      "name": "flag cura ao",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇨🇽",
      "skin tone support": false,
      "name": "flag christmas island",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇨🇾",
      "skin tone support": false,
      "name": "flag cyprus",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇨🇿",
      "skin tone support": false,
      "name": "flag czechia",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇩🇪",
      "skin tone support": false,
      "name": "flag germany",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇩🇬",
      "skin tone support": false,
      "name": "flag diego garcia",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇩🇯",
      "skin tone support": false,
      "name": "flag djibouti",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇩🇰",
      "skin tone support": false,
      "name": "flag denmark",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇩🇲",
      "skin tone support": false,
      "name": "flag dominica",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇩🇴",
      "skin tone support": false,
      "name": "flag dominican republic",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇩🇿",
      "skin tone support": false,
      "name": "flag algeria",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇪🇦",
      "skin tone support": false,
      "name": "flag ceuta melilla",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇪🇨",
      "skin tone support": false,
      "name": "flag ecuador",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇪🇪",
      "skin tone support": false,
      "name": "flag estonia",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇪🇬",
      "skin tone support": false,
      "name": "flag egypt",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇪🇭",
      "skin tone support": false,
      "name": "flag western sahara",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇪🇷",
      "skin tone support": false,
      "name": "flag eritrea",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇪🇸",
      "skin tone support": false,
      "name": "flag spain",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇪🇹",
      "skin tone support": false,
      "name": "flag ethiopia",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇪🇺",
      "skin tone support": false,
      "name": "flag european union",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇫🇮",
      "skin tone support": false,
      "name": "flag finland",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇫🇯",
      "skin tone support": false,
      "name": "flag fiji",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇫🇰",
      "skin tone support": false,
      "name": "flag falkland islands",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇫🇲",
      "skin tone support": false,
      "name": "flag micronesia",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇫🇴",
      "skin tone support": false,
      "name": "flag faroe islands",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇫🇷",
      "skin tone support": false,
      "name": "flag france",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇬🇦",
      "skin tone support": false,
      "name": "flag gabon",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇬🇧",
      "skin tone support": false,
      "name": "flag united kingdom",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇬🇩",
      "skin tone support": false,
      "name": "flag grenada",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇬🇪",
      "skin tone support": false,
      "name": "flag georgia",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇬🇫",
      "skin tone support": false,
      "name": "flag french guiana",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇬🇬",
      "skin tone support": false,
      "name": "flag guernsey",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇬🇭",
      "skin tone support": false,
      "name": "flag ghana",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇬🇮",
      "skin tone support": false,
      "name": "flag gibraltar",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇬🇱",
      "skin tone support": false,
      "name": "flag greenland",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇬🇲",
      "skin tone support": false,
      "name": "flag gambia",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇬🇳",
      "skin tone support": false,
      "name": "flag guinea",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇬🇵",
      "skin tone support": false,
      "name": "flag guadeloupe",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇬🇶",
      "skin tone support": false,
      "name": "flag equatorial guinea",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇬🇷",
      "skin tone support": false,
      "name": "flag greece",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇬🇸",
      "skin tone support": false,
      "name": "flag south georgia south sandwich islands",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇬🇹",
      "skin tone support": false,
      "name": "flag guatemala",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇬🇺",
      "skin tone support": false,
      "name": "flag guam",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇬🇼",
      "skin tone support": false,
      "name": "flag guinea bissau",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇬🇾",
      "skin tone support": false,
      "name": "flag guyana",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇭🇰",
      "skin tone support": false,
      "name": "flag hong kong sar china",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇭🇲",
      "skin tone support": false,
      "name": "flag heard mcdonald islands",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇭🇳",
      "skin tone support": false,
      "name": "flag honduras",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇭🇷",
      "skin tone support": false,
      "name": "flag croatia",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇭🇹",
      "skin tone support": false,
      "name": "flag haiti",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇭🇺",
      "skin tone support": false,
      "name": "flag hungary",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇮🇨",
      "skin tone support": false,
      "name": "flag canary islands",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇮🇩",
      "skin tone support": false,
      "name": "flag indonesia",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇮🇪",
      "skin tone support": false,
      "name": "flag ireland",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇮🇱",
      "skin tone support": false,
      "name": "flag israel",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇮🇲",
      "skin tone support": false,
      "name": "flag isle of man",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇮🇳",
      "skin tone support": false,
      "name": "flag india",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇮🇴",
      "skin tone support": false,
      "name": "flag british indian ocean territory",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇮🇶",
      "skin tone support": false,
      "name": "flag iraq",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇮🇷",
      "skin tone support": false,
      "name": "flag iran",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇮🇸",
      "skin tone support": false,
      "name": "flag iceland",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇮🇹",
      "skin tone support": false,
      "name": "flag italy",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇯🇪",
      "skin tone support": false,
      "name": "flag jersey",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇯🇲",
      "skin tone support": false,
      "name": "flag jamaica",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇯🇴",
      "skin tone support": false,
      "name": "flag jordan",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇯🇵",
      "skin tone support": false,
      "name": "flag japan",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇰🇪",
      "skin tone support": false,
      "name": "flag kenya",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇰🇬",
      "skin tone support": false,
      "name": "flag kyrgyzstan",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇰🇭",
      "skin tone support": false,
      "name": "flag cambodia",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇰🇮",
      "skin tone support": false,
      "name": "flag kiribati",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇰🇲",
      "skin tone support": false,
      "name": "flag comoros",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇰🇳",
      "skin tone support": false,
      "name": "flag st kitts nevis",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇰🇵",
      "skin tone support": false,
      "name": "flag north korea",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇰🇷",
      "skin tone support": false,
      "name": "flag south korea",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇰🇼",
      "skin tone support": false,
      "name": "flag kuwait",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇰🇾",
      "skin tone support": false,
      "name": "flag cayman islands",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇰🇿",
      "skin tone support": false,
      "name": "flag kazakhstan",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇱🇦",
      "skin tone support": false,
      "name": "flag laos",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇱🇧",
      "skin tone support": false,
      "name": "flag lebanon",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇱🇨",
      "skin tone support": false,
      "name": "flag st lucia",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇱🇮",
      "skin tone support": false,
      "name": "flag liechtenstein",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇱🇰",
      "skin tone support": false,
      "name": "flag sri lanka",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇱🇷",
      "skin tone support": false,
      "name": "flag liberia",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇱🇸",
      "skin tone support": false,
      "name": "flag lesotho",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇱🇹",
      "skin tone support": false,
      "name": "flag lithuania",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇱🇺",
      "skin tone support": false,
      "name": "flag luxembourg",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇱🇻",
      "skin tone support": false,
      "name": "flag latvia",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇱🇾",
      "skin tone support": false,
      "name": "flag libya",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇲🇦",
      "skin tone support": false,
      "name": "flag morocco",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇲🇨",
      "skin tone support": false,
      "name": "flag monaco",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇲🇩",
      "skin tone support": false,
      "name": "flag moldova",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇲🇪",
      "skin tone support": false,
      "name": "flag montenegro",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇲🇫",
      "skin tone support": false,
      "name": "flag st martin",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇲🇬",
      "skin tone support": false,
      "name": "flag madagascar",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇲🇭",
      "skin tone support": false,
      "name": "flag marshall islands",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇲🇰",
      "skin tone support": false,
      "name": "flag north macedonia",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇲🇱",
      "skin tone support": false,
      "name": "flag mali",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇲🇲",
      "skin tone support": false,
      "name": "flag myanmar",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇲🇳",
      "skin tone support": false,
      "name": "flag mongolia",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇲🇴",
      "skin tone support": false,
      "name": "flag macao sar china",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇲🇵",
      "skin tone support": false,
      "name": "flag northern mariana islands",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇲🇶",
      "skin tone support": false,
      "name": "flag martinique",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇲🇷",
      "skin tone support": false,
      "name": "flag mauritania",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇲🇸",
      "skin tone support": false,
      "name": "flag montserrat",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇲🇹",
      "skin tone support": false,
      "name": "flag malta",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇲🇺",
      "skin tone support": false,
      "name": "flag mauritius",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇲🇻",
      "skin tone support": false,
      "name": "flag maldives",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇲🇼",
      "skin tone support": false,
      "name": "flag malawi",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇲🇽",
      "skin tone support": false,
      "name": "flag mexico",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇲🇾",
      "skin tone support": false,
      "name": "flag malaysia",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇲🇿",
      "skin tone support": false,
      "name": "flag mozambique",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇳🇦",
      "skin tone support": false,
      "name": "flag namibia",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇳🇨",
      "skin tone support": false,
      "name": "flag new caledonia",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇳🇪",
      "skin tone support": false,
      "name": "flag niger",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇳🇫",
      "skin tone support": false,
      "name": "flag norfolk island",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇳🇬",
      "skin tone support": false,
      "name": "flag nigeria",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇳🇮",
      "skin tone support": false,
      "name": "flag nicaragua",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇳🇱",
      "skin tone support": false,
      "name": "flag netherlands",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇳🇴",
      "skin tone support": false,
      "name": "flag norway",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇳🇵",
      "skin tone support": false,
      "name": "flag nepal",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇳🇷",
      "skin tone support": false,
      "name": "flag nauru",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇳🇺",
      "skin tone support": false,
      "name": "flag niue",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇳🇿",
      "skin tone support": false,
      "name": "flag new zealand",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇴🇲",
      "skin tone support": false,
      "name": "flag oman",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇵🇦",
      "skin tone support": false,
      "name": "flag panama",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇵🇪",
      "skin tone support": false,
      "name": "flag peru",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇵🇫",
      "skin tone support": false,
      "name": "flag french polynesia",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇵🇬",
      "skin tone support": false,
      "name": "flag papua new guinea",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇵🇭",
      "skin tone support": false,
      "name": "flag philippines",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇵🇰",
      "skin tone support": false,
      "name": "flag pakistan",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇵🇱",
      "skin tone support": false,
      "name": "flag poland",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇵🇲",
      "skin tone support": false,
      "name": "flag st pierre miquelon",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇵🇳",
      "skin tone support": false,
      "name": "flag pitcairn islands",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇵🇷",
      "skin tone support": false,
      "name": "flag puerto rico",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇵🇸",
      "skin tone support": false,
      "name": "flag palestinian territories",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇵🇹",
      "skin tone support": false,
      "name": "flag portugal",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇵🇼",
      "skin tone support": false,
      "name": "flag palau",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇵🇾",
      "skin tone support": false,
      "name": "flag paraguay",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇶🇦",
      "skin tone support": false,
      "name": "flag qatar",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇷🇪",
      "skin tone support": false,
      "name": "flag r union",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇷🇴",
      "skin tone support": false,
      "name": "flag romania",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇷🇸",
      "skin tone support": false,
      "name": "flag serbia",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇷🇺",
      "skin tone support": false,
      "name": "flag russia",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇷🇼",
      "skin tone support": false,
      "name": "flag rwanda",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇸🇦",
      "skin tone support": false,
      "name": "flag saudi arabia",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇸🇧",
      "skin tone support": false,
      "name": "flag solomon islands",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇸🇨",
      "skin tone support": false,
      "name": "flag seychelles",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇸🇩",
      "skin tone support": false,
      "name": "flag sudan",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇸🇪",
      "skin tone support": false,
      "name": "flag sweden",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇸🇬",
      "skin tone support": false,
      "name": "flag singapore",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇸🇭",
      "skin tone support": false,
      "name": "flag st helena",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇸🇮",
      "skin tone support": false,
      "name": "flag slovenia",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇸🇯",
      "skin tone support": false,
      "name": "flag svalbard jan mayen",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇸🇰",
      "skin tone support": false,
      "name": "flag slovakia",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇸🇱",
      "skin tone support": false,
      "name": "flag sierra leone",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇸🇲",
      "skin tone support": false,
      "name": "flag san marino",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇸🇳",
      "skin tone support": false,
      "name": "flag senegal",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇸🇴",
      "skin tone support": false,
      "name": "flag somalia",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇸🇷",
      "skin tone support": false,
      "name": "flag suriname",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇸🇸",
      "skin tone support": false,
      "name": "flag south sudan",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇸🇹",
      "skin tone support": false,
      "name": "flag s o tom pr ncipe",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇸🇻",
      "skin tone support": false,
      "name": "flag el salvador",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇸🇽",
      "skin tone support": false,
      "name": "flag sint maarten",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇸🇾",
      "skin tone support": false,
      "name": "flag syria",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇸🇿",
      "skin tone support": false,
      "name": "flag eswatini",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇹🇦",
      "skin tone support": false,
      "name": "flag tristan da cunha",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇹🇨",
      "skin tone support": false,
      "name": "flag turks caicos islands",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇹🇩",
      "skin tone support": false,
      "name": "flag chad",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇹🇫",
      "skin tone support": false,
      "name": "flag french southern territories",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇹🇬",
      "skin tone support": false,
      "name": "flag togo",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇹🇭",
      "skin tone support": false,
      "name": "flag thailand",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇹🇯",
      "skin tone support": false,
      "name": "flag tajikistan",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇹🇰",
      "skin tone support": false,
      "name": "flag tokelau",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇹🇱",
      "skin tone support": false,
      "name": "flag timor leste",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇹🇲",
      "skin tone support": false,
      "name": "flag turkmenistan",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇹🇳",
      "skin tone support": false,
      "name": "flag tunisia",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇹🇴",
      "skin tone support": false,
      "name": "flag tonga",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇹🇷",
      "skin tone support": false,
      "name": "flag turkey",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇹🇹",
      "skin tone support": false,
      "name": "flag trinidad tobago",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇹🇻",
      "skin tone support": false,
      "name": "flag tuvalu",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇹🇼",
      "skin tone support": false,
      "name": "flag taiwan",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇹🇿",
      "skin tone support": false,
      "name": "flag tanzania",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇺🇦",
      "skin tone support": false,
      "name": "flag ukraine",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇺🇬",
      "skin tone support": false,
      "name": "flag uganda",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇺🇲",
      "skin tone support": false,
      "name": "flag u s outlying islands",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇺🇳",
      "skin tone support": false,
      "name": "flag united nations",
      "unicode version": "6.0",
      "emoji version": "4.0"
    },
    {
      "emoji": "🇺🇸",
      "skin tone support": false,
      "name": "flag united states",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇺🇾",
      "skin tone support": false,
      "name": "flag uruguay",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇺🇿",
      "skin tone support": false,
      "name": "flag uzbekistan",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇻🇦",
      "skin tone support": false,
      "name": "flag vatican city",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇻🇨",
      "skin tone support": false,
      "name": "flag st vincent grenadines",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇻🇪",
      "skin tone support": false,
      "name": "flag venezuela",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇻🇬",
      "skin tone support": false,
      "name": "flag british virgin islands",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇻🇮",
      "skin tone support": false,
      "name": "flag u s virgin islands",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇻🇳",
      "skin tone support": false,
      "name": "flag vietnam",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇻🇺",
      "skin tone support": false,
      "name": "flag vanuatu",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇼🇫",
      "skin tone support": false,
      "name": "flag wallis futuna",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇼🇸",
      "skin tone support": false,
      "name": "flag samoa",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇽🇰",
      "skin tone support": false,
      "name": "flag kosovo",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇾🇪",
      "skin tone support": false,
      "name": "flag yemen",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇾🇹",
      "skin tone support": false,
      "name": "flag mayotte",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇿🇦",
      "skin tone support": false,
      "name": "flag south africa",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇿🇲",
      "skin tone support": false,
      "name": "flag zambia",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🇿🇼",
      "skin tone support": false,
      "name": "flag zimbabwe",
      "unicode version": "6.0",
      "emoji version": "2.0"
    },
    {
      "emoji": "🏴󠁧󠁢󠁥󠁮󠁧󠁿",
      "skin tone support": false,
      "name": "flag england",
      "unicode version": "7.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🏴󠁧󠁢󠁳󠁣󠁴󠁿",
      "skin tone support": false,
      "name": "flag scotland",
      "unicode version": "7.0",
      "emoji version": "5.0"
    },
    {
      "emoji": "🏴󠁧󠁢󠁷󠁬󠁳󠁿",
      "skin tone support": false,
      "name": "flag wales",
      "unicode version": "7.0",
      "emoji version": "5.0"
    }
  ]
	}
]
}
`
