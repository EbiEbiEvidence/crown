# cource.go_server

## Background
Crowns is a simplified clone of [Reigns](https://www.devolverdigital.com/games/reigns).

It is a simple kingdom simulation. Your goal, as King, is simply to stay alive as long as you can.

There are 4 constituencies (aka Factions) in your kingdom.  

1. The Church
1. The People
1. The Military
1. The Merchants (aka your Treasury)  

You stay alive by balancing the needs of each of the Factions and avoiding you or your kingdom falling into ruin or being overthrown.

## Game play
### Faction power levels
The Factions have power levels that range from 0 to 100. 
If the power level of any Faction reaches 100, you die.
If the power level of any Faction reaches 0, you die.

Each Factions is visible as icons on the screen, with their power levels.

![factions](https://i.imgur.com/dg7Uwki.png, "factions")

### Event cards
The state of the game world transitions based on cards that you are dealt.

Each card presents a new situation that affects the kingdom. You are then presented two choices about what to do in response to that situation. 

You can “swipe left” or “swipe right” (like "Tinder for politics") to choose. 

Each choice will have some impact on the game state. Usually the impact is in the form of changing one or more Faction power levels (although sometimes it may just be a fun outcome like petting a dog).
As you begin to swipe left or right, the effect of that proposed choice will display over the Faction icons as a small or large dot, depending on whether the result will decrease or increase the energy level for that Faction.

![icon dot](https://i.imgur.com/DLAp1Ff.png, "icon dot")

### "Scenarios" 
If every card is totally random, the whole experience might feel too disjointed and not be very fun to play.  Therefore, some cards are associated with each other in a sub-tree called a Scenario.  Events unfold depending on your choices within a self-contained sub-narrative.  

![senario](https://i.imgur.com/mDkO5X4.png, "senario")

Here's a more-concrete Scenario example:

![concrete senario](https://i.imgur.com/Hf3AA79.png, "concrete_senario")


### Intermixing of Scenarios and one-off event cards

After a Scenario is complete, the system will either choose another Scenario, or it may choose an individual card to play.

I.e,, full Scenarios and one-off event cards intermix, like this:

![inter_mixing](https://i.imgur.com/ogmVEfg.png, "inter_mixing")

### Age of your reign

The more cards you play, the older you get (i.e., the longer your reign).

### Start age
Your age when you start is always 14 years old. 

### Ratio of cards to years
There is a fixed ratio of cards-played to years elapsed.  **3 cards is 1 year**.

### Situation bonuses
In addition, each Situation resolution awards an age bonus (which varies per Situation).

For example, Situation #6 may represent a protracted border conflict, and play out over a year. If you reach a non-Death outcome, you get an extra year added to your age.

## Game ending
### If you die...
There are a variety of unpleasant deaths, depending on the choices you made.

Each death is a Situation card as well, but with no choices.

### If you stay alive long enough...
The game is "won" when you live to be at least 94 years old. 

You can continue playing if you wish, to get a higher score.  But in any case, you will die for sure when you hit 104 years old.

## Leaderboard
The game maintains a leaderboard of reigns, sorted by length in descending order. 

In the case of a tie in number of years, the reigns are secondary-sorted by the total health points across all 4 constituencies.

## FAQ
Is this a multiplayer game?  Or will it ever be? 
No and no.