package consts

//									global
//-----------------------------------------------------------------------------

//Const containing string to be sent if decoding fails
const DecodingFailed string = "Something wrong happened when decoding data"

//								  handlers.go
//-----------------------------------------------------------------------------

//Help contains all the commands available
const Help string = "```Current commands are:\n\tping\n\tcard <card name>\n\tdice <die sides>\n\tinsult\n\tadvice\n\tkanye"
const MusicHelp string = "\n\nMusic commands:\n\tjoin\n\tleave\n\tplay <youtube url/query>\n\tskip\n\tstop```"

//									card.go
//-----------------------------------------------------------------------------

//Const containing the root of the url
const ScryfallBaseURL string = "https://api.scryfall.com/cards/named?fuzzy="

//Const containing string to be sent if scryfall API is unavailable
const ScryfallNotAvailable string = "Scryfall API not available at the moment."

//									advice.go
//-----------------------------------------------------------------------------

//contains url to adviceslip API
const AdviceSlipURL string = "https://api.adviceslip.com/advice"

//Const containing string to be sent if adviceslip API is unavailable
const AdviceslipNotAvailable string = "Adviceslip API not available at the moment."

//									insults.go
//-----------------------------------------------------------------------------

//insultURL contains the url for the API generating insults
const InsultURL string = "https://evilinsult.com/generate_insult.php?lang=en&type=json"

//String to be sent if Evil Insult API isn't available
const EvilInsultNotAvailable string = "Evil Insult API not available at the moment. Please try again later."

//									youtube.go
//-----------------------------------------------------------------------------

const YoutubeSearchEndpoint string = "https://www.googleapis.com/youtube/v3/search?part=snippet&type=video&key="
const YoutubeFindEndpoint string = "https://www.googleapis.com/youtube/v3/videos?part=snippet&key="

const YtVideoUrl string = "https://www.youtube.com/watch?v="



//									kanye.go
//-----------------------------------------------------------------------------
const KanyeRestEndpoint string = "https://api.kanye.rest/"
const KanyeRestUnavailable string = "Oops, something went wrong when getting Kanye quote"
