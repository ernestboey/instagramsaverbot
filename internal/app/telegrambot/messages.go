package telegrambot

const (
	errMsg       = "%s. Please try sending an Instagram link that contains a photo. We currently do not support videos 🙇‍♂️."
	foundLinkMsg = "I found at least one Instagram post from your message."
	helpMsg      = `You looked for help!

	The available commands I can handle are:
	- /save https://instagram.com/p/link - Gives you an album of image(s)
	- /help - You are viewing the help command now

	Happy saving! 😄
	`
	welcomeMsg = `Hello 👋! Hi!

	To utilize me, send me a message like:
	/save https://instagram.com/p/link
	I will then return you an album of image(s).

	If you need help, use the /help command. 😄
	`
)
