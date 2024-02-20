## Task

Everyone knows Pokémon. So far 7 generations of pokémon released in last 20 years. In pokémon universe there are 18 types of these pocket monsters. 
The types are normal, fight, flying, poison, ground, rock, bug, ghost, steel, fire, water, grass, steel, fire, water, grass, electric, psychc, ice, dragon, dark, and, fairy. Each one of them has advanteges and disadvantes to others. For example; a **fire** type pokémon hits 2x to a **grass** type but hits 0.5x to a **rock** pokémon. You can find the complete list at [this url](http://unrealitymag.com/wp-content/uploads/2014/11/fylyCdC.png).

In this task you are going to create a **pokedex** which will be accessible with the HTTP api. The data source will be given to you. You can use Python or Go programming languages.


### HTTP API

**List Api**

    $ curl http://localhost:8080/list?type=Grass

Lists the given type of pokémons. A sample output is given below. You can (and should) change the output. You should consider the next and previous evolutions and pokemon's second types.  
_Sample Output_:

    Bulbasaur
        Weight: 6,9 kg
        Height: 0,7 m
        BaseAttack: 118
        BaseDefense: 118
        BaseStamina: 90
        Next evolutions:
            Ivysaur
            Venusaur
        ...
        ...

    Oddish:
        Weight: 5,4 kg
        Height: 0,5 m
        BaseAttack: 131
        BaseDefense: 116
        BaseStamina: 90
        Next evolutions:
            Gloom
            Vileplume
            Bellossom
        ...
        ...

_How to improve_:

You can list pokemons by their attacks, powers or anything. 
Also you can list types for better user experience like below

    $ curl http://localhost:8080/list/types

You can and sorting to the list api.
    
    $ curl http://localhost:8080/list?type=Bug?sortby=BaseAttack


**Get Api**

    $ curl http://localhost:8080/get?name=Pikachu

    or
    
    $curl http://localhost:8080/get/Pikachu
    
    or 

    $ curl http://locahost:8080/Pikachu


The get api prints information abput `Pokemon`, `Type` or `Move`. The get api form can be any of the above.
_Sample Output_:

    Pokemon Type: Bug
    Effective Against:
    - Dark
    - Grass
    - Psychic
    Weak Against:
    - Fairy
    - Fighting
    - Fire
    - Flying
    - Ghost
    - Poison
    - Steel
    Example Pokemons:
    - Metapod
    - Butterfree

_How to improve_:

The **Get Api** should be easy to use and easy to understand the result. You can add properties like the worst enemies and lovely enemies. 


### Data Format

The given json data has the following structure.

    {
        "types": [...],
        "pokemons": [...],
        "moves": [...]
    }


**types** are the pokemon types. A type object has `name`, `effectiveAgainst` and `weakAgainst` fields. `effectiveAgainst` types are the types that hit by this pokemon type 2x. `weakAgainst` types are the types that hit by this pokemon type 0.5x.

    {
        "name": "Ground",
        "effectiveAgainst": [
            "Electric",
            "Fire",
            "Poison",
            "Rock",
            "Steel"
        ],
        "weakAgainst": [
            "Bug",
            "Flying",
            "Grass"
        ]
    }

**pokemons** are the pokemon objects. It has all the necessary information about a pokemon. The three properties `BaseAttack` `BaseDefense`, and `BaseStamina` determines the pokemon's power.
`Fast Attack(s)` and `Special Attack(s)` are moves of this pokemon.

**moves** aret he attack objects of the pokemon.

    {
        "id": 13,           // ID of the move
        "name": "Wrap",     // Name of the move
        "type": "Ice",      // Type of the move
        "damage": 60,       // Damage of the move
        "energy": 33,       // Energy requirement of the move
        "dps": 20.69,       // Damage Per Second
        "duration": 2900    // Cooldown Duration
    }

http://www.pokemongodb.net/ has much more information about attacks, types and pokemons.

### Challenges

- Read Pokemon, Type and Move data from given json

A file with containing all the necessary information will be given to you as JSON data format.

- Create a HTTP server for input

All the input and output with the user must be with HTTP API. All error handling and input must be considered.
You can/should change and improve the API given below for better user experience

- User Experience should be awesome.

Although we use the app within the terminal; the API must be easy to use and easy the understand the results (unlike this task). What is a user friendly api? For example; the following api call is friendly for the developer because the developer knowns the resource id and resource type.

    $ curl http://localhost:8080/getDetailedInformation?id=062&type=pokemon

But it is hard the use by the user. Instead, create an api like below: 

    $ curl http://localhost:8080/Poliwrath

With this api call; user thinks that he or she knowns the resource name and just wants it.

- Comments and Documentation

We should be able to understand your code without trouble. Try to write comments to tell us whats in your mind. 

## Installing Go Programming Language
Although using go optional, it is recommended.  
Install go: https://golang.org/dl/
or if you are using MacOS use brew to install go by just: `brew install go`  
**Tip:** Use "golang" instead of "go" for searching on google
Best resource to learn Go: https://gobyexample.com/  
Read gobyexample website until **"Methods"** chapter (Methods including).

Also look for below links for this task:

JSON: https://gobyexample.com/json

Reading files: https://gobyexample.com/reading-files

Writing files: https://gobyexample.com/writing-files

Url parsing: https://gobyexample.com/url-parsing

String to int: https://gobyexample.com/number-parsing

String utils: https://golang.org/pkg/strings/

For more resources following link has amazing content: http://dave.cheney.net/resources-for-new-go-programmers 

A video from google (for mac and linux): https://www.youtube.com/watch?v=2KmHtgtEZ1s

**Tip:** Use https://code.visualstudio.com/ writing go code with "https://github.com/Microsoft/vscode-go" extension. And use gofmt tool to format your code.

