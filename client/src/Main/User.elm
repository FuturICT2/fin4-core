module Main.User exposing (User, noUserYet, userDecoder)

import Json.Decode as JD
import Json.Decode.Pipeline as JP
import Json.Encode as JE


type alias User =
    { id : Int
    , email : String
    }


noUserYet : User
noUserYet =
    { id = 0
    , email = ""
    }


userDecoder : JD.Decoder User
userDecoder =
    JP.decode User
        |> JP.required "id" JD.int
        |> JP.required "email" JD.string
