module Main.User exposing (User, Users, noUserYet, userDecoder, usersDecoder)

import Json.Decode as JD
import Json.Decode.Pipeline as JP
import Json.Encode as JE


type alias User =
    { id : Int
    , name : String
    }


type alias Users =
    { entries : List User
    , count : Int
    , page : Int
    , limit : Int
    }


noUserYet : User
noUserYet =
    { id = 0
    , name = ""
    }


usersDecoder : JD.Decoder Users
usersDecoder =
    JP.decode Users
        |> JP.required "Entries" (JD.list userDecoder)
        |> JP.required "Count" JD.int
        |> JP.required "Limit" JD.int
        |> JP.required "Page" JD.int


userDecoder : JD.Decoder User
userDecoder =
    JP.decode User
        |> JP.required "id" JD.int
        |> JP.required "name" JD.string
