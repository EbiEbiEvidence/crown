# The Crowns Server specification

The Crowns server will implement a subset of functionalities for the Crowns game. In particular, the server is able to:
- Store each player's name, which is used to login and identify the players. There is no password.
- Generate a starting deck of cards for each crowns game.
- Record the highscore for each game the user plays.
- Generate a list of highscores for each player.
- Finally, generate a list of all highscores by all players that played on this server.

## data shapes

The Crowns server in general returns the following data shapes, using JSON:
* User

```
{
	"name": <string>,
	"token": <string>
}
```

The user object contains the name of the user. The name is case-sensitive, meaning that it should retain the same capitalization that the user enters. However, once a user account is made, no other user can use a name that differs with the existing name only by capitalizations.

i.e. If the user 'Tiger' exists, there cannot be another user with the name 'tiger'.

The token is a secret token assigned to the user as an API key. Certain request will only be processed if the correct token is included. As a JSON field in the user object, token is optional; the server does not need to include the token in its response every time.

* Highscore

```
{
	"name": <string>, // name of the player
	"submitted_at": <timestamp>, // when this score was submitted
	"start": <int>, // the year, i.e. 1998, that the reign started
	"age": <int>, // the age of the king when he is dead
	"scores": {
		"church": <int>, // the total church score
		"commoners": <int>,
		"merchants": <int>,
		"military": <int>,
		"bonus": <int> // some scenarios may give bonus scores to the player
	}
}
```

* Card

```
{
	"id": <int>,
	"scenario_id": <int>, // the id of the root card of this scenario tree
	"card_text": <string>, // the text on this card
	"card_image": <string>,
	"card_type": <string>, // can be either 'EVENT' or 'DEATH'
	"left_choice": <choice object>,
	"right_choice": <choice object>
}
```

* Choice

```
{
	"id": <int>,
	"next": <card object> or null, // the next card following this choice
	"choice_text": <string>, // short text describing the choice
	"church_modifier": <int>, // the score to add onto church score
	"commoners_modifier": <int>,
	"merchants_modifier": <int>,
	"military_modifier": <int>,
	"bonus_modifier": <int>
}
```

As can be seen here, the Card and Choice object are the nodes and edges in a tree data structure. When a choice leads to a NULL card, the scenario is complete, and the next scenario in the game (or single card) can be dealt.

And finally:

* Game

```
{
	"cards": [<card object>] // either single card or scenario root
}
```

* Highscores List

```
{
	"highscores": [<highscore object>] // a sorted list of highscores
}
```



## endpoints

For a first version, the server will implement the following interface:

* Create a user: `POST /user/create`
The body of the POST request should be:
```
{
	"name": <string>
}
```
The server responds with a `<user object>` which includes the secret token.
If there is a user with the same name, respond 400 with a message "user already exists".

* Getting a user: `POST /user/get`
The body of the POST request should be one of the following two:
```
{
	"name": <string>
}
```
```
{
	"token": <string>
}
```
The server responds with a `<user object>` which includes the secret token.
If there isn't a user with a matching name (case insensitive) or matching token, respond 404 with a message "user not found".

* Getting highscores from a user: `POST /user/highscores`
The body of the POST request should be:
```
{
	"token": <string>
}
```
The server responds with a `<highscores list object>` which includes the top 5 highscores by this user. The highscores must be sorted by total scores in descending order.
If there isn't a user with a matching token, respond 404 with a message "user not found".

* Getting highscores from all users: `GET /highscores`
The server responds with a `<highscores list object>` which includes the top 5 highscores by all users. The highscores must be sorted by total scores in descending order.

* Submitting a highscore: `POST /highscores/submit`
The body of the POST request should be:
```
{
	"token": <string>
	"highscore": <highscore object>
}
```
The server responds 200 with a message "ok".
If there isn't a user with a matching token, respond 404 with a message "user not found".

* Getting a new game: `POST /game/new`
The body of the POST request should be:
```
{
	"token": <string>
}
```
The server responds with a `<game object>` which includes a deck of at least 270 cards.
If there isn't a user with a matching token, respond 404 with a message "user not found".