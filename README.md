# matchingservice

Simple Go Matching service. 

My initial idea was to create this using web sockets, as I think it's the right way to go. Having said that, I didn't have the biggest experience with sockets in golang,
only used them a few times in Node.js applications.

So there are actually 2 versions of this. REST and socket. Both can work at the same time but they run at different ports. REST is running on 8080 and sockets are on 8000. But the data is shared so if you put a player in a game inside rest, it will be shown on the socket part as well.

One more important thing to note is that I created this all several hours, because the task said it should not take longer. So because of that, I didn't want to lose much time creating databases and storing games there, dockerizing, etc. I used 2 maps, one for games and one for users and they are being filled as long as the app is running. Database would not change much, would just take much longer to setup, so I decided not to bother with that. Also there was some optimization to be done, with for example `getCurrentGame` function that loops through all, and finds the first that is not filled up. Due to limit of users being 100 (so games being limited to 25), there was no harm done with one for loop. If the number would get much bigger, keeping track of id of current game being filled would be my first optimization. 

Now for the versions:

# REST 

Rest version has 3 routes `/join` `/session` and `/leave` that do what they are supposed to, and you can test that.

These are the curl requests

## Session -> Prints all rooms, and players inside them, if the room has 4 players it will say `Started` else it will say `Waiting`

```
curl --location --request GET 'localhost:8080/session'
```

## Join -> Joins to the first available room.

```
curl --location --request POST 'localhost:8080/join' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "Test"
}'
```

## Leave -> Player is removed from the room, and that room is again waiting for players

```
curl --location --request POST 'localhost:8080/leave' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "Test",
    "game_id": "Game 1"
}'
```

# Socket 

Socket version is not as stable but shows how people would be connecting to games via sockets. Once you uncomment socket version, there is a index.html that you can visit on `localhost:8000/` and it has 2 input forms. First input form adds user to games and other drops a user from games. The one that adds, just uses one input field, and you just provide it with username.

Other takes in 2 fields: username and game_id, and deletes the user from that game. Of course if the username is not valid, it's handled, and same goes for game_id.


# Conclusion

Again, only took me about few hours. Of course with bigger time limit I would maybe refactor the whole thing, but I still feel that for one day task, I kept the clone relatively clean. 
